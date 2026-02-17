package nasatest

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

type Server struct {
	URL              string
	StatusCode       int
	ResponseBody     string
	ResponseBodyJSON bool
}

func NewServer() *httptest.Server {
	return httptest.NewServer(&Server{
		StatusCode:       http.StatusOK,
		ResponseBody:     getAPODSuccessfulResponse(),
		ResponseBodyJSON: true,
	})
}

func (s Server) ServeHTTP(w http.ResponseWriter, rq *http.Request) {
	if strings.Contains(rq.URL.String(), "/planetary/apod") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(s.StatusCode)

		_, err := w.Write([]byte(s.ResponseBody))
		if err != nil {
			return
		}
	}
}

func getAPODSuccessfulResponse() string {
	rp := `{
				"copyright":"test",
				"date":"2006-01-01",
				"explanation":"test",
				"hdurl":"test",
				"media_type":"test",
				"service_version":"test",
				"title":"test",
				"url":"test"
			}`
	return rp
}
