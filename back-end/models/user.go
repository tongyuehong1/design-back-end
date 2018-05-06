/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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

type User struct {
	Id       int8   `orm:"id"         json:"id"`
	Name     string `orm:"name"       json:"name"`
	Role     string `orm:"role"       json:"role"`
	Class    string `orm:"class"      json:"class"`
	PassWord string `orm:"password"   json:"password"`
}

type UserServiceProvider struct {
}

var UserServer *UserServiceProvider

func (this *UserServiceProvider) Register(user User) error {
	o := orm.NewOrm()
	hash, err := utility.GenerateHash(user.PassWord)

	if err != nil {
		return err
	}
	password := string(hash)
	sql := "INSERT INTO design.user(name,role,class,password) VALUES(?,?,?,?)"
	values := []interface{}{user.Name,user.Role,user.Class,password}
	raw := o.Raw(sql, values)
	_, err = raw.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (this *UserServiceProvider) Login(name string, class string, pass string) (bool, string, error) {
	o := orm.NewOrm()
	var info User

	err := o.Raw("SELECT * FROM design.user WHERE name=? AND class=? LIMIT 1 LOCK IN SHARE MODE", name,class).QueryRow(&info)
	if err != nil {
		return false, "", err
	} else if !utility.CompareHash([]byte(pass), info.PassWord) {
		return false, "", nil
	}

	return true, info.Class, nil
}

func (this *UserServiceProvider) GetClasses() ([]string, error) {
	o := orm.NewOrm()
	var classes []string
	teacher := "班主任"
	_, err := o.Raw("SELECT class FROM design.user WHERE role = ?", teacher).QueryRows(&classes)
	if err != nil{
		return nil, err
	}
	return classes, nil
}
