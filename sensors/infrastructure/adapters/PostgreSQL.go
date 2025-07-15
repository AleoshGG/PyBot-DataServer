package adapters

import (
	"PyBot-DataServer/database/conn"
	"PyBot-DataServer/sensors/domain/models"
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

func (postgres *PostgreSQL)	WasteCollectionRegister(wc models.WasteCollection) (int, error) {
	query := `INSERT INTO waste_collection (period_id, amount, waste_id)
			  VALUES ($1, DEFAULT, $2)
			  RETURNING waste_collection_id`

	var id int		  
	
	err := postgres.conn.DB.QueryRow(query, wc.Period_id, wc.Waste_id).Scan(&id)
	if err != nil {
		fmt.Printf("Error al ejecutar WasteCollectionRegister: %v", err)
		return 0, err
	}

	return id, nil
}

func (postgres *PostgreSQL) UpdateWasteCollection(id int64) (error) {
	query := `UPDATE waste_collection 
			  SET amount = amount + 1
			  WHERE waste_collection_id = $1`

	_, err := postgres.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		fmt.Printf("Error al ejecutar UpdateWasteCollection: %v", err)
		return err
	}
	
	return nil
}

func (postgres *PostgreSQL)	WeightRegister(w models.WeightData) (int, error) {
	query := `INSERT INTO weight_data (period_id, hour_period, weight)
			  VALUES ($1, $2, $3)
			  RETURNING weight_data_id` 

	var id int
	
	const layout = "2006-01-02T15:04:05.999999Z07:00"
	startT, err := time.Parse(layout, w.Hour_period)
	if err != nil {
		fmt.Printf("Error al ejecutar parsear la hora de periodo: %v", err)
		return 0, err
	}

	err = postgres.conn.DB.QueryRow(query, w.Period_id, startT, w.Weight).Scan(&id)
	if err != nil {
		fmt.Printf("Error al ejecutarr la consulta WeightRegister: %v", err)
		return 0, err
	}
	
	return id, nil
}

func (postgres *PostgreSQL)	GPSRegister(gps models.GPSData) (int, error) {
	query := `INSERT INTO gps_data (period_id, latitude, longitude, altitude, speed, date_gps, hour_UTC)
			  VALUES ($1, $2, $3, $4, $5, $6, $7)
			  RETURNING gps_data_id`

	var id int

	const layout = "2006-01-02T15:04:05.999999Z07:00"
	startT, err := time.Parse(layout, gps.Hour_UTC)
	if err != nil {
		fmt.Printf("Error al ejecutar parsear la hora UTC: %v", err)
		return 0, err
	}

	err = postgres.conn.DB.QueryRow(
		query, 
		gps.Period_id,
		gps.Latitude,
		gps.Longitude,
		gps.Altitude,
		gps.Speed,
		gps.Date_gps,
		startT,
	).Scan(&id)	

	if err != nil {
		fmt.Printf("Error al ejecutarr GPSRegister: %v", err)
		return 0, err
	}

	return id, nil
}

func (postgres *PostgreSQL)	GetWasteTypes() ([]models.WasteType, error) {
	query := `SELECT * FROM waste_types`

	var WasteTypes []models.WasteType

	rows, err := postgres.conn.DB.Query(query)
	if err != nil {
		fmt.Printf("Error al ejecutar GetWasteTypes: %v", err)
		return []models.WasteType{}, err
	}
	
	defer rows.Close()

	for rows.Next() {
		var wT models.WasteType

		if err := rows.Scan(&wT.Waste_id, &wT.Type); err != nil {
			fmt.Printf("Error al escanear la fila: %v", err)
			return []models.WasteType{}, err
		}

		WasteTypes = append(WasteTypes, wT)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error al reccorrer las filas: %v", err)
		return []models.WasteType{}, err
	}
	
	return WasteTypes, nil
}