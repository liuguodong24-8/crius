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
	Crius struct {
		Address string
		Wechat  string
	}
	Log struct {
		Channel string
		Level   int
		Output  string
		Stack   bool
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SslModel string `toml:"sslmode"`
		Log      bool
	}
	Sms struct {
		Debug bool
	}
	Migrate struct {
		Table  string
		Schema string
	}
}
