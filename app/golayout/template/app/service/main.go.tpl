package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"{{.ModName}}"
	"{{.ModName}}/api"
	"{{.ModName}}/store"
	"{{.ModName}}/user"
)

var cfgFile string
var s {{.ProjName}}.Store

func main() {
	log.Info("Hello, {{.ProjName}} service is running!")

	flag.StringVar(&cfgFile, "c", "", "config file path")
	flag.Parse()

	initConfig()
	{{.ProjName}}.GlobalCfg.RunMod = viper.GetString("app.mod")
	if {{.ProjName}}.GlobalCfg.RunMod == "" {
		{{.ProjName}}.GlobalCfg.RunMod = "production"
	}
	setupLogger()

	initStore()	
	uSrv := user.NewInfoImpl(user.Cfg{
		Store: s,
	})

	apiCfg := api.Cfg{
		Port:        viper.GetString("server.port"),
		UserInfoSrv: uSrv,
	}
	api.Init(apiCfg)

	log.Fatal(api.Run())
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".{{.ProjName}}" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".{{.ProjName}}")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error when read config file: %s", err)
	} else {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}

func setupLogger() {
	if {{.ProjName}}.GlobalCfg.RunMod == "production" {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.TraceLevel)
	}
}

func initStore() {
	var err error
	scfg := store.StoreConfig{
		Use:            viper.GetString("store.use"),
	}
	log.Infof("Init store, use %s", scfg.Use)
	switch scfg.Use {
	case "postgres":
		scfg.DBCfg = store.SQLConfig{
			Host:           viper.GetString("store.postgres.host"),
			Port:           viper.GetInt("store.postgres.port"),
			DB:             viper.GetString("store.postgres.db"),
			User:           viper.GetString("store.postgres.user"),
			Pass:           viper.GetString("store.postgres.pass"),
			MaxOpenCnn:     viper.GetInt("store.postgres.maxopencnn"),
			MaxCnnLifeTime: time.Duration(viper.GetInt("store.postgres.maxcnnlifetime")) * time.Minute,
		}
	case "imdb":
	default:
		log.Fatalf("must set store in the config file")
	}

	s, err = store.New(scfg)
	if err != nil {
		log.WithError(err).Fatal("init Store failed")
	}
}