package routes

import "net/http"

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	apiRoutes(mux)
	uiRoutes(mux)
	wsRoute(mux)
	return mux
}
