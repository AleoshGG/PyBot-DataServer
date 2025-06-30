package models

type Sensor struct {
	Sensor_id    int    `json:"sensor_id"`
	Sensor_type  string `json:"sensor_type"`
	Model        string `json:"model"`
	Prototype_id int    `json:"Propotype_id"`
}