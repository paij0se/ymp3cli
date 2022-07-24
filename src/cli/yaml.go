package cli

import (
	"os"

	"gopkg.in/yaml.v3"
)

func CreateConfigDirectory() {
	os.MkdirAll(os.Getenv("HOME")+"/.ymp3cli/", 0755)
	if _, err := os.Stat(os.Getenv("HOME") + "/.ymp3cli/config.yaml"); os.IsNotExist(err) {
		file, err := os.Create(os.Getenv("HOME") + "/.ymp3cli/config.yaml")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		config := `
# port: the port that the server will listen on
port: 8888
`
		file.WriteString(config)
	}

}
func ReadFromYaml(variable string) int {
	file, err := os.Open(os.Getenv("HOME") + "/.ymp3cli/config.yaml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var config map[string]int
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return int(config[variable])
}
