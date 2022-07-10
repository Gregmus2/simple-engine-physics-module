package physics

import (
	"github.com/creasty/defaults"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type physicsConfig struct {
	Physics Config
}

type Config struct {
	Scale   float64 `default:"100"`
	Gravity struct {
		X float64 `default:"0"`
		Y float64 `default:"-5"`
	}
}

func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(c); err != nil {
		return errors.Wrap(err, "error setting defaults")
	}

	type plain Config
	if err := unmarshal((*plain)(c)); err != nil {
		return errors.Wrap(err, "error unmarshalling")
	}

	return nil
}

func NewConfig() (*Config, error) {
	cfg := &physicsConfig{}

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if !os.IsNotExist(err) {
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(yamlFile, cfg)
		if err != nil {
			return nil, err
		}
	}

	return &cfg.Physics, nil
}
