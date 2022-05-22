package server

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// Option -.
type Option func(s *server)

// Port -.
func Port(port string) Option {
	return func(s *server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

// ReadTimeout -.
func ReadTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.server.ReadTimeout = timeout
	}
}

// WriteTimeout -.
func WriteTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.server.WriteTimeout = timeout
	}
}

// ShutdownTimeout -.
func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *server) {
		s.shutdownTimeout = timeout
	}
}

// SetupTLS -.
func SetupTLS(certFile, keyFile string) Option {
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	return func(s *server) {
		s.server.TLSConfig = cfg
		s.server.TLSNextProto = make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0)
		s.certFileLoc = certFile
		s.keyFileLoc = keyFile
	}
}
