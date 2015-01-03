// mountMgo access mongodb databases through fusefs
package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "mountMgo"
	app.Usage = "mount a mongodb database"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "verbose", Usage: "verbose logging"},
		cli.StringFlag{Name: "dbhost, d", Value: "localhost", Usage: "the mongodb host to connect to"},
	}

	app.Action = func(c *cli.Context) {

		if len(c.Args()) != 2 {
			log.Fatal("Usage: mountMgo <dbname> <mountpoint>")
		}

		if !c.Bool("verbose") {
			log.SetOutput(ioutil.Discard)
		}

		dbName := c.Args()[0]
		dbHost := c.String("dbhost")
		mountPoint := c.Args()[1]

		initDb(dbHost, dbName)
		mount(mountPoint)
	}

	app.Run(os.Args)
}
