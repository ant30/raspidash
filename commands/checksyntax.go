package commands

import (
    "log"
    "flag"
    "github.com/ant30/raspidash/models"
)

var (
    checkCommand = flag.NewFlagSet("check", flag.ExitOnError)
)


func CheckSyntax(args []string) {
    var settings models.Settings

    configFilename := checkCommand.String("config", "settings.json", "This is the filename of json settings, like settings.json")
    checkCommand.Parse(args)
    log.Printf("Prepare to open the json file %v \n", *configFilename)

    settings.ReadFromJsonFile(*configFilename)

    log.Printf("Settings parsed:  %#v", settings)
}
