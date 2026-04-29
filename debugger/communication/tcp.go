package communication

import (
	"fmt"
	"net"
	"time"
)

type TcpClient struct {
	conn net.Conn
}

type TcpServer struct {
	listener net.Listener
	client   net.Conn
}

func NewTcpClient() *TcpClient {
	return &TcpClient{}
}

func (tc *TcpClient) ConnectToServer(ip string, port string) error {
	address := net.JoinHostPort(ip, port)

	var err error
	tc.conn, err = net.DialTimeout("tcp", address, 10*time.Second)
	if err != nil {
		return fmt.Errorf("unable to connect to the server: %w", err)
	}

	return nil
}

func (tc *TcpClient) SendMessage(sendBuf []byte) error {
	if tc.conn == nil {
		return fmt.Errorf("connection is not established")
	}

	_, err := tc.conn.Write(sendBuf)
	if err != nil {
		return fmt.Errorf("send failed: %w", err)
	}

	return nil
}

func (tc *TcpClient) ReceiveMessage(maxBuffLen int) ([]byte, error) {
	if tc.conn == nil {
		return nil, fmt.Errorf("connection is not established")
	}

	recvBuf := make([]byte, maxBuffLen)
	n, err := tc.conn.Read(recvBuf)
	if err != nil {
		return nil, fmt.Errorf("receive failed: %w", err)
	}

	if n == 0 {
		return nil, nil
	}

	return recvBuf[:n], nil
}

func (tc *TcpClient) ShutdownConnection() error {
	if tc.conn == nil {
		return nil
	}

	err := tc.conn.Close()
	if err != nil {
		return fmt.Errorf("shutdown failed: %w", err)
	}

	tc.conn = nil
	return nil
}

func (tc *TcpClient) IsConnected() bool {
	return tc.conn != nil
}

func NewTcpServer() *TcpServer {
	return &TcpServer{}
}

func (ts *TcpServer) CreateServerAndWaitForClient(port string) error {
	var err error
	ts.listener, err = net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("listen failed: %w", err)
	}

	fmt.Printf("listening on port %s...\n", port)

	ts.client, err = ts.listener.Accept()
	if err != nil {
		ts.listener.Close()
		return fmt.Errorf("accept failed: %w", err)
	}

	remoteAddr := ts.client.RemoteAddr().String()
	fmt.Printf("connected to : %s\n", remoteAddr)

	return nil
}

func (ts *TcpServer) ReceiveMessage(recvBufLen int) ([]byte, error) {
	if ts.client == nil {
		return nil, fmt.Errorf("no client connected")
	}

	recvBuf := make([]byte, recvBufLen)
	n, err := ts.client.Read(recvBuf)
	if err != nil {
		return nil, fmt.Errorf("recv failed: %w", err)
	}

	if n == 0 {
		return nil, nil
	}

	return recvBuf[:n], nil
}

func (ts *TcpServer) SendMessage(sendBuf []byte) error {
	if ts.client == nil {
		return fmt.Errorf("no client connected")
	}

	_, err := ts.client.Write(sendBuf)
	if err != nil {
		return fmt.Errorf("send failed: %w", err)
	}

	return nil
}

func (ts *TcpServer) Close() error {
	var errs []error

	if ts.client != nil {
		if err := ts.client.Close(); err != nil {
			errs = append(errs, err)
		}
		ts.client = nil
	}

	if ts.listener != nil {
		if err := ts.listener.Close(); err != nil {
			errs = append(errs, err)
		}
		ts.listener = nil
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors during close: %v", errs)
	}

	return nil
}

func (ts *TcpServer) HasClient() bool {
	return ts.client != nil
}
