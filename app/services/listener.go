package services

import (
	"fmt"
	"log"
	"os"
	"sensors/app/helpers"
	"sensors/app/models"
	"strconv"
	"time"

	"github.com/ambelovsky/gosf"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/spi"
	"gobot.io/x/gobot/platforms/raspi"
)

var master = gobot.NewMaster()

func ListenSensors() {
	sensors, err := models.GetSensors()
	if err != nil {
		log.Println(err)
	}
	for _, sensor := range sensors {
		master.AddRobot(ReaderBot(sensor.Id, sensor))
	}
	master.Start()
}

func AddSensorRuntime(id string, sensor models.Sensor) {
	master.AddRobot(ReaderBot(id, sensor)).Start()
}

func ReaderBot(id string, sensor models.Sensor) *gobot.Robot {
	adaptor := raspi.NewAdaptor()
	driver := spi.NewMCP3008Driver(adaptor)
	address, err := strconv.Atoi(sensor.Address[1:2])

	if err != nil {
		log.Println("An error ocurred reading on ", sensor.Address, " worker")
		return nil
	}

	f, err := os.OpenFile(sensor.Address+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	work := func() {
		gobot.Every(time.Second, func() {
			result, _ := driver.Read(address)
			updatedSensor, _ := models.GetSensorById(id)
			voltage := helpers.ConvertToVoltage(result)
			message := new(gosf.Message)
			message.Success = true
			message.Body = map[string]interface{}{"computed": fmt.Sprint(helpers.ComputeFormula(updatedSensor.Formula, voltage)), "voltage": fmt.Sprint(voltage)}
			f.WriteString(fmt.Sprint(message.Body) + "\n")
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
