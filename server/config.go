package server

import "net"

// Config for our application
type Config struct {
	Host       string
	Port       string
	ConfigFile string
	AppName    string
	AppVersion string
}

// Address returns the configured host:port
func (c Config) Address() string {
	return net.JoinHostPort(c.Host, c.Port)
}
