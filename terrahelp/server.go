package terrahelp

import "github.com/kataras/iris"

func StartServer() {
    iris.Get("/:name", func(ctx *iris.Context) {
		ctx.Text(iris.StatusOK, ctx.Param("name"))
	})
    iris.Listen(":8080")
}
