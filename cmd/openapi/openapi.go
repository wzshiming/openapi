package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/wzshiming/openapi/ui"
	"github.com/wzshiming/openapi/ui/redoc"
	"github.com/wzshiming/openapi/ui/swaggerui"
)

var port int

func init() {
	flag.IntVar(&port, "p", 9000, "port")
}

func main() {

	mux0 := mux.NewRouter()
	OpenAPI(mux0)
	mux := handlers.RecoveryHandler()(mux0)
	mux = handlers.CombinedLoggingHandler(os.Stdout, mux)
	p := fmt.Sprintf(":%d", port)

	fmt.Printf("Open http://127.0.0.1:%d/swagger/ or http://127.0.0.1:%d/redoc/ with your browser.\n", port, port)
	err := http.ListenAndServe(p, mux)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//  OpenAPI
func OpenAPI(router *mux.Router) *mux.Router {
	router.PathPrefix("/openapi.json").HandlerFunc(ui.Openapi)
	router.PathPrefix("/swagger/openapi.json").HandlerFunc(ui.Openapi)
	router.PathPrefix("/redoc/openapi.json").HandlerFunc(ui.Openapi)
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger", ui.HandleURL(swaggerui.Asset)))
	router.PathPrefix("/redoc/").Handler(http.StripPrefix("/redoc", ui.HandleURL(redoc.Asset)))
	return router
}
