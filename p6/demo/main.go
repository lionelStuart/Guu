package main

/*
(1) render array
$ curl http://localhost:9999/date
<html>
<body>
    <p>hello, gee</p>
    <p>Date: 2019-08-17</p>
</body>
</html>
*/

/*
(2) custom render function
$ curl http://localhost:9999/students
<html>
<body>
    <p>hello, gee</p>
    <p>0: Geektutu is 20 years old</p>
    <p>1: Jack is 22 years old</p>
</body>
</html>
*/

/*
(3) serve static files
$ curl http://localhost:9999/assets/css/geektutu.css
p {
    color: orange;
    font-weight: 700;
    font-size: 20px;
}
*/

import (
	"Guu/p6/mid"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"Guu/p6/base1"
)

type student struct {
	Name string
	Age  int8
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	guu := base1.New()
	guu.Use(mid.Logger())
	guu.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	guu.LoadHtmlFromGlob("templates/*")
	guu.Static("/assets", "./static")

	stu1 := &student{Name: "jim", Age: 20}
	stu2 := &student{Name: "green", Age: 22}
	guu.GET("/", func(c *base1.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	guu.GET("/students", func(c *base1.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", base1.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	guu.GET("/date", func(c *base1.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", base1.H{
			"title": "guu",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})

	guu.Run(":9999")
}
