package main

import (
	"github.com/sony/gobreaker"
	"math/rand"
	"github.com/pkg/errors"
	"log"
	"strconv"
	"time"
)

func main() {

	set := gobreaker.Settings{}
	set.Name = "test.cb"
	//set.Interval = time.Second
	set.MaxRequests = 1
	set.Timeout = time.Millisecond * 300
	set.ReadyToTrip = func(counts gobreaker.Counts) bool {
		log.Println("req", counts.Requests)
		log.Println("f", counts.ConsecutiveFailures)
		if counts.ConsecutiveFailures > 2 {
			return true
		}
		return false
	}
	set.OnStateChange = func(name string, from gobreaker.State, to gobreaker.State) {
		log.Println(name, from, ">>", to)
	}
	breaker := gobreaker.NewCircuitBreaker(set)

	executeBiz(breaker, 1)
	time.Sleep(1000 * time.Millisecond)
	executeBiz(breaker, 2)
	time.Sleep(1000 * time.Millisecond)
	executeBiz(breaker, 3)
	executeBiz(breaker, 3)
	executeBiz(breaker, 3)
	time.Sleep(300 * time.Millisecond)
	executeBiz(breaker, 4)
	executeBiz(breaker, 4)

	//c := make(chan R)
	//for i := 0; i < 20; i++ {
	//	time.Sleep(100 * time.Millisecond)
	//	go func(n int) {
	//		r := R{}
	//		res, index, e := executeBiz(breaker, n)
	//		r.res = res
	//		r.index = index
	//		r.err = e
	//		c <- r
	//	}(i)
	//}
	//
	//for i := 0; i < 20; i++ {
	//	v := <-c
	//	log.Println("index", v.index, "res", v.res, "e", v.err)
	//}

}

type R struct {
	res   int
	err   error
	index int
}

func executeBiz(cb *gobreaker.CircuitBreaker, index int) (int, int, error) {
	result, err := cb.Execute(func() (interface{}, error) {

		res := biz()

		if index > 3 {
			return res, nil
		}
		//time.Sleep(time.Second * 5)
		return nil, errors.New("biz error=" + strconv.Itoa(index))
	})

	if err != nil {
		log.Println("index", index, "err", err)
		return -1, index, err
	} else {
		log.Println("index", index, "result", result)
		return result.(int), index, nil
	}
}

func biz() int {
	return rand.Intn(100)
}
