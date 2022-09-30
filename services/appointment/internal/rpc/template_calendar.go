package rpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-module/carbon"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gorm.io/gorm"
)

// UpdateTemplateCalendar 修改模版日历
func (s *Server) UpdateTemplateCalendar(ctx context.Context, req *proto.UpdateTemplateCalendarRequest) (*proto.Response, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateTemplateCalendar")
	if len(req.BranchId) == 0 || req.BeginDate == 0 || req.EndDate == 0 || len(req.Settings) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("修改模版日历参数错误")
		return &proto.Response{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	begin := req.BeginDate
	end := req.EndDate
	duration := int64(24 * 60 * 60)
	metadata := pkgs.GetMetadata(ctx)

	calendar := model.TableTemplateCalendar{
		MerchantID: metadata.MerchantID,
		BranchID:   uuid.FromStringOrNil(req.BranchId),
	}

	tx := s.database.Conn.Begin()

	for true {
		t := carbon.CreateFromTimestamp(begin)

		for _, setting := range req.Settings {
			// 判断星期是否符合
			if !judgeWeek(int32(t.Week()), setting.Weeks) {
				continue
			}

			// 符合 设置
			calendar.BusinessDate = fields.UnixToDateTime(begin)
			calendar.TemplateID = uuid.FromStringOrNil(setting.TemplateId)
			calendar.Category = model.StringToCalendarCategory(setting.Category)
			themeIDs, err := fields.StringArrToUUIDArr(setting.ThemeIds)
			if err != nil {
				tx.Rollback()
				util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("修改模版日历参数错误")
				return &proto.Response{
					ErrorCode:    pkgs.ErrUnprocessableEntity,
					ErrorMessage: "参数错误",
				}, nil
			}
			calendar.ThemeIDs = &themeIDs

			if err := updateOrCreateTemplateCalendar(tx, calendar); nil != err {
				tx.Rollback()

				util.Logger.WithMetadata(ctx).WithFields("calendar", logger.MakeFields(calendar)).WithError(err).Error("保存或修改模版日历错误")
				return &proto.Response{
					ErrorCode:    pkgs.ErrInternal,
					ErrorMessage: fmt.Sprintf("保存或修改模版日历错误:%s", err.Error()),
				}, nil
			}
		}

		begin += duration
		if begin > end {
			break
		}
	}

	tx.Commit()

	// 保存操作记录
	before := pkgs.MakeParams(req)
	go s.SaveSnapshot(ctx, model.TableSnapshot{
		SnapShotTableName: model.TableTemplateCalendar{}.TableName(),
		TableID:           nil,
		Method:            model.UpdateMethod,
		Before:            &before,
	})

	return &proto.Response{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
	}, nil
}

// TemplateCalendar 模版日历
type TemplateCalendar struct {
	CalendarID       string          `json:"calendar_id"`
	BusinessDate     fields.DateTime `json:"business_date"`
	CalendarCategory string          `json:"calendar_category"`
	TemplateID       string          `json:"template_id"`
	TemplateName     string          `json:"template_name"`
	TemplateColor    string          `json:"template_color"`
	ThemeIDs         *fields.UUIDArr `json:"theme_ids" gorm:"column:theme_ids"`
}

// GetTemplateCalendar 获取模版日历
func (s *Server) GetTemplateCalendar(ctx context.Context, req *proto.GetTemplateCalendarRequest) (*proto.GetTemplateCalendarResponse, error) {
	defer util.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetTemplateCalendar")

	if len(req.BranchId) == 0 {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Error("查询模版日历参数错误")
		return &proto.GetTemplateCalendarResponse{
			ErrorCode:    pkgs.ErrUnprocessableEntity,
			ErrorMessage: "参数错误",
		}, nil
	}

	// 不传 默认当前年
	year := req.Year
	if year == 0 {
		year = int32(carbon.Now().Year())
	}

	var items []TemplateCalendar

	//s.database.Conn.Model(&model.TableTemplateCalendar{}).Scopes(util.ColumnEqualScope("branch_id", req.BranchId), util.ColumnEqualScope("EXTRACT(YEAR from business_date)", year)).Order("business_date").Find(&calendars).Error;
	if err := s.database.Conn.Model(&model.TableTemplateCalendar{}).
		Select("appointment.appointment_template_calendar.id as calendar_id, "+
			"appointment.appointment_template_calendar.business_date as business_date, "+
			"appointment.appointment_template_calendar.category as calendar_category,"+
			"appointment.appointment_template_calendar.template_id as template_id,"+
			"appointment.appointment_template.name as template_name,"+
			"appointment.appointment_template.color as template_color,"+
			"appointment.appointment_template_calendar.theme_ids as theme_ids").
		Scopes(util.ColumnEqualScope("appointment.appointment_template_calendar.branch_id", req.BranchId), util.ColumnEqualScope("EXTRACT(YEAR from appointment.appointment_template_calendar.business_date)", year)).
		Order("appointment.appointment_template_calendar.business_date").
		Joins("left join appointment.appointment_template on appointment.appointment_template.id = appointment.appointment_template_calendar.template_id and appointment.appointment_template.status = ?", util.StatusOpened).Scan(&items).Error; nil != err {
		util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).WithError(err).Error("查询模版日历错误")
		return &proto.GetTemplateCalendarResponse{
			ErrorCode:    pkgs.ErrInternal,
			ErrorMessage: fmt.Sprintf("查询模版日历错误:%s", err.Error()),
		}, nil
	}

	var calendars []*proto.GetTemplateCalendarResponse_Calendar

	for _, item := range items {
		calendars = append(calendars, &proto.GetTemplateCalendarResponse_Calendar{
			CalendarId:       item.CalendarID,
			CalendarCategory: item.CalendarCategory,
			BusinessDate:     item.BusinessDate.ToUnix(),
			TemplateId:       item.TemplateID,
			TemplateName:     item.TemplateName,
			TemplateColor:    item.TemplateColor,
			ThemeIds:         item.ThemeIDs.ToStringArr(),
		})
	}

	return &proto.GetTemplateCalendarResponse{
		ErrorCode:    pkgs.Success,
		ErrorMessage: "",
		Data:         calendars,
	}, nil
}

func judgeWeek(week int32, weeks []int32) bool {
	for _, w := range weeks {
		if w == week {
			return true
		}
	}

	return false
}

func updateOrCreateTemplateCalendar(tx *gorm.DB, calendar model.TableTemplateCalendar) error {
	var data model.TableTemplateCalendar

	if err := tx.Model(&model.TableTemplateCalendar{}).Scopes(util.ColumnEqualScope("branch_id", calendar.BranchID), util.ColumnEqualScope("business_date", calendar.BusinessDate)).First(&data).Error; nil != err {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// 不存在 创建
		calendar.ID = uuid.NewV4()
		return tx.Create(&calendar).Error
	}

	// 之前是节假日 现在是普通 跳过
	if data.Category == model.CalendarHoliday && calendar.Category == model.CalendarNormal {
		return nil
	}

	// 其他情况 节-节 普-节 普-普 都是直接更新
	data.Category = calendar.Category
	data.TemplateID = calendar.TemplateID
	data.ThemeIDs = calendar.ThemeIDs

	return tx.Save(&data).Error
}
