/*
 * Revision History:
 *     Initial: 2018/05/21        Chen Yanchen
 */

package util

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return hashedPassword, err
	}

	return hashedPassword, nil
}

func CompareHash(hashPwd []byte, pwd string) bool {
	hex := []byte(pwd)
	if err := bcrypt.CompareHashAndPassword(hashPwd, hex); err == nil {
		return true
	}

	return false
}
