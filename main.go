package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web/config"
	"web/jobs"
	"web/logger"
	"web/repository/cache"
	"web/repository/pg"
	"web/utils"
	"web/web/router"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func main() {
	// flag
	path := flag.String("config", "./conf.json", "config path")

	flag.Parse()

	logger.InitLogger()

	// init validator config
	config.InitConfig(*path)

	// main context
	mainCtx, cancel := context.WithCancel(context.TODO())

	// init db
	pg.InitPg(config.Configure)

	// init cache
	cache.InitMMCache()

	// init utils
	utils.InitUtils()

	// run web server
	go func() {
		err := runHttpServer(mainCtx)
		if err != nil {
			logger.Error("http server run got err", zap.Error(err))
		}
		panic(err)
	}()

	// run job
	go jobs.RunJob(mainCtx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	sg := <-quit
	logger.Infof("Receive signal %v and shutdown...", sg)

	cancel()
	
	logger.Infof("delay cancel in %+v ", config.Configure.ServerSetting.ShutDownTimeout)
	time.Sleep(config.Configure.ServerSetting.ShutDownTimeout)
}

func runHttpServer(ctx context.Context) error {
	routersInit := router.InitRouter()
	readTimeout := config.Configure.ServerSetting.ReadTimeout * time.Second
	writeTimeout := config.Configure.ServerSetting.WriteTimeout * time.Second
	endPoint := fmt.Sprintf(":%d", config.Configure.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	svr := &http.Server{
		Addr:           endPoint,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		Handler:        routersInit,
		MaxHeaderBytes: maxHeaderBytes,
	}

	stop := make(chan error)
	go func() {
		logger.Infof("Start http server listening %s", endPoint)
		if err := svr.ListenAndServe(); err == nil || err == http.ErrServerClosed {
			stop <- nil
		} else {
			stop <- errors.Wrap(err, "server serve err")
		}
	}()

	select {
	case <-ctx.Done():
		logger.Warnf("Canceled, stop server %s", endPoint)
		c, cancel := context.WithTimeout(context.TODO(), config.Configure.ServerSetting.ShutDownTimeout)
		defer cancel()
		return errors.Wrap(svr.Shutdown(c), "server shutdown err")
	case err := <-stop:
		return err
	}
}
