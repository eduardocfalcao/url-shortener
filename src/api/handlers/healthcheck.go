package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/eduardocfalcao/url-shortener/src/api/model"
)

type HealthcheckHandler struct {
}

func (h HealthcheckHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	hostname, err := os.Hostname()
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	m := model.HealthcheckStatus{
		Health: "alive",
		Host:   hostname,
	}

	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//w.WriteHeader(http.StatusOK)
}
