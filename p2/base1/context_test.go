package base1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func httpServeDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"msg":"mock test",code:200"}`)

}

func TestHttpTest(t *testing.T) {
	httptest.NewServer(http.HandlerFunc(httpServeDemo))
	t.Log("finish test ")
}

type User struct {
	Username string `json:"username"`
	Password string `json:"Password"`
}

func TestContext_Json(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := newContext(w, r)

		fmt.Printf("ctx %+v \n", c)
		c.Json(http.StatusOK, H{
			"username": c.PostForm("username"),
			"Password": c.PostForm("password")})
		fmt.Printf("ctx new %+v \n", c)

	}))
	defer s.Close()

	resp, err := http.PostForm(s.URL,
		url.Values{"username": {"jim"}, "password": {"1234"}})
	if err != nil {
		t.Fatal("fail :", err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("result body:%s\n", string(body))

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		t.Fatal("failure parse body ", body)
	}
	fmt.Printf("result user:%+v\n", user)

}
