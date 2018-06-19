/*
 * Revision History:
 *     Initial: 2018/05/05        Tong Yuehong
 */

package models

import (
	"github.com/astaxie/beego/orm"

	"github.com/tongyuehong1/design-back-end/back-end/common"
	"fmt"
)

func init() {
	orm.RegisterModel(new(Student))
}

// Student -
type (
	Student struct {
		ID        uint32 `orm:"column(id)"        json:"id"`
		Class     string `orm:"column(class)"     json:"class"`
		Name      string `orm:"column(name)"      json:"name"`
		StudentID string `orm:"column(studentid)" json:"studentid"`
		Avatar    string `orm:"column(avatar)"    json:"avatar"`
		Sex       string `orm:"column(sex)"       json:"sex"`
		Age       string `orm:"column(age)"       json:"age"`
		Phone     string `orm:"column(phone)"     json:"phone"`
		Address   string `orm:"column(address)"   json:"address"`
		Duty      string `orm:"column(duty)"      json:"duty"`
		Isonly    string `orm:"column(isonly)"    json:"isonly"`
		Status    int8   `orm:"column(status)"    json:"status"`
	}
	// Info -
	Info struct {
		ID        uint32 `json:"id"`
		Name      string `json:"name"`
		Class     string `json:"className"`
		Studentid string `json:"studentid"`
		Phone     string `json:"phone"`
		Address   string `json:"address"`
	}
)

var Avatar string
// StudentServiceProvider -
type StudentServiceProvider struct {
}

// StudentServer -
var StudentServer *StudentServiceProvider

// Insert -
func (sp *StudentServiceProvider) Insert(student Student) error {
	o := orm.NewOrm()
	sql := "INSERT INTO design.student(name,sex,class,studentid,avatar,age,phone,address,duty,isonly,status) VALUES(?,?,?,?,?,?,?,?,?,?,?)"
	if student.Sex == "女" {
		Avatar = "http://192.168.43.218:21001/avatar/girl.jpg"
	} else {
		Avatar = "http://192.168.43.218:21001/avatar/common.jpg"
	}
	values := []interface{}{student.Name, student.Sex, student.Class, student.StudentID, Avatar, student.Age, student.Phone, student.Address, student.Duty, student.Isonly, common.DefStatus}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	if err != nil {
		return err
	}
	return nil
}

// ModifyStudent -
func (sp *StudentServiceProvider) ModifyStudent(student Info) error {
	o := orm.NewOrm()
	sql := "UPDATE design.student SET name=?,studentid=?,phone=?,address=? WHERE name=? AND status=? AND class=? LIMIT 1"
	values := []interface{}{student.Name,student.Studentid, student.Phone,student.Address, student.Name, common.DefStatus, student.Class}
	raw := o.Raw(sql, values)
	result, err := raw.Exec()
	if err == nil {
		if row, _ := result.RowsAffected(); row == 0 {
			return common.ErrNotFound
		}
	}
	return err
}

// GetLeaders -
func (sp *StudentServiceProvider) GetLeaders(classes string) ([]Student, error) {
	var student []Student
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.student WHERE class=? AND duty!=? AND status=?", classes, "无", common.DefStatus).QueryRows(&student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

// GetAll -
func (sp *StudentServiceProvider) GetAll(classes string) ([]Student, error) {
	var student []Student
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.student WHERE class=? AND status=?", classes, common.DefStatus).QueryRows(&student)
	if err != nil {
		return nil, err
	}

	return student, nil
}

// GetOne -
func (sp *StudentServiceProvider) GetOne(name, class string) (*[]Student, error) {
	var student []Student
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.student WHERE name=? AND class=?", name, class).QueryRows(&student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

// UpAvatar -
func (sp *StudentServiceProvider) UpAvatar(name, path,class string) error {
	o := orm.NewOrm()
	fmt.Println("aaa--->", name, "path-->", path)
	sql := "UPDATE design.student SET avatar=? WHERE name=? AND status=? AND class=? LIMIT 1"
	values := []interface{}{path, name, common.DefStatus,class}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}

// UpAvatar -
func (sp *StudentServiceProvider) Delete(name, class string) error {
	o := orm.NewOrm()
	sql := "UPDATE design.student SET status=? WHERE name=? AND class=? LIMIT 1"
	values := []interface{}{0, name, class}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}
