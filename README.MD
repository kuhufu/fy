```go
var (
    code        int
    body        string
    contentType string
)

err := fy.GET("https://starmicro.happyelements.cn/v1/idol/forumdetail").
    SetHeader("Authorization", "{token}").
    AddCookie(&http.Cookie{Name:"{name}", Value:"{value}"}).
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