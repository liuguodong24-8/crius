# redis

封装 [`redigo`](github.com/gomodule/redigo/redis) redis处理方法

接口形式封装，对外不暴露内部包任何方法调用

### 配置

```
logger.Config {
    IP       string     //地址
    Port     int        //端口
    Password string     //密码
    Database int        //数据库
}
```

### 使用方法

- 实例化
```
cfg := redis.Config{
    IP:       "127.0.0.1",
    Port:     6379,
    Password: "",
    Database: 0,
}

entity, err := redis.NewEntity(cfg)
```

- 可使用方法

```
entity.HSET(key, field, value interface{}) error
entity.HGET(key, field interface{}) ([]byte, error)
entity.HEXISTS(key, field interface{}) (bool, error)
entity.HDEL(key, field interface{}) error
entity.HGETALL(key interface{}) ([]interface{}, error)
entity.HVALS(key interface{}) ([]interface{}, error)
```

### 使用示例

```
cfg := redis.Config{
    IP:       "127.0.0.1",
    Port:     6379,
    Password: "",
    Database: 0,
}

entity, _ := redis.NewEntity(cfg)

fmt.Println(entity.HSET("test", "hyc", "123"))

fmt.Println(entity.HGETALL("test"))
```