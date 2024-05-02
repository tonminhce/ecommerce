package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/tonminhce/auth-microservice/models"
	"github.com/tonminhce/auth-microservice/routers"
)

var (
	port = flag.String("port", "8080", "port to be used")
	ip   = flag.String("ip", "localhost", "ip to be used")
)

const (
	host     = "localhost"
	portDB   = 5432
	user     = "go_ecommerce_user"
	password = "tonminh123"
	dbname   = "go_ecommerce"
)

func main() {

	// Construct PostgreSQL connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portDB, user, password, dbname)

	// Initialize PostgreSQL database connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	fmt.Println("Connected to the database")
	flag.Parse()
	flags := models.NewFlags(*ip, *port)

	fmt.Println("Starting Api")

	logger := log.New(os.Stdout, "auth", 1)
	route := routers.NewRoute(logger, flags, db)
	engine := route.RegisterRoutes()

	url, _ := flags.GetApplicationUrl()
	engine.Run(*url)
}
