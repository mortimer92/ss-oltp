package main

import (
	"database/sql"
	"fmt"
	"log"
	"mortimer92/ss-oltp/internal/models"

	_ "github.com/lib/pq" // PostgreSQL driver
)


func main() {

	host := "localhost"
	port := "5432"
	user := "myuser"
	password := "mypassword"
	dbname := "mydatabase"

	//construct the conn string
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil{
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	sensor := models.NewSensor("funky_blipper", "a hobbit hole")

	sensorReadings := models.NewSensorReading(*sensor)

	fmt.Println("Sensor ID:", sensor.Id)
	fmt.Println("Sensor Name:", sensor.Name)
	fmt.Println("Sensor Location:", sensor.Location)
	fmt.Println("Sensor Created At:", sensor.CreatedAt)
	fmt.Println("Sensor Reading ID:", sensorReadings.TempValue)
}
