package fy

import "net/http"

type RouterGroup struct {
	cli     *http.Client
	err     error
	path    string
	header  http.Header
	cookies []*http.Cookie
}

func (g *RouterGroup) Group(path string) *RouterGroup {
	cookiesCopy := make([]*http.Cookie, len(g.cookies))
	copy(cookiesCopy, g.cookies)
	return &RouterGroup{
		cli:     g.cli,
		err:     g.err,
		path:    pathJoin(g.path, path),
		header:  g.header.Clone(),
		cookies: cookiesCopy,
	}
}

func (g *RouterGroup) SetHeader(key, val string) *RouterGroup {
	g.header.Set(key, val)
	return g
}

func (g *RouterGroup) AddCookie(cookie *http.Cookie) *RouterGroup {
	g.cookies = append(g.cookies, cookie)
	return g
}

func (g *RouterGroup) GET(path string) *Req {
	return g.Request("GET", path)
}

func (g *RouterGroup) POST(path string) *Req {
	return g.Request("POST", path)
}

func (g *RouterGroup) PUT(path string) *Req {
	return g.Request("PUT", path)
}

func (g *RouterGroup) DELETE(path string) *Req {
	return g.Request("DELETE", path)
}

func (g *RouterGroup) PATCH(path string) *Req {
	return g.Request("PATCH", path)
}

func (g *RouterGroup) HEAD(path string) *Req {
	return g.Request("HEAD", path)
}

func (g *RouterGroup) OPTIONS(path string) *Req {
	return g.Request("OPTIONS", path)
}

func (g *RouterGroup) TRACE(path string) *Req {
	return g.Request("TRACE", path)
}

func (g *RouterGroup) Request(method string, path string) *Req {
	cookiesCopy := make([]*http.Cookie, len(g.cookies))
	copy(cookiesCopy, g.cookies)
	return &Req{
		err:     g.err,
		cli:     g.cli,
		url:     pathJoin(g.path, path),
		method:  method,
		header:  g.header.Clone(),
		cookies: cookiesCopy,
		body:    nil,
	}
}

func pathJoin(p1, p2 string) string {
	i := len(p2)
	if p2[i-1] == '/' {
		p2 = p2[:i-1] //去除path尾部的最后一个 '/'
	}
	if p1 != "" && p2[0] != '/' {
		return p1 + "/" + p2
	}
	return p1 + p2
}
