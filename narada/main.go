package main

import (
	"os"

	"github.com/m1ome/narada"
	"github.com/m1ome/narada/narada/commands"

	"github.com/sirupsen/logrus"

	"github.com/urfave/cli"
)

const (
	ConsoleToolName    = "narada"
	ConsoleToolVersion = "0.1"
)

func main() {
	// Creating logger system
	logger := logrus.New()

	// Creating instance of Narada
	n := narada.New(ConsoleToolName, ConsoleToolVersion)

	// Creating urfave
	app := cli.NewApp()
	app.Name = ConsoleToolName
	app.Version = ConsoleToolVersion
	app.Description = "Narada CLI toolchain"
	app.Author = "Pavel Makarenko <cryfall@gmail.com>"
	app.Commands = []cli.Command{
		commands.Migrate(n),
		commands.CreateMigration(n),
	}

	if err := app.Run(os.Args); err != nil {
		logger.Fatalf("error starting: %v", err)
	}
}
