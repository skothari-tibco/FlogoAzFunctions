package shim

import (
	"os"
	"fmt"
	_ "github.com/project-flogo/core/data/expression/script"
	"github.com/project-flogo/core/engine"
	"github.com/project-flogo/core/support/log"
)

var (
	cfgJson       string
	cfgCompressed bool
)

func init() {
	log.SetLogLevel(log.RootLogger(), log.ErrorLevel)
	fmt.Println("Loading Engine")
	cfg, err := engine.LoadAppConfig(cfgJson, cfgCompressed)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println("Starting Engine")
	_, err = engine.New(cfg, engine.DirectRunner)
	if err != nil {
		log.RootLogger().Errorf("Failed to create engine: %s", err.Error())
		os.Exit(1)
	}
}
