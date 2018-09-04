package main

import (
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"time"
	"log"
)

func main() {

	hystrix.ConfigureCommand("test", hystrix.CommandConfig{Timeout: 10000, MaxConcurrentRequests: 10, RequestVolumeThreshold: 2})

	for i := 0; i < 10; i++ {
		r, ii, e := executeBiz(i)
		log.Println(r, ii, e)
	}

	//c := make(chan int)
	//for i := 0; i < 100000; i++ {
	//	go func(num int) {
	//		r, ii, e := executeBiz(num)
	//		log.Println(r, ii, e)
	//		c <- r
	//	}(i)
	//}

	//for i := 0; i < 100; i++ {
	//	<-c
	//}

	//errc := make(chan error)
	//hystrixStreamHandler := hystrix.NewStreamHandler()
	//hystrixStreamHandler.Start()
	//go func() {
	//	errc <- http.ListenAndServe(net.JoinHostPort("", "9000"), hystrixStreamHandler)
	//}()
	//
	//log.Println("exit", <-errc)
}

func executeBiz(index int) (res int, i int, err error) {

	var r int
	e := hystrix.Do("test", func() error {

		r, err = biz(index)
		return err
	}, nil)

	if e != nil {
		return -1, index, e
	} else {
		return r, index, nil
	}
}

func biz(num int) (int, error) {

	time.Sleep(1000 * time.Millisecond)
	if num < 1 {
		return num, nil
	}
	return -1, errors.New("num too small")
}
