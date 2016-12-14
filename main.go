package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kyokomi/amber_lambda/crawl"
	"gopkg.in/xmlpath.v2"
	"gopkg.in/yaml.v2"
)

func readHeadersByConfig(configFilePath string) (map[string]string, error) {
	headers := map[string]string{}
	if len(configFilePath) > 0 {
		data, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			return nil, err
		}
		config := map[string]map[string]string{}
		if err := yaml.Unmarshal(data, &config); err != nil {
			return nil, err
		}
		headers = config["headers"]
	}
	return headers, nil
}

func main() {
	// input

	var targetURL, xpath, configFilePath string
	flag.StringVar(&targetURL, "i", "", "target page url")
	flag.StringVar(&xpath, "x", "", "target page playlist xpath")
	flag.StringVar(&configFilePath, "c", "sample_config.yaml", "crawler setting config (yaml)")
	flag.Parse()

	// setup
	headers, err := readHeadersByConfig(configFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	crawler := crawlhtml.New(http.DefaultTransport)
	for k, v := range headers {
		crawler.SetHeader(k, v)
	}

	// execute

	reader, err := crawler.HTML(targetURL)
	if err != nil {
		log.Fatalln(err)
	}

	node, err := xmlpath.ParseHTML(reader)
	if err != nil {
		log.Fatalln(err)
	}

	path, err := xmlpath.Compile(xpath)
	if err != nil {
		log.Fatalln(err)
	}

	// output

	it := path.Iter(node)
	for it.Next() {
		scrapText := it.Node().String()

		fmt.Println(scrapText)
	}
}
