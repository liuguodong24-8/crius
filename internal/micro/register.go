package micro

import (
	"encoding/json"
	"time"
)

// Register 服务注册
func (r *Register) Register(req RegisterRequest) error {
	server := Server{
		IP:              req.IP,
		Port:            req.Port,
		Desc:            req.Desc,
		Weight:          req.Weight,
		Name:            req.Name,
		ArriveTime:      time.Now().Unix(),
		CurrentWeight:   0,
		EffectiveWeight: req.Weight,
	}

	exists, err := r.redis.HEXISTS(RedisMicroKey, server.Name)
	if err != nil {
		return err
	}

	var servers Servers
	if exists {
		caches, err := r.getServers(server.Name)
		if err != nil {
			return err
		}

		for _, s := range caches {
			// 已存在
			if s.IP == server.IP && s.Port == server.Port {
				// 服务权重未改变，不更改之前的权重值
				if s.Weight == server.Weight {
					server.CurrentWeight = s.CurrentWeight
					server.EffectiveWeight = s.EffectiveWeight
				}

				continue
			}

			servers = append(servers, s)
		}
	}

	servers = append(servers, server)

	return r.saveServers(server.Name, servers)
}

func (r *Register) getServers(key string) (Servers, error) {
	var servers Servers
	res, err := r.redis.HGET(RedisMicroKey, key)
	if err != nil {
		return servers, err
	}

	err = json.Unmarshal(res, &servers)

	return servers, err
}

func (r *Register) saveServers(key string, servers Servers) error {
	j, err := servers.ToJSON()
	if err != nil {
		return err
	}

	return r.redis.HSET(RedisMicroKey, key, j)
}
