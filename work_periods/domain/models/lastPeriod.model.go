package models

type LastPeriod struct {
	Period_id int    `json:"period_id"`
	LastHour  string `json:"last_hour"`
}