package discovery

import (
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "strconv"
    "time"
    "github.com/micro/mdns"
)


type ServiceDescription struct {
    host string
    port int
    ipV4 []net.IP
}

type Service struct {
    Service ServiceDescription
    Server mdns.Server
    mdnsservice mdns.MDNSService
}

const (
    serviceTag = "_raspidash._tcp"
)

func getIPs() ([]net.IP) {
	var IPs []net.IP
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
            // interface up with no addrs
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			IPs = append(IPs, ip)
		}
	}

	return IPs
}


func (s Service) DoDiscoverable(hostname string, hostport string) {
    s.Service.host = hostname

    _, listenPort, _ := net.SplitHostPort(hostport)
    intListenPort, _ := strconv.Atoi(listenPort)
    s.Service.port = intListenPort
    info := []string{"A raspidash device"}
    // TODO: Using first by default, but this should be configurable
    log.Printf("IPs %v", getIPs()[:1])

    mdnsservice, err := mdns.NewMDNSService(s.Service.host,
                                            serviceTag,
                                            "local.",
                                            fmt.Sprintf("%s.local.", s.Service.host),
                                            s.Service.port,
                                            getIPs()[:1],
                                            info)
    if err != nil {
        log.Fatal(err)
    }
    server, err := mdns.NewServer(&mdns.Config{Zone: mdnsservice})
    if err != nil {
        log.Fatal(err)
    }
    defer server.Shutdown()
    wait()
}


func wait() {
    ch := make(chan os.Signal)
    signal.Notify(ch, os.Interrupt, os.Kill)
    <-ch
    log.Println("Shutting down dns server")
    os.Exit(0)
}

func ListDevices() ([]ServiceDescription) {
    // Make a channel for results and start listening
    entriesCh := make(chan *mdns.ServiceEntry, 4)
    //serviceNameMap := make(map[string]bool)
    defer close(entriesCh)

    go func() {
        for entry := range entriesCh {
            log.Printf("Got new entry: %#v\n", entry)
        }
    }()

    // Start the lookup
    mdns.Query(
        &mdns.QueryParam{
            Service: serviceTag,
            Domain: "local",
            Timeout: 10*time.Second,
            Entries: entriesCh,
            WantUnicastResponse: false,
        })
    return nil
}
