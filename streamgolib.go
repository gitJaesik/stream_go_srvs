package streamgolib

import (
	"errors"
	"math/rand"
	"net"
	"os"
	"sync"
	"time"
)

const timeFormat = "2006-01-02T15:04:05.000-07:00"

var productName = ""
var productNameOnce sync.Once
var localAddress = ""
var localAddressOnce sync.Once
var hostname = ""
var hostnameOnce sync.Once

func init() {
	rand.Seed(time.Now().UnixNano())
}

func SetProductName(name string) {
	productNameOnce.Do(func() {
		productName = name
	})
}

func ProductName() string {
	return productName
}

func getLocalAddress() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
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
			return "", err
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
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func LocalAddress() string {
	localAddressOnce.Do(func() {
		addr, err := getLocalAddress()
		if err != nil {
			// log.Error("get local address failed: %v", err.Error())
			addr = "127.0.0.1"
		}

		localAddress = addr
	})

	return localAddress
}

func Hostname() string {
	hostnameOnce.Do(func() {
		name, err := os.Hostname()
		if err != nil {
			name = "hostname"
		}

		hostname = name
	})

	return hostname
}

func TimeFormatISO8601(t time.Time) string {
	local := t.Local()
	return local.Format(timeFormat)
}
