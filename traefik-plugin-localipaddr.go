package traefik_plugin_localipaddr

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
)

// Config the plugin configuration.
type Config struct {
	Enabled bool   `json:"enabled,omitempty"`
	Domain  string `json:"domain,omitempty"`
	Ipv4    bool   `json:"ipv4,omitempty"` // false, get ipv6
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Domain: "[2400:3200::1]:53",
		Ipv4:   false,
	}
}

type plugin struct {
	next   http.Handler
	name   string
	config *Config
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &plugin{
		next:   next,
		name:   name,
		config: config,
	}, nil
}

func (r *plugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if !r.config.Enabled {
		r.next.ServeHTTP(rw, req)
		return
	}

	rw.Header().Set("Content-Type", "text/plan")

	ip, err := getNetworkIpv6Addr(r.config.Domain, r.config.Ipv4)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = os.Stderr.WriteString(fmt.Sprintf("failed to get ipaddr. err:%+v", err))
		_, _ = fmt.Fprintf(rw, "failed to get ipaddr. err:%+v", err)
		return
	}

	_, err = rw.Write([]byte(ip))
}

func getNetworkIpv6Addr(domain string, ipv4 bool) (string, error) {
	network := "udp6"
	if ipv4 {
		network = "udp"
	}

	conn, err := net.Dial(network, domain)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
