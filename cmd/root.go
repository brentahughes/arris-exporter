package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/brentahughes/arris-exporter/pkg/arris"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "arris-exporter",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			bindAddress := fmt.Sprintf(":%d", viper.GetInt("prometheus.port"))
			log.Printf("listening on %s", bindAddress)
			http.Handle("/metrics", promhttp.Handler())
			if err := http.ListenAndServe(bindAddress, nil); err != nil {
				log.Fatal(err)
			}
		}()

		scrapper := arris.NewScrapper(viper.GetString("host"))

		if err := scrapper.Parse(context.TODO()); err != nil {
			log.Fatal(err)
		}

		for {
			select {
			case <-time.After(viper.GetDuration("interval")):
				if err := scrapper.Parse(context.TODO()); err != nil {
					log.Println(err)
				}
			}
		}

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().String("host", "192.168.100.1", "IP of arris modem")
	rootCmd.Flags().DurationP("interval", "i", 10*time.Second, "How often to scrape the interface")
	rootCmd.Flags().Int("prometheus.port", 9300, "Port to expose prometheus metrics on")
	viper.BindPFlags(rootCmd.Flags())
}
