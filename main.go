package main

import (
	"fmt"
	"log"
	"net/http"
	"potholder/models"
	"sync"

	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var lock = sync.RWMutex{}

func main() {
	models.MigrateDB()
	go executeTasksLoop()
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/", getStatus),
		rest.Get("/server", getServers),
		rest.Post("/server", addServer),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe("127.0.0.1:7788", api.MakeHandler()))
}

func getStatus(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson("Ok")
}

func getServers(w rest.ResponseWriter, r *rest.Request) {
	var servers = models.GetAllServers()
	w.WriteJson(servers)
}

func addServer(w rest.ResponseWriter, r *rest.Request) {
	server := models.Server{}
	err := r.DecodeJsonPayload(&server)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// if country.Code == "" {
	// 	rest.Error(w, "country code required", 400)
	// 	return
	// }
	lock.Lock()
	models.WriteServerToDB(server)
	lock.Unlock()
	w.WriteJson(&server)
}

var queue = make([]string, 0)

func appendTaskToQueue() {
	fmt.Println("Some task appended")
	queue = append(queue, "some task")
}

func executeTasksLoop() {
	// for {
	// 	fmt.Println("Trying to execute loop")
	// 	fmt.Println(queue)
	// 	if len(queue) == 0 {
	// 		time.Sleep(1000000000)
	// 	} else {
	// 		executeTask(queue[0])
	// 		queue = append(queue[:0], queue[1:]...)
	// 	}
	// }
}

func executeTask(name string) {
	fmt.Println(name)
}
