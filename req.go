package fy

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Req struct {
	err      error
	cli      *http.Client
	url      string
	method   string
	rawQuery string
	header   http.Header
	cookies  []*http.Cookie
	body     io.Reader
}

func (r *Req) SetHeader(key, val string) *Req {
	r.header.Set(key, val)
	return r
}

func (r *Req) AddCookie(cookie *http.Cookie) *Req {
	r.cookies = append(r.cookies, cookie)
	return r
}

func (r *Req) AddQuery(key, val string) *Req {
	if r.rawQuery == "" {
		r.rawQuery += key + "=" + val
		return r
	}
	r.rawQuery += "&" + key + "=" + val
	return r
}

func (r *Req) SetBody(body interface{}) *Req {
	var reader io.Reader
	switch body := body.(type) {
	case []byte:
		reader = bytes.NewReader(body)
	case string:
		reader = strings.NewReader(body)
	case io.Reader:
		reader = body
	default:
		panic(fmt.Errorf("unsupport type: %T", body))
	}
	r.body = reader
	return r
}

func (r *Req) Do() *Result {
	if r.err != nil {
		return &Result{
			resp: nil,
			err:  r.err,
		}
	}

	req, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		return &Result{
			resp: nil,
			err:  r.err,
		}
	}
	if r.header != nil {
		req.Header = r.header
	}

	for _, c := range r.cookies {
		req.AddCookie(c)
	}

	if r.rawQuery != "" {
		req.URL.RawQuery = r.rawQuery
	}

	resp, err := r.cli.Do(req)
	return &Result{
		resp: resp,
		err:  err,
	}
}
