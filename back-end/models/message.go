/*
 * Revision History:
 *     Initial: 2018/05/07        Tong Yuehong
 */

package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/tongyuehong1/design-back-end/back-end/common"
)

func init() {
	orm.RegisterModel(new(Message))
}

// Message -
type Message struct {
	ID      int32  `orm:"column(id);pk"     json:"id"`
	Class   string `orm:"column(class)"     json:"className"`
	Title   string `orm:"column(title)"     json:"title"`
	Content string `orm:"column(content)"   json:"content"`
	Status  string `orm:"column(status)"    json:"status"`
}

// MessageServiceProvider -
type MessageServiceProvider struct {
}

// MessageServer -
var MessageServer *MessageServiceProvider

// Publish -
func (sp *MessageServiceProvider) Publish(message Message) error {
	o := orm.NewOrm()
	sql := "INSERT INTO design.message(title,class,content,status) VALUES(?,?,?,?)"
	values := []interface{}{message.Title, message.Class, message.Content, common.DefStatus}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	if err != nil {
		return err
	}
	return nil
}

// Delete -
func (sp *MessageServiceProvider) Delete(id int8) error {
	o := orm.NewOrm()
	sql := "DELETE * FROM design.message WHERE id=? AND status=? LIMIT 1"
	values := []interface{}{id, common.DefStatus}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	if err != nil {
		return err
	}
	return nil
}

// Show -
func (sp *MessageServiceProvider) Show(class string) (*[]Message, error) {
	var message []Message
	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM design.message WHERE class=? AND status=?", class, common.DefStatus).QueryRows(&message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
