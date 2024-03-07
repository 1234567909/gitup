package config

import (
	"github.com/hashicorp/consul/api"
	"log"
	"net/http"
)

func ConsulIn() {
	config := api.DefaultConfig()
	consul, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	regis := &api.AgentServiceRegistration{
		Name: "czh",
		Tags: []string{"java8", "api"},
		Port: 8080,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://127.0.0.1:8080/health",
			Interval: "10s",
		},
	}

	err = consul.Agent().ServiceRegister(regis)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	log.Println("8080...")
	http.ListenAndServe(":8080", nil)
}
