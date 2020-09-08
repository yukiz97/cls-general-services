package apiservices

import (
	"github.com/gorilla/mux"
	"github.com/yukiz97/cls-general-services/lcservices"
	"github.com/yukiz97/utils/restapi"
	"log"
	"net/http"
	"strconv"
)

func home(response http.ResponseWriter, _ *http.Request) {
	restapi.RespondWithJSON(response, http.StatusOK, "Welcome to restful API of cls - general services")
}

func getProductList(response http.ResponseWriter, _ *http.Request) {
	listProduct := lcservices.GetProductList()

	restapi.RespondWithJSON(response, http.StatusOK, listProduct)
}

//InitRestfulAPIServices init customer restfull api
func InitRestfulAPIServices(listenPort int) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home)
	router.HandleFunc("/getProductList/", getProductList).Methods("GET")

	println("Running CLS - general services.... - Listen to port :" + strconv.Itoa(listenPort))

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(listenPort), router))
}
