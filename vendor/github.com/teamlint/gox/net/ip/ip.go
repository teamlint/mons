package ip

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

// GetSegment 获得ip的网段，例如：192.168.2.102 -> 192.168.2
func GetSegment(ip string) string {
	r := `^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`
	reg, err := regexp.Compile(r)
	if err != nil {
		return ""
	}
	ips := reg.FindStringSubmatch(ip)
	if ips == nil {
		return ""
	}
	return fmt.Sprintf("%s.%s.%s", ips[1], ips[2], ips[3])
}

// ParseAddress 解析地址，形如：192.168.1.1:80 -> 192.168.1.1, 80
func ParseAddress(addr string) (string, int) {
	r := `^(.+):(\d+)$`
	reg, err := regexp.Compile(r)
	if err != nil {
		return "", 0
	}
	result := reg.FindStringSubmatch(addr)
	if result != nil {
		i, _ := strconv.Atoi(result[2])
		return result[1], i
	}
	return "", 0
}

// IntranetIP 获取本地局域网ip列表
func IntranetIP() (ips []string, err error) {
	ips = make([]string, 0)
	ifaces, e := net.Interfaces()
	if e != nil {
		return ips, e
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}

		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		// ignore warden bridge
		if strings.HasPrefix(iface.Name, "w-") {
			continue
		}

		addrs, e := iface.Addrs()
		if e != nil {
			return ips, e
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

			ipStr := ip.String()
			if IsIntranet(ipStr) {
				ips = append(ips, ipStr)
			}
		}
	}
	return ips, nil
}

// IsIntranet 判断所给ip是否为局域网ip
// A类 10.0.0.0--10.255.255.255
// B类 172.16.0.0--172.31.255.255
// C类 192.168.0.0--192.168.255.255
func IsIntranet(ipStr string) bool {
	// ip协议保留的局域网ip
	if strings.HasPrefix(ipStr, "10.") || strings.HasPrefix(ipStr, "192.168.") {
		return true
	}
	if strings.HasPrefix(ipStr, "172.") {
		// 172.16.0.0 - 172.31.255.255
		arr := strings.Split(ipStr, ".")
		if len(arr) != 4 {
			return false
		}

		second, err := strconv.ParseInt(arr[1], 10, 64)
		if err != nil {
			return false
		}

		if second >= 16 && second <= 31 {
			return true
		}
	}

	return false
}

// IsIPV4  判断所给地址是否是一个IPv4地址
func IsIPV4(ip string) bool {
	reg, err := regexp.Compile(`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`)
	if err != nil {
		return false
	}
	return reg.Match([]byte(ip))
}

// ParseUint 字符串转为整形
func ParseUint(ipstr string) (ip uint32) {
	reg, _ := regexp.Compile(`^(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})$`)
	ips := reg.FindStringSubmatch(ipstr)
	if ips == nil {
		return
	}

	ip1, _ := strconv.Atoi(ips[1])
	ip2, _ := strconv.Atoi(ips[2])
	ip3, _ := strconv.Atoi(ips[3])
	ip4, _ := strconv.Atoi(ips[4])

	if ip1 > 255 || ip2 > 255 || ip3 > 255 || ip4 > 255 {
		return
	}

	ip += uint32(ip1 * 0x1000000)
	ip += uint32(ip2 * 0x10000)
	ip += uint32(ip3 * 0x100)
	ip += uint32(ip4)
	return
}

// Convert 整形转为IP字符串
func Convert(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip>>24, ip<<8>>24, ip<<16>>24, ip<<24>>24)
}
