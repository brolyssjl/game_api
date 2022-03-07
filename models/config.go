package models

import "time"

// ServiceConfig defines the game server configurations
type ServiceConfig struct {
	Port                int
	ReadTimeout         time.Duration
	WriteTimeout        time.Duration
	IdleTimeout         time.Duration
	ReadHeaderTimeout   time.Duration
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
}
