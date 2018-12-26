/*
 * @Description:
 * @Author: EagleXiang
 * @Github: https://github.com/eaglexiang
 * @Date: 2018-12-26 11:09:54
 * @LastEditors: EagleXiang
 * @LastEditTime: 2018-12-26 11:22:02
 */

package eagletunnel

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

// WhitelistDomains 需要被智能解析的DNS域名列表
var WhitelistDomains []string

// ResolvIPv4ByLocal 本地解析IPv4
func ResolvIPv4ByLocal(e *NetArg) error {
	addrs, err := net.LookupHost(e.domain)
	if err == nil {
		var ok bool
		for _, addr := range addrs {
			ip := net.ParseIP(addr)
			if ip.To4() != nil {
				e.IP = addr
				ok = true
				break
			}
		}
		if !ok {
			err = errors.New("ipv4 not found")
		}
	}
	return err
}

// ResolvIPv6ByLocal 本地解析IPv6
func ResolvIPv6ByLocal(e *NetArg) error {
	addrs, err := net.LookupHost(e.domain)
	if err == nil {
		var ok bool
		for _, addr := range addrs {
			ip := net.ParseIP(addr)
			if ip.To16() != nil && ip.To4() == nil {
				e.IP = addr
				ok = true
				break
			}
		}
		if !ok {
			err = errors.New("ipv4 not found")
		}
	}
	return err
}

// CheckPrivateIPv4 检查是否是保留IPv4地址
func CheckPrivateIPv4(ip string) bool {
	items := strings.Split(ip, ".")
	itemsInt := make([]int, 4)
	for ind, item := range items {
		itemInt, _ := strconv.ParseInt(item, 10, 32)
		itemsInt[ind] = int(itemInt)
	}
	// 0.0.0.0–0.255.255.255
	if itemsInt[0] == 0 {
		return true
	}
	// 10.0.0.0–10.255.255.255
	if itemsInt[0] == 10 {
		return true
	}
	// 100.64.0.0–100.127.255.255
	if itemsInt[0] == 100 {
		if 64 <= itemsInt[1] && itemsInt[1] <= 127 {
			return true
		}
	}
	// 127.0.0.0–127.255.255.255
	if itemsInt[0] == 127 {
		return true
	}
	// 169.254.0.0–169.254.255.255
	if itemsInt[0] == 169 && itemsInt[1] == 254 {
		return true
	}
	// 172.16.0.0–172.31.255.255
	if itemsInt[0] == 172 {
		if 16 <= itemsInt[1] && itemsInt[1] <= 31 {
			return true
		}
	}
	// 192.0.0.0–192.0.0.255
	if itemsInt[0] == 192 && itemsInt[1] == 0 && itemsInt[2] == 0 {
		return true
	}
	// 192.0.2.0–192.0.2.255
	if itemsInt[0] == 192 && itemsInt[1] == 0 && itemsInt[2] == 2 {
		return true
	}
	// 192.88.99.0–192.88.99.255
	if itemsInt[0] == 192 && itemsInt[1] == 88 && itemsInt[2] == 99 {
		return true
	}
	// 192.168.0.0–192.168.255.255
	if itemsInt[0] == 192 && itemsInt[1] == 168 {
		return true
	}
	// 198.18.0.0–198.19.255.255
	if itemsInt[0] == 198 {
		if 18 <= itemsInt[1] && itemsInt[1] <= 19 {
			return true
		}
	}
	// 198.51.100.0–198.51.100.255
	if itemsInt[0] == 198 && itemsInt[1] == 51 && itemsInt[2] == 100 {
		return true
	}
	// 203.0.113.0–203.0.113.255
	if itemsInt[0] == 203 && itemsInt[1] == 0 && itemsInt[2] == 113 {
		return true
	}
	// 224.0.0.0–239.255.255.255
	if 224 <= itemsInt[0] && itemsInt[0] <= 239 {
		return true
	}
	// 240.0.0.0–255.255.255.254
	// 255.255.255.255
	if 240 <= itemsInt[0] {
		return true
	}
	return false
}

// IsWhiteDomain 判断域名是否是白名域名
func IsWhiteDomain(host string) (isWhite bool) {
	var white bool
	for _, line := range WhitelistDomains {
		if strings.HasSuffix(host, line) {
			white = true
			break
		}
	}
	return white
}
