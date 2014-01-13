package main


import "fmt"
import "os"

import (
  "io/ioutil"
  "encoding/json";
)

var config_location = "/Users/pete/.xrebug.json"
var config Configuration
var ini_location string


type Configuration struct {
    Inifile string `json:"inifile"`
}


func main() {
    config := getConfig(config_location)
    ini_location = config.Inifile
    if isXdebugEnabled(ini_location) {
        fmt.Println("Disabling xdebug")
        disableXdebug(ini_location)
    } else {
        fmt.Println("Enabling xdebug")
        enableXdebug(ini_location)
    }
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
    // needs to warn if xdebug isn't in its own file
// }

func enableXdebug(location string) {
    disabled_location := getDisabledLocation(location)
    if !isXdebugEnabled(location) {
        os.Rename(disabled_location, location)
    }
}

func disableXdebug(location string) {
    disabled_location := getDisabledLocation(location)
    if isXdebugEnabled(location) {
        os.Rename(location, disabled_location)
    }
}

func getDisabledLocation(location string) string {
    return location + ".disabled"
}

func toggleXdebug(location string) {
    if !isXdebugEnabled(location) {
        enableXdebug(location)
    } else {
        disableXdebug(location)
    }
}

func isXdebugEnabled(location string) (bool) {
    config, err := os.Open(location)
    config = config // hmm ... how else do we "not use" config?
    return !os.IsNotExist(err)
}

