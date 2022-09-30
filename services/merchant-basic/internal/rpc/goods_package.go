package rpc

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs"
	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/logger"
	crius "gitlab.omytech.com.cn/micro-service/Crius/util"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model"
	centreData "gitlab.omytech.com.cn/micro-service/merchant-basic/internal/model/centre_data"
	"gitlab.omytech.com.cn/micro-service/merchant-basic/proto"
	"gorm.io/gorm"
)

// SearchGoodsAndPackage 商品和套餐
func (s *Server) SearchGoodsAndPackage(ctx context.Context, request *proto.SearchGoodsAndPackageRequest) (*proto.SearchGoodsAndPackageResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("SearchGoodsAndPackage")
	resp := &proto.SearchGoodsAndPackageResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	goods, err := centreData.SearchGoods(request.NameOrCode, merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetGoods 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}
	packages, err := centreData.SearchPackages(request.NameOrCode, merchantID)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetPackages 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}
	var protoGoods, protoPackages []*proto.GoodsAndPackageItem
	if len(goods) > 0 {
		for _, v := range goods {
			protoGoods = append(protoGoods, toProtoGoodsItem(v))
		}
	}
	if len(packages) > 0 {
		for _, v := range packages {
			protoPackages = append(protoPackages, toProtoPackageItem(v))
		}
	}

	resp.Data = &proto.SearchGoodsAndPackageResponse_Result{
		Goods:    protoGoods,
		Packages: protoPackages,
	}
	return resp, nil
}

// SearchGoodsOrPackage 搜索商品or套餐
func (s *Server) SearchGoodsOrPackage(ctx context.Context, request *proto.SearchGoodsOrPackageRequest) (*proto.SearchGoodsOrPackageResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("SearchGoodsOrPackage")
	resp := &proto.SearchGoodsOrPackageResponse{
		ErrorCode: pkgs.Success,
	}
	merchantID := pkgs.GetMetadata(ctx).MerchantID
	var protoData []*proto.GoodsAndPackageItem
	if request.Category == "goods" {
		goods, err := centreData.SearchGoods(request.NameOrCode, merchantID)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("GetGoods 数据库错误, %v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = err.Error()
			return resp, nil
		}
		if len(goods) > 0 {
			for _, v := range goods {
				protoData = append(protoData, toProtoGoodsItem(v))
			}
		}
	} else if request.Category == "package" {
		packages, err := centreData.SearchPackages(request.NameOrCode, merchantID)
		if err != nil {
			crius.Logger.Error(fmt.Sprintf("GetPackages 数据库错误, %v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = err.Error()
			return resp, nil
		}
		if len(packages) > 0 {
			for _, v := range packages {
				protoData = append(protoData, toProtoPackageItem(v))
			}
		}
	}

	resp.Data = protoData
	return resp, nil
}

// ShowGoods 商品详情
func (s *Server) ShowGoods(ctx context.Context, request *proto.ShowGoodsRequest) (*proto.ShowGoodsResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowGoods")
	resp := &proto.ShowGoodsResponse{
		ErrorCode: pkgs.Success,
	}
	id := uuid.FromStringOrNil(request.Id)
	branchID := uuid.FromStringOrNil(request.BranchId)
	if uuid.Equal(id, uuid.Nil) {
		crius.Logger.Error(fmt.Sprintf("ShowGoods 参数错误, %v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}

	goods, err := centreData.ShowGoods(id, branchID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowGoods 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}

	if goods.Price == 0 {
		lowestPrice, err := centreData.ShowBranchGoodsLowestPrice(id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			crius.Logger.Error(fmt.Sprintf("ShowBranchGoodsLowestPrice 数据库错误, %v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = err.Error()
			return resp, nil
		}
		if lowestPrice.Price != 0 {
			goods.Price = lowestPrice.Price
		}

	}

	resp.Data = toProtoGoodsItem(*goods)

	return resp, nil
}

// ShowPackage 套餐详情
func (s *Server) ShowPackage(ctx context.Context, request *proto.ShowPackageRequest) (*proto.ShowPackageResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowPackage")
	resp := &proto.ShowPackageResponse{
		ErrorCode: pkgs.Success,
	}
	id := uuid.FromStringOrNil(request.Id)
	branchID := uuid.FromStringOrNil(request.BranchId)
	if uuid.Equal(id, uuid.Nil) {
		crius.Logger.Error(fmt.Sprintf("ShowPackage 参数错误, %v", request))
		resp.ErrorCode = pkgs.ErrUnprocessableEntity
		resp.ErrorMessage = "参数错误"
		return resp, nil
	}

	pkg, err := centreData.ShowPackage(id, branchID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, nil
		}
		crius.Logger.Error(fmt.Sprintf("ShowPackage 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}

	if pkg.Price == 0 {
		lowestPrice, err := centreData.ShowBranchPackageLowestPrice(id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			crius.Logger.Error(fmt.Sprintf("ShowBranchPackageLowestPrice 数据库错误, %v", err))
			resp.ErrorCode = pkgs.ErrInternal
			resp.ErrorMessage = err.Error()
			return resp, nil
		}
		if lowestPrice.Price != 0 {
			pkg.Price = lowestPrice.Price
		}

	}

	resp.Data = toProtoPackageItem(*pkg)

	return resp, nil
}

// MultiGetGoodsAndPackages 批量获取商品、套餐
func (s *Server) MultiGetGoodsAndPackages(ctx context.Context, request *proto.MultiGetGoodsAndPackagesRequest) (*proto.MultiGetGoodsAndPackagesResponse, error) {
	defer crius.CatchException()
	crius.Logger.WithMetadata(ctx).WithFields("params", logger.MakeFields(request)).Info("ShowPackage")
	resp := &proto.MultiGetGoodsAndPackagesResponse{
		ErrorCode: pkgs.Success,
	}
	var goodsIDs, packagesIDs []uuid.UUID
	if len(request.GoodsIds) > 0 {
		for _, id := range request.GoodsIds {
			goodsIDs = append(goodsIDs, uuid.FromStringOrNil(id))
		}
	}
	if len(request.PackageIds) > 0 {
		for _, id := range request.PackageIds {
			packagesIDs = append(packagesIDs, uuid.FromStringOrNil(id))
		}
	}
	branchID := uuid.FromStringOrNil(request.BranchId)

	//处理商品
	var goods []centreData.GoodsWithPrice
	err := model.DatabaseConn().Raw(
		`SELECT
				pg.*,
				bg.price 
			FROM
				centre_data.product_b_goods pg
				LEFT JOIN centre_data.branch_b_goods bg ON pg.goods_id = bg.goods_id and bg.branch_id = ?
			WHERE
				pg.goods_id in (?)`, branchID, goodsIDs).
		Find(&goods).Error
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetGoodsWithPrice 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}
	goodsLowestPrice, err := centreData.GetBranchGoodsLowestPrice(goodsIDs)
	fmt.Println(goodsLowestPrice)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranchGoodsLowestPrice 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}
	goodsLowestPriceMap := make(map[uuid.UUID]int32)
	for _, p := range goodsLowestPrice {
		goodsLowestPriceMap[p.GoodsID] = p.Price
	}

	//处理套餐
	var packages []centreData.PackageWithPrice
	err = model.DatabaseConn().Raw(
		`SELECT
				pp.*,
				bp.price 
			FROM
				centre_data.promotion_b_package pp
			LEFT JOIN centre_data.branch_b_package bp ON bp.package_id = pp.package_id and bp.branch_id = ?
			WHERE
				pp.package_id in (?);`, branchID, packagesIDs).
		Find(&packages).Error
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetPackageWithPrice 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}

	packageLowestPrice, err := centreData.GetBranchPackagesLowestPrice(packagesIDs)
	if err != nil {
		crius.Logger.Error(fmt.Sprintf("GetBranchPackagesLowestPrice 数据库错误, %v", err))
		resp.ErrorCode = pkgs.ErrInternal
		resp.ErrorMessage = err.Error()
		return resp, nil
	}
	packageLowestPriceMap := make(map[uuid.UUID]int32)
	for _, p := range packageLowestPrice {
		packageLowestPriceMap[p.PackageID] = p.Price
	}

	var protoGoods, protoPackages []*proto.GoodsAndPackageItem
	if len(goods) > 0 {
		for _, v := range goods {
			price := v.Price
			if price == 0 {
				lowPrice, ok := goodsLowestPriceMap[v.GoodsID]
				if ok {
					price = lowPrice
				}
			}
			v.Price = price
			protoGoods = append(protoGoods, toProtoGoodsItem(v))
		}
	}
	if len(packages) > 0 {
		for _, v := range packages {
			price := v.Price
			if price == 0 {
				lowPrice, ok := packageLowestPriceMap[v.PackageID]
				if ok {
					price = lowPrice
				}
			}
			v.Price = price
			protoPackages = append(protoPackages, toProtoPackageItem(v))
		}
	}

	resp.Data = &proto.MultiGetGoodsAndPackagesResponse_Result{
		Goods:    protoGoods,
		Packages: protoPackages,
	}

	return resp, nil
}

func toProtoGoodsItem(goods centreData.GoodsWithPrice) *proto.GoodsAndPackageItem {
	return &proto.GoodsAndPackageItem{
		Id:       goods.GoodsID.String(),
		PosCode:  goods.PosCode,
		Code:     goods.Code,
		Name:     goods.CnName,
		Price:    goods.Price,
		UnitName: goods.UnitName,
	}
}

func toProtoPackageItem(pkg centreData.PackageWithPrice) *proto.GoodsAndPackageItem {
	return &proto.GoodsAndPackageItem{
		Id:      pkg.PackageID.String(),
		PosCode: pkg.PosCode,
		Code:    pkg.Code,
		Name:    pkg.PackageName,
		Price:   pkg.Price,
	}
}
