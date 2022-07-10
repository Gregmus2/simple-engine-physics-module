package physics

import (
	"github.com/ByteArena/box2d"
)

const velocityIterations = 8
const positionIterations = 2
const timeStep = 1.0 / 40

func NewWorld(cfg *Config) *box2d.B2World {
	gravity := box2d.MakeB2Vec2(cfg.Gravity.X, cfg.Gravity.Y)
	world := box2d.MakeB2World(gravity)

	//World.Step(timeStep, velocityIterations, positionIterations)

	return &world
}

/*
todo
	world SetContactListener in main to scene

*/
