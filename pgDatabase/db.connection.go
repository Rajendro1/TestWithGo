package pgdatabase

import (
	"database/sql"
	"log"
	"os"

	"github.com/Rajendro1/AccuKnox/config"
	"github.com/gin-gonic/gin"
)

var (
	DB  *sql.DB
	err error
)

func Connect(c *gin.Context) {
	var POSTGRES_URL_WITH_DATASBE string = "postgres://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=" + os.Getenv("DB_SSL_MODE") + ""

	var POSTGRES_URL_WITHOUT_DATASBE string = "postgres://" + os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/?sslmode=" + os.Getenv("DB_SSL_MODE") + ""

	log.Println(POSTGRES_URL_WITH_DATASBE)
	log.Println(POSTGRES_URL_WITHOUT_DATASBE)

	createpgdatabase(c, POSTGRES_URL_WITHOUT_DATASBE)

	DB, err = sql.Open("postgres", POSTGRES_URL_WITH_DATASBE)
	if err != nil {
		log.Println("Error To Connect Databae")
		c.JSON(config.SUCCESS, gin.H{
			"error": "Error To Connect Databae",
		})
	}
	if _, tableExecErr := DB.Exec(CreateTableQuery); tableExecErr != nil {
		log.Println("**************Table****************")
		log.Println(tableExecErr.Error())
		log.Println("***************Table***************")
	} else {
		log.Println("create table successfully")
	}
}
func createpgdatabase(c *gin.Context, url string) {
	pgdatabaseCon, err := sql.Open("postgres", url)
	if err != nil {
		log.Println("Error To Connect Databae")
		c.JSON(config.SUCCESS, gin.H{
			"error": "Error To Connect Databae",
		})
	}
	if _, dbExecErr := pgdatabaseCon.Exec(CreatePlatformDatabaseQuery); dbExecErr != nil {
		log.Println("**********pgdatabase********************")
		log.Println(dbExecErr.Error())
		log.Println("********* pgdatabase********************")
	}

}
