package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "embed"
)

//go:embed public/*
var f embed.FS

func main() {
	router := gin.Default()

	router.StaticFile("/", "./public/index.html")					// http://localhost:3000/

	router.Static("/public", "./public")									// http://localhost:3000/public/employee.html

	router.StaticFS("/fs", http.FileSystem(http.FS(f)))		// http://localhost:3000/fs/public/employee.html

	log.Fatal(router.Run(":3000"))
}
