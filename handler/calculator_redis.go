package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pakawatkung/calculator_api/service"
)

type calculatorApi struct {
	redisSrv service.CalculatorService
}

func NewCalculatorApi(redisSrv service.CalculatorService) CalculatorApi {
	return calculatorApi{redisSrv}
}

func (api calculatorApi) PlusCalculate(c *fiber.Ctx) error {

	numbers := setRedis{}
	if err := c.BodyParser(&numbers); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result, err := api.redisSrv.Plus(numbers.Setter, numbers.Action)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Operator":   "Plus",
		"expression": fmt.Sprintf("%v+%v", numbers.Setter, numbers.Action),
		"result":     result,
	})
}

func (api calculatorApi) MinusCalculate(c *fiber.Ctx) error {

	numbers := setRedis{}
	if err := c.BodyParser(&numbers); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result, err := api.redisSrv.Minus(numbers.Setter, numbers.Action)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Operator":   "Minus",
		"expression": fmt.Sprintf("%v-%v", numbers.Setter, numbers.Action),
		"result":     result,
	})
}

func (api calculatorApi) MultiplyCalculate(c *fiber.Ctx) error {

	numbers := setRedis{}
	if err := c.BodyParser(&numbers); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result, err := api.redisSrv.Multiply(numbers.Setter, numbers.Action)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Operator":   "Multiply",
		"expression": fmt.Sprintf("%v*%v", numbers.Setter, numbers.Action),
		"result":     result,
	})
}

func (api calculatorApi) DivideCalculate(c *fiber.Ctx) error {

	numbers := setRedis{}
	if err := c.BodyParser(&numbers); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	result, err := api.redisSrv.Divide(numbers.Setter, numbers.Action)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Operator":   "Divide",
		"expression": fmt.Sprintf("%v/%v", numbers.Setter, numbers.Action),
		"result":     result,
	})
}