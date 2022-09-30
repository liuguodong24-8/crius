## 账户

### 提供方法
- 列表 `GetAccounts`
- 详情 `ShowAccount`
- 冻结/解冻账户 `UpdateAccountStatus`
- 流水列表 `GetBills`
- 更改账户余额 `UpdateAccountBalance`
- 增加扣款 `AddAccountDeduction`
- 会员ID查账户 `GetAccountByMemberID`
- 卡号查账户 `GetAccountByCardCode`

#### 列表
GetAccounts 请求参数:
```request
    branch_id            uuid          //门店
    status               string        //账户状态

    //必填参数
    offset               int          //分区
    limit                int          //分区
```
GetAccounts 返回参数:
```response
	error_code          int32               //返回码
    error_message       string              //错误信息
	data                array           	
    data.accounts       []account           //账户对象集合
    data.count          int                 //查询总数
```

#### 账户详情
ShowAccount 请求参数:
```request
    //必填参数
    id                  uuid          //账户ID
```
ShowAccount 返回参数:
```response
	error_code          int32               //返回码
    error_message       string              //错误信息
	data                account对象           	
```

#### 冻结/解冻账户
UpdateAccountStatus 请求参数:
```request
    //必填参数
    id                      uuid          //账户ID
    action                  string        //操作 freeze-冻结;refreeze-解冻
    reason                  string        //原因
```
UpdateAccountStatus 返回参数:
```response
	error_code          int32               //返回码
    error_message       string              //错误信息
```

#### 账单流水
GetBills 请求参数:
```request
    category         string        //账单分类 消费:consume 充值:recharge 添加扣款:deduction 修改余额:change 

    //必填参数
    account_id              uuid         //账户id
    offset                  int          //分区
    limit                   int          //分区
```
GetBills 返回参数:
```response
	error_code          int32               //返回码
    error_message       string              //错误信息
	data                array           	
    data.bills          []bill              //账单对象集合
    data.count          int                 //查询总数
```

#### 更改账户余额
UpdateAccountBalance 请求参数:
```request
    //必填参数
    account_id           uuid                  //账户
    base_value           int32                 //修改后本金
    gift_value           int32                 //修改后赠金
    packages             array                 //修改后赠品
    tickets              array                 //修改后优惠券
    reason               string                //修改原因
```
UpdateCardStatus 返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```

#### 增加扣款
AddAccountDeduction 请求参数:
```request
    //必填参数
    bill_number           string                  //关联账单号
    branch_id             uuid                    //操作门店
    cost_base_value       int32                   //扣除本金
    cost_gift_value       int32                   //扣除赠金
    cost_packages         array                   //扣除赠品
    cost_tickets          array                   //扣除优惠券
    reason                string                  //修改原因
```
AddAccountDeduction 返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```

### 会员ID查账户 
GetAccountByMemberID 请求参数
```request
    //必填参数
    member_id           string             //会员id
```
GetAccountByMemberID 返回参数:
```response
	error_code          int32                  //返回码
    error_message       string                 //错误信息
	data                MemberAccount对象数组   
```

### 卡号查账户 
GetAccountByCardCode 请求参数
```request
    //必填参数
    card_code           string             //卡号
```
GetAccountByCardCode 返回参数:
```response
	error_code          int32                  //返回码
    error_message       string                 //错误信息
	data                MemberAccount对象数组   
```


account对象
```
     id                     uuid                    
     member_id              uuid              //会员
     branch_id              uuid              //门店
     base_value             int               //本金
     gift_value             int               //赠金
     products               array             //赠品集合
     packages               array             //优惠方案集合
     tickets                array             //优惠券集合
     status                 string            //状态
     tag_ids                array             //标签
     create_at              timestamp         //账户时间
```

bill对象
```
     bill_code              string              //订单号                    
     create_at              timestamp           //创建时间
     branch                 uuid                //门店
     packages               array               //优惠方案集合（开卡有）
     cost_base_value        int                 //消费本金
     cost_gift_value        int                 //消费赠金
     cost_products          array               //消费赠品
     cost_tickets           array               //消费优惠券
     curr_base_value        int                 //剩余本金
     curr_gift_value        int                 //剩余赠金
     curr_products          array               //剩余赠品
     curr_tickets           array               //剩余优惠券
     change_category        string              //账单类型
     payments               array               //支付方式集合
     staff_id               uuid                //操作人
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

MemberAccount 对象
```
  account                   account对象                
  cards                     card对象数组
  member_id                 uuid                     //会员ID
```
  