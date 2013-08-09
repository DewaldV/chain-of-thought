package crucible

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Conf *CoreConfigStruct

func LoadConfig(path string) error {
	configFile, err := os.Open(path)
	if err != nil {
		return &ConfigError{path, "Could not open config file for reading.", err.Error()}
	}
	defer configFile.Close()

	reader := bufio.NewReader(configFile)
	contents, err := ioutil.ReadAll(reader)
	if err != nil {
		return &ConfigError{path, "Error reading config file", err.Error()}
	}

	Conf = newCoreConfigStruct()

	err = json.Unmarshal(contents, Conf)
	if err != nil {
		return &ConfigError{path, "Error parsing config file", err.Error()}
	}

	fmt.Println("Loaded configuration:")
	fmt.Println(Conf.printConfig())

	return nil
}

type SiteConfigStruct struct {
	Location       string
	AllowedOrigins map[string]bool
}

type CoreConfigStruct struct {
	HttpPort      int
	HttpsPort     int
	WorkerThreads int
	RootContext   string
	//Sites map[string]SiteConfigStruct
}

func (c *CoreConfigStruct) printConfig() (s string) {
	s = fmt.Sprintf("HttpPort: %d\n", c.HttpPort)
	s += fmt.Sprintf("HttpsPort: %d\n", c.HttpsPort)
	s += fmt.Sprintf("RootContext: %s\n", c.RootContext)
	s += fmt.Sprintf("WorkerThreads: %d\n", c.WorkerThreads)
	return
}

func newCoreConfigStruct() (c *CoreConfigStruct) {
	c = new(CoreConfigStruct)
	c.HttpPort = 8787
	c.HttpsPort = 44443
	c.RootContext = "/"
	c.WorkerThreads = 1
	return
}

func newSiteConfigStruct() (s *SiteConfigStruct) {
	s = new(SiteConfigStruct)
	s.Location = "/"
	return
}

type ConfigError struct {
	ConfigFile string
	What       string
	Err        string
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("Configuration error in %s > %s > %s\n", e.ConfigFile, e.What, e.Err)
}
