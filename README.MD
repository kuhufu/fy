```go
var (
    code        int
    body        string
    contentType string
)

v1 := fy.Group("https://starmicro.happyelements.cn/v1").
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
    fmt.Println(err)
}

fmt.Println(code)
fmt.Println(body)
fmt.Println(contentType)
```