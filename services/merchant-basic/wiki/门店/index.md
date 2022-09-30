## 门店管理

### 提供方法
- 列表 `GetBranches`
- 删除 `DeleteBranch`
- 新增 `CreateBranch`
- 修改 `UpdateBranch`
- 查询 `ShowBranch`

#### 列表
GetBranches请求参数:
```request
	name       string           		//门店名称
	province_id string           		//省ID
	city_id     string           		//城市ID
	district_id string           		//区ID

	//必传参数
	offset     int32            		//分页查询offset
	limit      int32            		//单页数量limit
```
GetBranches返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
	data			BranchData         	//门店信息列表
```

#### 删除
DeleteBranch请求参数:
```request
	//必传参数
	id         string       			//被删除门店ID
	operater_id string       			//操作人ID
```
DeleteBranch返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
```

#### 新增
CreateBranch请求参数:
```request
	branch     *BranchInfo      		//门店信息

	//必传参数
	BranchInfo.name
	BranchInfo.province_id
	BranchInfo.city_id
	BranchInfo.district_id
	BranchInfo.address
	BranchInfo.phone
	BranchInfo.longitude
	BranchInfo.latitude
	BranchInfo.location
	operater_id string           		//操作人ID
```
CreateBranch返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
```

#### 修改
UpdateBranch请求参数:
```request
	branch     *BranchInfo      		//门店信息

	//必传参数
	BranchInfo.id
	BranchInfo.name
	BranchInfo.province_id
	BranchInfo.city_id
	BranchInfo.district_id
	BranchInfo.address
	BranchInfo.phone
	BranchInfo.longitude
	BranchInfo.latitude
	BranchInfo.location
	operater_id string           		//操作人ID
```
UpdateBranch返回参数:
```response
    error_code		int32				//错误码
    error_message	string				//错误信息
```

#### 查询
ShowBranch请求参数:
```
	id			string					//门店ID
```
ShowBranch返回参数:
```
    error_code		int32				//错误码
    error_message	string				//错误信息
    data			BranchInfo			//门店信息
```

#### 其他
BranchInfo门店信息结构:
```branchInfo
	id         string					//门店ID
	name       string					//门店名称
	province_id string					//省ID
	city_id     string 					//城市ID
	district_id string 					//区ID
	address    string 					//详细地址
	phone      string 					//门店联系电话
	created_at  int32  					//开业时间
	extra      string					//其他参数
	code       string 					//门店编码
	longitude  float					//经度
	latitude   float					//纬度
	location   string					//门店地址
```

BranchData门店信息列表返回结构:
```branchData
	branches	[]BranchInfo			//门店信息
    count		int32					//门店数量
```