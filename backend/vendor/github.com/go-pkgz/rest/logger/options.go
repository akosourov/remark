package logger

import (
	"net/http"
)

// Option func type
type Option func(l *Middleware)

// WithBody triggers request body logging. Body size is limited (default 1k)
func WithBody(l *Middleware) {
	l.logBody = true
}

// MaxBodySize sets size of the logged part of the request body.
func MaxBodySize(max int) Option {
	return func(l *Middleware) {
		if max >= 0 {
			l.maxBodySize = max
		}
	}
}

// Prefix sets log line prefix.
func Prefix(prefix string) Option {
	return func(l *Middleware) {
		l.prefix = prefix
	}
}

// IPfn sets IP masking function. If ipFn is nil then IP address will be logged as is.
func IPfn(ipFn func(ip string) string) Option {
	return func(l *Middleware) {
		l.ipFn = ipFn
	}
}

// UserFn triggers user name logging if userFn is not nil.
func UserFn(userFn func(r *http.Request) (string, error)) Option {
	return func(l *Middleware) {
		l.userFn = userFn
	}
}

// SubjFn triggers subject logging if subjFn is not nil.
func SubjFn(subjFn func(r *http.Request) (string, error)) Option {
	return func(l *Middleware) {
		l.subjFn = subjFn
	}
}

// Log sets logging backend. 
func Log(log Backend) Option {
	return func(l *Middleware) {
		l.log = log
	}
}
