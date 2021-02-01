package main

import (
	"context"
	"encoding/json"
	"github.com/anishj0shi/client-go-sample/pkg"
	"google.golang.org/appengine/log"
	"net/http"
)

func main() {
	http.HandleFunc("/create", handle)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		request := &pkg.CreateNSPayload{}
		err := json.NewDecoder(r.Body).Decode(request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Errorf(context.Background(), "unable to unmarshal, err: %+v", err)
		}
		err = pkg.GetClient().CreateNamespace(request.NAME)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Sorry, only POST method is supported."))
	}
}
