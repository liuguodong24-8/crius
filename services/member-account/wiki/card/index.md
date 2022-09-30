## 卡

### 提供方法
- 开卡 `ActiveCard`
- 充值 `RechargeCard`
- 绑卡 `BindCard`
- 卡列表 `GetCards`
- 卡详情 `ShowCard`
- 卡划账 `TransferCardBalance`
- 卡挂失/找回 `UpdateCardStatus`
- 补卡 `ReplaceCard`
- 注销卡 `CancelCard`

#### 开卡
ActiveCard请求参数:
```request
    birthday            string          //生日
    operator            string          //经办人

    //必填参数
    phone               string          //电话
    phone_code          string          //区号
    name                string          //名字
    card_code           string          //卡号
    gender              int32           //性别
    operator_id         string          //操作人id
    branch_id           string          //门店id
    payments_wechat     int32           //支付方式微信
    payments_cash       int32           //支付方式现金
    payments_alipay     int32           //支付方式支付宝
    payments_card       int32           //支付方式银行卡
    password            string          //密码
    category            string          //主副卡礼品卡
```
ActiveCard返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
	data         []RolesData           	//角色信息列表
```

#### 充值
RechargeCard请求参数:
```request
    discount_id     []string            //优惠id
    recommender     string              //推荐人

    //必填参数
    account_id      string              //账户id
    branch_id       string              //门店id
    payments_wechat int32               //支付方式微信
    payments_cash   int32               //支付方式现金
    payments_alipay int32               //支付方式支付宝
    payments_card   int32               //支付方式银行卡
    operator_id     string              //操作人id
```
RechargeCard返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```

#### 绑卡
BindCard请求参数:
```request
    operator    string                  //经办人

    //必填参数
    card_code   string                  //卡号
    account_id  string                  //账户id
    branch_id   string                  //门店id
    password    string                  //密码
    operator_id string                  //操作人id
```
BindCard返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```

#### 卡列表
GetCards 请求参数:
```request
    branch_id    uuid                  //门店
    category     string                //卡类别
    status       string                //卡状态
    //必填参数
    offset      int                    //分页                         
    limit       int                    //分页                         
  
```
GetCards 返回参数:
```response
	error_code      int32                  //返回码
    error_message   string                 //错误信息
    data            array                  //结果
    data.cards      []card                 //卡对象集合
    data.count      int                    //查询总数 
```


#### 根据卡号查询卡信息
ShowCard 请求参数:
```request
    //必填参数
    card_code    string                  //卡号
```
ShowCard 返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
    data          card对象                //卡信息
```

#### 卡划账
TransferCardBalance 请求参数:
```request
    //必填参数
    source_card_id    uuid                  //转出卡
    dest_card_id      uuid                  //转入卡
    amount            int32                 //转账金额（分）
    verified          bool                  //是否已验证身份
    password          string                //未验证身份需要卡密码
    operate_branch_id uuid                  //操作门店
```
TransferCardBalance 返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```

#### 卡挂失/找回
UpdateCardStatus 请求参数:
```  request
    //必填参数
    card_id           uuid                  //卡
    action            string                //操作 挂失-lost;找回-find
```
UpdateCardStatus 返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```

#### 补卡
ReplaceCard 请求参数:
```request
    //必填参数
    curr_card_id            uuid             //旧卡
    new_card_id             uuid             //新卡
    replacement_cost        int32            //手续费
    password                string           //新卡密码
    pay_method              string           //支付方式
    open_branch_id          uuid             //补卡开卡门店
```
ReplaceCard 返回参数:
```response
	error_code    int32                //返回码
    error_message string                 //错误信息
```

#### 注销卡
CancelCard 请求参数:
```request
    //必填参数
    card_id                 uuid                  //卡
    //账户余额不为0时必填
    bank_account            string                //银行账号
    bank_name               string                //开户行
    money_receiver          string                //收款人
    reason                  string                //注销原因
```
CancelCard 返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```








card对象
```
     id                 uuid                    
     category           string              //卡分类
     code               string              //卡号
     create_branch_id   uuid                //制卡门店
     create_staff_id    uuid                //制卡人
     status             string              //状态
     open_at            timestamp           //开卡时间
     open_operator      uuid                //开卡人
     account_id         uuid                //账户ID
     open_branch_id     uuid                //开卡门店
```