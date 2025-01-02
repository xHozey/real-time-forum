package routes

import "net/http"

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	apiRoutes(mux)
	appRoutes(mux)
	return mux
}
