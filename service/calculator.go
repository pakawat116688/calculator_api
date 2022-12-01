package service

type Calculate struct {
	Operator string `json:"operator"`
	Number1 float64 `json:"num1"`
	Number2 float64 `json:"num2"`
}

type CalculatorService interface {
	Plus(setter, action float64) (result float64, err error)
	Minus(setter, action float64) (result float64, err error)
	Multiply(setter, action float64) (result float64, err error)
	Divide(setter, action float64) (result float64, err error)
}

type CalculateOneService interface {
	Calculate(cal Calculate) (result float64, err error)
}
