package api

import (
	"expvar"
	"fmt"
	"github.com/newpurr/easy-go/application"

	"github.com/gin-gonic/gin"
)

func Expvar(c *gin.Context) {
	for i := 0; i < 1; i++ {
		_ = application.AntsGoroutinePool.Submit(func() {
			fmt.Println("AntsGoroutinePool process")
		})
	}

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	first := true
	report := func(key string, value interface{}) {
		if !first {
			fmt.Fprintf(c.Writer, ",\n")
		}
		first = false
		if str, ok := value.(string); ok {
			fmt.Fprintf(c.Writer, "%q: %q", key, str)
		} else {
			fmt.Fprintf(c.Writer, "%q: %v", key, value)
		}
	}

	fmt.Fprintf(c.Writer, "{\n")
	expvar.Do(func(kv expvar.KeyValue) {
		report(kv.Key, kv.Value)
	})
	fmt.Fprintf(c.Writer, "\n}\n")
}
