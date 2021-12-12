package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	Init()
	r := gin.Default()
	register(r)
	debug := true
	var err error
	if debug {
		err = r.Run("0.0.0.0:1988")
	} else {
		err = r.RunTLS("0.0.0.0:1988", "./6182282_asueeer.com.pem", "./6182282_asueeer.com.key") // listen and serve on 0.0.0.0:9990
	}
	if err != nil {
		panic(err)
	}
}
