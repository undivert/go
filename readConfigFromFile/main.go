package main

import (
  "fmt"
  "log"

  "gopkg.in/yaml.v2"
)

// Configuration represents the structure of our configuration file
type Configuration struct {
  Port int `yaml:"port"`
  Host string `yaml:"host"`
}

// ReadConfig reads the configuration from a YAML file and returns a Configuration struct
func ReadConfig(filename string) (Configuration, error) {
  // Load the configuration file into a byte slice
  configData, err := ioutil.ReadFile(filename)
  if err != nil {
    return Configuration{}, err
  }

  // Initialize a new Configuration struct
  var config Configuration

  // Unmarshal the YAML data into the Configuration struct
  err = yaml.Unmarshal(configData, &config)
  if err != nil {
    return Configuration{}, err
  }

  return config, nil
}

func main() {
  // Read the configuration from the YAML file
  config, err := ReadConfig("config.yml")
  if err != nil {
    log.Fatal(err)
  }

  // Print the configuration data
  fmt.Println(config.Port)
  fmt.Println(config.Host)
}
