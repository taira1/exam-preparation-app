package controller

import "net/http"

type controller interface {
	process(w http.ResponseWriter, r *http.Request) map[string]interface{}
}
