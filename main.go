package main

import (
	"log"
	"net/http"
	"sensors/app"
	"sensors/app/helpers"
	"sync"

	"github.com/ambelovsky/gosf"
	"github.com/judwhite/go-svc"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

type program struct {
	wg   sync.WaitGroup
	quit chan struct{}
}

func main() {
	prg := &program{}

	if err := svc.Run(prg); err != nil {
		log.Fatal(err)
	}
}

func (p *program) Init(env svc.Environment) error {
	return nil
}

func (p *program) Start() error {
	p.wg.Add(1)
	go func() {
		log.Println("Starting...")
		config()
		app := app.New()
		tcp := ":" + viper.GetString("port.tcp")
		socket := viper.GetInt("port.socket")

		http.HandleFunc("/", app.Router.ServeHTTP)

		log.Println("Running sensors backend on port", tcp)

		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
		})

		go gosf.Startup(map[string]interface{}{"port": socket})

		handler := c.Handler(app.Router)
		err := http.ListenAndServe(tcp, handler)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

func (p *program) Stop() error {
	log.Println("Quit signal received...")
	p.wg.Done()
	return nil
}

func config() {
	viper.AddConfigPath(helpers.Currentdir())
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
