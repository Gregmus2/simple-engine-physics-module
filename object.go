package physics

import (
	"github.com/ByteArena/box2d"
)

type Object struct {
	Body    *box2d.B2Body
	Fixture *box2d.B2Fixture
}

type ObjectFactory struct {
	Cfg *Config
}

func NewObjectFactory(cfg *Config) *ObjectFactory {
	return &ObjectFactory{Cfg: cfg}
}

func (o *ObjectFactory) NewObject(body *box2d.B2Body, shape box2d.B2CircleShape, density float64) *Object {
	return &Object{
		Body:    body,
		Fixture: body.CreateFixture(&shape, density),
	}
}
