## 门店管理

### 提供方法
- 列表 `GetRoles`
- 删除 `DeleteRole`
- 新增 `CreateRole`
- 修改 `UpdateRole`
- 修改状态 `UpdateRoleStatus`
- 角色权限操作记录 `GetRoleHistories`
- 权限列表 `GetPermissions`
- 详情 `ShowRole`

#### 列表
GetRoles请求参数:
```request
	name       string           		//角色名称

	//必传参数
	offset     int32            		//数据条数offset
	limit      int32            		//单页数量limit
```
GetRoles返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
	data         []RolesData           	//角色信息列表
```

#### 删除
DeleteRole请求参数:
```request
	//必传参数
	id         string       			//被删除角色ID
	operater_id string       			//操作人ID
```
DeleteRole返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```

#### 新增
CreateRole请求参数:
```request
	//必传参数
    name        string                  //角色名称
    property    int32                   //角色公私有，公有1 私有2
    permissions []int32                 //权限列表
    operaterId  string                  //操作人ID
```
CreateRole返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```

#### 修改
UpdateRole请求参数:
```request
	//必传参数
    id          string                  //角色ID
    name        string                  //角色名称
    property    int32                   //角色公私有，公有1 私有2
    permissions []int32                 //权限列表
    operater_id string                  //操作人ID
```
UpdateRole返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```

#### 修改状态
UpdateRoleStatus请求参数:
```request
	//必传参数
    id          string                  //角色ID
    status      string                  //权限列表
    operater_id string                  //操作人ID
```
UpdateRoleStatus返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
```

#### 角色权限操作记录
GetRoleHistories请求参数:
```request
	//必传参数
    id          string                  //角色ID
```
GetRoleHistories返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
	data         RoleHistoriesData		//角色操作记录
```

#### 权限列表
GetPermissions请求参数:无
GetPermissions返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
	data         []PermissionInfo		//角色操作记录
```

#### 详情
ShowRole请求参数:
```request
	//必传参数
    id          string                  //角色ID
```
ShowRole返回参数:
```response
	error_code    int32                  //返回码
    error_message string                 //错误信息
	data         RolePermissionData		//角色操作记录
```

#### 其他
RoleInfo角色信息结构:
```roleInfo
	id         string					//角色ID
	name       string					//角色名称
	status     string                   //角色状态opened  closed
	property   int32                    //角色公私有，公有1 私有2
	staff_id   string                   //角色创建人
	staff_name string                   //角色创建人名字
```

RoleHistoriesData角色操作记录:
```RoleHistoriesData
	role     RoleInfo					//角色信息
	snapshot Snapshot					//操作记录
```

Snapshot操作记录快照:
```Snapshot
	staff_id   string					//操作人id
	staff_name string					//操作人名字
	before    string					//操作前数据
	after     string					//操作后数据
	created_at int32					//操作时间
	method    string					//操作方式(create,update,delete)
```

PermissionInfo权限信息:
```PermissionInfo
	id			int32					//权限id
	permission	string					//权限详细
	service		string					//权限所属系统
```

RolePermissionData角色权限信息:
```RolePermissionData
	role		RoleInfo				//角色信息
	permissions []int32					//权限id
```

RolesData角色信息列表:
```RolesData
	roles		[]RoleInfo				//角色信息
	count		int32					//数量
```