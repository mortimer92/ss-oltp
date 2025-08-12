package models

import (
	"math/rand"
	"time"
)

type Sensors struct {
	Id        int64 //`db:"id" json:"id"` //db tag tells your ORM/DB library which column to map. json tag controls JSON output.
	Name      string
	Location  string
	CreatedAt time.Time
}

// func (s Sensors) NewSensor(param1 string, param2 string) {
// 	panic("unimplemented")
// }

type SensorReadings struct {
	Id         int64
	SensorId   int64
	TempValue  float64
	RecordedAt time.Time
}

func NewSensor(name, location string) *Sensors {
	return &Sensors{
		Id:        0,
		Name:      name,
		Location:  location,
		CreatedAt: time.Now(),
	}
}

func NewSensorReading(sensor Sensors) *SensorReadings {
	// Seed the random number generator with current time
	// rand.Seed(time.Now().UnixNano())//deprecated
	source := rand.NewSource(time.Now().UnixNano())

	rand.New(source)
	// Generate random temperature between -10 and 40
	temp := rand.Float64()*50 - 10 // 50 range → (-10 to 40)

	return &SensorReadings{
		Id:         0,
		SensorId:   sensor.Id,
		TempValue:  temp,
		RecordedAt: time.Now(),
	}
}
