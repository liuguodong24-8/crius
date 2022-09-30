## 员工管理

### 提供方法
- 列表 `GetStaffs`
- 删除 `DeleteStaff`
- 新增 `CreateStaff`
- 修改 `UpdateStaff`
- 查询 `ShowStaff`
- 更新状态 `UpdateStaffStatus`
- 登录 `SignIn`
- 修改密码 `UpdatePassword`
- 重置密码 `ResetPassword`
- 通过电话查询 `ShowStaffByPhone`

#### 列表
GetStaffs请求参数:
```request
	name      string        			//名字
	phone     string        			//电话
	gender    int32         			//性别1 男,2 女
	date_start int32         			//查询时间开始
	date_end   int32         			//查询时间结束
	branches   []string      			//门店ID

	//必传参数
	offset    int32         			//分页查询offset
	limit     int32         			//单页数量limit
```
GetStaffs返回参数:
```response
	error_code		int32				//错误码
    error_message	string				//错误信息
	data 			StaffsData			//用户信息列表
```

#### 删除
DeleteStaff请求参数:
```request
	//必传参数
	id         string       			//被删除员工ID
	operater_id string       			//操作人ID
```
DeleteStaff返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
```

#### 新增
CreateStaff请求参数:
```request
	staff      *StaffInfo       		//员工信息

	//必传参数
	StaffInfo.name         string 		//员工名称
	StaffInfo.employeeCode string 		//员工工号
	StaffInfo.phone        string 		//手机号
	StaffInfo.phone_code    string 		//区号
	StaffInfo.gender       int32  		//性别
	StaffInfo.entry_at      int32  		//入职时间
	roles      []string         		//员工所属角色ID
	branches    []string         		//员工所属门店ID
	operater_id string           		//操作人ID
```
CreateStaff返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
```

#### 修改
UpdateStaff请求参数:
```request
	staff      *StaffInfo       		//员工信息
	Staff.quitAt		   int32		//员工离职时间

	//必传参数
	StaffInfo.id           string		//员工ID
	StaffInfo.name         string 		//员工名称
	StaffInfo.employee_code string 		//员工工号
	StaffInfo.phone        string 		//手机号
	StaffInfo.phone_code    string 		//区号
	StaffInfo.gender       int32  		//性别
	StaffInfo.entry_at      int32  		//入职时间
	roles      []string         		//员工所属角色ID
	branches    []string         		//员工所属门店ID
	operater_id string           		//操作人ID
```
UpdateStaff返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
```

#### 查询
ShowStaff请求参数:
```request
	//必传参数
	id string                 			//员工id
```
ShowStaff返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
	data   	StaffFullInfo              	//员工信息
```

#### 更新状态
UpdateStaffStatus请求参数:
```request
	//必传参数
	id 				string              //员工id
	status			string				//状态
	operater_id		string				//操作人id
```
UpdateStaffStatus返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
```

#### 登录
SignIn请求参数:
```request
	//必传参数
	username string                 	//账号(工号/编号)
	password string                 	//密码
```
SignIn返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
	data   	SignInData              	//员工信息
```

#### 修改密码
UpdatePassword请求参数:
```request
	//必传参数
	id          string              	//修改密码员工ID
	new_password string              	//新密码
```
UpdatePassword返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
```

#### 重置密码
ResetPassword请求参数:
```request
	//必传参数
	id          string              	//修改密码员工ID
	operater_id  string              	//操作人ID
```
ResetPassword返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
```

#### 通过电话查询
ShowStaffByPhone请求参数:
```request
	//必传参数
	phone		string					//电话
	phone_code	string					//电话区号
```
ShowStaffByPhone返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
	data   	StaffFullInfo              	//员工信息
```
#### 其他
StaffInfo员工信息结构:
```staffInfo
	ID           string					//员工ID
	Name         string 				//员工名称
	EmployeeCode string 				//员工工号
	Phone        string 				//手机号
	PhoneCode    string 				//区号
	Gender       int32  				//性别
	Status       string 				//状态
	EntryAt      int32  				//入职时间
	QuitAt       int32  				//离职时间
```

StaffFullInfo员工信息(包括门店角色)结构:
```staffFullInfo
	Staff   StaffInfo					//员工基本信息
	Branches []BranchInfo				//员工所属门店信息
	Roles   []RoleInfo					//员工拥有角色信息
```

StaffsData员工列表返回结构:
```staffsData
    staffs	[]StaffFullInfo
    count	int32
```

SignInData登录返回数据结构:
```signInData
    staff_full			StaffFullInfo	//员工信息
    default_password	bool			//是否初始密码
```