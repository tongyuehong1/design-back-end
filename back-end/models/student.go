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
 *     Initial: 2018/05/05        Tong Yuehong
 */

package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"

	"github.com/tongyuehong1/design-back-end/back-end/common"
)

func init() {
	orm.RegisterModel(new(Student))
}

type Student struct {
	Id      int8      `orm:"column(id)"        json:"id"`
	Class   string    `orm:"column(class)"     json:"class"`
	Name    string    `orm:"column(name)"      json:"name"`
	Avatar  string    `orm:"column(name)       json:"avatar""`
	Sex     string    `orm:"column(sex)"       json:"sex"`
	Age     int8      `orm:"column(age)"       json:"age"`
	Phone   string    `orm:"column(phone)"     json:"phone"`
	Address string    `orm:"column(address)"   json:"address"`
	Duty    string    `orm:"column(duty)"      json:"duty"`
	Isonly  string    `orm:"column(isonly)"    json:"isonly"`
	Status  int8      `orm:"column(status)"    json:"status"`
}

type StudentServiceProvider struct {
}

var StudentServer *StudentServiceProvider

func createTable() {
	name := "student"
	force := true
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		beego.Error(err)
	}
}

func (this *StudentServiceProvider) Insert(student Student) error {
	o := orm.NewOrm()
	sql := "INSERT INTO design.student(name,sex,age,phone,address,duty,isonly,status) VALUES(?,?,?,?,?,?,?)"
	values := []interface{}{student.Name,student.Sex,student.Age,student.Phone,student.Address,student.Duty,student.Isonly,common.DefStatus}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (this *StudentServiceProvider) ModifyStudent(student Student) error {
	o := orm.NewOrm()
	sql := "UPDATE design.student SET design.student(name,sex,age,phone,address,duty,isonly,status) WHERE id=? AND status=? LIMIT 1"
	values := []interface{}{student.Name,student.Sex,student.Age,student.Phone,student.Address,student.Duty,student.Isonly,common.DefStatus}
	raw := o.Raw(sql, values)
	result, err := raw.Exec()
	if err == nil {
		if row, _ := result.RowsAffected(); row == 0 {
			return common.ErrNotFound
		}
	}
	return err
}

func (this *StudentServiceProvider) GetLeaders(classes string) ([]Student, error) {
	var student []Student
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.student WHERE classes=? AND duty!=? AND status=?", classes, "", common.DefStatus).QueryRows(&student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (this *StudentServiceProvider) GetAll(classes string) ([]Student,error) {
	var student []Student
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.student WHERE classes=? AND status=?", classes, common.DefStatus).QueryRows(&student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (this *StudentServiceProvider) GetOne(name,class string) (*Student, error) {
	var student Student
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.student WHERE name=? AND class=?", name, class).QueryRows(&student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}
