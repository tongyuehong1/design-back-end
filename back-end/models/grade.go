/*
 * Revision History:
 *     Initial: 2018/06/06        Tong Yuehong
 */

package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

func init() {
	orm.RegisterModel(new(Grade))
}

// Grade -
type Grade struct {
	ID      int8   `orm:"column(id);pk"        json:"id"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Subject string `json:"subject"`
	Grade   string `json:"grade"`
}

// GradeServiceProvider -
type GradeServiceProvider struct {
}

var GradeServer *GradeServiceProvider

// Insert -
func (sp *GradeServiceProvider) Insert(grade Grade) error {
	o := orm.NewOrm()
	sql := "INSERT INTO design.grade(name,class,subject,grade) VALUES(?,?,?,?)"
	values := []interface{}{grade.Name, grade.Class, grade.Subject, grade.Grade}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (sp *GradeServiceProvider) GetOne(class string, name string) (*[]Grade, error) {
	var grade []Grade
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.grade WHERE name=? AND class=?", name, class).QueryRows(&grade)
	fmt.Println("7777", grade)
	if err != nil {
		return nil, err
	}

	return &grade, nil
}

func (sp *GradeServiceProvider) GetAll(class string) (*[]Grade, error) {
	var grade []Grade
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.grade WHERE class=?", class).QueryRows(&grade)
	if err != nil {
		return nil, err
	}

	return &grade, nil
}
