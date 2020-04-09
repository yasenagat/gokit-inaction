package main

import (
	"errors"
	"fmt"
	"github.com/streadway/handy/breaker"
	"log"
	"time"
)

func main() {

	b := breaker.NewBreaker(0.5)

	if b.Allow() {
		fmt.Println("ok")
	}

	for i := 0; i < 100; i++ {
		executeBiz(b, i)
	}
}

func executeBiz(cb breaker.Breaker, index int) (res int, i int, err error) {
	if !cb.Allow() {
		log.Println(index, breaker.ErrCircuitOpen)
		return -1, index, breaker.ErrCircuitOpen
	}

	defer func(start time.Time) {
		//log.Println("e", err)
		if err != nil {
			cb.Failure(time.Since(start))
		} else {
			cb.Success(time.Since(start))
		}

	}(time.Now())

	time.Sleep(200 * time.Millisecond)

	res, err = biz(index)
	i = index
	if err != nil {
		log.Println("index", index, "err", err)
	} else {
		log.Println("index", index, "result", res)
	}
	return
}

func biz(num int) (int, error) {

	if num < 3 {
		return num, nil
	}
	return -1, errors.New("num too small")
}
