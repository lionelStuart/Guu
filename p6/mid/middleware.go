package mid

import (
	. "Guu/p6/base1"
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v ", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func MidForV2() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v MidFor V2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}

}
