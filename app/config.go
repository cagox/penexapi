package app

import (
	"database/sql"
	"log"
	"os"

	"github.com/cagox/config"
	"github.com/gorilla/mux"
)

var Config *ConfigStruct

type ConfigStruct struct {
	config.ConfigurationStruct //Most of the obvious stuff is already in here.

	SiteName string
	LogPath  string

	Port string

	MaximumNameLength int

	//The following items are no in the JSON file
	Router  *mux.Router
	Logger  *log.Logger
	LogFile *os.File

	//Database Related
	Database *sql.DB
}

func init() {
	Config = &ConfigStruct{}
	loadConfigs()
	Config.Router = mux.NewRouter()
}

func loadConfigs() {
	config.LoadConfigs(Config, "FICSITECONF")
}
