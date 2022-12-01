package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/pakawatkung/calculator_api/errs"
)

type calculate struct {
	redisClint *redis.Client
}

func NewCalculate(redisClint *redis.Client) CalculateOneService {
	return calculate{redisClint}
}

func (s calculate) Calculate(cal Calculate) (result float64, err error) {

	var key string

	switch cal.Operator {
	case "+":
		key = fmt.Sprintf("plus::%v+%v", cal.Number1, cal.Number2)
		result = cal.Number1 + cal.Number2
	case "-":
		key = fmt.Sprintf("minus::%v-%v", cal.Number1, cal.Number2)
		result = cal.Number1 - cal.Number2
	case "*":
		key = fmt.Sprintf("multi::%v*%v", cal.Number1, cal.Number2)
		result = cal.Number1 * cal.Number2
	case "/":
		key = fmt.Sprintf("divide::%v/%v", cal.Number1, cal.Number2)
		result = cal.Number1 / cal.Number2
	default:
		return 0, errs.ErrOperatorNotFound()
	}

	// Redis
	fmt.Println(key)
	if resultStr, err := s.redisClint.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(resultStr), &result) == nil {
			fmt.Println("check unmarshal")
			return result, nil
		}
	}

	// Set Redis
	if resultByte, err := json.Marshal(result); err == nil {
		s.redisClint.Set(context.Background(), key, string(resultByte), time.Second*10)
	}
	fmt.Println("Set Redis")

	return result, nil
}
