/*
 * Revision History:
 *     Initial: 2018/05/14        Tong Yuehong
 */

package utility

import (
	"encoding/base64"
	"io/ioutil"
	"strconv"

	"github.com/fengyfei/gu/libs/logger"
)

const (
	avatar    = "./avatar"
	avasuffix = "jpg"
	File = "./file"
	Filesuffix = "xlsx"
)

func fileName(userID uint32, diff int) string {
	//timestamp := time.Now().Unix()

	//time := time.Unix(timestamp, 0).Format("2006-01-02 03:04:05 PM")
	//time = strings.Replace(time, " ", "", 2)
	id := strconv.FormatUint(uint64(userID), 10)

	if diff == 0 {
		return avatar + id + "." + avasuffix
	}

	return file + id + "." + filesuffix
}

// SaveFile -
func SaveAvatar(userID uint32, image string, diff int) (string, error) {
	fileName := fileName(userID, diff)

	img, _ := base64.StdEncoding.DecodeString(image)
	err := ioutil.WriteFile(fileName, []byte(img), 0777)
	if err != nil {
		logger.Error(err)
	}

	return fileName, err
}
