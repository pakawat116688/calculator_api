package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
)

type calculatorRedis struct {
	redisClint *redis.Client
}

func NewCalculatorRedis(redisClint *redis.Client) CalculatorService {
	return calculatorRedis{redisClint}
}

func (s calculatorRedis) Plus(setter, action float64) (result float64, err error) {

	key := fmt.Sprintf("plus::%v+%v", setter, action)
	fmt.Println(key)
	fmt.Println(setter, action)

	if resultStr, err := s.redisClint.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(resultStr), &result) == nil {
			fmt.Println("check unmarshal plus")
			return result, nil
		}
	}

	result = setter + action
	if resultByte, err := json.Marshal(result); err == nil{
		s.redisClint.Set(context.Background(), key, string(resultByte), time.Second * 10)
	}
	fmt.Println("Set Redis plus")
	return result, nil
}

func (s calculatorRedis) Minus(setter, action float64) (result float64, err error) {

	key := fmt.Sprintf("minus::%v-%v", setter, action)
	fmt.Println(key)
	fmt.Println(setter, action)

	if resultStr, err := s.redisClint.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(resultStr), &result) == nil {
			fmt.Println("check unmarshal minus")
			return result, nil
		}
	}

	result = setter - action
	if resultByte, err := json.Marshal(result); err == nil{
		s.redisClint.Set(context.Background(), key, string(resultByte), time.Second * 10)
	}
	fmt.Println("Set Redis minus")
	return result, nil
}

func (s calculatorRedis) Multiply(setter, action float64) (result float64, err error) {

	key := fmt.Sprintf("multi::%v*%v", setter, action)
	fmt.Println(key)

	if resultStr, err := s.redisClint.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(resultStr), &result) == nil {
			fmt.Println("check unmarshal muti")
			return result, nil
		}
	}

	result = setter * action
	if resultByte, err := json.Marshal(result); err == nil{
		s.redisClint.Set(context.Background(), key, string(resultByte), time.Second * 10)
	}
	fmt.Println("Set Redis muti")
	return result, nil
}

func (s calculatorRedis) Divide(setter, action float64) (result float64, err error) {

	key := fmt.Sprintf("divide::%v/%v", setter, action)
	fmt.Println(key)

	if resultStr, err := s.redisClint.Get(context.Background(), key).Result(); err == nil {
		if json.Unmarshal([]byte(resultStr), &result) == nil {
			fmt.Println("check unmarshal muti")
			return result, nil
		}
	}

	result = setter / action
	if resultByte, err := json.Marshal(result); err == nil{
		s.redisClint.Set(context.Background(), key, string(resultByte), time.Second * 10)
	}
	fmt.Println("Set Redis divide")
	return result, nil
}