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

	key := fmt.Sprintf("cal::%v%v%v", cal.Number1, cal.Operator, cal.Number2)

	// Redis
	fmt.Println(key)
	if resultStr, err := s.redisClint.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(resultStr), &result) == nil {
			fmt.Println("check unmarshal")
			return result, nil
		}
	}

	switch cal.Operator {
	case "+":
		result = cal.Number1 + cal.Number2
	case "-":
		result = cal.Number1 - cal.Number2
	case "*":
		result = cal.Number1 * cal.Number2
	case "/":
		result = cal.Number1 / cal.Number2
	default:
		return 0, errs.ErrOperatorNotFound()
	}

	// Set Redis
	if resultByte, err := json.Marshal(result); err == nil {
		s.redisClint.Set(context.Background(), key, string(resultByte), time.Second*10)
	}
	fmt.Println("Set Redis")

	return result, nil
}
