package main

import (
	"embed"
	"log"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	KeyPort = "port"
)

//go:embed nuxtjs-sample/dist/*
var nuxtjs embed.FS

type hdr struct {
	h http.Handler
}

func (h hdr) ServeHTTP(r http.ResponseWriter, req *http.Request) {
	var a http.Request = *req
	var url url.URL = *req.URL
	url.Path = "nuxtjs-sample/dist/" + req.URL.Path
	a.URL = &url
	h.h.ServeHTTP(r, &a)
}

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", hdr{h: http.FileServer(http.FS(nuxtjs))})))

	log.Printf("start listening at %s", viper.GetString(KeyPort))
	log.Fatal(e.Start(viper.GetString(KeyPort)))
}

func init() {
	pf := RootCmd.PersistentFlags()
	pf.StringP("port", "p", ":8080", "Port to listen")
	viper.BindPFlag("port", pf.Lookup("port"))
	viper.BindEnv("port", "PORT")
}

var RootCmd = &cobra.Command{
	Use: "app",
	Run: func(cmd *cobra.Command, args []string) { Start() },
}

func main() {
	RootCmd.Execute()
}
