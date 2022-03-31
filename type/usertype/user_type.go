package Usertype

// User
// @Description: user表对象
// @Author: Zhenwei Huo
// @Date: 2022-03-30 16:36:47
type User struct {
	Id       int    `orm:"id" c:"id" desc:"主键"`
	Name     string `orm:"name" c:"name" desc:"姓名"`
	Password string `orm:"password" c:"password" desc:"密码"`
	Age      int    `orm:"age" c:"age" desc:"年龄"`
	Phone    string `orm:"phone" c:"phone" desc:"手机号"`
	Address  string `orm:"address" c:"address" desc:"地址"`
}

// TableName
// @Description: 获取表名
// @Author: Zhenwei Huo
// @Date: 2022-03-30 16:31:33
// @Return: string
func TableName() string {
	return "t_sys_user"
}

// AddReq
// @Description: 用户新增请求体
// @Author: Zhenwei Huo
// @Date: 2022-03-30 16:44:05
type AddReq struct {
	Name     string `json:"name" form:"name" desc:"姓名" v:"required#请输入用户名称"`
	Password string `json:"password" form:"password" desc:"密码" v:"required#请输入密码"`
	Age      int    `json:"age" form:"age" desc:"年龄"`
	Phone    string `json:"phone" form:"phone" desc:"手机号" v:"required#请输入手机号"`
	Address  string `json:"address" form:"address" desc:"地址"`
}

// UpdateReq
// @Description: 用户修改请求体
// @Author: Zhenwei Huo
// @Date: 2022-03-30 16:47:31
type UpdateReq struct {
	Id       int    `json:"id" form:"id" desc:"主键" v:"required#请输入id"`
	Name     string `json:"name" form:"name" desc:"姓名" v:"required#请输入用户名称"`
	Password string `json:"password" form:"password" desc:"密码" v:"required#请输入密码"`
	Age      int    `json:"age" form:"age" desc:"年龄"`
	Phone    string `json:"phone" form:"phone" desc:"手机号" v:"required#请输入手机号"`
	Address  string `json:"address" form:"address" desc:"地址"`
}
