package wire

import (
	"airbnb-auth-be/internal/pkg/env"
	"airbnb-auth-be/internal/pkg/env/tool"

	"github.com/google/wire"
)

var PackageSet = wire.NewSet(
	env.ProvideEnv,
	tool.ExtractServerConfig,
	tool.ExtractDBConfig,
	tool.ExtractOauthGoogleConfig,
	tool.ExtractOauthFacebookConfig,
	tool.ExtractKafkaConfig,
	tool.ExtractKafkaConsumerConfig,
	tool.ExtractKafkaRouterConfig,
	tool.ExtractServiceUserConfig,
)
