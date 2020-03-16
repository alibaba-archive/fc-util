package service

import (
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/alibabacloud-go/tea/utils"
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
