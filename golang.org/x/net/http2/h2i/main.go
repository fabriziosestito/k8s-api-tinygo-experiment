package main

import (
	"bytes"
	"crypto/tls"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/hpack"
)

type h2i struct {
	host        string
	tc          *tls.Conn
	framer      *http2.Framer
	term        interface{}
	streamID    uint32
	hbuf        bytes.Buffer
	henc        *hpack.Encoder
	peerSetting map[http2.SettingID]uint32
	hdec        *hpack.Decoder
}

type command struct {
	run      func(*h2i, []string) error
	complete func() []string
}

func (app *h2i) Main() error {
	panic("stub")
}

type Embedme interface{}
