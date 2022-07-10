package physics

import (
	"github.com/ByteArena/box2d"
	"github.com/Gregmus2/simple-engine/common"
)

const velocityIterations = 8
const positionIterations = 2
const timeStep = 1.0 / 40

func NewWorld(cfg *Config) (*box2d.B2World, common.UpdateActionOut) {
	gravity := box2d.MakeB2Vec2(cfg.Gravity.X, cfg.Gravity.Y)
	world := box2d.MakeB2World(gravity)

	return &world, common.UpdateActionOut{
		Action: func() {
			world.Step(timeStep, velocityIterations, positionIterations)
		},
	}
}
