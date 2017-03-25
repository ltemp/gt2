package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"log"
	"bufio"
	"gett2/models"
	"encoding/json"
	"github.com/astaxie/beego"
	_ "gett2/routers"
	"path/filepath"
)

const (
	DB_USER     = "gotest"
	DB_PASSWORD = "gotest"
	DB_NAME     = "gotest"
)

type Env struct {
	db *sql.DB
}

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	models.DBCon, err = sql.Open("postgres", dbinfo)
	checkErr(err)
	defer models.DBCon.Close()
	prepareDB()
	readFiles()

	beego.Run()
}

func readFiles() {
	log.Println("Reading files")
	absPath, _ := filepath.Abs("./data/drivers.json")
	readFile(absPath, createDriver)
	absPath, _ = filepath.Abs("./data/metrics.json")
	readFile(absPath, createMetric)
	log.Println("Finished reading files")
}

func prepareDB() {
	log.Println("Dropping metric table")
	execDbOp("drop table if exists metric")

	log.Println("Dropping driver table")
	execDbOp("drop table if exists driver")

	// SQL commands where added here for convenience, full script is in create-tables.sql
	log.Println("Creating table driver")
	execDbOp("CREATE TABLE driver(id INTEGER PRIMARY KEY,name VARCHAR(30) NOT NULL, license_number VARCHAR(10));")

	log.Println("Creating table metric")
	execDbOp("CREATE TABLE metric(name VARCHAR(30) NOT NULL," +
		"value VARCHAR(30) NOT NULL, lon DECIMAL, lat DECIMAL, timestamp BIGINT, " +
		"driver_id INTEGER REFERENCES driver(id) ON DELETE CASCADE);")

	log.Println("Creating table metric indexes")
	execDbOp("CREATE INDEX METRIC_NAME_IDX ON metric(NAME);")
	execDbOp("CREATE INDEX METRIC_VALUE_IDX ON metric(value);")
	execDbOp("CREATE INDEX METRIC_LON_IDX ON metric(lon);")
	execDbOp("CREATE INDEX METRIC_LAT_IDX ON metric(lat);")
	execDbOp("CREATE INDEX METRIC_TIMESTAMP_IDX ON metric(timestamp);")
	execDbOp("CREATE INDEX METRIC_DRIVER_ID_IDX ON metric(driver_id);")
}

func execDbOp(op string) {
	var db = models.DBCon
	_, err := db.Exec(op)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string, createObj func(row string)) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		createObj(scanner.Text())
	}
}

func createDriver(jsonDriver string) {
	if jsonDriver[0] == '[' {
		jsonDriver = jsonDriver[1:]
	}

	var length int = len(jsonDriver)
	if jsonDriver[length-1] != '}' {
		jsonDriver = jsonDriver[0:length-1]
	}
	res := models.Driver{}
	json.Unmarshal([]byte(jsonDriver), &res)
	models.InsertDriver(res)
}

func createMetric(jsonMetric string) {
	res := models.Metric{}
	json.Unmarshal([]byte(jsonMetric), &res)
	models.InsertMetric(res)
}
