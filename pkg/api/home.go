package api

import (
	"fmt"
	"net/http"
)

// Home ...
var Home = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "/")
}
