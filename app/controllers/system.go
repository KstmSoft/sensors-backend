package controllers

import (
	"log"
	"net/http"
	"sensors/app/helpers"
)

func PowerOff() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("PowerOff System")
		helpers.RunCommand("sudo poweroff")
	}
}
func Restart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Restarting System")
		helpers.RunCommand("sudo reboot")
	}
}
