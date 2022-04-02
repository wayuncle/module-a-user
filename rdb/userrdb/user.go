package userrdb

import (
	"fmt"
	Usertype "github.com/wayuncle/module-a-user/type/usertype"
	"msp-git.connext.com.cn/connext-go-third/third-db/pdb"
	"msp-git.connext.com.cn/connext-go-third/third-log/plog"
)

// Create
// @Description: 新增用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 09:38:47
// @Param user *Usertype.AddReq
func Create(user *Usertype.AddReq) (int64, error) {
	orm, err := pdb.GetDBInstance()
	if err != nil {
		plog.Error("", "%v", err)
		return -1, err
	}
	result, err := orm.Table(Usertype.TableName()).Data(user).Insert()
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// Save
// @Description: 修改用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:02:58
// @Param user *Usertype.UpdateReq
// @Return: error
func Save(user *Usertype.UpdateReq) (int64, error) {
	orm, err := pdb.GetDBInstance()
	if err != nil {
		plog.Error("", "%v", err)
		return -1, err
	}
	result, err := orm.Table(Usertype.TableName()).Data(user).Save()
	if err != nil {
		plog.Error("", "%v", err)
		return -1, err
	}
	return result.LastInsertId()
}

// Delete
// @Description: 根据id删除用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:05:49
// @Param id int
// @Return: error
func Delete(id int) error {
	orm, err := pdb.GetDBInstance()
	if err != nil {
		plog.Error("", "%v", err)
		return err
	}
	result, err := orm.Table(Usertype.TableName()).Delete("id", id)
	fmt.Println("err", result)
	if err != nil {
		plog.Error("", "%v", err)
		return err
	}
	return nil
}

// Query
// @Description: 根据id查询用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:15:57
// @Param id int
func Query(id int) (*Usertype.User, error) {
	orm, err := pdb.GetDBInstance()
	if err != nil {
		plog.Error("", "%v", err)
		return nil, err
	}
	user := (*Usertype.User)(nil)
	errs := orm.Table(Usertype.TableName()).Where("id", id).Structs(&user)
	fmt.Println("data", "err", user, errs)
	return user, errs
}
