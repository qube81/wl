package main

import (
	"github.com/codegangsta/cli"
	"github.com/qube81/wunderlist-api-go"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	newApp().Run(os.Args)
}

func newApp() *cli.App {

	app := cli.NewApp()
	app.Name = "wl"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{{Name: "qube81"}}
	app.Usage = "Wunderlist CLI client"
	app.Commands = commands

	return app
}

func newWLClient() *wunderlist.Client {
	clientID := os.Getenv("WL_CLIENT_ID")
	clientSecret := os.Getenv("WL_ACCESS_TOKEN")
	return wunderlist.NewClient(clientID, clientSecret)
}

var commands = []cli.Command{
	taskCommand,
}
