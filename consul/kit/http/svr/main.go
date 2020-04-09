package main

import (
	"flag"
	yt "gitee.com/godY/tcgetway/time"
	"github.com/go-kit/kit/sd/consul"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//nohup ./time -checkhttpurl=http://192.168.3.125:14444 -consul.regist.addr=http://192.168.3.125 -id=ts125 >time.out &
//nohup ./time -checkhttpurl=http://192.168.10.208:14444 -consul.regist.addr=http://192.168.10.208 -id=ts208 >time.out &

//nohup ./time -checkhttpurl=http://192.168.10.210:14444 -consul.regist.addr=http://192.168.10.210 -id=ts210 >time.out &
func main() {

	id := flag.String("id", "", "")
	name := flag.String("name", "TimeSvr", "")
	addr := flag.String("http.addr", ":14444", "")
	consuladdr := flag.String("consul.addr", "localhost:8500", "")
	address := flag.String("consul.regist.addr", ":14444", "")
	checkhttpurl := flag.String("checkhttpurl", "http://192.168.10.37:14444", "")
	port := flag.Int("port", 14444, "")

	flag.Parse()

	s := kithttp.NewServer(func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//if r, ok := request.(string); ok {
		//	log.Println(r)
		//	if r == "1" {
		//		return nil, errors.New(*id + " : " + "Error Req 1")
		//	} else if r == "2" {
		//		return *id + " : " + "Error Req 2", nil
		//	}
		//}
		return request, nil
	}, func(i context.Context, req *http.Request) (request interface{}, err error) {

		bytes, e := ioutil.ReadAll(req.Body)

		if e != nil {
			return nil, e
		}
		return string(bytes), nil
	}, func(i context.Context, writer http.ResponseWriter, v interface{}) error {

		if v != nil {
			if r, ok := v.(string); ok {

				if r == "1" {
					writer.WriteHeader(http.StatusBadRequest)
					//io.WriteString(writer, r)
					return errors.New(*id + " Err req 1")
				} else {
					time.Sleep(500 * time.Millisecond)
					io.WriteString(writer, *id+" => "+yt.Now())
				}
				return nil
			} else {
				time.Sleep(1000 * time.Millisecond)
				io.WriteString(writer, *id+" => "+yt.Now())
			}
		} else {
			return errors.New(*id + "Error")
		}

		return nil
	})

	cfg := api.DefaultConfig()
	cfg.Address = *consuladdr
	c, e := api.NewClient(cfg)

	if e != nil {
		log.Println(e)
		os.Exit(-1)
	}

	kitc := consul.NewClient(c)
	r := &api.AgentServiceRegistration{Name: *name, Port: *port, Address: *address, Check: &api.AgentServiceCheck{HTTP: *checkhttpurl, Interval: "2s"}}
	kitc.Register(r)

	http.ListenAndServe(*addr, s)
}
