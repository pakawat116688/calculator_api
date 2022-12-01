package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/calculator_api/service"
)

type calculateApi struct {
	redisOneSrv service.CalculateOneService
}

func NewCalculateApi(redisOneSrv service.CalculateOneService) CalculateApi {
	return calculateApi{redisOneSrv}
}

func (h calculateApi) Calculate(c *fiber.Ctx) error {

	numbers := service.Calculate{}
	if err := c.BodyParser(&numbers); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result, err := h.redisOneSrv.Calculate(numbers)
	if err != nil {
		return c.JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": result,
	})

}
