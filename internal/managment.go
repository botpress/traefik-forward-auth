package tfa

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type ManagmentServer struct {
	Enabled    bool
	Port       int
	HealthPath string
}

func NewManagmentServer(conf *Config) *ManagmentServer {
	return &ManagmentServer{
		Enabled:    conf.ManagmentEnabled,
		Port:       conf.ManagmentPort,
		HealthPath: conf.HealthPath,
	}
}

func (managmentServer *ManagmentServer) Serve() {
	if !managmentServer.Enabled {
		return
	}

	router := mux.NewRouter()

	router.HandleFunc(managmentServer.HealthPath, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", managmentServer.Port), router)

	if err != nil {
		log.Fatal(errors.Wrap(err, "unable to start managment server"))
	}
}
