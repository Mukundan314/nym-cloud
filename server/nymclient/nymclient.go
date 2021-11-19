package nymclient

import (
	"encoding/binary"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	sendRequestTag        = 0x00
	replyRequestTag       = 0x01
	selfAddressRequestTag = 0x02

	errorResponseTag       = 0x00
	receivedResponseTag    = 0x01
	selfAddressResponseTag = 0x02

	errCodeEmptyRequest     = 0x01
	errCodeTooShortRequest  = 0x02
	errCodeUnknownRequest   = 0x03
	errCodeMalformedRequest = 0x03

	errCodeEmptyResponse     = 0x80
	errCodeTooShortResponse  = 0x81
	errCodeUnknownResponse   = 0x82
	errCodeMalformedResponse = 0x83

	errCodeOther = 0xff
)

type Client struct {
	SelfAddress []byte
	conn        *websocket.Conn
	writeMutex  sync.Mutex
	readMutex   sync.Mutex
}

func New(clientUri string) (*Client, error) {
	conn, _, err := websocket.DefaultDialer.Dial(clientUri, nil)
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

func (c *Client) Send(recipient []byte, message []byte, withReplySurb bool) error {
	messageLen := make([]byte, 8)
	binary.BigEndian.PutUint64(messageLen, uint64(len(message)))

	surbByte := byte(0)
	if withReplySurb {
		surbByte = 1
	}

	data := []byte{sendRequestTag, surbByte}
	data = append(data, recipient...)
	data = append(data, messageLen...)
	data = append(data, message...)

	c.writeMutex.Lock()
	err := c.conn.WriteMessage(websocket.BinaryMessage, data)
	c.writeMutex.Unlock()

	return err
}

func (c *Client) Reply(replySURB []byte, message []byte) error {
	messageLen := make([]byte, 8)
	binary.BigEndian.PutUint64(messageLen, uint64(len(message)))

	surbLen := make([]byte, 8)
	binary.BigEndian.PutUint64(surbLen, uint64(len(replySURB)))

	data := []byte{replyRequestTag}
	data = append(data, surbLen...)
	data = append(data, replySURB...)
	data = append(data, messageLen...)
	data = append(data, message...)

	c.writeMutex.Lock()
	err := c.conn.WriteMessage(websocket.BinaryMessage, data)
	c.writeMutex.Unlock()

	return err
}

func (c *Client) Recv() {
	// c.readMutex.Lock()
	// messageType, rawResponse, err := c.conn.ReadMessage()
	// c.readMutex.Unlock()
}

func (c *Client) Close() error {
	return c.conn.Close()
}
