package micro

import (
	"encoding/json"
	"fmt"
	"sort"

	"gitlab.omytech.com.cn/micro-service/Crius/pkgs/redis"
)

// RedisMicroKey 服务缓存key
const RedisMicroKey = "micro_servers"

type (
	// Server 服务结果
	Server struct {
		IP              string `json:"ip"`
		Name            string `json:"name"`
		Port            int64  `json:"port"`
		Desc            string `json:"desc"`
		Weight          int32  `json:"weight"`
		ArriveTime      int64  `json:"arrive_time"`
		CurrentWeight   int32  `json:"current_weight"`
		EffectiveWeight int32  `json:"effective_weight"`
	}

	// Servers servers
	Servers []Server
)

// SortByWeight 按权重排序
func (s Servers) SortByWeight() Servers {
	sort.Slice(s, func(i, j int) bool {
		return s[i].Weight < s[j].Weight
	})

	return s
}

// ToJSON to json
func (s Servers) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

// UniqueKey 唯一值
func (s Server) UniqueKey() string {
	return fmt.Sprintf("%s:%s:%d", s.Name, s.IP, s.Port)
}

type (
	// BasicServer 服务
	BasicServer struct {
		Name   string `json:"name"`
		IP     string `json:"ip"`
		Port   int64  `json:"port"`
		Desc   string `json:"desc"`
		Weight int32  `json:"weight"`
	}
	// RegisterRequest 服务注册请求
	RegisterRequest BasicServer
	// DiscoverRequest 发现请求
	DiscoverRequest struct {
		Name string `json:"name"`
	}
	// DiscoverResponse 发现响应
	DiscoverResponse BasicServer

	// DiscoverServersRequest 批量发现
	DiscoverServersRequest struct {
		Names []string `json:"names"`
	}
	// DiscoverServersResponse 批量发现返回
	DiscoverServersResponse struct {
		Servers []BasicServer `json:"servers"`
	}
)

// 强制要求该接口方法必须实现
var _ DiscoverInterface = (*Discover)(nil)
var _ RegisterInterface = (*Register)(nil)

// RegisterInterface 服务注册接口
type RegisterInterface interface {
	Register(req RegisterRequest) error
}

// DiscoverInterface 服务发现接口
type DiscoverInterface interface {
	Discover(req DiscoverRequest) (*DiscoverResponse, error)
	DiscoverServers(req DiscoverServersRequest) (*DiscoverServersResponse, error)
}

// Register 注册
type Register struct {
	redis *redis.Entity
}

// Discover 发现
type Discover struct {
	redis *redis.Entity
}

// NewRegister 实例化注册
func NewRegister(r *redis.Entity) *Register {
	return &Register{
		redis: r,
	}
}

// NewDiscover 实例化发现
func NewDiscover(r *redis.Entity) *Discover {
	return &Discover{
		redis: r,
	}
}
