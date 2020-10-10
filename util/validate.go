package util

import "regexp"

// CheckEmail return ture if the email is valid or false if invalid
func CheckEmail(email string) bool {
	if email == "" {
		return false
	}
	if matched, _ := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", email); matched {
		return true
	}
	return false
}

// CheckPassword return ture if the password is valid or false if invalid
func CheckPassword(password string) bool {
	if password == "" {
		return false
	}
	if matched, _ := regexp.MatchString("^[0-9a-zA-Z@.]{6,30}$", password); matched {
		return true
	}
	return false
}

// CheckPhoneNumber return ture if the phoneNumber is valid or false if invalid
func CheckPhoneNumber(phoneNumber string) bool {
	if phoneNumber == "" {
		return false
	}
	if matched, _ := regexp.MatchString("^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])d{8}$", phoneNumber); matched {
		return true
	}
	return false
}

// CheckUsername return ture if the username is valid or false if invalid
func CheckUsername(username string) bool {
	if username == "" {
		return false
	}
	// 字母开头，允许5-16字节，允许字母数字下划线
	if matched, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_]{4,15}$", username); matched {
		return true
	}
	return false
}
