package pkgs

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"Crius/pkgs/fields"
	uuid "github.com/satori/go.uuid"

	"google.golang.org/grpc/metadata"
)

const (
	// UserAgentWechat 微信
	UserAgentWechat = "wechat"
	// UserAgentWeb 网页
	UserAgentWeb = "web"
)

// MetadataSleuthContent rpc调用使用metadata组装唯一调用标识
func MetadataSleuthContent() context.Context {
	// http header 的生命周期是一次 http 请求， Metadata 的生命周期则是一次 RPC 调用。
	md := metadata.New(map[string]string{
		"sleuth_code": fmt.Sprintf("%d", time.Now().Unix()),
	})
	return metadata.NewOutgoingContext(context.Background(), md)
}

// MetadataContent rpc调用使用metadata组装唯一调用标识
func MetadataContent(m Metadata) context.Context {
	// http header 的生命周期是一次 http 请求， Metadata 的生命周期则是一次 RPC 调用。
	md := metadata.New(map[string]string{
		"sleuth_code": fmt.Sprintf("%d", time.Now().Unix()),
		"merchant_id": m.MerchantID.String(),
		"staff_id":    m.StaffID.String(),
		"is_admin":    strconv.Itoa(m.IsAdmin),
		"branch_ids":  m.BranchIDs.ToMetadataString(),
		"user_agent":  m.UserAgent,
	})
	return metadata.NewOutgoingContext(context.Background(), md)
}

// NewMetadataSleuthContent rpc调用使用metadata组装唯一调用标识
func NewMetadataSleuthContent(code string) context.Context {
	// http header 的生命周期是一次 http 请求， Metadata 的生命周期则是一次 RPC 调用。
	md := metadata.New(map[string]string{
		"sleuth_code": code,
	})
	return metadata.NewOutgoingContext(context.Background(), md)
}

// Metadata metadata metadata 商户员工调用链
type Metadata struct {
	SleuthCode string         `json:"sleuth_code"`
	MerchantID uuid.UUID      `json:"merchant_id"`
	StaffID    uuid.UUID      `json:"staff_id"`
	IsAdmin    int            `json:"is_admin"`
	BranchIDs  fields.UUIDArr `json:"branch_ids"`
	UserAgent  string         `json:"user_agent"`
}

// GetMetadata 从grpc metadata content 中获取链路信息
func GetMetadata(ctx context.Context) Metadata {
	var res Metadata
	if in, exists := metadata.FromIncomingContext(ctx); exists {
		code := in.Get("sleuth_code")
		if len(code) > 0 {
			res.SleuthCode = code[0]
		}

		merchant := in.Get("merchant_id")
		if len(merchant) > 0 {
			res.MerchantID = uuid.FromStringOrNil(merchant[0])
		}

		admin := in.Get("is_admin")
		if len(admin) > 0 {
			res.IsAdmin, _ = strconv.Atoi(admin[0])
		}

		staff := in.Get("staff_id")
		if len(staff) > 0 {
			res.StaffID = uuid.FromStringOrNil(staff[0])
		}

		branch := in.Get("branch_ids")
		if len(branch) > 0 {
			res.BranchIDs, _ = fields.MetadataStringToUUIDArr(branch[0])
		}

		userAgent := in.Get("user_agent")
		if len(userAgent) > 0 {
			res.UserAgent = userAgent[0]
		}
	}

	return res
}
