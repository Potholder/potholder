package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	go executeTasksLoop()
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/", getStatus),
		rest.Get("/server", getServers),
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
	// var country *Country

	w.WriteJson("country")
}

var queue = make([]string, 0)

func appendTaskToQueue() {
	fmt.Println("Some task appended")
	queue = append(queue, "some task")
}

func executeTasksLoop() {
	for {
		fmt.Println("Trying to execute loop")
		fmt.Println(queue)
		if len(queue) == 0 {
			time.Sleep(1000000000)
		} else {
			executeTask(queue[0])
			queue = append(queue[:0], queue[1:]...)
		}
	}
}

func executeTask(name string) {
	fmt.Println(name)
}

// package main
//
// import (
// 	"fmt"
// 	"os"
//
// 	"./models"
// 	"github.com/codegangsta/cli"
// 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// )
//
// func main() {
// 	app := cli.NewApp()
// 	app.Name = "Potholder"
// 	app.Usage = "Super easy web servers administration."
// 	app.Flags = []cli.Flag{
// 		cli.StringFlag{
// 			Name:  "webservices",
// 			Usage: "Show all web sercives",
// 		},
// 		cli.StringFlag{
// 			Name:  "servers",
// 			Usage: "webservices WEBSERVICENAME to SERVER",
// 		},
// 		cli.StringFlag{
// 			Name:  "movewebservice",
// 			Usage: "movewebservice WEBSERVICENAME to SERVER",
// 		},
// 	}
// 	app.Action = func(c *cli.Context) error {
// 		fmt.Println("Hello friend!")
// 		return nil
// 	}
//
// 	app.Run(os.Args)
//
// 	models.MigrateDB()
// 	models.Create("192.168.1.1", "Some server")
// 	// models.PaintServers()
// }
