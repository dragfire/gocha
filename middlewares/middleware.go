package middleware

import (
	"github.com/dragfire/gocha/logger"
	"net"
	"net/http"
)

func getIpAddress() string {
	var ip net.IP
	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
		}
	}
	return ip.String()
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info.Println(getIpAddress())
		logger.Info.Println(r.UserAgent())
		logger.Info.Println(r.Referer())
		logger.Info.Println(r.RequestURI)
		logger.Info.Println(r.Host)
		next.ServeHTTP(w, r)
	})
}
