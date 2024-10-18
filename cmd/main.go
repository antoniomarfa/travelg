package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"travel/app/api"
	"travel/app/async"
	"travel/config"

	"github.com/hashicorp/go-multierror"
	"github.com/joho/godotenv"
)

// @title Go Hexagonal API
// @description Powered by scv-go-tools - https://github.com/sergicanet9/scv-go-tools

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {

	var opts struct {
		Version       string `long:"ver" description:"Version" required:"true"`
		Environment   string `long:"env" description:"Environment" choice:"local" choice:"dev" required:"true"`
		Port          int    `long:"port" description:"Running port" required:"true"`
		Database      string `long:"db" description:"The database adapter to use" choice:"mongo" choice:"postgres" required:"true"`
		DSN           string `long:"dsn" description:"DSN of the selected database" required:"true"`
		FlowApikey    string `long:"flowapikey" description:"APIKEY"  required:"true"`
		FlowSecretkey string `long:"flowsecretkey" description:"SECRETKEY" required:"true"`
		FlowApiurl    string `long:"flowapurl" description:"APIURL" required:"true"`
		FLoWBaseurl   string `long:"flowbaseurl" description:"BASEURL" required:"true"`
	}
	env := os.Getenv("ENVIRONMENT")

	if env == "local" {
		err := godotenv.Load(".env")

		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	opts.Version = os.Getenv("VERSION")
	opts.Environment = os.Getenv("ENVIRONMENT")
	opts.Port, _ = strconv.Atoi(os.Getenv("PORT"))
	opts.Database = os.Getenv("DATABASE")
	opts.DSN = os.Getenv("DSN")

	opts.FlowApikey = os.Getenv("FLOW_APIKEY")
	opts.FlowSecretkey = os.Getenv("FLOW_SECRETKEY")
	opts.FlowApiurl = os.Getenv("FLOW_APIURL")
	opts.FLoWBaseurl = os.Getenv("FLOW_BASEURL")

	cfg, err := config.ReadConfig(opts.Version, opts.Environment, opts.Port, opts.Database, opts.DSN, opts.FlowApikey, opts.FlowSecretkey, opts.FlowApiurl, opts.FLoWBaseurl) // , "config")
	if err != nil {
		log.Fatal(fmt.Errorf("no se puede analizar el archivo de configuración ENV %s: %w", opts.Environment, err))
	}

	var g multierror.Group
	ctx, cancel := context.WithCancel(context.Background())

	a := api.New(ctx, cfg)
	g.Go(a.Run(ctx, cancel))

	if cfg.Async.Run {
		async := async.New(cfg)
		g.Go(async.Run(ctx, cancel))
	}

	if err := g.Wait().ErrorOrNil(); err != nil {
		log.Fatal(err)
	}
}
