package userrdb

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	Usertype "github.com/wayuncle/module-a-user/type/usertype"
)

// Create
// @Description: 新增用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 09:38:47
// @Param user *Usertype.AddReq
func Create(user *Usertype.AddReq) error {
	_, err := g.DB().Table(Usertype.TableName()).Data(user).Insert()
	fmt.Println("err", err)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

// Save
// @Description: 修改用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:02:58
// @Param user *Usertype.UpdateReq
// @Return: error
func Save(user *Usertype.UpdateReq) error {
	_, err := g.DB().Table(Usertype.TableName()).Data(user).Save()
	fmt.Println("err", err)
	if err != nil {
		fmt.Println("err", err)
		return err
	}
	return nil
}

// Delete
// @Description: 根据id删除用户
// @Author: Zhenwei Huo
// @Date: 2022-03-31 10:05:49
// @Param id int
// @Return: error
func Delete(id int) error {
	_, err := g.DB().Table(Usertype.TableName()).Delete("id", id)
	fmt.Println("err", err)
	if err != nil {
		fmt.Println("err", err)
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
	user := (*Usertype.User)(nil)
	err := g.DB().Table(Usertype.TableName()).Where("id", id).Structs(&user)
	fmt.Println("data", "err", user, err)
	return user, err
}
