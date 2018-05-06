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

	"github.com/tongyuehong1/design-back-end/back-end/common"
)

func init() {
	orm.RegisterModel(new(Teacher))
}

type Teacher struct {
	Id    int8   `orm:"id"    json:"id"`
	Name  string `orm:"name"  json:"name"`
	Class string `orm:"class" json:"class"`
	Sex   string `orm:"sex"   json:"sex"`
	Phone string `orm:"phone" json:"phone"`
	Age   int8   `orm:"age"   json:"age"`
}

type TeacherServiceProvider struct {
}

var TeacherServer *TeacherServiceProvider

func (this *TeacherServiceProvider) AddTeacher(teacher Teacher) error {
	o := orm.NewOrm()
	sql := "INSERT INTO design.teacher(name,class,sex,phone,age) VALUES(?,?,?,?,?)"
	values := []interface{}{teacher.Name,teacher.Class,teacher.Sex,teacher.Phone,teacher.Age}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (this *TeacherServiceProvider) ChangeTeacher(teacher Teacher) error {
	o := orm.NewOrm()
	sql := "UPDATE design.student SET design.teacher(name,class,sex,phone,age) WHERE id=?LIMIT 1"
	values := []interface{}{teacher.Name,teacher.Class,teacher.Sex,teacher.Phone,teacher.Age}
	raw := o.Raw(sql, values)
	result, err := raw.Exec()
	if err == nil {
		if row, _ := result.RowsAffected(); row == 0 {
			return common.ErrNotFound
		}
	}
	return err
}
