package service

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"net/url"
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
	stringToSign := buildRpcStringToSign(request, queriesToSign)
	signature := sign(stringToSign, accessKeySecret, "&")
	return "FC " + accessKeyId + ":" + signature
}

func Use(condition bool, a string, b string) string {
	if condition {
		return a
	}
	return b
}

func Is4XXor5XX(code int) bool {
	return code >= 400 && code < 600
}

func buildRpcStringToSign(request *tea.Request, queriesToSign map[string]string) (stringToSign string) {
	stringToSign = getUrlFormedMap(queriesToSign)
	stringToSign = strings.Replace(stringToSign, "+", "%20", -1)
	stringToSign = strings.Replace(stringToSign, "*", "%2A", -1)
	stringToSign = strings.Replace(stringToSign, "%7E", "~", -1)
	stringToSign = url.QueryEscape(stringToSign)
	stringToSign = request.Method + "&%2F&" + stringToSign
	return
}

func getUrlFormedMap(source map[string]string) (urlEncoded string) {
	urlEncoder := url.Values{}
	for key, value := range source {
		urlEncoder.Add(key, value)
	}
	urlEncoded = urlEncoder.Encode()
	return
}

func sign(stringToSign, accessKeySecret, secretSuffix string) string {
	secret := accessKeySecret + secretSuffix
	signedBytes := shaHmac1(stringToSign, secret)
	signedString := base64.StdEncoding.EncodeToString(signedBytes)
	return signedString
}

func shaHmac1(source, secret string) []byte {
	key := []byte(secret)
	hmac := hmac.New(sha1.New, key)
	hmac.Write([]byte(source))
	return hmac.Sum(nil)
}
