package youhaosuda

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	_ "fmt"
	"net/url"
	"sort"
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
