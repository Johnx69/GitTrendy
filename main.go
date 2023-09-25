package main

import (
	"GitTrendy/db"
	"GitTrendy/handler"
	"GitTrendy/repository/repo_impl"
	"GitTrendy/router"

	"github.com/labstack/echo/v4"
)

func main() {

    sql := &db.Sql{
        Host: "localhost",
        Port: 5432,
        UserName: "hoanganh692004",
        Password: "postgres",
        DbName: "GitTrendy",
    }

    sql.Connect()
    defer sql.Close()

    e := echo.New()

    userHandler := handler.UserHandler{
        UserRepo: repo_impl.NewUserRepo(sql),
    }

    api := router.API{
        Echo: e,
        UserHandler: userHandler,
    }
    
    api.SetupRouter()

    e.Logger.Fatal(e.Start(":3000"))
}


