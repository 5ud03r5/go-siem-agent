package main

import (
	"flag"
	"sync"

	"github.com/5ud03r5/go-siem-agent/logprocessor"
	"github.com/5ud03r5/go-siem-agent/utils"
)

func main() {
	var yaml string
	var destinationHost string
	var destinationPort int

	flag.StringVar(&yaml, "yaml", "example_config.yaml", "Yaml config")
	flag.StringVar(&destinationHost, "destination-host", "127.0.0.1", "Destination host")
	flag.IntVar(&destinationPort, "destination-port", 80, "Destination port")

	utils.CreateDirectory("backlog")

	var wg sync.WaitGroup
	config := utils.RetrieveParsedYaml(yaml)
	logprocessor.RunProcessors(&config, &wg)
	wg.Wait()

}