package dict

import "net/textproto"

type Defn struct {
	Dict Dict
	Word string
	Text []byte
}

type Client struct{ text *textproto.Conn }

type Dict struct {
	Name string
	Desc string
}

func Dial(network, addr string) (*Client, error) {
	panic("stub")
}

func (c *Client) Close() error {
	panic("stub")
}

func (c *Client) Dicts() ([]Dict, error) {
	panic("stub")
}

func (c *Client) Define(dict, word string) ([]*Defn, error) {
	panic("stub")
}

type Embedme interface{}
