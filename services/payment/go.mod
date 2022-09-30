module gitlab.omytech.com.cn/micro-service/payment

go 1.15

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/allegro/bigcache/v2 v2.2.5 // indirect
	github.com/allegro/bigcache/v3 v3.0.0
	github.com/fideism/golang-wechat v0.0.0-20210727060019-4ea7f434ed8f
	github.com/gin-gonic/gin v1.6.3
	github.com/lib/pq v1.10.0
	github.com/rubenv/sql-migrate v0.0.0-20210408115534-a32ed26c37ea
	github.com/satori/go.uuid v1.2.0
	gitlab.omytech.com.cn/micro-service/Crius v0.0.0-20210702025405-7e31e52fdb76
	gitlab.omytech.com.cn/micro-service/merchant-basic v0.0.0-20210702032434-28f02212a02c
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.26.0
	gorm.io/gorm v1.20.11
)
