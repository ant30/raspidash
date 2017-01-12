package discovery

import (
    "os"
)

func GetHostname() (string) {
    host, _ := os.Hostname()
    return host
}
