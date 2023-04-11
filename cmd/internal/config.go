package internal

import (
	"flag"
)

type Configuration interface {
	UsesMetrics() bool
	UsesTracing() bool
	UsesProfiling() bool
	UsesKubeconfig() bool
	UsesPolicyExceptions() bool
	UsesConfigMapCaching() bool
	UsesCosign() bool
	UsesRegistryClient() bool
	FlagSets() []*flag.FlagSet
}

func NewConfiguration(options ...ConfigurationOption) Configuration {
	c := &configuration{}
	for _, option := range options {
		option(c)
	}
	return c
}

type ConfigurationOption func(c *configuration)

func WithMetrics() ConfigurationOption {
	return func(c *configuration) {
		c.usesMetrics = true
	}
}

func WithTracing() ConfigurationOption {
	return func(c *configuration) {
		c.usesTracing = true
	}
}

func WithProfiling() ConfigurationOption {
	return func(c *configuration) {
		c.usesProfiling = true
	}
}

func WithKubeconfig() ConfigurationOption {
	return func(c *configuration) {
		c.usesKubeconfig = true
	}
}

func WithPolicyExceptions() ConfigurationOption {
	return func(c *configuration) {
		c.usesPolicyExceptions = true
	}
}

func WithConfigMapCaching() ConfigurationOption {
	return func(c *configuration) {
		c.usesConfigMapCaching = true
	}
}

func WithCosign() ConfigurationOption {
	return func(c *configuration) {
		c.usesCosign = true
	}
}

func WithRegistryClient() ConfigurationOption {
	return func(c *configuration) {
		c.usesRegistryClient = true
	}
}

func WithFlagSets(flagsets ...*flag.FlagSet) ConfigurationOption {
	return func(c *configuration) {
		c.flagSets = append(c.flagSets, flagsets...)
	}
}

type configuration struct {
	usesMetrics          bool
	usesTracing          bool
	usesProfiling        bool
	usesKubeconfig       bool
	usesPolicyExceptions bool
	usesConfigMapCaching bool
	usesCosign           bool
	usesRegistryClient   bool
	flagSets             []*flag.FlagSet
}

func (c *configuration) UsesMetrics() bool {
	return c.usesMetrics
}

func (c *configuration) UsesTracing() bool {
	return c.usesTracing
}

func (c *configuration) UsesProfiling() bool {
	return c.usesProfiling
}

func (c *configuration) UsesKubeconfig() bool {
	return c.usesKubeconfig
}

func (c *configuration) UsesPolicyExceptions() bool {
	return c.usesPolicyExceptions
}

func (c *configuration) UsesConfigMapCaching() bool {
	return c.usesConfigMapCaching
}

func (c *configuration) UsesCosign() bool {
	return c.usesCosign
}

func (c *configuration) UsesRegistryClient() bool {
	return c.usesRegistryClient
}

func (c *configuration) FlagSets() []*flag.FlagSet {
	return c.flagSets
}
