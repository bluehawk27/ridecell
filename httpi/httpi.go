package httpi

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/bluehawk27/ridecell/service"
)

// "key := AIzaSyBUxxG1-YiU6YJx9tEywaRGUVaj3EIRJIU"

var s = service.NewService()

// List : http handler for the List endpoint
func List(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	q := req.URL.Query()
	lat := q["lat"][0]
	long := q["long"][0]
	rad := q["radius"][0]

	res, err := s.List(ctx, lat, long, rad)
	if err != nil {
		log.Error("Error from the store: ", err)
	}

	respBytes, jerr := json.Marshal(res)
	if jerr != nil {
		log.Error("Error Marshaling Response", jerr)
	}

	jsonString := string(respBytes)

	io.WriteString(w, jsonString)
}
