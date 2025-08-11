package internal

import (
	"math/rand"
	"time"
)

type Sensors struct {
	id        int64 //`db:"id" json:"id"` //db tag tells your ORM/DB library which column to map. json tag controls JSON output.
	name      string
	location  string
	createdAt time.Time
}

type SensorReadings struct {
	id         int64
	sensorId   int64
	tempValue  float64
	recordedAt time.Time
}

func NewSensor(name, location string) *Sensors {
	return &Sensors{
		id:        0,
		name:      name,
		location:  location,
		createdAt: time.Now(),
	}
}

func NewSensorReading(sensor Sensors) *SensorReadings {
	// Seed the random number generator with current time
	// rand.Seed(time.Now().UnixNano())//deprecated
	rand.New()
	// Generate random temperature between -10 and 40
	temp := rand.Float64()*50 - 10 // 50 range → (-10 to 40)

	return &SensorReadings{
		id:        0,
		sensorId:  sensor.id,
		tempValue: temp,
	}
}
