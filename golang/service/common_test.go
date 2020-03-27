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
	req.Query["fc"] = "test"
	req.Headers["x-fc-key"] = "key"
	req.Headers["x-fc-value"] = "value"

	sign := GetSignature("accessKeyId", "accessKeySecret", req, "version")
	utils.AssertEqual(t, "FC accessKeyId:kLwSLdTyh317hUm7lChbT3FHVfB3MsQgaXINQNnUgZ0=", sign)

	req.Pathname = "version/proxy/"
	sign = GetSignature("accessKeyId", "accessKeySecret", req, "version")
	utils.AssertEqual(t, "FC accessKeyId:qaUk3ESvwnBOcX7186Bq5Niww86dPv6i4MjpWDu+IoA=", sign)
}

func Test_Sorter(t *testing.T) {
	tmp := map[string]string{
		"key":   "ccp",
		"value": "ok",
	}
	sort := newSorter(tmp)
	sort.Sort()

	len := sort.Len()
	utils.AssertEqual(t, len, 2)

	isLess := sort.Less(0, 1)
	utils.AssertEqual(t, isLess, true)

	sort.Swap(0, 1)
	isLess = sort.Less(0, 1)
	utils.AssertEqual(t, isLess, false)
}

func Test_Use(t *testing.T) {
	str := Use(false, "1", "2")
	utils.AssertEqual(t, "2", str)

	str = Use(true, "1", "2")
	utils.AssertEqual(t, "1", str)
}
