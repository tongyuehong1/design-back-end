/*
 * Revision History:
 *     Initial: 2018/05/06        Tong Yuehong
 */

package models

import (
	"github.com/astaxie/beego/orm"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"fmt"
)

func init() {
	orm.RegisterModel(new(Teacher))
}

// Teacher -
type Teacher struct {
	ID     uint8   `orm:"column(id);pk"    json:"id"`
	Name   string `orm:"column(name)"  json:"name"`
	Class  string `orm:"column(class)" json:"className"`
	Avatar string `orm:"column(avatar)" json:"avatar"`
	Sex    string `orm:"column(sex)"   json:"gender"`
	Phone  string `orm:"column(phone)" json:"phone"`
	Office string   `orm:"column(office)"   json:"office"`
}

// TeacherServiceProvider -
type TeacherServiceProvider struct {
}

// TeacherServer -
var TeacherServer *TeacherServiceProvider
const (
	TAvatar = "http://10.0.0.43:21001/avatar/teacher.jpg"
)

// AddTeacher -
func (sp *TeacherServiceProvider) AddTeacher(name,class string) error {
	o := orm.NewOrm()
	sql := "INSERT INTO design.teacher(name,class,avatar) VALUES(?,?,?)"
	values := []interface{}{name, class, TAvatar}
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
	fmt.Println("111111", teacher)
	sql := "UPDATE design.teacher SET name=?,sex=?,phone=?,office=? WHERE class= ? LIMIT 1"
	values := []interface{}{teacher.Name, teacher.Sex, teacher.Phone, teacher.Office,teacher.Class}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()

	return err
}

// GetOne -
func (sp *TeacherServiceProvider) GetOne(class string) (*[]Teacher, error) {
	var teacher []Teacher
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.teacher WHERE class= ?", class).QueryRows(&teacher)
	if err != nil {
		return nil, err
	}

	return &teacher, nil
}

// UpAvatar -
func (sp *TeacherServiceProvider) TechAvatar(name, path string) error {
	o := orm.NewOrm()
	sql := "UPDATE design.teacher SET avatar=? WHERE name=? AND status=? LIMIT 1"
	values := []interface{}{path, name, common.DefStatus}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}
