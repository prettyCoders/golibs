package utils

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type Request struct {
	Header      map[string]string
	Url         string
	RequestBody interface{}
}

func Launch(request *Request) (string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(request.Url)

	for k, v := range request.Header {
		req.Header.Set(k, v)
	}
	if jsonRequestBody, err := json.Marshal(request.RequestBody); err != nil {
		return "", err
	} else {
		req.SetBodyString(string(jsonRequestBody))
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return "", err
	}

	b := resp.Body()
	return string(b), nil
}

func PostWithContentType(contentType string, url string, requestBody interface{}) (string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType(contentType)
	req.Header.SetMethod("POST")

	if jsonRequestBody, err := json.Marshal(requestBody); err != nil {
		return "", err
	} else {
		req.SetBodyString(string(jsonRequestBody))
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return "", err
	}

	b := resp.Body()
	return string(b), nil
}

func Post(url string, requestBody interface{}) (string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)

	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	if jsonRequestBody, err := json.Marshal(requestBody); err != nil {
		return "", err
	} else {
		req.SetBodyString(string(jsonRequestBody))
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		return "", err
	}

	b := resp.Body()
	return string(b), nil
}

func Get(url string) (string, error) {
	_, resp, err := fasthttp.Get(nil, url)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}
