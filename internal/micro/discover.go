package micro

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gitlab.omytech.com.cn/micro-service/Crius/internal/config"
	"gitlab.omytech.com.cn/micro-service/Crius/util"
)

// DiscoverServers 批量发现
func (d *Discover) DiscoverServers(req DiscoverServersRequest) (*DiscoverServersResponse, error) {
	var result []BasicServer

	for _, name := range req.Names {
		server, err := d.Discover(DiscoverRequest{Name: name})
		if err != nil {
			return nil, fmt.Errorf("批量发现服务:%s, 错误：%s", name, err.Error())
		}
		result = append(result, BasicServer{
			Name:   server.Name,
			IP:     server.IP,
			Port:   server.Port,
			Desc:   server.Desc,
			Weight: server.Weight,
		})
	}

	return &DiscoverServersResponse{Servers: result}, nil
}

// Discover discover
func (d *Discover) Discover(req DiscoverRequest) (*DiscoverResponse, error) {
	servers, serversErr := d.getServers(req.Name)

	if serversErr != nil {
		return nil, serversErr
	}
	res, err := d.rotationServers(req.Name, servers)
	if err != nil {
		return nil, fmt.Errorf("服务名：%s, 错误信息：%s", req.Name, err.Error())
	}

	return res, nil
}

func (d *Discover) rotationServers(serverName string, servers Servers) (*DiscoverResponse, error) {
	if len(servers) == 0 {
		return nil, errors.New("获取服务发现：信息为空")
	}

	if len(servers) == 1 {
		return &DiscoverResponse{
			Name: servers[0].Name,
			IP:   servers[0].IP,
			Port: servers[0].Port,
		}, nil
	}

	server := d.nginxSmoothWeighted(serverName, servers.SortByWeight())
	if nil == server {
		return nil, errors.New("获取服务信息错误")
	}
	return &DiscoverResponse{
		Name: server.Name,
		IP:   server.IP,
		Port: server.Port,
	}, nil
}

func (d *Discover) getServers(name string) (Servers, error) {
	var servers Servers
	res, err := d.redis.HGET(RedisMicroKey, name)
	if err != nil {
		return servers, err
	}

	return judgeServerArrive(res)
}

func judgeServerArrive(res []byte) (Servers, error) {
	var servers Servers
	if err := json.Unmarshal(res, &servers); err != nil {
		return servers, err
	}

	var result Servers
	// 服务存活 1分钟判断
	t := time.Now().Unix() - config.Setting.Grpc.Arrive
	for _, s := range servers {
		// 活跃时间判断
		if s.ArriveTime < t {
			continue
		}

		result = append(result, s)
	}

	return result, nil
}

func (d *Discover) nginxSmoothWeighted(serverName string, servers Servers) (next *Server) {
	var total int32

	serverMap := make(map[string]Server)

	for i := 0; i < len(servers); i++ {

		s := servers[i]
		s.CurrentWeight += s.EffectiveWeight
		total += s.EffectiveWeight

		if s.EffectiveWeight < s.Weight {
			s.EffectiveWeight++
		}

		serverMap[s.UniqueKey()] = s
		if next == nil || s.CurrentWeight > next.CurrentWeight {
			next = &s
		}
	}

	if next == nil {
		return nil
	}

	next.CurrentWeight -= total

	serverMap[next.UniqueKey()] = *next

	// 异步保存服务权重
	go d.refreshServers(serverName, serverMap)

	return next
}

func (d *Discover) refreshServers(serverName string, servers map[string]Server) {
	var cache Servers
	for _, s := range servers {
		cache = append(cache, s)
	}

	if cache == nil {
		util.Logger.Error("重新刷新服务，map信息错误")
		return
	}

	j, err := cache.ToJSON()
	if err != nil {
		util.Logger.WithError(err).Error("重新刷新服务,json错误")
		return
	}

	if err := d.redis.HSET(RedisMicroKey, serverName, j); nil != err {
		util.Logger.WithError(err).Error("重新刷新服务，保存redis错误")
	}
}
