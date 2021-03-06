//Additional service for test getting YAML (such as Python DB module works).
//1. Listen requests on port http :8083 path /add
//2. Decode body from YAML.
//3. Print structure on the screen.

package main

import (
	"fmt"
	"log"
	"log/syslog"
	"net/http"

	"github.com/Kv-062-DevOps/post_conv_jtoy/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v3"
)

func main() {

	syslogger, err := syslog.New(syslog.LOG_INFO, "syslog_test")
	if err != nil {
		log.Fatalln(err)
	}

	log.SetOutput(syslogger)
	log.Println("service getyamlsrv started")

	router := NewRouter()
	//log.Fatal(http.ListenAndServe(":8083", router))
	log.Fatal(http.ListenAndServe("127.0.0.1:8083", router))

}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	var message handlers.Employ
	if err := yaml.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "400 Bad Request from POST converting service", 400)
		fmt.Println(err)
		return
	}

	fmt.Println(message)
	w.WriteHeader(200)
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routesArray {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

var routesArray = []Route{
	Route{
		Name:        "Root",
		Method:      "POST",
		Pattern:     "/add",
		HandlerFunc: RootHandler,
	},
}
