package graphcontroller

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/romanov-is-here/godeviz/internal/graph/depgraph"
)

// TODO refac ?

func strFromQuery(w http.ResponseWriter, r *http.Request, name string, required bool) (string, bool) {
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

func boolFromQuery(w http.ResponseWriter, r *http.Request, name string, required bool) (bool, bool) {
	fromQuery := r.URL.Query().Get(name)

	if fromQuery == "" {
		if required {
			http.Error(w, "Missing required'"+name+"' query parameter", http.StatusBadRequest)
		}
		return false, false
	}

	value, err := strconv.ParseBool(fromQuery)
	if err != nil {
		http.Error(w, "Failed to parse bool param '"+name+"'", http.StatusBadRequest)
		return false, false
	}

	return value, true
}

func intFromQuery(w http.ResponseWriter, r *http.Request, name string, required bool) (int, bool) {
	fromQuery := r.URL.Query().Get(name)

	if fromQuery == "" {
		if required {
			http.Error(w, "Missing required'"+name+"' query parameter", http.StatusBadRequest)
		}
		return 0, false
	}

	value, err := strconv.Atoi(fromQuery)
	if err != nil {
		http.Error(w, "Failed to parse int param '"+name+"'", http.StatusBadRequest)
		return 0, false
	}

	return value, true
}

func getFilter(w http.ResponseWriter, r *http.Request) (*depgraph.Filter, bool) {
	filter := depgraph.NewDefaultFilter()

	if showStandard, ok := boolFromQuery(w, r, "show_standard", false); ok {
		filter.ShowStandard = showStandard
	}

	if showPlatform, ok := boolFromQuery(w, r, "show_platform", false); ok {
		filter.ShowPlatform = showPlatform
	}

	if showOuter, ok := boolFromQuery(w, r, "show_outer", false); ok {
		filter.ShowOuter = showOuter
	}

	if focusPackage, ok := strFromQuery(w, r, "focusPackage", false); ok {
		filter.FocusPackage = focusPackage

		if focusType, okType := intFromQuery(w, r, "focusType", false); okType {
			filter.FocusType = depgraph.FocusType(focusType)
		}
	}

	return filter, true
}
