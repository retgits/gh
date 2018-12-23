// Package util implements utility methods
package util

import (
	"strings"
)

// LeftPad pads a string with a padStr for a certain number of times
func LeftPad(s string, padStr string, pLen int) string {
	return strings.Repeat(padStr, pLen) + s
}

// RightPadToLen pads a string to a certain length
func RightPadToLen(s string, padStr string, overallLen int) string {
	padCountInt := 1 + ((overallLen - len(padStr)) / len(padStr))
	retStr := s + strings.Repeat(padStr, padCountInt)
	return retStr[:overallLen]
}
