package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/robfig/cron/v3"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/crius"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/db"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	mqMessage "gitlab.omytech.com.cn/micro-service/Crius/pkgs/message"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/redis"
	cutil "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/cache"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/config"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gitlab.omytech.com.cn/micro-service/appointment/util"
	merchantBasic "gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	message "gitlab.omytech.com.cn/micro-service/message/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

// Server server
type Server struct {
	proto.UnimplementedAppointmentServerServer
	// 服务发现注册
	crius *crius.Client
	// 数据库
	database *db.Entity
	// redis
	redis *redis.Entity
	// cache
	cache *cache.Entity
	// 队列
	mq *mqMessage.MqttEntity
}

type messageField struct {
	customer      string //客户
	gender        string //性别
	branch        string //门店
	roomType      string //房型
	arriveDate    string //到店日期
	week          string //星期
	arriveAt      string //到店时间
	num           string //预约人数
	branchPhone   string //门店电话
	branchAddress string //门店地址
	navigation    string //门店导航
	deposit       string //订金
}

var (
	// merchant-basic服务
	merchantBasicClient merchantBasic.MerchantBasicServiceClient
	// message服务
	messageClient message.MessageServiceClient
)

// 用于时间时分秒比较
const timeFormate = "20060102150405"

// NewAppointmentServe 实例化
func NewAppointmentServe() (*Server, error) {
	client, err := crius.NewClient(context.Background(), crius.ClientConfig{Address: config.Setting.Crius.Address})
	if err != nil {
		return nil, fmt.Errorf("crius client失败:%s", err.Error())
	}

	entity, err := model.DatabaseConnection()
	if err != nil {
		return nil, fmt.Errorf("数据库连接错误:%s", err.Error())
	}

	redisClient, err := model.RedisConnection()
	if err != nil {
		panic(fmt.Sprintf("redis连接失败:%s", err.Error()))
	}

	cacheEntity, err := cache.NewEntity()
	if err != nil {
		return nil, fmt.Errorf("cache初始化错误:%v", err)
	}

	mq, err := mqMessage.NewMqtt(mqMessage.MqttConfig{
		Client:   config.Setting.Mqtt.Client,
		UserName: config.Setting.Mqtt.Username,
		Password: config.Setting.Mqtt.Password,
		Broker:   config.Setting.Mqtt.Broker,
		ReconnectingHandler: func(c mqtt.Client, options *mqtt.ClientOptions) {
		},
		ConnectionLostHandler: func(client mqtt.Client, err error) {
		},
	})
	if err != nil {
		return nil, fmt.Errorf("mqtt初始化错误:%v", err)
	}

	s := &Server{
		crius:    client,
		database: entity,
		redis:    redisClient,
		cache:    cacheEntity,
		mq:       mq,
	}

	c := cron.New()
	c.AddJob("*/1 * * * *", s)
	c.Start()
	return s, nil
}

// Heart 心跳
func (s *Server) Heart(ctx context.Context) error {
	cfg := config.Setting.App
	return s.crius.Heart(ctx, crius.HeartRequest{
		Name:   cfg.Name,
		Desc:   cfg.Desc,
		IP:     cfg.IP,
		Port:   cfg.Port,
		Weight: cfg.Weight,
	})
}

// SaveSnapshot 保存快照
func (s *Server) SaveSnapshot(ctx context.Context, snapshot model.TableSnapshot) {
	snapshot.ID = uuid.NewV4()
	metadata := pkgs.GetMetadata(ctx)
	snapshot.SleuthCode = metadata.SleuthCode
	snapshot.StaffID = metadata.StaffID
	if err := s.database.Conn.Create(&snapshot).Error; nil != err {
		cutil.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(snapshot)).Info("保存快照信息错误")
	}
}

func (s *Server) merchantBasic() merchantBasic.MerchantBasicServiceClient {
	if merchantBasicClient == nil {
		resp, err := s.crius.Discover(context.Background(), crius.DiscoverRequest{Name: config.Setting.Crius.MerchantBasic})
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("merchantBasic发现服务错误:%v, 服务名:%v", err, config.Setting.Crius.MerchantBasic))
			return nil
		}
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", resp.IP, resp.Port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("merchantBasic rpc连接错误err:%v, ip:port:%v", err, fmt.Sprintf("%s:%d", resp.IP, resp.Port)))
			return nil
		}
		merchantBasicClient = merchantBasic.NewMerchantBasicServiceClient(conn)
	}
	return merchantBasicClient
}

func (s *Server) message() message.MessageServiceClient {
	if messageClient == nil {
		resp, err := s.crius.Discover(context.Background(), crius.DiscoverRequest{Name: config.Setting.Crius.Message})
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("message发现服务错误:%v, 服务名:%v", err, config.Setting.Crius.Message))
			return nil
		}
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", resp.IP, resp.Port), grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("message rpc client错误err:%v, ip:port:%v", err, fmt.Sprintf("%s:%d", resp.IP, resp.Port)))
			return nil
		}
		messageClient = message.NewMessageServiceClient(conn)
	}
	return messageClient
}

// Run 定时job实现接口
func (s *Server) Run() {
	defer cutil.CatchException()
	var appointments []model.TableAppointment
	if err := s.database.Conn.Scopes(cutil.ColumnInScope("status", []interface{}{model.AppointmentStatusAppointed, model.AppointmentStatusArrearage})).Find(&appointments).Error; err != nil {
		cutil.Logger.Error(fmt.Sprintf("预约定时任务获取预约列表数据库错误:%v", err))
		return
	}

	for _, v := range appointments {
		if v.KeepAt != nil && time.Now().After(*v.KeepAt) && v.Status == model.AppointmentStatusArrearage {
			before := pkgs.MakeParams(v)
			v.Status = model.AppointmentStatusCancelled
			v.CancelledReason = model.AppointmentChangeCancelReason
			v.Operator = nil
			after := pkgs.MakeParams(v)
			if err := s.cronCancelAppointment(v); err == nil {
				go s.SaveSnapshot(pkgs.MetadataContent(pkgs.Metadata{MerchantID: *v.MerchantID}), model.TableSnapshot{
					SnapShotTableName: model.TableAppointment{}.TableName(),
					TableID:           nil,
					Method:            model.CancelMethod,
					Before:            &before,
					After:             &after,
					RelatedID:         v.RelatedID,
				})
			}
			continue
		} else if v.ExpireAt != nil && time.Now().After(*v.ExpireAt) && v.Status == model.AppointmentStatusAppointed {
			before := pkgs.MakeParams(v)
			v.Status = model.AppointmentStatusExpired
			v.Operator = nil
			v.Breach = true
			v.BreachReason = model.AppointmentBreachReasonNotArrive
			after := pkgs.MakeParams(v)
			if err := s.database.Conn.Select("status", "operator", "breach", "breach_reason").Updates(&v).Error; err != nil {
				cutil.Logger.Error(fmt.Sprintf("预约定时任务更新预约数据库错误:%v", err))
			} else {
				go s.SaveSnapshot(pkgs.MetadataContent(pkgs.Metadata{MerchantID: *v.MerchantID}), model.TableSnapshot{
					SnapShotTableName: model.TableAppointment{}.TableName(),
					TableID:           nil,
					Method:            model.UpdateMethod,
					Before:            &before,
					After:             &after,
					RelatedID:         v.RelatedID,
				})
			}
			continue
		}
		config, err := s.getAppointmentConfig(*v.MerchantID)
		if err != nil {
			cutil.Logger.Error(fmt.Sprintf("预约定时任务获取预约配置错误:%v", err))
			continue
		}
		if v.AppointmentAt != nil && time.Now().Add(time.Duration(config.RemindTime*60)*time.Minute).Format(timeFormate) == v.AppointmentAt.Format(timeFormate) {
			template, err := s.getMessageTemplate(*v.MerchantID, *v.BranchID)
			if err != nil {
				cutil.Logger.Error(fmt.Sprintf("预约定时任务获取短信模板错误:%v", err))
				continue
			}
			contents := s.buildMessage(v, template)
			go s.sendJobs(v, contents, template)
		}
	}
}

func timeToWeekDesc(t time.Time) string {
	switch t.Weekday() {
	case time.Monday:
		return "星期一"
	case time.Tuesday:
		return "星期二"
	case time.Wednesday:
		return "星期三"
	case time.Thursday:
		return "星期四"
	case time.Friday:
		return "星期五"
	case time.Saturday:
		return "星期六"
	case time.Sunday:
		return "星期日"
	}

	return ""
}

func (s *Server) buildMessage(appointment model.TableAppointment, template *message.GetBranchTemplateResponse_GetBranchTemplateResponseDate) []string {
	branchResp, err := s.merchantBasic().ShowBranch(pkgs.MetadataContent(pkgs.Metadata{MerchantID: *appointment.MerchantID}), &merchantBasic.ShowBranchRequest{Id: appointment.BranchID.String()})
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("预约定时任务获取门店信息错误:%v", err))
		return nil
	}
	if branchResp.ErrorCode != pkgs.Success {
		cutil.Logger.Error(fmt.Sprintf("预约定时任务获取门店信息失败:%v", branchResp.ErrorMessage))
		return nil
	}

	roomResp, err := s.merchantBasic().ShowRoomType(pkgs.MetadataContent(pkgs.Metadata{MerchantID: *appointment.MerchantID}), &merchantBasic.ShowRoomTypeRequest{Id: appointment.RoomTypeID.String()})
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("预约定时任务获取房型信息错误:%v", err))
		return nil
	}
	if branchResp.ErrorCode != pkgs.Success {
		cutil.Logger.Error(fmt.Sprintf("预约定时任务获取房型信息失败:%v", roomResp.ErrorMessage))
		return nil
	}

	gender := "先生"
	if appointment.Gender != 1 {
		gender = "女士"
	}
	field := messageField{
		customer:      appointment.Name,
		gender:        gender,
		branch:        branchResp.Data.Name,
		roomType:      roomResp.Data.Name,
		arriveDate:    appointment.AppointmentAt.Format("2006年01月02日"),
		week:          timeToWeekDesc(*appointment.AppointmentAt),
		arriveAt:      appointment.AppointmentAt.Format("15:04"),
		num:           fmt.Sprintf("%d", appointment.CustomerNum),
		branchPhone:   branchResp.Data.Phone,
		branchAddress: branchResp.Data.Address,
		navigation:    branchResp.Data.Location,
		deposit:       strconv.Itoa(int(appointment.DepositFee)),
	}
	content := template.SmsTemplate.Content
	replaceMessage(field, &content)
	contents := strings.Split(content, "[分页]")
	replaceMessage(field, &template.WechatTemplate.Content.First.Value)
	replaceMessage(field, &template.WechatTemplate.Content.Remark.Value)
	for i := 0; i < len(template.WechatTemplate.Content.Detail); i++ {
		replaceMessage(field, &template.WechatTemplate.Content.Detail[i].Value)
	}
	return contents
}

func replaceMessage(field messageField, content *string) {
	*content = strings.ReplaceAll(*content, "[换行]", "\n")
	*content = strings.ReplaceAll(*content, "[客户]", field.customer)
	*content = strings.ReplaceAll(*content, "[性别]", field.gender)
	*content = strings.ReplaceAll(*content, "[门店]", field.branch)
	*content = strings.ReplaceAll(*content, "[房型]", field.roomType)
	*content = strings.ReplaceAll(*content, "[到店日期]", field.arriveDate)
	*content = strings.ReplaceAll(*content, "[星期]", field.week)
	*content = strings.ReplaceAll(*content, "[到店时间]", field.arriveAt)
	*content = strings.ReplaceAll(*content, "[预约人数]", field.num)
	*content = strings.ReplaceAll(*content, "[门店电话]", field.branchPhone)
	*content = strings.ReplaceAll(*content, "[门店地址]", field.branchAddress)
	*content = strings.ReplaceAll(*content, "[门店导航]", field.navigation)
	*content = strings.ReplaceAll(*content, "[订金]", field.deposit)
}

func (s *Server) sendJobs(v model.TableAppointment, contents []string, template *message.GetBranchTemplateResponse_GetBranchTemplateResponseDate) {
	for _, cc := range template.Cc {
		for _, content := range contents {
			resp, err := s.message().SendSms(pkgs.MetadataContent(pkgs.Metadata{MerchantID: *v.MerchantID}), &message.SendSmsRequest{
				AreaCode:    cc.Code,
				Phone:       cc.Phone,
				Message:     content,
				System:      "预约系统",
				BranchId:    v.BranchID.String(),
				Sign:        template.SmsTemplate.Sign,
				MessageType: util.AppointmentExpirationRemind,
			})
			if err != nil || resp.ErrorCode != pkgs.Success {
				cutil.Logger.Error(fmt.Sprintf("预约定时任务发送短信失败,err:%v, resp:%v", err, resp))
			}
		}
	}
	for _, content := range contents {
		resp, err := s.message().SendSms(pkgs.MetadataContent(pkgs.Metadata{MerchantID: *v.MerchantID}), &message.SendSmsRequest{
			AreaCode: v.AppointmentCode,
			Phone:    v.AppointmentPhone,
			Message:  content,
			System:   "预约系统",
			BranchId: v.BranchID.String(),
			Sign:     template.SmsTemplate.Sign,
		})
		if err != nil || resp.ErrorCode != pkgs.Success {
			cutil.Logger.Error(fmt.Sprintf("预约定时任务发送短信失败,err:%v, resp:%v", err, resp))
		}
	}
	// 微信预约
	if v.Way&1 == 1 {
		memberResp, err := s.merchantBasic().ShowWechatUserByMember(pkgs.MetadataContent(pkgs.Metadata{MerchantID: *v.MerchantID}), &merchantBasic.ShowWechatUserByMemberRequest{MemberId: v.MemberID.String()})
		if err != nil || memberResp.ErrorCode != pkgs.Success {
			cutil.Logger.Error(fmt.Sprintf("预约定时任务发送微信查询微信用户信息失败,err:%v, resp:%v", err, memberResp))
			return
		}

		messageResp, err := s.message().SendWechatTemplate(pkgs.MetadataContent(pkgs.Metadata{MerchantID: *v.MerchantID}), &message.SendWechatTemplateRequest{
			System:     "appointment_server",
			Channel:    "appointment",
			TemplateId: template.WechatTemplate.Id,
			Content: &message.WechatTemplateContent{First: &message.WechatTemplateContentBase{
				Value: template.WechatTemplate.Content.First.Value},
				Detail: template.WechatTemplate.Content.Detail,
				Remark: template.WechatTemplate.Content.Remark,
			},
			BranchId:    v.BranchID.String(),
			MessageType: util.AppointmentExpirationRemind,
			WechatUser: &message.SendWechatTemplateRequest_WechatUser{
				MemberId:       v.MemberID.String(),
				MemberWechatId: memberResp.Data.Id,
				MemberOpenId:   memberResp.Data.Openid,
			},
			OfficialLink: template.WechatTemplate.OfficialLink,
		})
		if err != nil || messageResp.ErrorCode != pkgs.Success {
			cutil.Logger.Error(fmt.Sprintf("预约定时任务发送微信失败,err:%v, resp:%v", err, messageResp))
			return
		}
	}
}

func (s *Server) getMessageTemplate(merchantID, branchID uuid.UUID) (*message.GetBranchTemplateResponse_GetBranchTemplateResponseDate, error) {
	template := new(message.GetBranchTemplateResponse_GetBranchTemplateResponseDate)

	// 获取缓存中的配置
	bs, err := s.cache.Get(fmt.Sprintf(model.MessageTemplateKey, merchantID.String(), branchID.String()))
	if err == nil {
		err = json.Unmarshal(bs, template)
		if err == nil {
			return template, nil
		}
	}
	msgResp, err := s.message().GetBranchTemplate(pkgs.MetadataContent(pkgs.Metadata{MerchantID: merchantID}), &message.GetBranchTemplateRequest{
		BranchId:    branchID.String(),
		MessageType: util.AppointmentExpirationRemind,
	})
	if err != nil {
		return nil, err
	}
	if msgResp.ErrorCode != pkgs.Success {
		return nil, errors.New(msgResp.ErrorMessage)
	}

	// 反写配置入缓存
	bs, _ = json.Marshal(msgResp.Data)
	s.cache.Set(fmt.Sprintf(model.MessageTemplateKey, merchantID.String(), branchID.String()), bs)
	return msgResp.Data, nil
}

func (s *Server) cronCancelAppointment(v model.TableAppointment) error {
	tx := s.database.Conn.Begin()
	if err := tx.Select("status", "cancelled_reason", "operator").Updates(&v).Error; err != nil {
		tx.Rollback()
		cutil.Logger.Error(fmt.Sprintf("预约定时任务更新预约数据库错误:%v", err))
		return err
	}
	if err := tx.Model(&model.TableAppointmentDateCounter{}).Scopes(cutil.ColumnEqualScope("id", v.DateCounterID)).Update("appoint_num", gorm.Expr("appoint_num-1")).Error; err != nil {
		tx.Rollback()
		cutil.Logger.Error(fmt.Sprintf("预约定时任务预约数量自减数据库错误:%v", err))
		return err
	}
	return tx.Commit().Error
}

// SetCrius ...
func (s *Server) SetCrius(c *crius.Client) {
	s.crius = c
}

//PublishDataChangeEvent 缓存更新
func (s *Server) PublishDataChangeEvent(message mqMessage.DataChangeMessage) {
	data, _ := json.Marshal(message)
	err := s.mq.Publish(mqMessage.DataChangeTopic, 1, false, data)
	if err != nil {
		cutil.Logger.Error(fmt.Sprintf("publishDataChangeEvent err:%v", err))
	}
}
