package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

/*
Server is structure to store server data
*/
type Server struct {
	gorm.Model
	IP          string
	Name        string
	SSHUser     string
	SSHPassword string
	SSHPort     int64
}

/*
MigrateDB migrates all schemas in database
*/
func MigrateDB() {
	getDB().AutoMigrate(&Server{})
}

/*
WriteServerToDB is for creation of new records for server
*/
func WriteServerToDB(server Server) {
	getDB().Create(&server)
}

func GetAllServers() []Server {
	var servers []Server
	getDB().Find(&servers)
	return servers
}

/*
PaintServers prints servers
*/
func PaintServers() {
	var server Server
	getDB().Find(&server)
	fmt.Printf("Hello %v\n", server)
}
