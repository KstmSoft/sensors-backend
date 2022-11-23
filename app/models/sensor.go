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
		err = rows.Scan(&sensor.Id, &sensor.Tag, &sensor.Enabled, &sensor.Address, &sensor.Refreshrate)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}
	return sensors, nil
}

func AddSensor(sensor Sensor) (int, error) {
	db, err := sql.Open("sqlite3", helpers.Currentdir()+"/database")
	if err != nil {
		return -1, err
	}
	stmt, err := db.Prepare("INSERT INTO sensors (tag, enabled, address, refreshrate) values(?,?,?,?)")
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(&sensor.Tag, sensor.Enabled, sensor.Address, sensor.Refreshrate)
	if err != nil {
		return -1, err
	}
	inserted, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(inserted), nil
}

func UpdateSensor(id string, sensor Sensor) error {
	db, err := sql.Open("sqlite3", helpers.Currentdir()+"/database")
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("UPDATE sensors SET tag=?, enabled=?, address=?, refreshrate=? WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(&sensor.Tag, sensor.Enabled, sensor.Address, sensor.Refreshrate, id)
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
