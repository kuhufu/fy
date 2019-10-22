package fy

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	var (
		code        int
		body        string
		contentType string
	)

	v1 := defaultGroup.Group("https://starmicro.happyelements.cn/v1").
		SetHeader("Authorization", "{token}").
		AddCookie(&http.Cookie{Name: "{name}", Value: "{value}"})

	err := v1.GET("/idol/forumdetail").
		AddQuery("id", "15180").
		Do().
		BindCode(&code).
		BindBody(&body).
		BindHeader("Content-Type", &contentType).
		Err()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(code)
	fmt.Println(body)
	fmt.Println(contentType)
}
