package main

import (
	"alpha/config"
	v "alpha/pkg/version"
	redis "alpha/repositories/data-mappers/go-redis"
	"alpha/repositories/data-mappers/model"
	"alpha/router"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "alpha config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

func main() {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	pflag.Parse()
	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		config.Logger.Info("endpoint",
			zap.String("httpMethod", httpMethod),
			zap.String("absolutePath", absolutePath),
			zap.String("handlerName", handlerName),
			zap.Int("nuHandlers", nuHandlers),
		)
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	//init redis
	redis.Client.Init()
	defer redis.Client.Close()
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(
		g,
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			config.Logger.Fatal("The router has no response, or it might took too long to start up.",
				zap.Error(err),
			)
		}
		config.Logger.Info("The router has been deployed successfully.")
	}()

	// Start to listening the incoming requests.
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			config.Logger.Info("Start to listening the incoming requests on https",
				zap.String("address", viper.GetString("tls.addr")),
			)
			config.Logger.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
	config.Logger.Info("Start to listening the incoming requests on http",
		zap.String("address", viper.GetString("addr")),
	)
	srv := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: g,
	}
	go func(log *zap.Logger) {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen: ",
				zap.Error(err),
			)
		}
	}(config.Logger)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	config.Logger.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		config.Logger.Fatal("Server Shutdown: ",
			zap.Error(err),
		)
	}
	config.Logger.Info("Server exiting")
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		config.Logger.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
