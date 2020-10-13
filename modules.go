package main

import (
	_ "github.com/l-herrmann/logspout-fluentd/fluentd"
	_ "github.com/gliderlabs/logspout/httpstream"
	_ "github.com/gliderlabs/logspout/transports/tcp"
	_ "github.com/gliderlabs/logspout/transports/udp"
)
