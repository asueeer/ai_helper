package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	r := gin.Default()
	register(r)
	var err error
	err = r.Run("0.0.0.0:1988")
	if err != nil {
		panic(err)
	}

}
