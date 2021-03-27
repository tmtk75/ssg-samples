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

//go:embed nuxtjs/dist/*
var nuxtjs embed.FS

//go:embed nextjs/out/*
var nextjs embed.FS

//go:embed vite-vue/dist/*
var vitevue embed.FS

//go:embed vite-react/dist/*
var vitereact embed.FS

type hdr struct {
	parent string
	h      http.Handler
}

func (h hdr) ServeHTTP(r http.ResponseWriter, req *http.Request) {
	var a http.Request = *req
	var url url.URL = *req.URL
	url.Path = h.parent + req.URL.Path
	//log.Printf("%v", url.Path)
	a.URL = &url
	h.h.ServeHTTP(r, &a)
}

func Start() {
	e := echo.New()
	e.Use(middleware.Logger())

	serve := func(prefix, parent string, fs embed.FS) {
		e.GET(prefix+"*", echo.WrapHandler(http.StripPrefix(prefix, hdr{parent: parent, h: http.FileServer(http.FS(fs))})))
	}
	serve("/nuxtjs", "nuxtjs/dist/", nuxtjs)
	serve("/nextjs", "nextjs/out/", nextjs)
	serve("/vite-vue", "vite-vue/dist/", vitevue)
	serve("/vite-react", "vite-react/dist/", vitereact)

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
