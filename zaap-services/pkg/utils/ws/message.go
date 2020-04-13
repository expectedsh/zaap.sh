package ws

import "encoding/json"

type MessageType int

const (
	MessageTypeUnknown = 0

	MessageTypeApplicationDeployment = iota + 1 // type applications.Deployment
	MessageTypeSchedulerToken                   // type scheduler.Token
	MessageTypeDockerEvent                      // type events.message
	MessageTypeApplicationRunning
)

type Message struct {
	MessageType MessageType
	Payload     []byte
}

func NewMessage(messageType MessageType, payload interface{}) (*Message, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return &Message{MessageType: messageType, Payload: payloadBytes}, nil
}

func NewMessageRaw(messageType MessageType, payload []byte) (*Message, error) {
	return &Message{MessageType: messageType, Payload: payload}, nil
}

func NewMessageFromBytes(message []byte) (*Message, error) {
	m := Message{}
	if err := json.Unmarshal(message, &m); err != nil {
		return nil, err
	}
	return &m, nil
}
