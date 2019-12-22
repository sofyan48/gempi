package config

import (
	"github.com/sofyan48/gempi/entity"
	"github.com/sofyan48/gempi/libs"
)

// Configs ...
type Configs struct {
	PathURL            string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
	AwsAPArea          string
	Backend            string
	Broker             string
}

// Configure call config entity
func Configure() *entity.AwsConfig {
	return &entity.AwsConfig{}
}

// NewConfig ...
func NewConfig() *Configs {
	return &Configs{}
}

// Credential client
func (cfg *Configs) Credential(awsCfg *entity.AwsConfig) *Configs {
	cfg.PathURL = awsCfg.PathURL
	cfg.AwsAccessKeyID = awsCfg.AwsAccessKeyID
	cfg.AwsSecretAccessKey = awsCfg.AwsSecretAccessKey
	cfg.AwsAPArea = awsCfg.APArea
	return cfg
}

// New create new config
func (cfg *Configs) New() *entity.NewClient {
	clients := &entity.NewClient{}
	awsLibs := &libs.Aws{}
	awsCfg := &entity.AwsConfig{}
	awsCfg.PathURL = cfg.PathURL
	awsCfg.AwsAccessKeyID = cfg.AwsAccessKeyID
	awsCfg.AwsSecretAccessKey = cfg.AwsSecretAccessKey
	awsCfg.APArea = cfg.AwsAPArea
	awsCfg.Backend = cfg.Backend
	awsCfg.Broker = cfg.Broker
	sqs := awsLibs.SQSession(awsCfg)
	clients.Sessions = sqs
	clients.Configs = awsCfg
	return clients
}
