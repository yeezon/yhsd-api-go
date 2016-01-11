package youhaosuda

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	_ "fmt"
	"io"
	"io/ioutil"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func URLEncode(data map[string]string) string {
	paras := url.Values{}
	for k, v := range data {
		paras.Add(k, v)
	}
	return paras.Encode()
}

func ConvertMap(data map[string][]string) map[string]string {
	h := make(map[string]string)
	for k, v := range data {
		h[k] = v[0]
	}
	return h
}

func CalAuthorization(key, secret string) string {
	arr := []string{key, secret}
	s := strings.Join(arr, ":")
	sEnc := b64.URLEncoding.EncodeToString([]byte(s))
	arr = []string{"Basic", sEnc}
	s = strings.Join(arr, " ")
	return s
}

func CalHMAC(secret string, data map[string]string) string {
	delete(data, "hmac")
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var paras []string
	for _, k := range keys {
		v := k + "=" + data[k]
		paras = append(paras, v)
	}
	s := strings.Join(paras, "&")
	sign := hmac.New(sha256.New, []byte(secret))
	sign.Write([]byte(s))
	return hex.EncodeToString(sign.Sum(nil))
}

func CalBase64HMAC(token string, data string) string {
	sign := hmac.New(sha256.New, []byte(token))
	sign.Write([]byte(data))
	return b64.StdEncoding.EncodeToString(sign.Sum(nil))
}

func CalBase64Aes(secret string, data string) string {
	key := []byte(secret)
	block, err := aes.NewCipher(key[0:16])
	if err != nil {
		panic(err)
	}
	blockSize := aes.BlockSize
	orig := []byte(data)

	padding := blockSize - len(orig)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	orig = append(orig, padtext...)

	cypted := make([]byte, len(orig))
	iv := key[16:32]
	encrypter := cipher.NewCBCEncrypter(block, iv)

	encrypter.CryptBlocks(cypted, orig)

	return b64.StdEncoding.EncodeToString(cypted)
}

func RemoveSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}

func VerifyHMAC(form_value map[string][]string) bool {
	res := ConvertMap(form_value)
	hmac = res["hmac"]
	cal_hmac := CalHMAC(res)
	return hmac == cal_hmac
}

// webhook and openpayment
func VerifyWebhook(webhook_token string, header map[string][]string, body io.Reader) bool {
	if v := header["X-YHSD-HMAC-SHA256"]; len(v) > 0 {
		hmac := v[0]
		body, err := ioutil.ReadAll(body)
		if err != nil {
			panic(err)
		}
		data := string(body)
		cal_hmac := CalBase64HMAC(webhook_token, data)
		return hmac == cal_hmac
	}
	return false
}

func RedirectUserUrl(domain, secret string, data map[string]string) string {
	jsonString, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	res := CalBase64Aes(secret, jsonString)
	d := RemoveSuffix(domain, "/")
	u := []string{d, "account/multipass/login", res}
	return strings.Join(u, "/")
}

func RedirectYouPayUrl(domain, you_pay_key, you_pay_secret string, data map[string]string) string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var paras []string
	for _, k := range keys {
		if v := data[k]; len(v) > 0 {
			res := []string{k, v}
			paras = append(paras, strings.Join(res, "="))
		}
	}
	sign_para := strings.Join(paras, "&") + you_pay_secret
	md := md5.New()
	io.WriteString(md, para)
	sign := md.Sum(nil)

	root_url := RemoveSuffix(domain, "/")

	var url_paras []string
	for _, k := range keys {
		if v := data[k]; len(v) > 0 {
			escaped := url.QueryEscape(v)
			res := []string{k, escaped}
			url_paras = append(url_paras, strings.Join(res, "="))
		}
	}
	res := []string{"sign", sign}
	url_paras = append(url_paras, strings.Join(res, "="))
	redirect_para := strings.Join(url_paras, "&")

	rdirect_url := []string{root_url, redirect_para}
	return strings.Join(rdirect_url, "?")
}
func VerifyEnoughBucket(resp_header map[string]string) bool {
	if v := resp_header["X-YHSD-SHOP-API-CALL-LIMIT"]; len(v) > 0 {
		s := strings.Split(v, "/")
		if len(s) == 2 {
			bucket, err := strconv.Atoi(s[0])
			total, err := strconv.Atoi(s[1])
			return bucket < total
		}
	}
	return false
}
