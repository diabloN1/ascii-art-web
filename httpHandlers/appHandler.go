package httpHandlers

import "net/http"

func AppHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Get(w, r)
	case "POST":
		Post(w, r)
	default:
		http.Error(w, "400 - Bad request (bad method used)", 400)
	}
}
