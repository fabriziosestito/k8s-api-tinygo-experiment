package net

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
)

type DialerFunc func(req *http.Request) (net.Conn, error)

type DialFunc func(ctx context.Context, net, addr string) (net.Conn, error)

type TLSClientConfigHolder interface {
	TLSClientConfig() *tls.Config
}

type Dialer interface {
	Dial(req *http.Request) (net.Conn, error)
}

type WarningHeader struct {
	Code  int
	Agent string
	Text  string
}

type RoundTripperWrapper interface {
	WrappedRoundTripper() interface{}
}

func JoinPreservingTrailingSlash(elem ...string) string {
	panic("stub")
}

func IsTimeout(err error) bool {
	panic("stub")
}

func IsProbableEOF(err error) bool {
	panic("stub")
}

func SetOldTransportDefaults(t interface{}) interface{} {
	panic("stub")
}

func SetTransportDefaults(t interface{}) interface{} {
	panic("stub")
}

func DialerFor(transport interface{}) (DialFunc, error) {
	panic("stub")
}

func CloseIdleConnectionsFor(transport interface{}) {
	panic("stub")
}

func TLSClientConfig(transport interface{}) (*tls.Config, error) {
	panic("stub")
}

func FormatURL(scheme string, host string, port int, path string) *url.URL {
	panic("stub")
}

func GetHTTPClient(req *http.Request) string {
	panic("stub")
}

func SourceIPs(req *http.Request) []net.IP {
	panic("stub")
}

func GetClientIP(req *http.Request) net.IP {
	panic("stub")
}

func AppendForwardedForHeader(req *http.Request) {
	panic("stub")
}

func NewProxierWithNoProxyCIDR(delegate func(req *http.Request) (*url.URL, error)) func(req *http.Request) (*url.URL, error) {
	panic("stub")
}

func (fn DialerFunc) Dial(req *http.Request) (net.Conn, error) {
	panic("stub")
}

func CloneRequest(req *http.Request) *http.Request {
	panic("stub")
}

func CloneHeader(in http.Header) http.Header {
	panic("stub")
}

func ParseWarningHeaders(headers []string) ([]WarningHeader, []error) {
	panic("stub")
}

func ParseWarningHeader(header string) (result WarningHeader, remainder string, err error) {
	panic("stub")
}

func NewWarningHeader(code int, agent, text string) (string, error) {
	panic("stub")
}

type RouteFile struct {
	name  string
	parse func(input io.Reader) ([]Route, error)
}

type networkInterface struct{}

type AddressFamilyPreference []AddressFamily

type noRoutesError struct{ message string }

type networkInterfacer interface {
	InterfaceByName(intfName string) (*interface{}, error)
	Addrs(intf *interface{}) ([]net.Addr, error)
	Interfaces() ([]interface{}, error)
}

type AddressFamily uint

type Route struct {
	Interface   string
	Destination net.IP
	Gateway     net.IP
	Family      AddressFamily
}

func (e noRoutesError) Error() string {
	panic("stub")
}

func IsNoRoutesError(err error) bool {
	panic("stub")
}

func ChooseHostInterface() (net.IP, error) {
	panic("stub")
}

func (_ networkInterface) InterfaceByName(intfName string) (*interface{}, error) {
	panic("stub")
}

func (_ networkInterface) Addrs(intf *interface{}) ([]net.Addr, error) {
	panic("stub")
}

func (_ networkInterface) Interfaces() ([]interface{}, error) {
	panic("stub")
}

func ResolveBindAddress(bindAddress net.IP) (net.IP, error) {
	panic("stub")
}

func ChooseBindAddressForInterface(intfName string) (net.IP, error) {
	panic("stub")
}

type PortRange struct {
	Base int
	Size int
}

func (pr *PortRange) Contains(p int) bool {
	panic("stub")
}

func (pr PortRange) String() string {
	panic("stub")
}

func (pr *PortRange) Set(value string) error {
	panic("stub")
}

func ParsePortRange(value string) (*PortRange, error) {
	panic("stub")
}

func ParsePortRangeOrDie(value string) *PortRange {
	panic("stub")
}

func SplitSchemeNamePort(id string) (scheme, name, port string, valid bool) {
	panic("stub")
}

func JoinSchemeNamePort(scheme, name, port string) string {
	panic("stub")
}

func IPNetEqual(ipnet1, ipnet2 *net.IPNet) bool {
	panic("stub")
}

func IsConnectionReset(err error) bool {
	panic("stub")
}

func IsHTTP2ConnectionLost(err error) bool {
	panic("stub")
}

func IsConnectionRefused(err error) bool {
	panic("stub")
}

type Embedme interface{}
