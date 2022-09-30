module gitlab.omytech.com.cn/micro-service/appointment

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/allegro/bigcache/v3 v3.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/eclipse/paho.mqtt.golang v1.3.5
	github.com/golang-module/carbon v1.3.7
	github.com/gomodule/redigo v1.8.3
	github.com/lib/pq v1.10.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/rubenv/sql-migrate v0.0.0-20210408115534-a32ed26c37ea
	github.com/satori/go.uuid v1.2.0
	gitlab.omytech.com.cn/micro-service/Crius v0.0.0-20220117075146-be0bd1e5be6a
	gitlab.omytech.com.cn/micro-service/merchant-basic v0.0.0-20211222034944-e91c7439c828
	gitlab.omytech.com.cn/micro-service/message v0.0.0-20210805071350-d17910c00d05
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.26.0
	gorm.io/gorm v1.20.11
)
