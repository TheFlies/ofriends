package marshal

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

func ParseRequest(r *http.Request, t interface{}) (error error) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		return errors.Wrap(err, "error when decode http request")
	}
	return
}
