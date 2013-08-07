package coreServ

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Conf CoreConfigStruct

func LoadConfig(path string) error {
	configFile, err := os.Open(path)
	if err != nil {
		return &ConfigError{path, 0, "Could not open config file for reading."}
	}
	defer configFile.Close()

	reader := bufio.NewReader(configFile)
	contents, err := ioutil.ReadAll(reader)

	Conf = newCoreConfigStruct()

	err = json.Unmarshal(contents, &Conf)
	if err != nil {
		return &ConfigError{path, 0, err.Error()}
	}

	fmt.Println("Loaded configuration:")
	fmt.Printf("HttpPort: %d\n", Conf.HttpPort)
	fmt.Printf("HttpsPort: %d\n", Conf.HttpsPort)
	fmt.Printf("RootContext: %s\n", Conf.RootContext)
	fmt.Printf("WorkerThreads: %d\n", Conf.WorkerThreads)

	return nil
}

type SiteConfigStruct struct {
	Location string
}

type CoreConfigStruct struct {
	HttpPort      int
	HttpsPort     int
	WorkerThreads int
	RootContext   string
	//Sites map[string]SiteConfigStruct
}

func newCoreConfigStruct() (c CoreConfigStruct) {
	c = CoreConfigStruct{
		HttpPort:      8787,
		HttpsPort:     44443,
		RootContext:   "/",
		WorkerThreads: 1}

	return
}

func newSiteConfigStruct() (s SiteConfigStruct) {
	s = SiteConfigStruct{
		Location: "/"}
	return
}

type ConfigError struct {
	ConfigFile string
	LineNumber int
	What       string
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("Configuration error %s:%d > %s\n", e.ConfigFile, e.LineNumber, e.What)
}
