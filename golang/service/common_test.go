package service

import (
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/aliyun/fc-util/golang/utils"
)

func Test_GetContentMD5(t *testing.T) {
	str := GetContentMD5([]byte("test"))
	utils.AssertEqual(t, "CY9rzUYh03PK3k6DJie09g==", str)
}

func Test_GetContentLength(t *testing.T) {
	str := GetContentLength([]byte("test"))
	utils.AssertEqual(t, "4", str)
}

func Test_GetSignature(t *testing.T) {
	req := tea.NewRequest()
	req.Query["test"] = "ok"

	sign := GetSignature("accessKeyId", "accessKeySecret", req, "version")
	utils.AssertEqual(t, "FC accessKeyId:NDLiuxe3uHGNaZAUJQ0Fm1zVhxY=", sign)

	req.Pathname = "version/proxy/"
	sign = GetSignature("accessKeyId", "accessKeySecret", req, "version")
	utils.AssertEqual(t, "FC accessKeyId:jHx/oHoHNrbVfhncHEvPdHXZwHU=", sign)
}

func Test_Use(t *testing.T) {
	str := Use(false, "1", "2")
	utils.AssertEqual(t, "2", str)

	str = Use(true, "1", "2")
	utils.AssertEqual(t, "1", str)
}

func Test_Is4XXor5XX(t *testing.T) {
	ok := Is4XXor5XX(500)
	utils.AssertEqual(t, true, ok)

	ok = Is4XXor5XX(300)
	utils.AssertEqual(t, false, ok)

	ok = Is4XXor5XX(600)
	utils.AssertEqual(t, false, ok)
}
