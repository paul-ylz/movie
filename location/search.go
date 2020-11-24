package location

import (
	"fmt"
	"net/http"
)

type dep struct {
	// 	for dependencies (e.g. the service)
}

// Search finds movies by location
func Search(d *dep) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Movies in %s", r.URL.Query().Get("location"))
	}
}
