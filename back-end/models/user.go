/*
 * Revision History:
 *     Initial: 2018/05/06        Tong Yuehong
 */

package models

import (
	"github.com/astaxie/beego/orm"

	"github.com/tongyuehong1/design-back-end/back-end/utility"
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
}

// UserServiceProvider -
type UserServiceProvider struct {
}

// UserServer -
var UserServer *UserServiceProvider

// Register -
func (sp *UserServiceProvider) Register(user User) error {
	o := orm.NewOrm()
	hash, err := utility.GenerateHash(user.PassWord)

	if err != nil {
		return err
	}
	password := string(hash)
	sql := "INSERT INTO design.user(name,role,class,password) VALUES(?,?,?,?)"
	values := []interface{}{user.Name, user.Role, user.Class, password}
	raw := o.Raw(sql, values)
	_, err = raw.Exec()
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

	err := o.Raw("SELECT * FROM design.user WHERE name=? AND class=? AND role=? LIMIT 1 LOCK IN SHARE MODE", name, class,role).QueryRow(&info)
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
	teacher := "班主任"
	_, err := o.Raw("SELECT class FROM design.user WHERE role = ?", teacher).QueryRows(&classes)
	if err != nil {
		return nil, err
	}
	return classes, nil
}
