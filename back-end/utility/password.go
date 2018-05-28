/*
 * Revision History:
 *     Initial: 2018/05/06        Tong Yuehong
 */

package utility

import "golang.org/x/crypto/bcrypt"

// CompareHash -
func CompareHash(digest []byte, password string) bool {
	hex := []byte(password)
	if err := bcrypt.CompareHashAndPassword(digest, hex); err == nil {
		return true
	}

	return false
}

// GenerateHash -
func GenerateHash(password string) ([]byte, error) {
	hex := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(hex, 10)

	if err != nil {
		return hashedPassword, err
	}

	return hashedPassword, nil
}
