package adapters

import (
	b "PyBot-DataServer/backup/domian/models"
	"PyBot-DataServer/database/conn"
	s "PyBot-DataServer/sensors/domain/models"
	wP "PyBot-DataServer/work_periods/domain/models"
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

func (postgres *PostgreSQL) GetData() ([]b.DataTable, error) {
	var backup []b.DataTable
	
	var work_periods b.DataTable

	work_periods_rows, err := getWorkPeriods(postgres)
	if err != nil {
		fmt.Printf("Error al obtener los datos de la tabla work_periods: %v", err)
	}

	work_periods.Table_name = "work_periods"
	work_periods.Data = work_periods_rows
	backup = append(backup, work_periods)

	var readings b.DataTable

	readings_rows, err := getReadings(postgres)
	if err != nil {
		fmt.Printf("Error al obtener los datos de la tabla readings: %v", err)
	}

	readings.Table_name = "readings"
	readings.Data = readings_rows
	backup = append(backup, readings)

	var waste_collection b.DataTable

	waste_collection_rows, err := getWasteCollection(postgres)
	if err != nil {
		fmt.Printf("Error al obtener los datos de la tabla waste_collection: %v", err)
	}

	waste_collection.Table_name = "waste_collection"
	waste_collection.Data = waste_collection_rows
	backup = append(backup, waste_collection)

	var weight_data b.DataTable

	weight_data_rows, err := getWeightData(postgres)
	if err != nil {
		fmt.Printf("Error al obtener los datos de la tabla weight_data: %v", err)
	}

	weight_data.Table_name = "weight_data"
	weight_data.Data = weight_data_rows
	backup = append(backup, weight_data)

    var gps_data b.DataTable

	gps_data_rows, err := getGPSData(postgres)
	if err != nil {
		fmt.Printf("Error al obtener los datos de la tabla gps_data: %v", err)
	}

	gps_data.Table_name = "gps_data"
	gps_data.Data = gps_data_rows
	backup = append(backup, gps_data)

	return backup, nil

}

func getWorkPeriods(postgres *PostgreSQL) ([]wP.WorkPeriod, error) {
	query := "SELECT * FROM work_periods WHERE backup = FALSE"

	var work_periods []wP.WorkPeriod

	rows, err := postgres.conn.DB.Query(query)
	if err != nil {
		fmt.Printf("Error al ejecutar getWorkPeriods: %v", err)
		return []wP.WorkPeriod{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var wp wP.WorkPeriod

		if err := rows.Scan(&wp.Period_id, &wp.Start_hour, &wp.End_hour, &wp.Day_work, &wp.Backup); err != nil {
			fmt.Printf("Error al escanear la fila: %v", err)
			return []wP.WorkPeriod{}, nil
		}

		work_periods = append(work_periods, wp)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error al reccorrer las filas: %v", err)
		return []wP.WorkPeriod{}, err
	}
	
	return work_periods, nil
}

func getReadings(postgres *PostgreSQL) ([]wP.Reading, error) {
	query := "SELECT * FROM readings WHERE backup = FALSE"

	var readings []wP.Reading

	rows, err := postgres.conn.DB.Query(query)
	if err != nil {
		fmt.Printf("Error al ejecutar getReading: %v", err)
		return []wP.Reading{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var r wP.Reading

		if err := rows.Scan(&r.Period_id, &r.Distance_traveled, &r.Weight_waste); err != nil {
			fmt.Printf("Error al escanear la fila: %v", err)
			return []wP.Reading{}, nil
		}

		readings = append(readings, r)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error al reccorrer las filas: %v", err)
		return []wP.Reading{}, err
	}
	
	return readings, nil
}

func getWasteCollection(postgres *PostgreSQL) ([]s.WasteCollection, error) {
	query := "SELECT * FROM waste_collection WHERE backup = FALSE"

	var waste_collection []s.WasteCollection

	rows, err := postgres.conn.DB.Query(query)
	if err != nil {
		fmt.Printf("Error al ejecutar getWasteCollection: %v", err)
		return []s.WasteCollection{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var c s.WasteCollection

		if err := rows.Scan(&c.Waste_collection_id, &c.Period_id, &c.Amount, &c.Waste_id); err != nil {
			fmt.Printf("Error al escanear la fila: %v", err)
			return []s.WasteCollection{}, nil
		}

		waste_collection = append(waste_collection, c)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error al reccorrer las filas: %v", err)
		return []s.WasteCollection{}, err
	}
	
	return waste_collection, nil
}

func getWeightData(postgres *PostgreSQL) ([]s.WeightData, error) {
	query := "SELECT * FROM weight_data WHERE backup = FALSE"

	var weight_data []s.WeightData

	rows, err := postgres.conn.DB.Query(query)
	if err != nil {
		fmt.Printf("Error al ejecutar getWeightData: %v", err)
		return []s.WeightData{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var wD s.WeightData

		if err := rows.Scan(&wD.Weight_data_id, &wD.Period_id, &wD.Hour_period, &wD.Weight); err != nil {
			fmt.Printf("Error al escanear la fila: %v", err)
			return []s.WeightData{}, nil
		}

		weight_data = append(weight_data, wD)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error al reccorrer las filas: %v", err)
		return []s.WeightData{}, err
	}
	
	return weight_data, nil
}

func getGPSData(postgres *PostgreSQL) ([]s.GPSData, error) {
	query := "SELECT * FROM gps_data WHERE backup = FALSE"

	var gps_data []s.GPSData

	rows, err := postgres.conn.DB.Query(query)
	if err != nil {
		fmt.Printf("Error al ejecutar getGPSData: %v", err)
		return []s.GPSData{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var g s.GPSData

		if err := rows.Scan(&g.Gps_data_id, &g.Period_id, &g.Latitude, &g.Longitude, &g.Altitude, &g.Speed, &g.Date_gps, &g.Hour_UTC); err != nil {
			fmt.Printf("Error al escanear la fila: %v", err)
			return []s.GPSData{}, nil
		}

		gps_data = append(gps_data, g)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf("Error al reccorrer las filas: %v", err)
		return []s.GPSData{}, err
	}
	
	return gps_data, nil
}

func (postgres *PostgreSQL) UpdateIdsBackupDone() error {
	query := `UPDATE work_periods 
			  SET backup = TRUE 
			  WHERE backup = FALSE`
	
	_, err := postgres.conn.ExecutePreparedQuery(query)
	if err != nil {
		fmt.Printf("Error al ejecutar UpdateIdsBackupDone: %v", err)
		return err
	}
	
	return nil
}
