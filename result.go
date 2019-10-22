package fy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

type Result struct {
	resp *http.Response
	err  error
}

func (r *Result) Bytes() (data []byte, err error) {
	if r.err != nil {
		return nil, r.err
	}
	defer r.resp.Body.Close()
	return ioutil.ReadAll(r.resp.Body)
}

func (r *Result) String() (data string, err error) {
	if r.err != nil {
		return "", r.err
	}
	defer r.resp.Body.Close()
	bytes, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (r *Result) JsonUnmarshal(v interface{}) (err error) {
	if r.err != nil {
		return r.err
	}

	defer r.resp.Body.Close()
	data, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	return nil
}

func (r *Result) Resp() (resp *http.Response, err error) {
	resp, err = r.resp, r.err
	return
}

func (r *Result) Err() (err error) {
	return r.err
}

func (r *Result) BindBody(v interface{}) *Result {
	if r.err != nil {
		return r
	}

	bytes, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		r.err = err
		return r
	}

	val := reflect.ValueOf(v).Elem()
	switch v.(type) {
	case *string:
		val.SetString(string(bytes))
	case *[]byte:
		val.SetBytes(bytes)
	default:
		r.err = json.Unmarshal(bytes, v) //缺省绑定为json
		return r
	}
	return r
}

func (r *Result) BindCode(v interface{}) *Result {
	if r.err != nil {
		return r
	}

	code := r.resp.StatusCode
	val := reflect.ValueOf(v).Elem()
	switch v.(type) {
	case *int, *int16, *int32, *int64:
		val.SetInt(int64(code))
	case *string:
		val.SetString(strconv.Itoa(code))
	default:
		r.err = fmt.Errorf("not match type: %T", v)
		return r
	}
	return r
}

func (r *Result) BindHeader(headerName string, v interface{}) *Result {
	if r.err != nil {
		return r
	}

	headerValue := r.resp.Header.Get(headerName)
	val := reflect.ValueOf(v).Elem()
	switch v.(type) {
	case *string:
		val.SetString(headerValue)
	default:
		r.err = fmt.Errorf("not match type: %T", v)
		return r
	}
	return r
}

func (r *Result) BindCookie(cookieName string, v interface{}) *Result {
	if r.err != nil {
		return r
	}

	var cookieValue string
	for _, cookie := range r.resp.Cookies() {
		if cookie.Name == cookieName {
			cookieValue = cookie.Value
			break
		}
	}
	val := reflect.ValueOf(v).Elem()
	switch v.(type) {
	case *string:
		val.SetString(cookieValue)
	default:
		r.err = fmt.Errorf("not match type: %T", v)
		return r
	}
	return r
}
