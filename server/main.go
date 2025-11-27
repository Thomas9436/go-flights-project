package main

import (
	"aggregator/controller"
	"aggregator/repo"
	svc "aggregator/services"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetEnvPrefix("FL")
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "3001")
	viper.SetDefault("J_SERVER1_URL", "http://j-server1:4001")
	viper.SetDefault("J_SERVER2_URL", "http://j-server2:4002")
}

func main() {
	initConfig()

	mux := http.NewServeMux()

	// Instantiate repos
	j1 := repo.NewJServer1Repo()
	j2 := repo.NewJServer2Repo()

	flightSvc := svc.NewFlightService(j1, j2)
	h := controller.NewHandler(flightSvc)
	h.RegisterRoutes(mux)

	port := viper.GetString("PORT")
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("Server running on http://localhost:%s\n", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server error:", err)
	}
}
