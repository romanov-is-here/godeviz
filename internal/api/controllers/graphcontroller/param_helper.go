package graphcontroller

import (
	"net/http"
	"net/url"
)

// TODO refac ?

func fromQuery(w http.ResponseWriter, r *http.Request, name string, required bool) (string, bool) {
	pathEnc := r.URL.Query().Get(name)
	path, err := url.QueryUnescape(pathEnc)
	if err != nil {
		http.Error(w, "Failed to decode '"+name+"'", http.StatusBadRequest)
		return "", false
	}

	if required && path == "" {
		http.Error(w, "Missing required'"+name+"' query parameter", http.StatusBadRequest)
		return "", false
	}

	return path, true
}
