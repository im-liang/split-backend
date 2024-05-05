package main

import (

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.Use(gin.Logger())

  // Recovery middleware recovers from any panics and writes a 500 if there was one.
  router.Use(gin.Recovery())

  router.Run(":8090")
}