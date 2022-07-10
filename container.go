package physics

import "go.uber.org/dig"

func buildContainer(c *dig.Container) error {
	if err := c.Provide(NewConfig); err != nil {
		return err
	}
	if err := c.Provide(NewObjectFactory); err != nil {
		return err
	}
	if err := c.Provide(NewWorld); err != nil {
		return err
	}
	if err := c.Provide(NewPhysics); err != nil {
		return err
	}

	return nil
}
