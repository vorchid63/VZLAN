package main

import (
	"os"

	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/docker/go-plugins-helpers/network"
	"github.com/vorchid63/vzlan/vzdriver"
)

const (
	version = "0.4.0"
)

func main() {
	fmt.Printf("Starting\n")

	var flagDebug = cli.BoolFlag{
		Name:  "debug, d",
		Usage: "enable debugging",
	}
        var sysGrp = cli.StringFlag {
                Name:  "sysGroup, g",
                Usage: "specifies system group",
		EnvVar: "SYSGROUP",
        }
	app := cli.NewApp()
	app.Name = "vzlan"
	app.Usage = "Docker Vzlan Networking"
	app.Version = version
	app.Flags = []cli.Flag{
		flagDebug,
                sysGrp,
	}
	app.Action = Run
	app.Run(os.Args)
}

// Run initializes the driver
func Run(ctx *cli.Context) {
	fmt.Printf("Starting1\n")
	fmt.Println("TOR_IP\n", os.Getenv("TOR_IP"))
	if ctx.Bool("debug") {
		log.SetLevel(log.DebugLevel)
		fmt.Printf("Starting2\n")
	}

	fmt.Printf("calling NewDriver\n")
	d, err := vzdriver.NewDriver(version, ctx)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	h := network.NewHandler(d)
	fmt.Printf("Starting3 gruop=%s\n", ctx.String("sysGroup"))
	h.ServeUnix(ctx.String("sysGroup"), "vzlan")
}
