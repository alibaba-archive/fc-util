package service

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
)

func GetContentMD5(a []byte) string {
	sum := md5.Sum(a)
	b64 := base64.StdEncoding.EncodeToString(sum[:])
	return b64
}

func GetContentLength(a []byte) string {
	return strconv.Itoa(len(a))
}

func GetSignature(accessKeyId string, accessKeySecret string, request *tea.Request, versionPrefix string) string {
	queriesToSign := make(map[string]string)
	if strings.HasPrefix(request.Pathname, versionPrefix+"/proxy/") {
		queriesToSign = request.Query
	}
	signature := getSignature(request, queriesToSign, accessKeySecret)
	return "FC " + accessKeyId + ":" + signature
}

func getSignature(request *tea.Request, queriesToSign map[string]string, accessKeySecret string) string {
	resource := request.Pathname
	if !strings.Contains(resource, "?") && len(queriesToSign) > 0 {
		resource += "?"
	}
	for key, value := range queriesToSign {
		if value != "" {
			if strings.HasSuffix(resource, "?") {
				resource = resource + key + "=" + value
			} else {
				resource = resource + "&" + key + "=" + value
			}
		}
	}
	return getSignedStr(request, resource, accessKeySecret)
}

// Sorter defines the key-value structure for storing the sorted data in signHeader.
type Sorter struct {
	Keys []string
	Vals []string
}

// newSorter is an additional function for function Sign.
func newSorter(m map[string]string) *Sorter {
	hs := &Sorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]string, 0, len(m)),
	}

	for k, v := range m {
		hs.Keys = append(hs.Keys, k)
		hs.Vals = append(hs.Vals, v)
	}
	return hs
}

// Sort is an additional function for function SignHeader.
func (hs *Sorter) Sort() {
	sort.Sort(hs)
}

// Len is an additional function for function SignHeader.
func (hs *Sorter) Len() int {
	return len(hs.Vals)
}

// Less is an additional function for function SignHeader.
func (hs *Sorter) Less(i, j int) bool {
	return bytes.Compare([]byte(hs.Keys[i]), []byte(hs.Keys[j])) < 0
}

// Swap is an additional function for function SignHeader.
func (hs *Sorter) Swap(i, j int) {
	hs.Vals[i], hs.Vals[j] = hs.Vals[j], hs.Vals[i]
	hs.Keys[i], hs.Keys[j] = hs.Keys[j], hs.Keys[i]
}

func getSignedStr(req *tea.Request, canonicalizedResource, accessKeySecret string) string {
	// Find out the "x-oss-"'s address in header of the request
	temp := make(map[string]string)

	for k, v := range req.Headers {
		if strings.HasPrefix(strings.ToLower(k), "x-fc-") {
			temp[strings.ToLower(k)] = v
		}
	}
	hs := newSorter(temp)

	// Sort the temp by the ascending order
	hs.Sort()

	// Get the canonicalizedOSSHeaders
	canonicalizedOSSHeaders := ""
	for i := range hs.Keys {
		canonicalizedOSSHeaders += hs.Keys[i] + ":" + hs.Vals[i] + "\n"
	}

	// Give other parameters values
	// when sign URL, date is expires
	date := req.Headers["date"]
	contentType := req.Headers["content-type"]
	contentMd5 := req.Headers["content-md5"]

	signStr := req.Method + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n" + canonicalizedOSSHeaders + canonicalizedResource
	h := hmac.New(func() hash.Hash { return sha256.New() }, []byte(accessKeySecret))
	io.WriteString(h, signStr)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return signedStr
}

func Use(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}
