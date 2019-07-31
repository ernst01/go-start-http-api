package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ernst01/go-start-http-api/internal/server"
	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

// Server parameters
type srvParameters struct {
	port int // port for listening
}

func main() {
	var prm srvParameters

	flag.IntVar(&prm.port, "port", 8080, "Http server port.")
	flag.Parse()

	apiSrv := server.Server{
		Router: mux.NewRouter(),
	}
	apiSrv.Routes()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%v", prm.port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      apiSrv.Router,
	}

	ctx, cancel := context.WithCancel(context.Background())

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	defer func() {
		glog.Info("Running defer function.")
		signal.Stop(stopChan)
		cancel()
	}()

	go func() {
		select {
		case <-stopChan:
			glog.Info("Running Interrupt cleanup.")
			srv.Shutdown(ctx)
			signal.Stop(stopChan)
			cancel()
		}
	}()

	glog.Infof("Starting to listen on port: %v.", prm.port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		glog.Errorf("HTTP Server ListenAndServe() error: %v", err)
		os.Exit(1)
	}

	glog.Info("Shutting down.")
	os.Exit(0)
}
