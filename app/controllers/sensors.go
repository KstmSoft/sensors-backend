package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sensors/app/models"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func GetSensors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sensors, err := models.GetSensors()
		if httpError(err, w) {
			return
		}
		json, err := json.Marshal(sensors)
		if httpError(err, w) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func GetSensorById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		sensor, err := models.GetSensorById(id)
		if httpError(err, w) {
			return
		}
		json, err := json.Marshal(sensor)
		if httpError(err, w) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func AddSensor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sensor models.Sensor
		err := json.NewDecoder(r.Body).Decode(&sensor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		inserted, err := models.AddSensor(sensor)
		if httpError(err, w) {
			return
		}
		json, err := json.Marshal(models.Response{Success: true, Id: fmt.Sprint(inserted)})
		if httpError(err, w) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func DeleteSensor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		err := models.DeleteSensor(id)
		if httpError(err, w) {
			return
		}
		response := models.Response{Success: true, Id: id}
		json, err := json.Marshal(response)
		if httpError(err, w) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func UpdateSensor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sensor models.Sensor
		err := json.NewDecoder(r.Body).Decode(&sensor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if sensor.Id <= 0 {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		err = models.UpdateSensor(sensor)
		if httpError(err, w) {
			return
		}
		response := models.Response{Success: true, Id: fmt.Sprint(sensor.Id)}
		json, err := json.Marshal(response)
		if httpError(err, w) {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}

func httpError(err error, w http.ResponseWriter) bool {
	if err != nil {
		log.Println(err.Error())
		json, _ := json.Marshal(models.Response{Success: false, Id: "-1"})
		http.Error(w, string(json), http.StatusInternalServerError)
	}
	return err != nil
}
