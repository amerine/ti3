package server

import (
	"net/http"
)

func NewServeMux() *http.ServeMux {
	prefix := "/api"
	sm := http.NewServeMux()
	wireUpGameResource(sm, prefix)

	return sm
}
