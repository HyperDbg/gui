package remote

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	CommunicationBufferSize = 4096
	TcpEndOfBufferChars     = "\r\n"
)

type NamedPipeServer struct {
	pipeName    string
	listener    net.Listener
	connections map[net.Conn]bool
	mu          sync.RWMutex
}

func NewNamedPipeServer(pipeName string) *NamedPipeServer {
	return &NamedPipeServer{
		pipeName:    pipeName,
		connections: make(map[net.Conn]bool),
	}
}

func (s *NamedPipeServer) Start() error {
	listener, err := net.Listen("tcp", s.pipeName)
	if err != nil {
		return fmt.Errorf("failed to create named pipe server: %w", err)
	}

	s.listener = listener
	return nil
}

func (s *NamedPipeServer) WaitForClient() (net.Conn, error) {
	conn, err := s.listener.Accept()
	if err != nil {
		return nil, fmt.Errorf("failed to accept client connection: %w", err)
	}

	s.mu.Lock()
	s.connections[conn] = true
	s.mu.Unlock()

	return conn, nil
}

func (s *NamedPipeServer) ReadMessage(conn net.Conn, buffer []byte) (int, error) {
	n, err := conn.Read(buffer)
	if err != nil {
		return 0, fmt.Errorf("failed to read from client: %w", err)
	}
	return n, nil
}

func (s *NamedPipeServer) SendMessage(conn net.Conn, message []byte) error {
	_, err := conn.Write(message)
	if err != nil {
		return fmt.Errorf("failed to send message to client: %w", err)
	}
	return nil
}

func (s *NamedPipeServer) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for conn := range s.connections {
		conn.Close()
		delete(s.connections, conn)
	}

	if s.listener != nil {
		return s.listener.Close()
	}

	return nil
}

type NamedPipeClient struct {
	conn net.Conn
	mu   sync.RWMutex
}

func NewNamedPipeClient(pipeAddress string) (*NamedPipeClient, error) {
	conn, err := net.Dial("tcp", pipeAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to named pipe: %w", err)
	}

	return &NamedPipeClient{
		conn: conn,
	}, nil
}

func (c *NamedPipeClient) ReadMessage(buffer []byte) (int, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	n, err := c.conn.Read(buffer)
	if err != nil {
		return 0, fmt.Errorf("failed to read from server: %w", err)
	}
	return n, nil
}

func (c *NamedPipeClient) SendMessage(message []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, err := c.conn.Write(message)
	if err != nil {
		return fmt.Errorf("failed to send message to server: %w", err)
	}
	return nil
}

func (c *NamedPipeClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

type RemoteConnection struct {
	serverSocket net.Conn
	listenSocket net.Listener
	clientSocket net.Conn
	isConnected  bool
	isServer     bool
	mu           sync.RWMutex
}

func NewRemoteConnection() *RemoteConnection {
	return &RemoteConnection{
		isConnected: false,
		isServer:    false,
	}
}

func (rc *RemoteConnection) Listen(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", port, err)
	}

	rc.listenSocket = listener
	rc.isServer = true

	rc.mu.Lock()
	rc.isConnected = true
	rc.mu.Unlock()

	return nil
}

func (rc *RemoteConnection) WaitForClient() error {
	conn, err := rc.listenSocket.Accept()
	if err != nil {
		return fmt.Errorf("failed to accept client connection: %w", err)
	}

	rc.clientSocket = conn
	return nil
}

func (rc *RemoteConnection) Connect(address string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", address, err)
	}

	rc.serverSocket = conn
	rc.isServer = false

	rc.mu.Lock()
	rc.isConnected = true
	rc.mu.Unlock()

	return nil
}

func (rc *RemoteConnection) SendMessage(message string) error {
	rc.mu.RLock()
	defer rc.mu.RUnlock()

	if !rc.isConnected {
		return fmt.Errorf("not connected")
	}

	var conn net.Conn
	if rc.isServer {
		conn = rc.clientSocket
	} else {
		conn = rc.serverSocket
	}

	if conn == nil {
		return fmt.Errorf("connection not established")
	}

	_, err := conn.Write([]byte(message + TcpEndOfBufferChars))
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (rc *RemoteConnection) ReceiveMessage(buffer []byte) (int, error) {
	rc.mu.RLock()
	defer rc.mu.RUnlock()

	if !rc.isConnected {
		return 0, fmt.Errorf("not connected")
	}

	var conn net.Conn
	if rc.isServer {
		conn = rc.clientSocket
	} else {
		conn = rc.serverSocket
	}

	if conn == nil {
		return 0, fmt.Errorf("connection not established")
	}

	conn.SetReadDeadline(time.Now().Add(30 * time.Second))
	n, err := conn.Read(buffer)
	if err != nil {
		return 0, fmt.Errorf("failed to receive message: %w", err)
	}

	return n, nil
}

func (rc *RemoteConnection) IsConnected() bool {
	rc.mu.RLock()
	defer rc.mu.RUnlock()
	return rc.isConnected
}

func (rc *RemoteConnection) Close() error {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	rc.isConnected = false

	if rc.serverSocket != nil {
		rc.serverSocket.Close()
		rc.serverSocket = nil
	}

	if rc.clientSocket != nil {
		rc.clientSocket.Close()
		rc.clientSocket = nil
	}

	if rc.listenSocket != nil {
		rc.listenSocket.Close()
		rc.listenSocket = nil
	}

	return nil
}

type TcpServer struct {
	listener    net.Listener
	connections map[net.Conn]bool
	mu          sync.RWMutex
}

func NewTcpServer(address string) (*TcpServer, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to create TCP server: %w", err)
	}

	return &TcpServer{
		listener:    listener,
		connections: make(map[net.Conn]bool),
	}, nil
}

func (s *TcpServer) Accept() (net.Conn, error) {
	conn, err := s.listener.Accept()
	if err != nil {
		return nil, fmt.Errorf("failed to accept connection: %w", err)
	}

	s.mu.Lock()
	s.connections[conn] = true
	s.mu.Unlock()

	return conn, nil
}

func (s *TcpServer) SendMessage(conn net.Conn, message []byte) error {
	_, err := conn.Write(message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

func (s *TcpServer) ReceiveMessage(conn net.Conn, buffer []byte) (int, error) {
	n, err := conn.Read(buffer)
	if err != nil {
		return 0, fmt.Errorf("failed to receive message: %w", err)
	}
	return n, nil
}

func (s *TcpServer) CloseConnection(conn net.Conn) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	conn.Close()
	delete(s.connections, conn)
	return nil
}

func (s *TcpServer) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for conn := range s.connections {
		conn.Close()
		delete(s.connections, conn)
	}

	if s.listener != nil {
		return s.listener.Close()
	}

	return nil
}

type TcpClient struct {
	conn net.Conn
	mu   sync.RWMutex
}

func NewTcpClient(address string) (*TcpClient, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to TCP server: %w", err)
	}

	return &TcpClient{
		conn: conn,
	}, nil
}

func (c *TcpClient) SendMessage(message []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, err := c.conn.Write(message)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}
	return nil
}

func (c *TcpClient) ReceiveMessage(buffer []byte) (int, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	n, err := c.conn.Read(buffer)
	if err != nil {
		return 0, fmt.Errorf("failed to receive message: %w", err)
	}
	return n, nil
}

func (c *TcpClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
