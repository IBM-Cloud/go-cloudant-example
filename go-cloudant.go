package main

import (
	//remove the following line to not have your deployment tracker
	"github.com/IBM-Bluemix/cf_deployment_tracker_client_go"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/fjl/go-couchdb"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

	type Note struct {
		Rev   string `json:"_rev,omitempty"`
		Field int64  `json:"field"`
		ID    string `json:_id,omitempty"`
	}

	dbName := "go-cloudant"

	//remove the following line to not have your deployment tracker
	cf_deployment_tracker.Track()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	err := godotenv.Load()
	if err != nil {
		log.Println(".env file does not exist")
	}

	appEnv, err := cfenv.Current()
	if err != nil {
		log.Fatal(err)
	}

	cloudantService, err := appEnv.Services.WithName("cloudant-go-cloudant")
	if err != nil {
		log.Fatal(err)
	}

	cloudantUrl, _ := cloudantService.Credentials["url"].(string)

	cloudant, err := couchdb.NewClient(cloudantUrl, nil)
	if err != nil {
		log.Println(err)
	}

	//ensure db exists
	//if the db exists the db will be returned anyway
	cloudant.CreateDB(dbName)

	var doc Note
	if err := cloudant.DB(dbName).Get("doc", &doc, nil); err != nil {
		log.Println(err)
	}
	if doc == (Note{}) {
		log.Println("nil")
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/hi", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hi",
		})
	})

	port := os.Getenv("VCAP_APP_PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
