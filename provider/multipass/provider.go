package multipass

import (
	stdcontext "context"
	"github.com/juju/clock"
	"github.com/juju/errors"
	"github.com/juju/jsonschema"
	"github.com/juju/juju/cloud"
	"github.com/juju/juju/environs"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/environs/context"
)

type environProvider struct {
	environs.ProviderCredentials
	environs.RequestFinalizeCredential
	environs.ProviderCredentialsRegister
	Clock clock.Clock
}

var mpCloud = cloud.Cloud{
	Name:        "multipass",
	Type:        "multipass",
	AuthTypes:   []cloud.AuthType{},
	Endpoint:    "",
	Regions:     []cloud.Region{},
	Description: cloud.DefaultCloudDescription("multipass"),
}

func (p *environProvider) DetectCloud(name string) (cloud.Cloud, error) {
	return mpCloud, nil
}

func (p *environProvider) DetectClouds() ([]cloud.Cloud, error) {
	return []cloud.Cloud{mpCloud}, nil
}

func (p *environProvider) Validate(cfg, old *config.Config) (valid *config.Config, _ error) {
	//TODO implement me
	return nil, errors.New("Validate FAILED")
}

func (p *environProvider) Version() int {
	return 0
}

func (p *environProvider) CloudSchema() *jsonschema.Schema {
	//TODO implement me
	return &jsonschema.Schema{}
}

func (p *environProvider) Ping(ctx context.ProviderCallContext, endpoint string) error {
	//TODO implement me
	return errors.New("Ping FAILED")
}

func (p *environProvider) PrepareConfig(params environs.PrepareConfigParams) (*config.Config, error) {
	//TODO implement me
	return nil, errors.New("PrepareConfig FAILED")
}

func (p *environProvider) Open(_ stdcontext.Context, params environs.OpenParams) (environs.Environ, error) {
	//TODO implement me
	return nil, errors.New("Open FAILED")
}

func NewProvider() environs.CloudEnvironProvider {
	return &environProvider{}
}

func (p *environProvider) CredentialSchemas() map[cloud.AuthType]cloud.CredentialSchema {
	return map[cloud.AuthType]cloud.CredentialSchema{}
}

func (p *environProvider) DetectCredentials(cloudName string) (*cloud.CloudCredential, error) {
	return nil, errors.New("DetectCredentials FAILED")
}

func (p *environProvider) FinalizeCredential(
	environs.FinalizeCredentialContext,
	environs.FinalizeCredentialParams,
) (*cloud.Credential, error) {
	return nil, errors.New("FinalizeCredential FAILED")
}
