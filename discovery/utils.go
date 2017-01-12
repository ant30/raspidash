package discovery

import (
    "os"
)

func getHostname() (string) {
    host, _ := os.Hostname()
    return host
}
