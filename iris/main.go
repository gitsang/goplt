package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)

type Info struct {
	ID   int64  `json:"id"`
	City string `json:"city"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Info Info
}

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello Iris [GET]")
	})
	app.Post("/", func(ctx iris.Context) {
		ctx.Write([]byte("Hello Iris [POST]"))
	})

	helloParty := app.Party("/hello")
	helloParty.Get("/world", func(ctx iris.Context) {
		ctx.WriteString("Hello World [GET]")
	})

	app.Post("/json", func(ctx iris.Context) {
		var user User
		ctx.ReadJSON(&user)
		fmt.Printf("[POST] %v", user)

		yujinB := User{
			Name: "Yujin B",
			Age:  14,
			Info: Info{
				ID:   7,
				City: "April",
			},
		}
		ctx.JSON(yujinB)
	})

	app.Get("/json", func(ctx iris.Context) {
		yujinA := User{
			Name: "Yujin A",
			Age:  14,
			Info: Info{
				ID:   7,
				City: "April",
			},
		}
		ctx.JSON(yujinA)
	})

	app.Run(iris.Addr(":8085"), iris.WithCharset("UTF-8"))

	for {
		time.Sleep(time.Hour)
	}
}
