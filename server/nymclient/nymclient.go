package nymclient

import (
	"encoding/binary"
	"errors"
	"fmt"
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

	ErrCodeEmptyRequest     = 0x01
	ErrCodeTooShortRequest  = 0x02
	ErrCodeUnknownRequest   = 0x03
	ErrCodeMalformedRequest = 0x03

	ErrCodeEmptyResponse     = 0x80
	ErrCodeTooShortResponse  = 0x81
	ErrCodeUnknownResponse   = 0x82
	ErrCodeMalformedResponse = 0x83

	ErrCodeOther = 0xff
)

type Client struct {
	SelfAddress []byte
	conn        *websocket.Conn
	writeMutex  sync.Mutex
	readMutex   sync.Mutex
}

type Address []byte

type ErrorResponse struct {
	Code    byte
	Message []byte
}

type ReceivedResponse struct {
	Message []byte
	Surb    []byte
}

type Response struct {
	Type        string
	SelfAddress Address
	Error       *ErrorResponse
	Received    *ReceivedResponse
}

func parseErrorResponse(rawResponse []byte) (errCode byte, msg []byte, err error) {
	if len(rawResponse) < 10 {
		err = errors.New("Length of error response is too short")
		return
	}

	errCode = rawResponse[1]
	msgLen := int(binary.BigEndian.Uint64(rawResponse[2:10]))
	msg = rawResponse[10:]

	if len(msg) != msgLen {
		err = errors.New("Length of message does not match msgLen")
		return
	}

	return
}

func parseReceivedResponse(rawResponse []byte) (msg []byte, surb []byte, err error) {
	if len(rawResponse) < 10 {
		err = errors.New("Length of received response is too short")
		return
	}

	withReply := rawResponse[1]
	if withReply > 1 {
		err = fmt.Errorf("Invald value for withReply received: %#02x", withReply)
		return
	}

	if withReply == 1 {
		surbLen := int(binary.BigEndian.Uint64(rawResponse[2:10]))
		surb = rawResponse[10 : 10+surbLen]
		if len(surb) != surbLen {
			err = errors.New("Length of surb does not match surbLen")
			return
		}

		if len(rawResponse) < 18+surbLen {
			err = errors.New("Length of received response is too short")
			return
		}

		msgLen := int(binary.BigEndian.Uint64(rawResponse[10+surbLen : 18+surbLen]))
		msg = rawResponse[18+surbLen:]
		if len(msg) != msgLen {
			err = errors.New("Length of msg does not match msgLen")
			return
		}
	}

	msgLen := int(binary.BigEndian.Uint64(rawResponse[2:10]))
	msg = rawResponse[10:]
	if len(msg) != msgLen {
		err = errors.New("Length of msg does not match msgLen")
		return
	}

	return
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

func (c *Client) Recv() (*Response, error) {
	c.readMutex.Lock()
	messageType, rawResponse, err := c.conn.ReadMessage()
	c.readMutex.Unlock()

	if err != nil {
		return nil, err
	}

	if messageType != websocket.BinaryMessage {
		return nil, errors.New("Recieved response is not of type binary")
	}

	if len(rawResponse) == 0 {
		return nil, errors.New("Received response is empty")
	}

	switch responseTag := rawResponse[0]; responseTag {
	case errorResponseTag:
		errCode, msg, err := parseErrorResponse(rawResponse)
		if err != nil {
			return nil, err
		}
		return &Response{
			Type: "error",
			Error: &ErrorResponse{
				Code:    errCode,
				Message: msg,
			},
		}, nil
	case receivedResponseTag:
		msg, surb, err := parseReceivedResponse(rawResponse)
		if err != nil {
			return nil, err
		}
		return &Response{
			Type: "received",
			Received: &ReceivedResponse{
				Message: msg,
				Surb:    surb,
			},
		}, nil
	case selfAddressResponseTag:
		return &Response{
			Type:        "selfAddress",
			SelfAddress: rawResponse[1:],
		}, nil
	default:
		return nil, fmt.Errorf("Unknown response tag received: %#02x", responseTag)
	}
}

func (c *Client) Close() error {
	return c.conn.Close()
}
