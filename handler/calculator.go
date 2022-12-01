package handler

import "github.com/gofiber/fiber/v2"

type setRedis struct {
	Setter float64 `json:"setter"`
	Action float64 `json:"action"`
}

type CalculatorApi interface {
	PlusCalculate(c *fiber.Ctx) error
	MinusCalculate(c *fiber.Ctx) error
	MultiplyCalculate(c *fiber.Ctx) error
	DivideCalculate(c *fiber.Ctx) error
}

type CalculateApi interface {
	Calculate(c *fiber.Ctx) error
}