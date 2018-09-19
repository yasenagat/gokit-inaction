package main

import (
	"net/http"
	"log"
	"github.com/hashicorp/consul/api"
	"os"
	"gitee.com/godY/gokit-inaction/consul/std/user"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"github.com/google/uuid"
	"strings"
)

func main() {

	cfg := api.DefaultConfig()
	cfg.Address = "192.168.10.208:8500"
	c, e := api.NewClient(cfg)

	if e != nil {
		log.Println(e)
		os.Exit(-1)
	}

	h := c.Health()

	if e != nil {
		log.Println(e)
		os.Exit(-1)
	}

	s, _, e := h.Service(user.SVRNAME, "", true, nil)
	if e != nil {
		log.Println(e)
	}

	for i, v := range s {
		log.Println(i, v.Service.Service, v.Service.Address, v.Service.Port)
	}

	url := "http://" + s[0].Service.Address + ":" + strconv.Itoa(s[0].Service.Port) + "/users"
	log.Println(url)

	//add

	req, e := http.NewRequest("POST", url, nil)
	u := user.User{}
	u.Id = uuid.New().String()
	u.Username = "yasenagat"

	byte, e := json.Marshal(u)
	if e != nil {
		log.Println(e)
	}
	req.Body = ioutil.NopCloser(strings.NewReader(string(byte)))

	if e != nil {
		log.Println(e)
	}

	res, e := http.DefaultClient.Do(req)
	if e != nil {
		log.Println(e)
	}
	bytes, e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Println(e)
	}
	log.Println(string(bytes))

	//query
	req, e = http.NewRequest("GET", url, nil)

	if e != nil {
		log.Println(e)
	}
	res, e = http.DefaultClient.Do(req)
	if e != nil {
		log.Println(e)
	}
	bytes, e = ioutil.ReadAll(res.Body)
	if e != nil {
		log.Println(e)
	}
	log.Println(string(bytes))

}
