package config

// Setting 配置对象
var Setting setting

// setting 配置
type setting struct {
	App struct {
		Name            string
		Desc            string
		IP              string
		Port            int64
		Weight          int32
		AppointmentSalt string `toml:"appointment_salt"`
	}
	Crius struct {
		Address       string
		MerchantBasic string
		Message       string
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
	Migrate struct {
		Table  string
		Schema string
	}
	Redis struct {
		IP       string
		Port     int
		Password string
		Database int
	}
	Mqtt struct {
		Client    string
		Username  string
		Password  string
		Broker    string
		TaskTopic string `toml:"task_topic"`
	}
}
