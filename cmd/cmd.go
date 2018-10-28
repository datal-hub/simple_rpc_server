package cmd

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/urfave/cli"

	"rpc-server/handlers"
	"rpc-server/pkg/settings"
)

var Srv = cli.Command{
	Name:        "srv",
	Usage:       "Start rpc server",
	Description: ``,
	Action:      runSrv,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name: "listen, l", Value: "127.0.0.1:8181",
			Usage: "Start server at specified address, port",
		},
	},
}

func updateSettings(c *cli.Context) {
	if c.IsSet("listen") {
		settings.ListenAddr = c.String("listen")
	}
}

func runSrv(c *cli.Context) error {
	updateSettings(c)

	srv := rpc.NewServer()

	srv.RegisterCodec(json.NewCodec(), "application/json")
	srv.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	api := new(handlers.UserApi)
	srv.RegisterService(api, "")

	r := mux.NewRouter()
	r.Handle("/api", srv)

	err := http.ListenAndServe(settings.ListenAddr, r)
	log.Fatal(err)

	return nil
}
