package usermodel

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/wayuncle/module-a-user/rdb/userrdb"
	Usertype "github.com/wayuncle/module-a-user/type/usertype"
	"msp-git.connext.com.cn/connext-go-third/third-log/plog"
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
	fmt.Println("user", user.Name)
	row, err := userrdb.Create(user)
	return row, err
}

// Save
// @Description: 修改用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:18:42
// @Param user *Usertype.UpdateReq
func Save(user *Usertype.UpdateReq) (int64, error) {
	fmt.Println("update user", user.Name)
	row, err := userrdb.Save(user)
	return row, err
}

// DeleteUserById
// @Description: 根据id删除用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:18:55
// @Param id int
func DeleteUserById(req *Usertype.IdReq) error {
	plog.Info("DeleteUserById user_id", "%v", req.Id)
	err := userrdb.DeleteUserById(req.Id)
	return err
}

// GetUserById
// @Description: 根据id查询用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:19:07
// @Param id int
func GetUserById(req *Usertype.IdReq) ([]*Usertype.User, error) {
	plog.Info("GetUserById user_id", "%v", req.Id)
	user, err := userrdb.GetUserById(req.Id)
	return user, err
}
