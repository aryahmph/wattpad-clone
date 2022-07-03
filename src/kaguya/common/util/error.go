package util

import (
	"strconv"
	"strings"
)

func GetErrorStatusCodeAndMessage(err error) (uint32, string) {
	errs := strings.Split(err.Error(), ":")
	errCode, _ := strconv.Atoi(errs[0])
	return uint32(errCode), errs[1]
}
