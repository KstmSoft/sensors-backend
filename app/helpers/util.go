package helpers

import (
	"log"
	"math"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

func Currentdir() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return cwd
}

func RunCommand(command string) (p *os.Process, err error) {
	args := strings.Fields(command)
	if args[0], err = exec.LookPath(args[0]); err == nil {
		var procAttr os.ProcAttr
		procAttr.Files = []*os.File{os.Stdin,
			os.Stdout, os.Stderr}
		p, err := os.StartProcess(args[0], args, &procAttr)
		if err == nil {
			return p, nil
		}
	}
	return nil, err
}

func MaximumValueBits() float64 {
	bits := viper.GetInt("bits")
	return math.Pow(2, float64(bits)) - 1
}

func ConvertToVoltage(value int) float64 {
	maxVolt := viper.GetInt("maxVolt")
	voltage := (float64(maxVolt) / MaximumValueBits()) * float64(value)
	return float64(voltage)
}
