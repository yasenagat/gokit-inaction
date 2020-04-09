package main

import (
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"log"
	"strconv"
)

const svrname = "newsvr"

func main() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	c, e := api.NewClient(cfg)
	if e != nil {
		log.Println(e)
	}

	//e = kv(c, e)
	//
	//registSvr(c, e)

	//getSvr(e, c.Agent())

	getHealthService(c, e)

	//getConsulHealthService(c, e)

	//removeAllcriticalSvr(c, e)
}

func removeAllcriticalSvr(c *api.Client, e error) {
	services, e := c.Agent().Services()
	if e != nil {
		log.Println(e)
	}
	for k, v := range services {
		log.Println(k, v.Service)
		checks, _, e := c.Health().Checks(v.Service, nil)
		if e != nil {
			log.Println(e)
		}
		if checks.AggregatedStatus() == "critical" {
			e = c.Agent().ServiceDeregister(k)
			if e != nil {
				log.Println(e)
			}
		}
	}
}

func getConsulHealthService(c *api.Client, e error) {
	h := c.Health()
	entries, _, e := h.Service("consul", "", true, nil)
	for _, e := range entries {
		log.Println(e.Service.Address, e.Service.Port)
	}
}

func getHealthService(c *api.Client, e error) {
	h := c.Health()
	entries, _, e := h.Service(svrname, "", true, nil)
	for _, e := range entries {
		log.Println(e.Node.Node, e.Node.Datacenter, e.Service.ID, e.Service.Address, e.Service.Port, e.Checks.AggregatedStatus())
	}
}

func getSvr(e error, agent *api.Agent) {
	services, e := agent.Services()
	if e != nil {
		log.Println(e)
	}
	for k, v := range services {
		log.Println(k, v)
	}
	//for i := 0; i < 5; i++ {
	//	a := services[svrname]
	//	if a != nil {
	//		log.Println(a.Address, a.Port)
	//	}
	//}
}

func registSvr(c *api.Client, e error) (*api.Agent, error) {
	agent := c.Agent()
	e = agent.ServiceRegister(&api.AgentServiceRegistration{ID: uuid.New().String(), Name: svrname, Port: 30000, Address: "192.168.10.208", Check: &api.AgentServiceCheck{HTTP: "http://192.168.10.208:30000/health", Interval: "2s"}})
	if e != nil {
		log.Println(e)
	}
	//e = agent.ServiceRegister(&api.AgentServiceRegistration{ID: "u2", Name: svrname, Port: 30000, Address: "192.168.10.208", Check: &api.AgentServiceCheck{HTTP: "http://192.168.10.208:30000/health", Interval: "2s"}})
	//if e != nil {
	//	log.Println(e)
	//}
	//e = agent.ServiceRegister(&api.AgentServiceRegistration{ID: "u3", Name: svrname, Port: 30000, Address: "192.168.3.125", Check: &api.AgentServiceCheck{HTTP: "http://192.168.3.125:30000/health", Interval: "2s"}})
	return agent, e
}

func kv(c *api.Client, e error) error {
	kv := c.KV()
	for i := 0; i < 3; i++ {
		kvp := api.KVPair{Key: "username" + strconv.Itoa(i), Value: []byte("tom")}
		_, e = kv.Put(&kvp, nil)
		//_, e = kv.Delete("username"+strconv.Itoa(i), nil)
		if e != nil {
			log.Println(e)
		}
	}
	return e
}
