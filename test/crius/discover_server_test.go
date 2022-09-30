package crius

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/Crius/proto"
)

// Test_DiscoverServer 发现服务
func Test_DiscoverServer(t *testing.T) {
	fmt.Println("begin Test_DiscoverServer")
	client := newClient()

	ctx := newContext()

	_, _ = client.client.Heart(ctx, &proto.Service{
		Name:   "test",
		Desc:   "测试",
		Ip:     "127.0.0.1",
		Port:   12345,
		Weight: 1,
	})

	_, _ = client.client.Heart(ctx, &proto.Service{
		Name:   "test",
		Desc:   "测试",
		Ip:     "127.0.0.1",
		Port:   23456,
		Weight: 2,
	})

	_, _ = client.client.Heart(ctx, &proto.Service{
		Name:   "test",
		Desc:   "测试",
		Ip:     "127.0.0.1",
		Port:   34567,
		Weight: 3,
	})

	servers := make(map[string]int)

	for i := 0; i < 1000; i++ {
		s, e := client.client.Discover(ctx, &proto.DiscoverRequest{Name: "test"})

		fmt.Println(e)
		fmt.Println(s)

		key := fmt.Sprintf("%s:%d", s.Ip, s.Port)

		if _, ok := servers[key]; ok {
			servers[key] = servers[key] + 1
			continue
		}

		servers[key] = 1
	}

	fmt.Println(servers)
}
