package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/priscila-albertini-silva/jaded-backend/config"
	"github.com/priscila-albertini-silva/jaded-backend/pkg/gormfx"
	"github.com/priscila-albertini-silva/jaded-backend/pkg/serverfx"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

func executeRun(cmd *cobra.Command, args []string) {
	log.Info("Startup")

	config.InitConfig()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go cleanup(sigs)

	fx.New(
		serverfx.ModuleServer,
		gormfx.Module,
	).Run()
}

func cleanup(c chan os.Signal) {
	sig := <-c
	log.Info(fmt.Sprintf("Signal %s received, proceeding with graceful stop.", sig))
	os.Exit(0)
}
