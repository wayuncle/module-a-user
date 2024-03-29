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
	fmt.Println("进入到rdb create")
	orm, err := pdb.GetDBInstance()
	fmt.Println("进入到rdb create orm", "err", orm, err)
	if err != nil {
		plog.Error("", "%v", err)
		return -1, err
	}
	fmt.Println("开始执行sql")
	result, err := orm.Table(Usertype.TableName()).Data(user).Insert()
	fmt.Println("进入到rdb create result", "err", result, err)
	if err != nil {
		plog.Error("", "%v", err)
		return -1, err
	}
	fmt.Println("返回")
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

// DeleteUserById
// @Description: 根据id删除用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:05:49
// @Param id int
// @Return: error
func DeleteUserById(id int) error {
	fmt.Println("rdb id", id)
	orm, err := pdb.GetDBInstance()
	fmt.Println("rdb err", err)
	if err != nil {
		plog.Error("", "%v", err)
		return err
	}
	result, err := orm.Table(Usertype.TableName()).Where("id", id).Delete()
	fmt.Println("result", "err", result, err)
	if err != nil {
		plog.Error("", "%v", err)
		return err
	}
	fmt.Println("删除结束")
	return nil
}

// GetUserById
// @Description: 根据id查询用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:15:57
// @Param id int
func GetUserById(id int) ([]*Usertype.User, error) {
	plog.Info("rdb id", "%v", id)
	orm, err := pdb.GetDBInstance()
	if err != nil {
		plog.Error("", "%v", err)
		return nil, err
	}
	result := make([]*Usertype.User, 0)
	errs := orm.Table(Usertype.TableName()).Where("id", id).Structs(&result)
	plog.Info("user", "%v", result, "  ", "err", "%v", errs)
	return result, errs
}
