package main


import "fmt"
import "os"

import (
  "io/ioutil"
  "encoding/json";
)

var config_location = "/Users/pete/.xrebug.json"
var config Configuration


type Configuration struct {
    Inifile string `json:"inifile"`
}


func main() {
    config := getConfig(config_location)
    fmt.Println(config.Inifile)
}


func getConfig(config_location string) Configuration {
    file, err := ioutil.ReadFile(config_location)
    if err != nil {
        file = makeConfigFile(config_location)
    }
    return parseConfig(file)
}


func parseConfig(contents []byte) Configuration {
    var configuration Configuration
    json.Unmarshal(contents, &configuration)
    return configuration
}


func makeConfigFile(config_location string) []byte {
    fo, err := os.Create(config_location)
    if err != nil { panic(err) }
    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()
    // this should use all the stuff below instead
    contents := "{\"inifile\":\"/usr/local/etc/php/5.5/conf.d/ext-xdebug.ini\"}"

    fo.WriteString(contents)
    file, err := ioutil.ReadFile(config_location)
    if err != nil { panic(err) }
    return file
}

// func findIniFiles() {

// }

// func findXdebugIniFile() {

// }

// func enableXdebug() {

// }

// func disableXdebug() {

// }

// func toggleXdebug() {

// }

// func isXdebugEnabled() {

// }

