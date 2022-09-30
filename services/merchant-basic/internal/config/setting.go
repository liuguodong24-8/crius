package config

// Setting 配置对象
var Setting setting

// setting 配置
type setting struct {
	App struct {
		Name   string
		Desc   string
		IP     string
		Port   int64
		Weight int32
	}
	Log struct {
		Channel string
		Level   int
		Output  string
		Stack   bool
	}
	Crius struct {
		Address string
	}
	DataBase struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SslModel string `toml:"sslmode"`
		Log      bool
	}
	Migrate struct {
		Table  string
		Schema string
	}
	MQTT struct {
		Client    string
		Username  string
		Password  string
		Broker    string
		TaskTopic string `toml:"task_topic"`
	} `toml:"mqtt"`
}
