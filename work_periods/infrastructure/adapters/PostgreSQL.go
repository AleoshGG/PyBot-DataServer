package adapters

import (
	"PyBot-DataServer/database/conn"
	"PyBot-DataServer/work_periods/domain/models"
	"fmt"
)

type PostgreSQL struct {
	conn *conn.ConnPostgreSQL
}

func NewPostgreSQL() *PostgreSQL {
	conn := conn.GetDBPool()

	if conn.Err != "" {
		fmt.Printf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &PostgreSQL{conn: conn}
}

func (postgre *PostgreSQL) CreatePeriod(wp models.WorkPeriod) (int, error) {
	query := `INSERT INTO work_periods (start_hour, end_hour, day_work, prototype_id)
	          VALUES ($1, $2, $3, $4)
			  RETURNING period_id`
	
	var id int
	
	err := postgre.conn.DB.QueryRow(query, wp.Start_hour, wp.End_hour, wp.Day_work, wp.Prototype_id).Scan(&id)	
	if err != nil {
		fmt.Printf("Error al ejecutar CreatePeriod: %v", err)
		return 0, err
	}
	
	return id, nil
}

func (postgre *PostgreSQL) UpdatePeriod(end_hour string, period_id int64) (error) {
	query := `UPDATE work_periods 
			  SET end_hour = $1 
			  WHERE period_id = $2`
	
	_, err := postgre.conn.ExecutePreparedQuery(query, end_hour, period_id)
	if err != nil {
		fmt.Printf("Error al ejecutar UpdatePeriod: %v", err)
		return err
	}
	
	return nil
}
	
func (postgre *PostgreSQL) GetDistanceAndWeight() (models.Reading, error) {
	// Llamadas a las funciones especiales de PostgreSQL

	return models.Reading{}, nil
}

func (postgre *PostgreSQL) ReadingsRegister(r models.Reading) (error) {
	query := `INSERT INTO readings (period_id, distance_traveled, weight_waste)
			  VALUES ($1, $2, $3)
			  RETURNING period_id`

	var id int

	err := postgre.conn.DB.QueryRow(query, r.Period_id, r.Distance_traveled, r.Weight_waste).Scan(&id)
	if err != nil {
		fmt.Printf("Error al ejecutar ReadingsRegister: %v", err)
		return err
	}

	return nil
}
	
	 
	