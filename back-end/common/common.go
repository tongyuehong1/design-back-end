/*
 * Revision History:
 *     Initial: 2018/05/05        Tong Yuehong
 */

package common

import "errors"

const (
	// DBdesign -
	DBdesign = "design"

	// DefStatus -
	DefStatus = 1

	// ErrSucceed -
	ErrSucceed = 0 // 成功
	// ErrMysqlQuery -
	ErrMysqlQuery = 500 // MySQL 错误
	// ErrInvalidParam -
	ErrInvalidParam = 1 // 参数错误
	// ErrMysqlNotFound -
	ErrMysqlNotFound = 501

	// ErrLoginRequired -
	ErrLoginRequired = 401
	// ErrInvalidUser -
	ErrInvalidUser = 1
	// ErrWrongPass -
	ErrWrongPass = 2

	// RespKeyStatus -
	RespKeyStatus = "status"
	// RespKeyData -
	RespKeyData = "data"
)

// ErrNotFound -
var ErrNotFound = errors.New("ErrNotFound")
