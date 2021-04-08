package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eduardocfalcao/url-shortener/src/api/entities"
	"github.com/eduardocfalcao/url-shortener/src/api/interfaces"
	"github.com/eduardocfalcao/url-shortener/src/api/model"
)

type ShortUrlHandler struct {
	repository interfaces.ShortUrlRepository
}

func (h ShortUrlHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var request model.ShorturlRequest
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error:     "The request sent couldn't be processed.",
			ErrorCode: model.InvalidRequest,
		})
		return
	}

	s := entities.ShortUrl{
		Name:     request.Name,
		ShortUrl: request.Shorturl,
		URL:      request.URL,
	}

	if _, err := h.repository.Create(s); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{
			Error:     "The request sent couldn't be processed.",
			ErrorCode: model.InvalidRequest,
		})
		log.Printf("An error occured when calling the Create method from the repository.")
		return
	}

	w.WriteHeader(http.StatusCreated)
}
