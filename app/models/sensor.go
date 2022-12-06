package models

import (
	"database/sql"
	"sensors/app/helpers"
)

type Sensor struct {
	Id          int    `json:"id"`
	Tag         string `json:"tag"`
	Enabled     bool   `json:"enabled"`
	Address     string `json:"address"`
	Refreshrate int    `json:"refreshrate"`
	Formula     string `json:"formula"`
	Symbol      string `json:"symbol"`
	MaxValue    int    `json:"max_value"`
	Color       string `json:"color"`
}

type Response struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
}

func GetSensors() ([]Sensor, error) {
	var sensor Sensor
	db, err := sql.Open("sqlite3", helpers.Currentdir()+"/database")
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM sensors")
	if err != nil {
		return nil, err
	}
	sensors := []Sensor{}
	for rows.Next() {
		err = rows.Scan(&sensor.Id, &sensor.Tag, &sensor.Enabled, &sensor.Address, &sensor.Refreshrate, &sensor.Formula, &sensor.Symbol, &sensor.MaxValue, &sensor.Color)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}
	return sensors, nil
}

func GetSensorById(id string) (Sensor, error) {
	var sensor Sensor
	db, err := sql.Open("sqlite3", helpers.Currentdir()+"/database")
	if err != nil {
		return Sensor{}, err
	}
	err = db.QueryRow("SELECT * FROM sensors WHERE id=?", id).Scan(&sensor.Id, &sensor.Tag, &sensor.Enabled, &sensor.Address, &sensor.Refreshrate, &sensor.Formula, &sensor.Symbol, &sensor.MaxValue, &sensor.Color)
	if err != nil {
		return Sensor{}, err
	}
	return sensor, nil
}

func AddSensor(sensor Sensor) (int, error) {
	db, err := sql.Open("sqlite3", helpers.Currentdir()+"/database")
	if err != nil {
		return -1, err
	}
	stmt, err := db.Prepare("INSERT INTO sensors (tag, enabled, address, refreshrate, formula, symbol, max_value, color) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(sensor.Tag, sensor.Enabled, sensor.Address, sensor.Refreshrate, sensor.Formula, sensor.Symbol, sensor.MaxValue, sensor.Color)
	if err != nil {
		return -1, err
	}
	inserted, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(inserted), nil
}

func UpdateSensor(sensor Sensor) error {
	db, err := sql.Open("sqlite3", helpers.Currentdir()+"/database")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("UPDATE sensors SET tag=?, enabled=?, address=?, refreshrate=?, formula=?, symbol=?, max_value=?, color=? WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sensor.Tag, sensor.Enabled, sensor.Address, sensor.Refreshrate, sensor.Formula, sensor.Symbol, sensor.MaxValue, sensor.Color, sensor.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSensor(id string) error {
	db, err := sql.Open("sqlite3", helpers.Currentdir()+"/database")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("DELETE FROM sensors WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
