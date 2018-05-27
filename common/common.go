package common

import (
	"fmt"
	"github.com/json-iterator/go"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "9876"
	CONN_TYPE_TCP = "tcp"
	CONN_TYPE_UDP = "udp"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// TcpConn
// Return Tcp Listener to accept connections
func TcpConn() (net.Listener, error) {
	return net.Listen(CONN_TYPE_TCP, fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT))
}
// UdpConn
// Return Udp Listener to parse data streams
func UdpConn() (net.PacketConn, error) {
	return net.ListenPacket(CONN_TYPE_UDP, fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT))
}

// AsJson
// Simplify naming convention for third party
func AsJson(b []byte) (MyObject, error) {
	obj := MyObject{}
	err := json.Unmarshal(b, &obj)
	return obj, err
}
// AsByte
// Simplify naming convention for third party
func AsByte(object MyObject) ([]byte, error) {
	return json.Marshal(object)
}
// HandleJSON
// Modify this function to do whatever you need to do
// with your struct.
func HandleJSON(b []byte) ([]byte, error) {
	obj, err := AsJson(b)
	if err != nil {
		return nil, err
	}
	// Do what you will with the JSON object here
	// Inject a response into the struct and send it back
	obj.Response = "Handled"
	// ...
	// ...
	// ...
	return AsByte(obj)
}

// SendJson
// Send the MyObject struct to the server
// return the connection and error object
func SendJson(object MyObject) (net.Conn, error) {
	data, err := AsByte(object)
	if err != nil {
		return nil, err
	}
	conn, err := net.Dial(CONN_TYPE_TCP, fmt.Sprintf("%s:%s", CONN_HOST, CONN_PORT))
	if err != nil {
		return conn, err
	}
	conn.Write(data)
	return conn, err
}