/*
 * Revision History:
 *     Initial: 2018/05/14        Tong Yuehong
 */

package utility

import (
	"encoding/base64"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/fengyfei/gu/libs/logger"
)

const (
	avatar    = "./avatar"
	suffix = "jpg"
)

func fileName(userID uint32) string {
	timestamp := time.Now().Unix()

	time := time.Unix(timestamp, 0).Format("2006-01-02 03:04:05 PM")
	time = strings.Replace(time, " ", "", 2)

	id := strconv.FormatUint(uint64(userID), 10)

	return avatar + time + id + "." + suffix
}

// SaveAvatar -
func SaveAvatar(userID uint32, image string) (string, error) {
	fileName := fileName(userID)

	img, _ := base64.StdEncoding.DecodeString(image)
	err := ioutil.WriteFile(fileName, []byte(img), 0777)
	if err != nil {
		logger.Error(err)
	}

	return fileName, err
}
