package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// запуск сервера
func MainServer() {

	router := routers()

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8888",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

// все пути
func routers() *mux.Router {
	router := mux.NewRouter()

	// params key = GUID
	router.HandleFunc("/getToken", getToken).Methods(http.MethodPost)
	// Headers key = Refresher, Authorization
	router.HandleFunc("/refreshToken", refreshToken).Methods(http.MethodPost)
	// Headers key = Authorization
	router.HandleFunc("/useToken", useToken).Methods(http.MethodPost)

	return router
}
