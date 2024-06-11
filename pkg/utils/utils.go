package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err1 := json.Unmarshal([]byte(body), x); err1 != nil {
			return
		}
	}
}
