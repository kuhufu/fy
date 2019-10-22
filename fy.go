package fy

import (
	"net/http"
)

var defaultGroup = New(nil)

func New(cli *http.Client) *RouterGroup {
	g := &RouterGroup{
		cli:    cli,
		header: http.Header{},
	}

	if g.cli == nil {
		g.cli = http.DefaultClient
	}
	return g
}

func Group(path string) *RouterGroup {
	return defaultGroup.Group(path)
}

func GET(path string) *Req {
	return defaultGroup.GET(path)
}

func POST(path string) *Req {
	return defaultGroup.POST(path)
}

func DELETE(path string) *Req {
	return defaultGroup.DELETE(path)
}

func PUT(path string) *Req {
	return defaultGroup.PUT(path)
}

func PATCH(path string) *Req {
	return defaultGroup.PATCH(path)
}

func HEAD(path string) *Req {
	return defaultGroup.HEAD(path)
}

func OPTIONS(path string) *Req {
	return defaultGroup.OPTIONS(path)
}

func TRACE(path string) *Req {
	return defaultGroup.TRACE(path)
}

func Request(method, path string) *Req {
	return defaultGroup.Request(method, path)
}
