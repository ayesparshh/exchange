package internal

import (
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

type PubSub struct {
	conn *nats.EncodedConn
}

func NewPubSub(addr, user, pass string) PubSub {

	conn, err := nats.Connect(addr, nats.UserInfo(user, pass))
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to nats")
	}

	enc, err := nats.NewEncodedConn(conn, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating encoded connection")
	}

	return PubSub{enc}
}

func (p *PubSub) Subscribe(event string, handler any) *nats.Subscription {
	sub, err := p.conn.Subscribe(event, handler)

	if err != nil {
		log.Error().Err(err).Str("event", event).Msg("PubSub.Subscribe")
	}

	return sub
}

func (p *PubSub) Publish(event string, payload any) {
	err := p.conn.Publish(event, payload)

	if err != nil {
		log.Error().Err(err).Str("event", event).Msg("PubSub.Publish")
	}
}

func (p PubSub) JetStream() nats.JetStream {
	js, err := p.conn.Conn.JetStream()

	if err != nil {
		log.Error().Err(err).Msg("PubSub.JetStream")
	}

	return js
}

func (p PubSub) Close() {
	p.conn.Close()
}
