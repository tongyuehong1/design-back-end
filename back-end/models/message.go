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

type Message struct {
	Id      int8   `orm:"column(id)"        json:"id"`
	Title   string `orm:"column(title)"     json:"title"`
	Content string `orm:"column(content)"   json:"content"`
	Status  string `orm:"column(status)"    json:"status"`
}

type MessageServiceProvider struct {
}

var MessageServer *MessageServiceProvider

func (this *MessageServiceProvider) Publish(message Message) error {
	o := orm.NewOrm()
	sql := "INSERT INTO design.message(title,content,status) VALUES(?,?,?,?,?,?,?)"
	values := []interface{}{message.Title, message.Content, common.DefStatus}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (this *MessageServiceProvider) Delete(id int8) error {
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
