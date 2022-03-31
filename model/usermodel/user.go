package usermodel

import (
	"github.com/gogf/gf/net/ghttp"
	"module-a-user/rdb/userrdb"
	Usertype "module-a-user/type/usertype"
)

// Index
// @Description: 测试接口
// @Author: Zhenwei Huo
// @Date: 2022-03-31 09:16:24
// @Param r *ghttp.Request
func Index(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}

// Create
// @Description: 创建用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:02:16
// @Param user *Usertype.User
func Create(user *Usertype.AddReq)  {
	userrdb.Create(user)
}

// Save
// @Description: 修改用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:18:42
// @Param user *Usertype.UpdateReq
func Save(user *Usertype.UpdateReq)  {
	userrdb.Save(user)
}

// Delete
// @Description: 根据id删除用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:18:55
// @Param id int
func Delete(id int)  {
	userrdb.Delete(id)
}

// Query
// @Description: 根据id查询用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:19:07
// @Param id int
func Query(id int)  {
	userrdb.Query(id)
}
