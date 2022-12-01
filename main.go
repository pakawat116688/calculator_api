package main

import (
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/calculator_api/handler"
	"github.com/pakawatkung/calculator_api/service"
)

func main()  {

	myRedis := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	redisSrv := service.NewCalculatorRedis(myRedis)
	redisApi := handler.NewCalculatorApi(redisSrv)

	redisOneSrv := service.NewCalculate(myRedis)
	redisOneApi := handler.NewCalculateApi(redisOneSrv)

	app := fiber.New()

	// 4 function
	app.Get("/plus", redisApi.PlusCalculate)
	app.Get("/minus", redisApi.MinusCalculate)
	app.Get("/multiply", redisApi.MultiplyCalculate)
	app.Get("/divide", redisApi.DivideCalculate)

	// 1 function
	app.Get("/calculate", redisOneApi.Calculate)

	app.Listen(":8000")
}