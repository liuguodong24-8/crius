package crius

import (
	"fmt"
	"testing"

	"gitlab.omytech.com.cn/micro-service/Crius/proto"
)

// Test_DiscoverServers 批量发现
func Test_DiscoverServers(t *testing.T) {
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

	s, e := client.client.DiscoverServers(ctx, &proto.DiscoverServersRequest{
		Names: []string{"test"},
	})
	fmt.Println(s)
	fmt.Println(e)
}
