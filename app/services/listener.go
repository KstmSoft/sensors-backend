package services

import (
	"fmt"
	"log"
	"sensors/app/helpers"
	"sensors/app/models"
	"strconv"
	"time"

	"github.com/ambelovsky/gosf"
	"github.com/Pramod-Devireddy/go-exprtk"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/spi"
	"gobot.io/x/gobot/platforms/raspi"
)

func ListenSensors() {
	master := gobot.NewMaster()
	sensors, err := models.GetSensors()

	// Create a new exprtk instance
	exprtkObj := exprtk.NewExprtk()
	defer exprtkObj.Delete()

	// Set the expression
	exprtkObj.SetExpression("(x + 2)*(y-2)")

	// Add variables of expression
	exprtkObj.AddDoubleVariable("x")
	exprtkObj.AddDoubleVariable("y")

	// Compile the expression after expression and variables declaration
	err := exprtkObj.CompileExpression()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Set values for the variables
	exprtkObj.SetDoubleVariableValue("x", 18)
	exprtkObj.SetDoubleVariableValue("y", 32)

	// Get the evaluated value
	fmt.Println(exprtkObj.GetEvaluatedValue())

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
	driver := spi.NewMCP3008Driver(adaptor)
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
			message.Text = fmt.Sprint(helpers.ConvertToVoltage(result))
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
