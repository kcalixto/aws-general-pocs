package sesClient

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/osteele/liquid"
	"github.com/pkg/errors"
)

type ExternalSESConfig struct {
	ExternalID      string `json:"external_id"`
	RoleArn         string `json:"role_arn"`
	RoleSessionName string `json:"role_session_name"`
	Region          string `json:"region"`
}

//go:generate mockery --name SESClient --output ./mocks
type SESClient interface {
	SendMail(ctx context.Context, input *sesv2.SendEmailInput) (*sesv2.SendEmailOutput, error)
	// SendTemplatedEmail(input *sesv2.SendTemplatedEmailInput) (*sesv2.SendTemplatedEmailOutput, error)
	ReadAndParseTemplate(ctx context.Context, templateName string, content map[string]any) (string, error)
}

type sesClient struct {
	client *sesv2.Client
	engine *liquid.Engine
}

func NewSESClient(ctx context.Context, optionalExternalSesConfig ...ExternalSESConfig) (SESClient, error) {
	engine := liquid.NewEngine()

	externalSesConfig := ExternalSESConfig{}
	if len(optionalExternalSesConfig) > 0 {
		externalSesConfig = optionalExternalSesConfig[0]
	}
	if externalSesConfig.RoleArn == "" {
		region := endpoints.SaEast1RegionID
		if externalSesConfig.Region != "" {
			region = externalSesConfig.Region
		}

		cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region),
			config.WithRetryer(func() aws.Retryer {
				return aws.NopRetryer{}
			}))
		if err != nil {
			return nil, err
		}

		return &sesClient{
			client: sesv2.NewFromConfig(cfg),
			engine: engine,
		}, nil
	}

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(externalSesConfig.Region),
		config.WithRetryer(func() aws.Retryer {
			return aws.NopRetryer{}
		}),
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load SDK config")
	}

	stsClient := sts.NewFromConfig(cfg)
	provider := stscreds.NewAssumeRoleProvider(stsClient, externalSesConfig.RoleArn, func(opt *stscreds.AssumeRoleOptions) {
		opt.ExternalID = &externalSesConfig.ExternalID
		opt.RoleSessionName = externalSesConfig.RoleSessionName
	})

	cfg.Credentials = aws.NewCredentialsCache(provider)

	cfgExternalRegion, err := config.LoadDefaultConfig(ctx, config.WithRegion(externalSesConfig.Region),
		config.WithCredentialsProvider(
			provider,
		),
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load SDK config")
	}

	return &sesClient{
		client: sesv2.NewFromConfig(cfgExternalRegion),
		engine: engine,
	}, nil
}

func (s *sesClient) SendMail(ctx context.Context, input *sesv2.SendEmailInput) (*sesv2.SendEmailOutput, error) {
	return s.client.SendEmail(ctx, input)
}

func (s *sesClient) ReadAndParseTemplate(ctx context.Context, templateName string, bindings map[string]any) (string, error) {
	input := &sesv2.GetEmailTemplateInput{
		TemplateName: &templateName,
	}

	result, err := s.client.GetEmailTemplate(ctx, input)
	if err != nil {
		return "", err
	}
	if result == nil {
		return "", errors.New("template not found")
	}

	parsedMessage, err := s.engine.ParseAndRenderString(*result.TemplateContent.Html, bindings)
	if err != nil {
		return "", err
	}

	return parsedMessage, nil
}
