package metatrader

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"time"
)

type MTFunctions struct {
	HOST                     string
	PORT                     int
	debug                    bool
	instrumentConversionList []string
	authorizationCode        string
	connected                bool
	socket_error_message     string
	timeout_value            int
	sock                     net.Conn
}

func (mt *MTFunctions) IsConnected() bool {
	return mt.connected
}

func (mt *MTFunctions) setTimeout(timeout_in_seconds int) {
	mt.timeout_value = timeout_in_seconds
	mt.sock.SetDeadline(time.Now().Add(time.Duration(mt.timeout_value) * time.Second))
	mt.sock.(*net.TCPConn).SetReadBuffer(SOCKET_BUFFER_SIZE)
}

func (mt *MTFunctions) Disconnect() bool {
	mt.sock.Close()
	return true
}

func (mt *MTFunctions) Connect() bool {
	// Connect to the server
	// err := mt.sock.Dial((mt.HOST, mt.PORT))
	// err := mt.sock.("tcp", fmt.Sprintf("%s:%d", mt.HOST, mt.PORT))
	var err error

	mt.sock, err = net.Dial("tcp", fmt.Sprintf("%s:%d", mt.HOST, mt.PORT))
	if err != nil {
		fmt.Printf("Could not connect with the socket-server: %s\n terminating program", err)
		mt.connected = false
		mt.socket_error_message = "Could not connect to server."
		mt.sock.Close()
		return false
	}
	mt.sock.(*net.TCPConn).SetNoDelay(true)
	mt.sock.(*net.TCPConn).SetKeepAlive(true)
	mt.setTimeout(mt.timeout_value)
	mt.connected = true
	return true
}

func (mt *MTFunctions) sendRequest(data string) error {
	_, err := mt.sock.Write([]byte(data))

	return err
}

func (mt *MTFunctions) recv(bufferSize int) ([]byte, error) {
	data := make([]byte, bufferSize)

	n, err := mt.sock.Read(data)
	if n == 0 && len(data) == 0 && err == nil { 
		return nil, fmt.Errorf("no data received from the server")
	}

	if err != nil {
		if errors.Is(err, os.ErrDeadlineExceeded) {
			return nil, fmt.Errorf("timeout reached while reading from the server")
		}
		return nil, fmt.Errorf("error reading data from the server: %v", err)
	}

	return data[:n-1], nil
}

func (mt *MTFunctions) getReply() (interface{}, error) {
	var msg interface{}

	buffer := ""

	for {
		data, err := mt.recv(SOCKET_BUFFER_SIZE)
		if err != nil {
			return nil, err
		}

		buffer += string(data)

		if err := json.Unmarshal([]byte(buffer), &msg); err == nil {
			// Successfully parsed JSON, return the message
			// fmt.Println("msg => ", msg)
			return msg, nil
		} else if _, ok := err.(*json.SyntaxError); ok {
			// Incomplete JSON, continue reading
			continue
		} else {
			// Other JSON decoding error, return an error
			return nil, err
		}
	}
}

func (mt *MTFunctions) SendCommand(command string) (interface{}, error) {
	request := command + "|" + mt.authorizationCode + "\r\n"
	err := mt.sendRequest(request)
	if err != nil {
		return nil, err
	}

	return mt.getReply()
}

func NewMTFunctions(host string, port int, debug bool, instrumentConversionList []string, authorizationCode string) *MTFunctions {
	mt := &MTFunctions{
		HOST:                     host,
		PORT:                     port,
		debug:                    debug,
		instrumentConversionList: instrumentConversionList,
		authorizationCode:        authorizationCode,
		connected:                false,
		socket_error_message:     "",
		timeout_value:            120,
		sock:                     nil,
	}
	return mt
}
