package gee

import (
	"log"
	"fmt"
	"time"
)

var TimeZone = time.FixedZone("UTC", int((0 * time.Hour).Seconds()))

func LoadTimeZone(x string, i int)(*time.Location){
	TimeZone = time.FixedZone(x, int((time.Duration(i) * time.Hour).Seconds()))
	return TimeZone
}

func LoadLocation(x *time.Location){
	TimeZone = x
}

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		log.SetPrefix(time.Now().In(TimeZone).Format("[Web] 2006-01-02 15:04:05 "))
		log.SetFlags(0)
		if c.StatusCode == 0 {
			c.StatusCode = 200
		}
		// Calculate resolution time
		log.Printf("[%s] [%d] %s in %v", c.Req.Method, c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func Error(x ...interface{}){
	log.SetPrefix(time.Now().In(TimeZone).Format("[Error] 2006-01-02 15:04:05 "))
	log.SetFlags(0)
	var s string
	for _, o := range x {
		s = s + fmt.Sprintf("%v",o)
	}
	log.Println(s)
}
