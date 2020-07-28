package main

import (
	"flag"
	"log"
	"net/http"

	"zero/core/conf"
	"zero/core/httpx"
	"zero/core/logx"
	"zero/core/service"
	"zero/example/tracing/remote/portal"
	"zero/ngin"
	"zero/rpcx"
)

var (
	configFile = flag.String("f", "config.json", "the config file")
	client     *rpcx.RpcClient
)

func handle(w http.ResponseWriter, r *http.Request) {
	conn, ok := client.Next()
	if !ok {
		log.Fatal("no server")
	}

	greet := portal.NewPortalClient(conn)
	resp, err := greet.Portal(r.Context(), &portal.PortalRequest{
		Name: "kevin",
	})
	if err != nil {
		httpx.WriteJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	} else {
		httpx.OkJson(w, resp.Response)
	}
}

func main() {
	flag.Parse()

	var c rpcx.RpcClientConf
	conf.MustLoad(*configFile, &c)
	client = rpcx.MustNewClient(c)
	engine := ngin.MustNewEngine(ngin.NgConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Port: 3333,
	})
	defer engine.Stop()

	engine.AddRoute(ngin.Route{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: handle,
	})
	engine.Start()
}