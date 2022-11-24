package services

import (
	"fmt"
	"log"
	"sensors/app/models"
	"strconv"
	"time"

	"github.com/ambelovsky/gosf"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/spi"
	"gobot.io/x/gobot/platforms/raspi"
)

func ListenSensors() {
	master := gobot.NewMaster()
	sensors, err := models.GetSensors()

	if err != nil {
		log.Println(err)
	}
	for _, sensor := range sensors {
		master.AddRobot(ReaderBot(sensor))
	}
	master.Start()
}

func ReaderBot(sensor models.Sensor) *gobot.Robot {
	adaptor := raspi.NewAdaptor()
	driver := spi.NewMCP3208Driver(adaptor)
	address, err := strconv.Atoi(sensor.Address[1:2])

	if err != nil {
		log.Println("An error ocurred reading on ", sensor.Address, " worker")
		return nil
	}

	work := func() {
		gobot.Every(time.Second, func() {
			result, _ := driver.Read(address)
			message := new(gosf.Message)
			message.Success = true
			message.Text = fmt.Sprint(result)
			gosf.Broadcast("", sensor.Address, message)
		})
	}
	robot := gobot.NewRobot(sensor.Address,
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
		work,
	)

	return robot
}
