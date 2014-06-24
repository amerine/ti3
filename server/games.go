package server

import (
	"net/http"
	"path"
	"io"
)

func wireUpGameResource(mux *http.ServeMux, prefix string) {
	res := path.Join(prefix, "games")
	games := gamesResource{}
	mux.Handle(res, &games)
}

type gamesResource struct {
}


func (games *gamesResource) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	io.WriteString(rw, "HI")
}
