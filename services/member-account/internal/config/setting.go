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
		Address         string
		Basic           string
		MemberPrivate   string
		MerchantBasic   string
		MemberExtension string
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
	Redis struct {
		IP       string
		Port     int
		Password string
		Database int
	}
	Migrate struct {
		Table  string
		Schema string
	}
}
