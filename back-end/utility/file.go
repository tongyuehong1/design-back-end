/*
 * Revision History:
 *     Initial: 2018/05/14        Tong Yuehong
 */

package utility

import (
	"encoding/base64"
	"io/ioutil"

	"github.com/fengyfei/gu/libs/logger"
)

const (
	avatar    = "./avatar/"
	avasuffix = ".jpg"
	File = "./file"
	Filesuffix = "xlsx"
	Grade = "./grade"
)

// SaveGrade -
func SaveGrade(class,subject string, file string) (string,error) {
	 filename := class + subject + File
	 grade, _ := base64.StdEncoding.DecodeString(file)
	 err := ioutil.WriteFile(filename, []byte(grade),0777)
	 if err != nil {
	 	return "", err
	 }

	 return filename, nil
}

func SaveFile(class,file string) (string, error) {
	filename := class + File
	info, _ := base64.StdEncoding.DecodeString(file)
	err := ioutil.WriteFile(filename, []byte(info),0777)
	if err != nil {
		return "", err
	}

	return filename, nil
}

// SaveFile -
func SaveAvatar(name,image string) (string, error) {
	filename := avatar + name + avasuffix

	img, _ := base64.StdEncoding.DecodeString(image)
	err := ioutil.WriteFile(filename, []byte(img), 0777)
	if err != nil {
		logger.Error(err)
	}

	return filename, err
}
