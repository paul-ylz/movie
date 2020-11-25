package location

import (
	"fmt"
	"net/http"
)

// Search finds movies by location
func Search() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Movies in %s", r.URL.Query().Get("location"))
	}
}
