package common

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	letterBytes   = "0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var cidrs []*net.IPNet

// GenerateReferenceID generates reference ID
func GenerateReferenceID(prefix string) string {
	now := time.Now().Format("20060102030405")
	buff := bytes.NewBufferString(now)
	buff.WriteString(GenerateRandomNumberString(8))

	return fmt.Sprintf("%s%s", prefix, buff.String())
}

// GenerateRandomNumberString generate random string number
func GenerateRandomNumberString(n int) string {
	src := rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// IPFromRequest return client's real public IP address from http request headers.
func IPFromRequest(r *http.Request) string {
	// Fetch header value
	xRealIP := r.Header.Get("X-Real-Ip")
	xForwardedFor := r.Header.Get("X-Forwarded-For")

	// If both empty, return IP from remote address
	if xRealIP == "" && xForwardedFor == "" {
		var remoteIP string

		// If there are colon in remote address, remove the port number
		// otherwise, return remote address as is
		remoteIP = r.RemoteAddr
		if strings.ContainsRune(r.RemoteAddr, ':') {
			remoteIP, _, _ = net.SplitHostPort(r.RemoteAddr)
		}

		return remoteIP
	}

	// Check list of IP in X-Forwarded-For and return the first global address
	for _, addr := range strings.Split(xForwardedFor, ",") {
		addr = strings.TrimSpace(addr)
		isPrivate, err := isPrivateAddress(addr)
		if !isPrivate && err == nil {
			return addr
		}
	}

	// If nothing succeed, return X-Real-IP
	return xRealIP
}

func isPrivateAddress(address string) (bool, error) {
	ipAddress := net.ParseIP(address)
	if ipAddress == nil {
		return false, errors.New("address is not valid")
	}

	for i := range cidrs {
		if cidrs[i].Contains(ipAddress) {
			return true, nil
		}
	}

	return false, nil
}
