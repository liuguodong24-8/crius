package config

// Setting 配置对象
var Setting setting

// setting 配置
type setting struct {
	Web struct {
		Address string
	}
	Grpc struct {
		Address string
		Arrive  int64
	}
	Log struct {
		Channel string
		Level   int
		Output  string
		Stack   bool
	}
	Redis struct {
		IP       string
		Port     int
		Password string
		Database int
	}
}
