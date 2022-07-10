package physics

import "github.com/ByteArena/box2d"

type Physics struct {
	World *box2d.B2World
	Cfg   *Config
}

func NewPhysics(w *box2d.B2World, cfg *Config) *Physics {
	return &Physics{World: w, Cfg: cfg}
}

func (p *Physics) SimpleBody(x, y float64, bodyType uint8) *box2d.B2Body {
	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Position = box2d.MakeB2Vec2(x/p.Cfg.Scale, y/p.Cfg.Scale)
	bodyDef.Type = bodyType

	return p.World.CreateBody(&bodyDef)
}
