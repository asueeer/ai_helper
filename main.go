package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	r := gin.Default()
	register(r)

	err := r.Run("0.0.0.0:1988") // listen and serve on 0.0.0.0:9990
	if err != nil {
		panic(err)
	}
}
