package coreServ

type configItem struct {
	Key string
	Value string
}

var ConfigMap map[string]string

func ReadConfig(path string) {
	ConfigMap = make(map([string]string))
}