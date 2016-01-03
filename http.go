package youhaosuda

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type request struct {
	url     string
	method  string
	headers map[string]string
	data    string
}

type response struct {
	status  int
	body    string
	headers map[string]string
}

func (c *Config) Get(token, path string) response {
	s := []string{c.ApiUrl, c.ApiVersion, path}
	url := strings.Join(s, "/")
	h := make(map[string]string)
	h["X-API-ACCESS-TOKEN"] = token
	r := request{
		url:     url,
		method:  "GET",
		headers: h,
	}
	res, _ := r.request()
	return res
}

func (c *Config) Post(token, path, data string) response {
	s := []string{c.ApiUrl, c.ApiVersion, path}
	url := strings.Join(s, "/")
	h := make(map[string]string)
	h["X-API-ACCESS-TOKEN"] = token
	h["Content-Type"] = "application/json"
	r := request{
		url:     url,
		method:  "POST",
		headers: h,
		data:    data,
	}
	res, _ := r.request()
	return res
}

func (c *Config) Put(token, path, data string) response {
	s := []string{c.ApiUrl, c.ApiVersion, path}
	url := strings.Join(s, "/")
	h := make(map[string]string)
	h["X-API-ACCESS-TOKEN"] = token
	h["Content-Type"] = "application/json"
	r := request{
		url:     url,
		method:  "PUT",
		headers: h,
		data:    data,
	}
	res, err := r.request()
	if err != nil {
		panic(err)
	}
	return res
}

func (c *Config) Delete(token, path string) response {
	s := []string{c.ApiUrl, c.ApiVersion, path}
	url := strings.Join(s, "/")
	h := make(map[string]string)
	h["X-API-ACCESS-TOKEN"] = token
	r := request{
		url:     url,
		method:  "DELETE",
		headers: h,
	}
	res, _ := r.request()
	return res
}

func (r *request) request() (response, error) {
	empty := response{}
	data := strings.NewReader(r.data)
	req, err := http.NewRequest(r.method, r.url, data)
	if err != nil {
		return empty, err
	}
	for k, v := range r.headers {
		req.Header.Add(k, v)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   DefaultTimeOut,
	}
	res, err := client.Do(req)
	if err != nil {
		return empty, err
	}
	defer res.Body.Close()
	c, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return empty, err
	}

	result := response{
		status:  res.StatusCode,
		body:    string(c),
		headers: ConvertMap(res.Header),
	}
	return result, nil
}

func (r *response) parseJson(result interface{}) {
	byt := []byte(r.body)
	json.Unmarshal(byt, &result)
}
