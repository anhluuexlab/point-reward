package main

import (
	"log"

	database "github.com/zett-8/go-clean-echo/db"
	_ "github.com/zett-8/go-clean-echo/docs"
	"github.com/zett-8/go-clean-echo/handlers"
	"github.com/zett-8/go-clean-echo/logger"
	"github.com/zett-8/go-clean-echo/services"
	"github.com/zett-8/go-clean-echo/stores"
	"github.com/zett-8/go-clean-echo/utils"
	"go.uber.org/zap"
)

func main() {
	err := logger.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New()
	if err != nil {
		logger.Fatal("failed to connect to the database", zap.Error(err))
	}
	// defer db.Close()

	e := handlers.Echo()
	structValidator := utils.NewStructValidator()
	structValidator.RegisterValidate()
	e.Validator = structValidator
	s := stores.New(db)
	ss := services.New(s)
	h := handlers.New(ss)

	// jwtCheck, err := middlewares.JwtMiddleware()
	// if err != nil {
	// 	logger.Fatal("failed to set JWT middleware", zap.Error(err))
	// }

	handlers.SetDefault(e)
	// handlers.SetApi(e, h, jwtCheck)
	handlers.SetApi(e, h)

	logger.Fatal("failed to start server", zap.Error(e.Start(":8080")))
}
