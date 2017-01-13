package discovery

import (
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "strconv"
    "strings"
    "time"
    "github.com/micro/mdns"
)


type ServiceDescription struct {
    Host string
    Name string
    Port int
    IPv4 net.IP
}

type Service struct {
    Service ServiceDescription
    Server mdns.Server
    mdnsservice mdns.MDNSService
}

type ListDevicesCache struct {
    services []ServiceDescription
    updated time.Time
}

const (
    serviceTag = "_raspidash._tcp"
    domain = "local"
)

var (
    listDevicesCache ListDevicesCache
    dnsCacheExpiration time.Duration = 60*time.Second
    dnsSearchTimeout time.Duration = 1*time.Second
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
    s.Service.Host = hostname

    _, listenPort, _ := net.SplitHostPort(hostport)
    intListenPort, _ := strconv.Atoi(listenPort)
    s.Service.Port = intListenPort
    s.Service.IPv4 = getIPs()[0]
    info := []string{"A raspidash device"}
    // TODO: Using first by default, but this should be configurable
    log.Printf("IPs %v", getIPs()[:1])

    mdnsservice, err := mdns.NewMDNSService(s.Service.Host,
                                            serviceTag,
                                            fmt.Sprintf("%s.", domain),
                                            fmt.Sprintf("%s.%s.", s.Service.Host, domain),
                                            s.Service.Port,
                                            []net.IP{s.Service.IPv4},
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

func (c ListDevicesCache) isValid() (bool) {
    if c.services == nil {
        return false
    }
    if time.Now().Sub(c.updated) > dnsCacheExpiration  {
        return false
    }
    return true
}

func listDevices(timeout time.Duration) ([]ServiceDescription) {
    // Make a channel for results and start listening
    var sd []ServiceDescription
    entriesCh := make(chan *mdns.ServiceEntry, 4)
    serviceNameMap := make(map[string]bool)

    defer close(entriesCh)

    go func() {
        for entry := range entriesCh {
            name := strings.TrimSuffix(entry.Name, fmt.Sprintf(".%s.%s.", serviceTag, domain))
            if !serviceNameMap[name] {
                serviceNameMap[name] = true
                sd = append(sd, ServiceDescription{
                    Host: fmt.Sprintf("%s.%s", name, domain),
                    Name: name,
                    Port: entry.Port,
                    IPv4: entry.AddrV4,
                })
            }
        }
    }()

    // Start the lookup
    mdns.Query(
        &mdns.QueryParam{
            Service: serviceTag,
            Domain: domain,
            Timeout: timeout,
            Entries: entriesCh,
            WantUnicastResponse: false,
        })
    return sd
}

func ListDevices() ([]ServiceDescription) {
    if ! listDevicesCache.isValid() {
        listDevicesCache.services = listDevices(dnsSearchTimeout)
        listDevicesCache.updated = time.Now()
    }
    return listDevicesCache.services
}
