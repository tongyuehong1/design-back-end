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

// Teacher -
type Teacher struct {
	ID    int8   `orm:"column(id);pk"    json:"id"`
	Name  string `orm:"column(name)"  json:"name"`
	Class string `orm:"column(class)" json:"class"`
	Sex   string `orm:"column(sex)"   json:"sex"`
	Phone string `orm:"column(phone)" json:"phone"`
	Age   int8   `orm:"column(age)"   json:"age"`
}

// TeacherServiceProvider -
type TeacherServiceProvider struct {
}

// TeacherServer -
var TeacherServer *TeacherServiceProvider

// AddTeacher -
func (sp *TeacherServiceProvider) AddTeacher(teacher Teacher) error {
	o := orm.NewOrm()
	sql := "INSERT INTO design.teacher(name,class,sex,phone,age) VALUES(?,?,?,?,?)"
	values := []interface{}{teacher.Name, teacher.Class, teacher.Sex, teacher.Phone, teacher.Age}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	if err != nil {
		return err
	}
	return nil
}

// ChangeTeacher -
func (sp *TeacherServiceProvider) ChangeTeacher(teacher Teacher) error {
	o := orm.NewOrm()
	sql := "UPDATE design.student SET design.teacher(name,class,sex,phone,age) WHERE id=?LIMIT 1"
	values := []interface{}{teacher.Name, teacher.Class, teacher.Sex, teacher.Phone, teacher.Age}
	raw := o.Raw(sql, values)
	result, err := raw.Exec()
	if err == nil {
		if row, _ := result.RowsAffected(); row == 0 {
			return common.ErrNotFound
		}
	}
	return err
}

// GetOne -
func (sp *TeacherServiceProvider) GetOne(class string) (*Teacher, error) {
	var teacher Teacher
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.teacher WHERE class= ?", class).QueryRows(&teacher)
	if err != nil {
		return nil, err
	}

	return &teacher, nil
}
