package adapters

import (
	"PyBot-DataServer/database/conn"
	"PyBot-DataServer/work_periods/domain/models"
	"fmt"
	"time"
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

func (postgre *PostgreSQL) GetLastPeriod() (models.LastPeriod, error) {
	query := `SELECT *
			  FROM getLastHourPeriod()`
	
	var lastPeriod models.LastPeriod
	
	rows, err := postgre.conn.FetchRows(query)
	if err != nil {
		fmt.Printf("error al ejecutar la consulta: %v", err)
		return models.LastPeriod{}, err
	}

	defer rows.Close()

	if !rows.Next() {
        fmt.Println("No se pudieron obtener los datos.")
        return models.LastPeriod{}, nil
    }

	if err := rows.Scan(&lastPeriod.Period_id, &lastPeriod.LastHour); err != nil {
		fmt.Printf("error al escanear el ultimo periodo: %v", err)
        return models.LastPeriod{}, err
    }

	
	return lastPeriod, nil
}

func (postgre *PostgreSQL) CreatePeriod(wp models.WorkPeriod) (int, error) {
	query := `INSERT INTO work_periods (start_hour, end_hour, day_work, prototype_id)
	          VALUES ($1, $2, $3, $4)
			  RETURNING period_id`
			  
	const layout = "2006-01-02T15:04:05.999999Z07:00"
	startT, err := time.Parse(layout, wp.Start_hour)
	if err != nil {
		fmt.Printf("Error al ejecutar parsear la hora de inicio: %v", err)
		return 0, err
	}
	
	var id int
	
	err = postgre.conn.DB.QueryRow(query, startT, startT, wp.Day_work, wp.Prototype_id).Scan(&id)	
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
	
func (postgre *PostgreSQL) GetDistanceAndWeight(period_id int64) (models.Reading, error) {
	// Llamadas a las funciones especiales de PostgreSQL
	query := `SELECT *
			  FROM getLastWeight()`
	
	var readings models.Reading
	
	rows, err := postgre.conn.FetchRows(query)
	if err != nil {
		fmt.Printf("error al ejecutar la consulta: %v", err)
		return models.Reading{}, err
	}

	defer rows.Close()

	if !rows.Next() {
        fmt.Println("No se pudieron obtener los datos.")
        return models.Reading{}, nil
    }

	if err := rows.Scan(&readings.Weight_waste); err != nil {
		fmt.Printf("error al escanear el weight: %v", err)
        return models.Reading{}, err
    }

	query = `SELECT *
			  FROM calcular_distancia_total($1)`

	rows, err = postgre.conn.FetchRows(query, period_id)
	if err != nil {
		fmt.Printf("error al ejecutar la consulta: %v", err)
		return models.Reading{}, err
	}

	defer rows.Close()

	if !rows.Next() {
        fmt.Println("No se pudieron obtener los datos.")
        return models.Reading{}, nil
    }

	if err := rows.Scan(&readings.Distance_traveled); err != nil {
		fmt.Printf("error al escanear el distance: %v", err)
        return models.Reading{}, err
    }

	readings.Period_id = int(period_id)
	
	return readings, nil
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
	
	 
	