/*
 * Revision History:
 *     Initial: 2018/05/06        Tong Yuehong
 */

package models

import (
	"github.com/astaxie/beego/orm"

	"github.com/tongyuehong1/design-back-end/back-end/utility"
	"github.com/tongyuehong1/design-back-end/back-end/common"
	"fmt"
)

func init() {
	orm.RegisterModel(new(User))
}

// User -
type User struct {
	ID       int8   `orm:"column(id);pk"         json:"id"`
	Name     string `orm:"column(name)"       json:"userName"`
	Role     string `orm:"column(role)"       json:"role"`
	Class    string `orm:"column(class)"      json:"classroom"`
	PassWord string `orm:"column(password)"   json:"password"`
	Status   int8   `orm:"column(status)"     json:"status"`
}

// UserServiceProvider -
type UserServiceProvider struct {
}

// UserServer -
var UserServer *UserServiceProvider

// Register -
func (sp *UserServiceProvider) Register(user User) error {
	var stu Student
	o := orm.NewOrm()
	hash, err := utility.GenerateHash(user.PassWord)

	if err != nil {
		return err
	}

	password := string(hash)

	err = o.Raw("SELECT *FROM design.student WHERE name=? AND status=? AND class=?", user.Name, common.DefStatus,user.Class).QueryRow(&stu)
	if err != nil {
		return err
	}

	sql := "INSERT INTO design.user(name,role,class,password) VALUES(?,?,?,?)"
	values := []interface{}{user.Name, user.Role, user.Class, password}
	raws := o.Raw(sql, values)
	_, err = raws.Exec()
	if err != nil {
		return err
	}
	return nil
}

// Login -
func (sp *UserServiceProvider) Login(name string, class string, pass string,role string) (bool, string, error) {
	o := orm.NewOrm()
	var (
		info User
	)

	err := o.Raw("SELECT * FROM design.user WHERE name=? AND class=? AND role=? AND status=? LIMIT 1 LOCK IN SHARE MODE", name, class,role,1).QueryRow(&info)
	if err != nil {
		return false, "", err
	} else if !utility.CompareHash([]byte(info.PassWord), pass) {
		return false, "", nil
	}

	return true, info.Class, nil
}

// GetClasses -
func (sp *UserServiceProvider) GetClasses() ([]string, error) {
	o := orm.NewOrm()
	var (
		classes []string
	)
	teacher := "teacher"
	_, err := o.Raw("SELECT class FROM design.user WHERE role = ?", teacher).QueryRows(&classes)
	if err != nil {
		return nil, err
	}
	return classes, nil
}

// UpAvatar -
func (sp *UserServiceProvider) Drop(name, class string) error {
	fmt.Println("222222", name, class)
	o := orm.NewOrm()
	sql := "UPDATE design.user SET status=? WHERE name=? AND class=? LIMIT 1"
	values := []interface{}{0, name, class}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}

