package server

import (
	"golang-from-zero/server/controller"

	"github.com/gin-gonic/gin"
)

func Backend() {
	r := gin.Default()
	ambil := controller.Panggil // untuk ambil data
	r.GET("/", ambil)
	r.Run()
}
