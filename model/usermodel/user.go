package usermodel

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/wayuncle/module-a-user/rdb/userrdb"
	Usertype "github.com/wayuncle/module-a-user/type/usertype"
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
func Create(user *Usertype.AddReq) (int64, error) {
	row, err := userrdb.Create(user)
	return row, err
}

// Save
// @Description: 修改用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:18:42
// @Param user *Usertype.UpdateReq
func Save(user *Usertype.UpdateReq) (int64, error) {
	row, err := userrdb.Save(user)
	return row, err
}

// DeleteUserById
// @Description: 根据id删除用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:18:55
// @Param id int
func DeleteUserById(id int) error {
	err := userrdb.DeleteUserById(id)
	return err
}

// GetUserById
// @Description: 根据id查询用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:19:07
// @Param id int
func GetUserById(id int) (*Usertype.User, error) {
	user, err := userrdb.GetUserById(id)
	return user, err
}
