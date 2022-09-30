package config

// Setting 配置对象
var Setting setting

type official struct {
	Channel   string `toml:"channel"`
	AppID     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
	Token     string `toml:"token"`
	AesKey    string `toml:"aes_key"`
}

// setting 配置
type setting struct {
	App struct {
		Name   string
		Desc   string
		IP     string
		Port   int64
		Weight int32
	}
	Crius struct {
		Address string
	}
	Redis struct {
		IP       string
		Port     int
		Password string
		Database int
	}
	Log struct {
		Channel string
		Level   int
		Output  string
		Stack   bool
	}
	Wechat struct {
		Official []official
	}
}
