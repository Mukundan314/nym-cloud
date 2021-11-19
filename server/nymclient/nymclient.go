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
)

type NymClient struct {
	SelfAddress []byte
	conn        *websocket.Conn
	writeMutex  sync.Mutex
}

func NewNymClient(clientUri string) (*NymClient, error) {
	conn, _, err := websocket.DefaultDialer.Dial(clientUri, nil)
	if err != nil {
		return nil, err
	}

	return &NymClient{
		conn: conn,
	}, nil
}

func (n *NymClient) Send(recipient []byte, message []byte, withReplySurb bool) error {
	messageLen := make([]byte, 8)
	binary.BigEndian.PutUint64(messageLen, uint64(len(message)))

	surbByte := byte(0)
	if withReplySurb {
		surbByte = 1
	}

	data := []byte{sendRequestTag, surbByte}
	data = append(out, recipient...)
	data = append(out, messageLen...)
	data = append(out, message...)

	n.writeMutex.Lock()
	err := n.conn.WriteMessage(websocket.BinaryMessage, data)
	n.writeMutex.Unlock()

	return err
}

func (n *NymClient) Reply(replySURB []byte, message []byte) {
	messageLen := make([]byte, 8)
	binary.BigEndian.PutUint64(messageLen, uint64(len(message)))

	surbLen := make([]byte, 8)
	binary.BigEndian.PutUint64(surbLen, uint64(len(replySURB)))

	data := []byte{replyRequestTag}
	data = append(out, surbLen...)
	data = append(out, replySURB...)
	data = append(out, messageLen...)
	data = append(out, message...)

	n.writeMutex.Lock()
	err := n.conn.WriteMessage(websocket.BinaryMessage, data)
	n.writeMutex.Unlock()

	return err
}

func (n *NymClient) Recv() {
	// TODO
}

func (n *NymClient) Close() error {
	return n.conn.Close()
}
