package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
	"time"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context){
		ctx.WriteString("Hello Iris [GET]")
	})
	app.Post("/", func(ctx iris.Context) {
		ctx.Write([]byte("Hello Iris [POST]"))
	})

	helloParty := app.Party("/hello")
	helloParty.Get("/world", func(ctx iris.Context) {
		ctx.WriteString("Hello World [GET]")
	})

	app.Handle(http.MethodGet, "/region-list", httpGetRegionList)

	app.Run(iris.Addr(":8085"), iris.WithCharset("UTF-8"))

	for {
		time.Sleep(time.Hour)
	}
}

func httpGetRegionList(ctx iris.Context) {
	output := "output string"
	fmt.Printf("GET")
	_, _ = ctx.WriteString(output)
}

