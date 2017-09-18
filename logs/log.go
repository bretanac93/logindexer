package logs

import "time"

//Log The log struct defined by the proxy.
type Log struct {
	// Request's user
	User string `json:"user"`
	// Client address
	Addr string `json:"address"`
	// HTTP method
	Meth string `json:"method"`
	// Accessed URI
	URI string `json:"uri"`
	// HTTP version
	Proto string `json:"proto"`
	// Response status code
	StatusCode int `json:"status_code"`
	// Response size
	RespSize uint64 `json:"response_size"`
	// Response date-time
	Time time.Time `json:"time"`
}
