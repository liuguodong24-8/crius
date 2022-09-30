## 会员管理

### 提供方法
- 列表 `GetMembers`
- 新增 `CreateMember`
- 修改 `UpdateStaff`
- 查询 `ShowMember`
- 手机号精准查询会员 `ShowMemberByAccuratePhone`
- 来电同步会员信息 `CreateOrUpdateCallingMember`

#### 列表
GetMembers请求参数:
```request
	name      string        			//名字
	phone     string        			//电话
	channel    string         			//渠道来源

	//必传参数
	offset    int32         			//分页查询offset
	limit     int32         			//单页数量limit
```
GetMembers返回参数:
```response
	error_code		int32				//错误码
    error_message	string				//错误信息
	data 			MembersData			//会员信息列表
                    count               //计数
```


#### 新增
CreateMember请求参数:
```request
	member      *MemberInfo       		//会员信息
   
    MemberInfo.avatar     string         //头像
 
	//必传参数
	MemberInfo.name         string 		//姓名
	MemberInfo.phone        string 		//手机号
	MemberInfo.phone_code   string 	    //区号
	MemberInfo.gender       int32  		//性别
	MemberInfo.first_branch_id    uuid  //交互门店
	MemberInfo.staff_id     uuid        //专属客服
    MemberInfo.birthday     date         //生日
    MemberInfo.channel      string        //渠道来源
```
CreateStaff返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
```

#### 查询
ShowMember请求参数:
```request
	//必传参数
	id string                 			//会员id
```
ShowMember返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
	data   	        MemberInfo          //会员信息
```

#### 通过电话查询
ShowMemberByAccuratePhone请求参数:
```request
	//必传参数
	phone		string					//电话
	phone_code	string					//电话区号
```
ShowStaffByPhone返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
	data   	MemberInfo              	//会员信息
```

#### 来电同步会员信息 
CreateOrUpdateCallingMember 请求参数:
```request
	//必传参数
    name        string                  //姓名
	phone		string					//电话
	phone_code	string					//电话区号
    gender      int                     //1男2女
    channel     string                  //call来电  call-appointment来电预约
    can_overwrite bool                  //是否可覆盖 来电列表true，新增预约false
```
