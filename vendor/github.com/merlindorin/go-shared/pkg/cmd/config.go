package cmd

import (
	"fmt"
	"path"

	"github.com/alecthomas/kong"
	kongyaml "github.com/alecthomas/kong-yaml"
)

type Option func(c *Config)

func (o Option) Apply(c *Config) {
	o(c)
}

func WithGroup(group string) Option {
	return func(c *Config) {
		c.group = group
	}
}

type Config struct {
	// name application config name
	name string

	// group application group name
	group string

	ConfigFile kong.ConfigFlag `json:"config" name:"config" short:"c" help:"Full path to a user-supplied config file"`
}

func NewConfig(name string, opts ...Option) *Config {
	c := &Config{name: name}

	for _, opt := range opts {
		opt.Apply(c)
	}

	return c
}

func (c *Config) BeforeResolve(k *kong.Kong) error {
	if c.name == "" {
		return fmt.Errorf("must specify an application name")
	}

	etcFileName := path.Join("/etc", c.group, fmt.Sprintf("%s.yaml", c.name))
	homeFileName := path.Join("~", fmt.Sprintf(".%s", c.group), fmt.Sprintf("%s.yaml", c.name))

	if c.group == "" {
		etcFileName = path.Join("/etc", c.name, "config.yaml")
		homeFileName = path.Join("~", fmt.Sprintf(".%s", c.name), "config.yaml")
	}

	return kong.Configuration(kongyaml.Loader, etcFileName, homeFileName).Apply(k)
}
