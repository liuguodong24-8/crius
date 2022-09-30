## 短信

调用纯K消息模块发送短信

### 方法
- 发送短信息 `SendSms`
- 短信历史记录 `GetSmsRecord`

### 数据库

```sql
create table sms_records(
    id uuid primary key,
    phone varchar(20) not null comment '电话号码',
    context text not null comment '消息内容',
    create_time timestamp,
    update_time timestamp
)
```
