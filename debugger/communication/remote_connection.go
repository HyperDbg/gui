package communication

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	TcpEndOfBufferCharsCount = 4
	CommunicationBufferSize  = 0x100000
)

var endOfBufferCheckTcp = [TcpEndOfBufferCharsCount]byte{0x00, 0x0F, 0xFF, 0x00}

type ConnectionState int

const (
	StateDisconnected ConnectionState = iota
	StateConnectedAsDebugger
	StateConnectedAsDebuggee
)

type RemoteConnection struct {
	mu                        sync.Mutex
	state                     ConnectionState
	serverSocket              net.Listener
	clientSocket              net.Conn
	isEndOfMessageReceived    bool
	endOfMessageReceivedEvent chan struct{}
	onReceiveMessage          func([]byte) bool
	onHandshakeComplete       func(string) bool
	buildSignature            string
}

func NewRemoteConnection(buildSignature string) *RemoteConnection {
	return &RemoteConnection{
		state:                     StateDisconnected,
		endOfMessageReceivedEvent: make(chan struct{}, 1),
		buildSignature:            buildSignature,
	}
}

func (rc *RemoteConnection) IsConnected() bool {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.state == StateConnectedAsDebugger || rc.state == StateConnectedAsDebuggee
}

func (rc *RemoteConnection) IsConnectedToDebuggee() bool {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.state == StateConnectedAsDebugger
}

func (rc *RemoteConnection) IsConnectedToDebugger() bool {
	rc.mu.Lock()
	defer rc.mu.Unlock()
	return rc.state == StateConnectedAsDebuggee
}

func (rc *RemoteConnection) Listen(port string) error {
	rc.mu.Lock()
	if rc.state != StateDisconnected {
		rc.mu.Unlock()
		return fmt.Errorf("already connected")
	}
	rc.mu.Unlock()

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", port, err)
	}

	rc.mu.Lock()
	rc.serverSocket = listener
	rc.mu.Unlock()

	conn, err := listener.Accept()
	if err != nil {
		listener.Close()
		return fmt.Errorf("failed to accept connection: %w", err)
	}

	rc.mu.Lock()
	rc.clientSocket = conn
	rc.mu.Unlock()

	buf := make([]byte, CommunicationBufferSize)
	n, err := conn.Read(buf)
	if err != nil || n == 0 {
		conn.Close()
		rc.mu.Lock()
		rc.serverSocket.Close()
		rc.serverSocket = nil
		rc.clientSocket = nil
		rc.mu.Unlock()
		return fmt.Errorf("handshake failed: %w", err)
	}

	receivedSignature := string(buf[:n])
	if receivedSignature != rc.buildSignature {
		conn.Write([]byte("NO"))
		conn.Close()
		rc.mu.Lock()
		rc.serverSocket.Close()
		rc.serverSocket = nil
		rc.clientSocket = nil
		rc.mu.Unlock()
		return fmt.Errorf("build signature mismatch")
	}

	conn.Write([]byte("OK"))

	rc.mu.Lock()
	rc.state = StateConnectedAsDebugger
	rc.mu.Unlock()

	go rc.listenForMessages()

	return nil
}

func (rc *RemoteConnection) Connect(ip, port string) error {
	rc.mu.Lock()
	if rc.state != StateDisconnected {
		rc.mu.Unlock()
		return fmt.Errorf("already connected")
	}
	rc.mu.Unlock()

	addr := net.JoinHostPort(ip, port)
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", addr, err)
	}

	_, err = conn.Write([]byte(rc.buildSignature))
	if err != nil {
		conn.Close()
		return fmt.Errorf("handshake send failed: %w", err)
	}

	response := make([]byte, 2)
	n, err := conn.Read(response)
	if err != nil || n < 2 || string(response[:n]) != "OK" {
		conn.Close()
		return fmt.Errorf("handshake rejected by server")
	}

	rc.mu.Lock()
	rc.clientSocket = conn
	rc.state = StateConnectedAsDebuggee
	rc.mu.Unlock()

	go rc.listenForMessages()

	return nil
}

func (rc *RemoteConnection) listenForMessages() {
	buf := make([]byte, CommunicationBufferSize)
	for {
		rc.mu.Lock()
		conn := rc.clientSocket
		rc.mu.Unlock()

		if conn == nil {
			break
		}

		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		if n > 0 && rc.onReceiveMessage != nil {
			rc.onReceiveMessage(buf[:n])
		}
	}

	rc.mu.Lock()
	rc.state = StateDisconnected
	if rc.clientSocket != nil {
		rc.clientSocket.Close()
		rc.clientSocket = nil
	}
	if rc.serverSocket != nil {
		rc.serverSocket.Close()
		rc.serverSocket = nil
	}
	rc.mu.Unlock()
}

func (rc *RemoteConnection) SendMessage(message []byte) error {
	rc.mu.Lock()
	conn := rc.clientSocket
	rc.mu.Unlock()

	if conn == nil {
		return fmt.Errorf("not connected")
	}

	_, err := conn.Write(message)
	return err
}

func (rc *RemoteConnection) ReceiveMessage(buffer []byte) (int, error) {
	rc.mu.Lock()
	conn := rc.clientSocket
	rc.mu.Unlock()

	if conn == nil {
		return 0, fmt.Errorf("not connected")
	}

	return conn.Read(buffer)
}

func (rc *RemoteConnection) ShutdownAndCleanup() {
	rc.mu.Lock()
	defer rc.mu.Unlock()

	if rc.clientSocket != nil {
		rc.clientSocket.Close()
		rc.clientSocket = nil
	}
	if rc.serverSocket != nil {
		rc.serverSocket.Close()
		rc.serverSocket = nil
	}
	rc.state = StateDisconnected
}

func CheckForTheEndOfTcpBuffer(currentLoopIndex *int, buffer []byte) bool {
	actualLen := *currentLoopIndex

	if actualLen <= 3 {
		return false
	}

	if buffer[actualLen] == endOfBufferCheckTcp[3] &&
		buffer[actualLen-1] == endOfBufferCheckTcp[2] &&
		buffer[actualLen-2] == endOfBufferCheckTcp[1] &&
		buffer[actualLen-3] == endOfBufferCheckTcp[0] {

		buffer[actualLen-3] = 0
		buffer[actualLen-2] = 0
		buffer[actualLen-1] = 0
		buffer[actualLen] = 0

		*currentLoopIndex = actualLen - 3
		return true
	}
	return false
}
