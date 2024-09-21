package pkg

import (
	"regexp"
	"strings"
)

// RemoveColorCodes 移除字符串中的 ANSI 颜色代码
func RemoveColorCodes(line string) string {
	re := regexp.MustCompile(`\033\[[0-9]*m`)
	return re.ReplaceAllString(line, "")
}

// 将字符串居中
func Center(s string, width int, fill string) string {
	if len(s) >= width {
		return s
	}
	padding := width - len(s)
	leftPadding := padding / 2
	rightPadding := padding - leftPadding
	return strings.Repeat(fill, leftPadding) + s + strings.Repeat(fill, rightPadding)
}

// 分离主机域名（IP）和端口号
//
// 示例：
// "example.com:8096" 					=> 	"example.com", "8096"
// "[240e:da8:a801:5a47::316]:8096" 	=> 	"240e:da8:a801:5a47::316" "8096"
// "192.168.1.1:8096" 					=> 	"192.168.1.1" "8096"
func SplitHostPort(hostPort string) (host string, port string) {
	host = hostPort

	colon := strings.LastIndexByte(host, ':')
	if colon != -1 && validOptionalPort(host[colon:]) {
		host, port = host[:colon], host[colon+1:]
	}

	// 对地址是IPv6地址进行处理
	if strings.HasPrefix(host, "[") && strings.HasSuffix(host, "]") {
		host = host[1 : len(host)-1]
	}

	return
}

func validOptionalPort(port string) bool {
	if port == "" {
		return true
	}
	if port[0] != ':' {
		return false
	}
	for _, b := range port[1:] {
		if b < '0' || b > '9' {
			return false
		}
	}
	return true
}
