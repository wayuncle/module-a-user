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
