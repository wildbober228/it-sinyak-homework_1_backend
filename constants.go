package main

import "time"

const (
	defaultPort       = "8080"
	defaultOrigins    = "http://localhost:5173,http://localhost:4173"
	readHeaderTimeout = 5 * time.Second
)
