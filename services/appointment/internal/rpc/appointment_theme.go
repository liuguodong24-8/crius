package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	mqMessage "gitlab.omytech.com.cn/micro-service/Crius/pkgs/message"
	"time"

	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/fields"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/appointment/internal/model"
	"gitlab.omytech.com.cn/micro-service/appointment/proto"
	"gorm.io/gorm"
)

// CreateAppointmentTheme 创建主题
func (s *Server) CreateAppointmentTheme(ctx context.Context, req *proto.CreateAppointmentThemeRequest) (*proto.Response, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("CreateAppointmentTheme")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	_, err := model.ShowAppointmentThemeByName(req.Theme.AppointmentTheme.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		util.Logger.Error(fmt.Sprintf("CreateAppointmentTheme 查询主题数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建主题失败"
		return resp, nil
	}
	if err == nil {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "主题名称已存在"
		return resp, nil
	}

	theme, packages := themeParam(req.Theme, uuid.NewV4())

	if err := model.CreateAppointmentTheme(theme, packages); err != nil {
		util.Logger.Error(fmt.Sprintf("CreateAppointmentTheme 创建主题数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "创建主题失败"
		return resp, nil
	}
	message := mqMessage.DataChangeMessage{
		Category: mqMessage.Appointment_Theme,
	}
	go s.PublishDataChangeEvent(message)
	return resp, nil
}

// UpdateAppointmentTheme 更新主题
func (s *Server) UpdateAppointmentTheme(ctx context.Context, req *proto.UpdateAppointmentThemeRequest) (*proto.Response, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointmentTheme")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Theme.AppointmentTheme.Id)

	theme, err := model.ShowAppointmentThemeByName(req.Theme.AppointmentTheme.Name)
	if err != nil && err != gorm.ErrRecordNotFound {
		util.Logger.Error(fmt.Sprintf("UpdateAppointmentTheme 查询主题数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新主题失败"
		return resp, nil
	}
	if err == nil && theme.ID != id {
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "主题名称已存在"
		return resp, nil
	}

	theme, packages := themeParam(req.Theme, id)
	if err := model.UpdateAppointmentTheme(theme, packages); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateAppointmentTheme 更新主题数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新主题失败"
		return resp, nil
	}
	message := mqMessage.DataChangeMessage{
		Category: mqMessage.Appointment_Theme,
	}
	go s.PublishDataChangeEvent(message)
	return resp, nil
}

// UpdateAppointmentThemeStatus 更新主题状态
func (s *Server) UpdateAppointmentThemeStatus(ctx context.Context, req *proto.UpdateAppointmentThemeStatusRequest) (*proto.Response, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("UpdateAppointmentThemeStatus")
	resp := &proto.Response{
		ErrorCode: pkgs.Success,
	}

	id := uuid.FromStringOrNil(req.Id)
	if err := model.UpdateAppointmentThemeStatus(id, util.Status(req.Status)); err != nil {
		util.Logger.Error(fmt.Sprintf("UpdateAppointmentThemeStatus 更新主题状态数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "更新主题状态失败"
		return resp, nil
	}
	message := mqMessage.DataChangeMessage{
		Category: mqMessage.Appointment_Theme,
	}
	go s.PublishDataChangeEvent(message)
	return resp, nil
}

// GetAppointmentThemes 获取主题列表
func (s *Server) GetAppointmentThemes(ctx context.Context, req *proto.GetAppointmentThemesRequest) (*proto.GetAppointmentThemesResponse, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentThemes")
	resp := &proto.GetAppointmentThemesResponse{
		ErrorCode: pkgs.Success,
	}

	themes, total, err := model.GetAppointmentThemes(req.Name, util.Status(req.Status), uuid.FromStringOrNil(req.CategoryId), req.Offset, req.Limit)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetAppointmentThemes 获取主题列表数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取主题列表失败"
		return resp, nil
	}

	resp.Data = &proto.GetAppointmentThemesResponse_Data{Total: int32(total)}
	for i := range themes {
		resp.Data.Themes = append(resp.Data.Themes, toProtoAppointmentTheme(&themes[i]))
	}

	return resp, nil
}

// ShowAppointmentTheme 获取主题
func (s *Server) ShowAppointmentTheme(ctx context.Context, req *proto.ShowAppointmentThemeRequest) (*proto.ShowAppointmentThemeResponse, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("ShowAppointmentTheme")
	resp := &proto.ShowAppointmentThemeResponse{
		ErrorCode: pkgs.Success,
	}

	theme, err := model.ShowAppointmentTheme(uuid.FromStringOrNil(req.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.ErrorCode = pkgs.ErrNotFound
			resp.ErrorMessage = "主题不存在"
			return resp, nil
		}
		util.Logger.Error(fmt.Sprintf("ShowAppointmentTheme 获取主题数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取主题失败"
		return resp, nil
	}

	packages, err := model.GetAppointmentThemePackagesByThemeID(theme.ID)
	if err != nil {
		util.Logger.Error(fmt.Sprintf("ShowAppointmentTheme 获取主题套餐数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取主题失败"
		return resp, nil
	}

	resp.Data = &proto.Theme{
		AppointmentTheme: toProtoAppointmentTheme(theme),
	}
	for i := range packages {
		resp.Data.ThemePackages = append(resp.Data.ThemePackages, toProtoAppointmentThemePackage(&packages[i]))
	}
	return resp, nil
}

// GetAppointmentThemesByRoomType 根据房型获取主题套餐列表
func (s *Server) GetAppointmentThemesByRoomType(ctx context.Context, req *proto.GetAppointmentThemesByRoomTypeRequest) (*proto.GetAppointmentThemesByRoomTypeResponse, error) {
	defer crius.CatchException()
	util.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(req)).Info("GetAppointmentThemesByRoomType")
	resp := &proto.GetAppointmentThemesByRoomTypeResponse{
		ErrorCode: pkgs.Success,
	}

	branchID := uuid.FromStringOrNil(req.BranchId)
	date, _ := time.Parse("20060102", time.Unix(int64(req.Date), 0).Format("20060102"))
	calendar, err := model.ShowTemplateCalendarByBranchIDDate(branchID, date)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return resp, nil
		}
		util.Logger.Error(fmt.Sprintf("GetAppointmentThemesByRoomType 获取门店日历模板数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取主题套餐列表失败"
		return resp, nil
	}
	if calendar.ThemeIDs == nil || len(*calendar.ThemeIDs) == 0 {
		return resp, nil
	}

	themes, err := model.GetAppointmentThemesByRoomType(uuid.FromStringOrNil(req.RoomTypeId), uuid.FromStringOrNil(req.ThemeId))
	if err != nil {
		util.Logger.Error(fmt.Sprintf("GetAppointmentThemesByRoomType 获取主题数据库错误:%v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = "获取主题套餐列表失败"
		return resp, nil
	}
	for i := range themes {
		for _, v := range calendar.ThemeIDs.Slice() {
			if themes[i].ID == v {
				p := toProtoAppointmentThemePackage(&model.TableAppointmentThemePackage{
					Name:          themes[i].PackageName,
					Packages:      themes[i].Packages,
					Decoration:    themes[i].Decoration,
					Staffing:      themes[i].Staffing,
					CustomConfigs: themes[i].CustomConfigs,
					RoomTypes:     themes[i].RoomTypes,
					ID:            themes[i].PackageID,
				})
				resp.Data = append(resp.Data, &proto.ThemeRoomType{AppointmentTheme: toProtoAppointmentTheme(&themes[i].TableAppointmentTheme), ThemePackage: p, CategoryName: themes[i].CategoryName})
			}
		}
	}

	return resp, nil
}

func toProtoAppointmentTheme(t *model.TableAppointmentTheme) *proto.AppointmentTheme {
	var contents []*proto.AppointmentTheme_Content

	if t.Contents != nil {
		err := json.Unmarshal([]byte(t.Contents.JSON()), &contents)
		if err != nil {
			util.Logger.Error(fmt.Sprintf("toProtoAppointmentTheme unmarshal contents 错误:%v", err))
		}
	}
	return &proto.AppointmentTheme{
		Id:         t.ID.String(),
		Color:      t.Color,
		FeatureIds: t.FeatureIDs.ToStringArr(),
		Contents:   contents,
		Style:      t.Style,
		Images:     t.Images.Slice(),
		Video:      t.Video,
		Details:    t.Details.Slice(),
		Weight:     t.Weight,
		Status:     t.Status.String(),
		Name:       t.Name,
		CategoryId: util.UUIDToString(t.CategoryID),
	}
}

func toProtoAppointmentThemePackage(p *model.TableAppointmentThemePackage) *proto.AppointmentThemePackage {
	var (
		packages      []*proto.AppointmentThemePackage_Package
		customConfigs []*proto.AppointmentThemePackage_CustomConfig
		roomTypes     []*proto.AppointmentThemePackage_RoomType
	)

	if p.Packages != nil {
		err := json.Unmarshal([]byte(p.Packages.JSON()), &packages)
		if err != nil {
			util.Logger.Error(fmt.Sprintf("toProtoAppointmentThemePackage unmarshal packages 错误:%v", err))
		}
	}
	if p.CustomConfigs != nil {
		err := json.Unmarshal([]byte(p.CustomConfigs.JSON()), &customConfigs)
		if err != nil {
			util.Logger.Error(fmt.Sprintf("toProtoAppointmentThemePackage unmarshal custom_configs 错误:%v", err))
		}
	}
	if p.RoomTypes != nil {
		err := json.Unmarshal([]byte(p.RoomTypes.JSON()), &roomTypes)
		if err != nil {
			util.Logger.Error(fmt.Sprintf("toProtoAppointmentThemePackage unmarshal room_types 错误:%v", err))
		}
	}
	return &proto.AppointmentThemePackage{
		Name:          p.Name,
		Packages:      packages,
		Decoration:    p.Decoration,
		Staffing:      p.Staffing,
		CustomConfigs: customConfigs,
		RoomTypes:     roomTypes,
		Id:            p.ID.String(),
	}
}

func themeParam(t *proto.Theme, id uuid.UUID) (*model.TableAppointmentTheme, []model.TableAppointmentThemePackage) {
	var (
		featureIDs *fields.UUIDArr
		contents   *pkgs.ParamsArr
		images     *fields.StringArr
	)

	if len(t.AppointmentTheme.FeatureIds) != 0 {
		featureIDs = new(fields.UUIDArr)
		var err error
		*featureIDs, err = fields.StringArrToUUIDArr(t.AppointmentTheme.FeatureIds)
		if err != nil {
			util.Logger.Error(fmt.Sprintf("themeParam featureID 参数错误:%v", t.AppointmentTheme.FeatureIds))
		}
	}
	if t.AppointmentTheme.Contents != nil {
		contents = new(pkgs.ParamsArr)
		*contents = pkgs.MakeParamsArr(t.AppointmentTheme.Contents)
	}
	if len(t.AppointmentTheme.Images) != 0 {
		images = new(fields.StringArr)
		*images = fields.StringArr(t.AppointmentTheme.Images)
	}
	categoryID := uuid.FromStringOrNil(t.AppointmentTheme.CategoryId)
	details := fields.StringArr(t.AppointmentTheme.Details)
	theme := &model.TableAppointmentTheme{
		ID:         id,
		Name:       t.AppointmentTheme.Name,
		Color:      t.AppointmentTheme.Color,
		Weight:     t.AppointmentTheme.Weight,
		Status:     util.Status(t.AppointmentTheme.Status),
		FeatureIDs: featureIDs,
		Contents:   contents,
		Style:      t.AppointmentTheme.Style,
		Images:     images,
		Video:      t.AppointmentTheme.Video,
		Details:    &details,
		CategoryID: &categoryID,
	}
	packages := make([]model.TableAppointmentThemePackage, 0)
	for _, v := range t.ThemePackages {
		var p, customConfigs, roomTypes *pkgs.ParamsArr
		if len(v.Packages) != 0 {
			p = new(pkgs.ParamsArr)
			*p = pkgs.MakeParamsArr(v.Packages)
		}
		if len(v.CustomConfigs) != 0 {
			customConfigs = new(pkgs.ParamsArr)
			*customConfigs = pkgs.MakeParamsArr(v.CustomConfigs)
		}
		if len(v.RoomTypes) != 0 {
			roomTypes = new(pkgs.ParamsArr)
			*roomTypes = pkgs.MakeParamsArr(v.RoomTypes)
		}

		packages = append(packages, model.TableAppointmentThemePackage{
			Name:          v.Name,
			Packages:      p,
			Decoration:    v.Decoration,
			Staffing:      v.Staffing,
			CustomConfigs: customConfigs,
			RoomTypes:     roomTypes,
			ThemeID:       &id,
			ID:            uuid.NewV4(),
		})
	}
	return theme, packages
}
