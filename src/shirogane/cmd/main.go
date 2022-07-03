package main

import (
	"fmt"
	"github.com/aryahmph/wattpad-clone/src/shirogane/common/env"
	httpCommon "github.com/aryahmph/wattpad-clone/src/shirogane/common/http"
	redisCommon "github.com/aryahmph/wattpad-clone/src/shirogane/common/redis"
	kaguyaDelivery "github.com/aryahmph/wattpad-clone/src/shirogane/internal/delivery/kaguya/http"
	authRepo "github.com/aryahmph/wattpad-clone/src/shirogane/internal/repository/auth/redis"
	authUc "github.com/aryahmph/wattpad-clone/src/shirogane/internal/usecase/auth"
	"github.com/gin-contrib/cors"
	"log"
)

func main() {
	cfg := env.LoadConfig()
	fmt.Println(cfg)
	redisClient := redisCommon.NewRedisClient(cfg.RedisAddress, cfg.RedisUsername, cfg.RedisPassword)

	authRepository := authRepo.NewRedisAuthRepositoryImpl(redisClient)
	authUsecase := authUc.NewAuthUsecaseImpl(authRepository)

	httpServer := httpCommon.NewHTTPServer()
	httpServer.Router.Use(httpCommon.MiddlewareErrorHandler())
	httpServer.Router.Use(cors.Default())

	api := httpServer.Router.Group("/api")

	_, err := kaguyaDelivery.RegisterKaguyaService(api, cfg.KaguyaAddress, authUsecase)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(httpServer.Router.Run(fmt.Sprintf(":%d", cfg.Port)))
}
