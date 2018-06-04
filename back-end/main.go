package main

import (
	"net/http"
	_ "github.com/tongyuehong1/design-back-end/back-end/routers"
	"github.com/astaxie/beego"
	mysql "github.com/tongyuehong1/design-back-end/back-end/init"
	"github.com/astaxie/beego/plugins/cors"
	"fmt"
)

type handler struct {
	h http.Handler
}

func main() {
	var (
		fileServe handler
	)
	go func() {
		fileServe.h = http.StripPrefix("/avatar/", http.FileServer(http.Dir("./avatar")))
		http.Handle("/avatar/", fileServe)

		if err := http.ListenAndServe(":21001", nil); err != nil {
			fmt.Println(err)
		}
	}()

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	mysql.InitSql()
	beego.Run()
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	//w.Header().Set("content-type", "application/json")

	h.h.ServeHTTP(w, r)
}
