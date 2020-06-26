package api

import (
	"errors"
	"log"
	"net"
	"net/http"
	"strconv"
)

const (
	addressDivider       string = ":"
	defaultServerAddress string = ":0"
)

type Server struct {
	port string
}

// NewServer .
func NewServer() (*Server, error) {
	port, err := getRandomPort()
	if err != nil {
		return nil, err
	}
	return &Server{
		port: getNetworkAddress(port),
	}, nil
}

// Run .
func (s *Server) Run() {
	http.Handle("/", http.FileServer(FS(false)))
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		log.Print("error starting http server")
	}
}

// Port .
func (s *Server) Port() string {
	return s.port
}

func getRandomPort() (string, error) {
	ln, err := net.Listen("tcp", defaultServerAddress)

	if err != nil {
		return "", errors.New("error listening on default address")
	}

	defer ln.Close()
	return strconv.Itoa(ln.Addr().(*net.TCPAddr).Port), nil
}

func getNetworkAddress(port string) string {
	return addressDivider + port
}
