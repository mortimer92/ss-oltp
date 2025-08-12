package main

import (
	"fmt"
	"mortimer92/ss-oltp/internal/models"
)

func main() {

	// sensor2 := models.NewSensor("funky_blipper", "a hobbit hole")
	// fmt.Println("Sensor Name:", sensor.sensorId)
	// sensorReadings := sensorreadings.NewSensor("funky_blipper", "a hobbit hole")

	// fmt.Println("Sensor Name:", sensorReadings.Name)

	// // sensor := internal.SensorReadingsSensors.NewSensor("funky_blipper", "a hobbit hole")
	// // sensorReadings := SensorReadings.NewSensorReading()

	// sensorReadings2 := models.NewSensorReading("funky_blipper", "a hobbit hole", 42.0, "2023-10-01T12:00:00Z")

	// var sensor models.Sensors
	sensor := models.NewSensor("funky_blipper", "a hobbit hole")

	sensorReadings := models.NewSensorReading(*sensor)

	fmt.Println("Sensor ID:", sensor.Id)
	fmt.Println("Sensor Name:", sensor.Name)
	fmt.Println("Sensor Location:", sensor.Location)
	fmt.Println("Sensor Created At:", sensor.CreatedAt)
	fmt.Println("Sensor Reading ID:", sensorReadings.TempValue)
}
