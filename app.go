package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	_ "main/configs"
	config "main/configs"
	"main/controller"
	"main/service"
)

func main() {

	ac := config.MakeAccessLog()
	defer ac.Close() // Close the underline file.
	app := iris.New()
	// Register the middleware (UseRouter to catch http errors too).
	app.UseRouter(ac.Handler)
	app.Logger().SetLevel("debug")
	mvc.Configure(app.Party("/user"), users)
	app.Listen(":4000")
}

func users(app *mvc.Application) {
	userService := service.NewUserService(config.Database)
	app.Register(userService)
	app.Handle(new(controller.UserController))
}
