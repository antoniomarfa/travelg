package config

import (
	"fmt"
	"path"

	"travel/tools/api/utils"
)

type Async struct {
	Run      bool
	Interval utils.Duration
}

type Config struct {
	// set in flags
	Version     string
	Environment string
	Port        int
	Database    string
	DSN         string

	// set in json config files
	config
}

type config struct {
	PostgresMigrationsDir string
	JWTSecret             string
	Timeout               utils.Duration
	Async                 Async
	FlowApikey            string
	Flowsecretkey         string
	FlowApiurl            string
	FlowBaseurl           string
}

// ReadConfig from the project´s JSON config files.
// Default values are specified in the default configuration file, config/config.json
// and can be overrided with values specified in the environment configuration files, config/config.{env}.json.
func ReadConfig(version, env string, port int, database, dsn, flowApikey, flowsecretkey, flowApiurl, flowBaseurl, configPath string) (Config, error) {
	var c Config
	c.Version = version
	c.Environment = env
	c.Port = port
	c.Database = database
	c.DSN = dsn

	c.FlowApikey = flowApikey
	c.Flowsecretkey = flowsecretkey
	c.FlowApiurl = flowApiurl
	c.FlowBaseurl = flowBaseurl

	// Asignar manualmente los valores a los campos del struct 'config'

	var cfg config

	if err := utils.LoadJSON(path.Join(configPath, "config.json"), &cfg); err != nil {
		return c, fmt.Errorf("error parsing configuration, %s", err)
	}

	if err := utils.LoadJSON(path.Join(configPath, "config.local.json"), &cfg); err != nil {
		//		if err := utils.LoadJSON(path.Join(configPath, "config."+env+".json"), &cfg); err != nil {
		return c, fmt.Errorf("error parsing environment configuration, %s", err)
	}

	// Asignar la configuración cargada al campo config
	c.config = cfg

	return c, nil
}
