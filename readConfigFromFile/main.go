package main

import (
  "fmt"
  "log"
  "os"

  "gopkg.in/yaml.v2"
)

// Configuration represents the structure of our configuration file
type Configuration struct {
  Port int    `yaml:"port" env:"PORT"`
  Host string `yaml:"host" env:"HOST"`
}

// ReadConfig reads the configuration from a YAML file and environment variables,
// and returns a Configuration struct
func ReadConfig(filename string) (Configuration, error) {
  // Initialize a new Configuration struct
  var config Configuration

  // Read the configuration from the YAML file
  configData, err := ioutil.ReadFile(filename)
  if err != nil {
    return Configuration{}, err
  }
  err = yaml.Unmarshal(configData, &config)
  if err != nil {
    return Configuration{}, err
  }

  // Read the configuration from environment variables
  err = envdecode.Decode(&config)
  if err != nil {
    return Configuration{}, err
  }

  return config, nil
}

func main() {
  // Read the configuration from the YAML file and environment variables
  config, err := ReadConfig("config.yml")
  if err != nil {
    log.Fatal(err)
  }

  // Print the configuration data
  fmt.Println(config.Port)
  fmt.Println(config.Host)
}
