// mountMgo access mongodb databases through fusefs
package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/cryptix/go/logging"
)

var slog = logging.Logger("mountMgo")

func main() {
	app := cli.NewApp()
	app.Name = "mountMgo"
	app.Usage = "mount a mongodb database"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "dbhost, d", Value: "localhost", Usage: "the mongodb host to connect to"},
	}

	app.Action = func(c *cli.Context) {

		if len(c.Args()) != 2 {
			slog.Fatal("Usage: mountMgo <dbname> <mountpoint>")
		}

		dbName := c.Args()[0]
		dbHost := c.String("dbhost")
		mountPoint := c.Args()[1]

		initDb(dbHost, dbName)
		mount(mountPoint)
	}

	app.Run(os.Args)
}
