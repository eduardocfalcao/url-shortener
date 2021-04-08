package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eduardocfalcao/url-shortener/src/api/entities"
	"github.com/eduardocfalcao/url-shortener/src/api/mocks"
	"github.com/eduardocfalcao/url-shortener/src/api/model"
)

func Handler(fn http.HandlerFunc) http.HandlerFunc {
	return fn
}
func Test_Create(t *testing.T) {

	s := &model.ShorturlRequest{
		Name:     "name",
		Shorturl: "sh",
		URL:      "http://something.com",
	}
	cases := map[string]struct {
		ShortUrl   *model.ShorturlRequest
		DecodeErr  error
		ReturnID   int
		CreateErr  error
		StatusCode int
	}{
		"Should works fine": {ShortUrl: s, DecodeErr: nil, ReturnID: 1, CreateErr: nil, StatusCode: http.StatusCreated},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			repositoryMock := &mocks.ShortUrlRepository{}
			sut := ShortUrlHandler{repositoryMock}

			s := entities.ShortUrl{
				Name:     tc.ShortUrl.Name,
				ShortUrl: tc.ShortUrl.Shorturl,
				URL:      tc.ShortUrl.URL,
			}
			repositoryMock.On("Create", s).Return(tc.ReturnID, tc.CreateErr)

			var jsonBytes *bytes.Buffer
			if tc.ShortUrl != nil {
				jsonString, _ := json.Marshal(*tc.ShortUrl)
				jsonBytes = bytes.NewBuffer(jsonString)
			}

			req, err := http.NewRequest("POST", "/", jsonBytes)
			if err != nil {
				log.Fatalf("Error trying to generate the post request. %s", err)
			}

			rr := httptest.NewRecorder()
			Handler(sut.Create).ServeHTTP(rr, req)

			if rr.Result().StatusCode != tc.StatusCode {
				t.Errorf("Expected StatusCode: %d. StatusCode returned: %d", tc.StatusCode, rr.Result().StatusCode)
			}
		})

	}
}
