package verificationutil

import (
	"errors"
	"regexp"
)

const emailPattern = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
const mobilePattern = `^1[3|4|5|7|8][0-9]{9}$`

// VerifyEmail checks if an email is in valid format
// if an empty string is provided, then false and an empty string error is returned
// if email syntax regex compilation false and an empty string error is returned
// otherwise the result for matching input string and nil for error is returned
func VerifyEmail(email string) (bool, error) {

	if email == "" {
		return false, errEmpty
	}

	reg, regErr := regexp.Compile(emailPattern)

	if regErr != nil {
		return false, regErr
	}
	return reg.MatchString(email), nil
}

// VerifyMobile checks if an mobile is in valid format
// if an empty string is provided, then false and an empty string error is returned
// if mobile syntax regex compilation false and an empty string error is returned
// otherwise the result for matching input string and nil for error is returned
// if an empty string is provided, then false and an empty string error is returned
// if mobile syntax regex compilation false and an empty string error is returned
// otherwise the result for matching input string and nil for error is returned
func VerifyMobile(mobile string) (bool, error) {

	if mobile == "" {
		return false, errEmpty
	}

	reg, regErr := regexp.Compile(mobilePattern)

	if regErr != nil {
		return false, regErr
	}
	return reg.MatchString(mobile), nil
}

var errEmpty = errors.New("empty string")
