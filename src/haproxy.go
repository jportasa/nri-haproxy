//go:generate goversioninfo -o ../resource.syso ../versioninfo.json
package main

import (
	"net/http"
	"os"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
)

const (
	integrationName    = "com.newrelic.haproxy"
	integrationVersion = "2.1.0"
)

var (
	args argumentList
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Username    string `default:"" help:"The HAProxy basic auth username."`
	Password    string `default:"" help:"The HAProxy basic auth password."`
	StatsURL    string `default:"" help:"The URL where HAProxy stats are available."`
	ClusterName string `default:"" help:"Cluster name to identify this HAProxy instance."`
}

func main() {
	haproxyIntegration, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	if err != nil {
		log.Error("Failed to create integration: %s", err)
		os.Exit(1)
	}

	log.SetupLogging(args.Verbose)

	if args.StatsURL == "" {
		log.Error("Must supply a URL pointing to the HAProxy stats page")
		os.Exit(1)
	}

	client := &http.Client{}

	// Create the http request
	req, err := createStatsRequest(args.Username, args.Password, addCSVtoURL(args.StatsURL))
	if err != nil {
		log.Error("Failed to create request: %s", err.Error())
		os.Exit(1)
	}

	// Collect the response
	resp, err := client.Do(req)
	if err != nil {
		log.Error("Failed to retrieve stats: %s", err.Error())
		os.Exit(1)
	}
	resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Error("Failed to retrieve stats with error code %s", resp.Status)
		os.Exit(1)
	}

	// Process CSV response into an array of metric:value maps
	haproxyObjects, err := processResponseToMap(resp.Body)
	if err != nil {
		log.Error("Failed to parse response: %s", err)
		os.Exit(1)
	}

	// Collect metrics and inventory for each row of the result
	for _, haproxyObject := range haproxyObjects {
		e, err := createEntity(haproxyObject, haproxyIntegration, args.StatsURL)
		if err != nil {
			log.Error("Failed to create entity: %s", err)
			os.Exit(1)
		}
		haproxyIntegration.AddEntity(e)

		if args.HasMetrics() {
			collectMetrics(haproxyObject, e)
		}

		if args.HasInventory() {
			collectInventory(haproxyObject, e)
		}
	}

	if err = haproxyIntegration.Publish(); err != nil {
		println("ERROR")
		log.Error("Failed to publish integration: %s", err.Error())
	}
}
