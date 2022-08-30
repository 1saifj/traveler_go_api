package main

import (
	"github.com/kataras/iris/v12"
	"net/http"
	"taveler/config"
	"taveler/infrastructure/datastore"
	"taveler/registry"
)

func main() {
	app := newApp()
	con, err := config.GetAppConfig()
	_, err = datastore.SetupDB(&con.DB)
	if err != nil {
		panic(err)
		return
	}
	_ = registry.NewRegistry()

	listenForHosts(app, con.Hosts)

}
func newApp() *iris.Application {
	app := iris.Default()
	return app
}

func listenForHosts(app *iris.Application, hosts []config.Host) {
	for i := 0; i < len(hosts)-1; i++ {
		host := hosts[i]
		go listen(app, host)
	}

	host := hosts[len(hosts)-1]
	err := app.Listen(host.GetHost())
	if err != nil {
		app.Logger().Error(err)
		panic(err)
	}
}

func listen(app *iris.Application, host config.Host) {
	err := app.NewHost(&http.Server{Addr: host.GetHost()}).ListenAndServe()
	if err != nil {
		app.Logger().Error(err)
		panic(err)
	}
}
