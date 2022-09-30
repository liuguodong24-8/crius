# 数据库连接封装

封装 [`gorm`](https://github.com/go-gorm/gorm) 

接口形式封装，对外不暴露内部包任何方法调用

### 配置

```
db.Config{
    Dialect:  db.DialectSqlserver,
    Server:   "192.168.3.194",
    Port:     1433,
    User:     "xx",
    Database: "xxx",
    Password: "xx",
    Debug:    true, //是否需要打印sql信息
}

db.Config{
    Dialect: db.DialectSqlite,
    Server:  "path.db",
    Debug:   true,
}

const DialectPostgre Dialect = "postgres"
const DialectMysql Dialect = "mysql"
const DialectSqlite Dialect = "sqlite"
const DialectSqlserver Dialect = "sqlserver"
```

### 使用方法

- 实例化
```
cfg := db.Config{
    Dialect:  db.DialectSqlserver,
    Server:   "192.168.3.194",
    Port:     1433,
    User:     "xx",
    Database: "xx",
    Password: "xx",
    Debug:    true,
}

entity, err := db.NewEntity(cfg)

entity.Conn.xxxxxxxx
```