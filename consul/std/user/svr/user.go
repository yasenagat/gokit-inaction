package main

import (
	"encoding/json"
	"flag"
	"gitee.com/godY/gokit-inaction/consul/std/user"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	"io"
	"log"
	"net/http"
)

var DB = make(map[string]user.User)

func main() {

	address := flag.String("address", ":30000", "")

	r := mux.NewRouter()

	r.Path("/health").Methods("GET").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "success")
	})

	r.Path("/users/{id}").Methods("GET").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		params := mux.Vars(request)

		if u, ok := DB[params["id"]]; ok {
			json.NewEncoder(writer).Encode(u)
		} else {
			writer.WriteHeader(http.StatusNotFound)
			io.WriteString(writer, "no this user")
		}

	})

	r.Path("/users/{id}").Methods("DELETE").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		params := mux.Vars(request)

		if _, ok := DB[params["id"]]; ok {
			delete(DB, params["id"])
			io.WriteString(writer, "success")
		} else {
			writer.WriteHeader(http.StatusNotFound)
			io.WriteString(writer, "no this user")
		}

	})

	r.Path("/users").Methods("GET").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		users := make([]user.User, 0)
		for _, v := range DB {
			users = append(users, v)
		}

		json.NewEncoder(writer).Encode(users)

	})

	r.Path("/users").Methods("POST").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		u := user.User{}
		json.NewDecoder(request.Body).Decode(&u)
		if u.Username == "" {
			writer.WriteHeader(http.StatusBadRequest)
			io.WriteString(writer, "require username")
			return
		}
		id := uuid.New().String()
		u.Id = id
		DB[id] = u
		io.WriteString(writer, "success")
	})
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.10.208:8500"
	c, e := api.NewClient(cfg)

	if e != nil {
		log.Println(e)
	}
	id := uuid.New().String()

	e = RegSvr(c, user.SVRNAME, "localhost", "http://192.168.10.208:30000/health", id, 30000)

	if e != nil {
		log.Println("reg fail", e)
	}

	http.ListenAndServe(*address, r)

	e = c.Agent().ServiceDeregister(id)

	if e != nil {
		log.Println("remove fail", e)
	}
}

func RegSvr(c *api.Client, name, address, checkhttpurl, id string, port int) error {
	agent := c.Agent()
	return agent.ServiceRegister(&api.AgentServiceRegistration{ID: id, Name: name, Port: port, Address: address, Check: &api.AgentServiceCheck{HTTP: checkhttpurl, Interval: "2s"}})
}
