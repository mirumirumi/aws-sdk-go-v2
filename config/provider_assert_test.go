// Code generated by github.com/aws/aws-sdk-go-v2/config. DO NOT EDIT.

package config

// apiOptionsProvider implementor assertions
var (
	_ apiOptionsProvider = &LoadOptions{}
)

// assumeRoleCredentialOptionsProvider implementor assertions
var (
	_ assumeRoleCredentialOptionsProvider = &LoadOptions{}
)

// clientLogModeProvider implementor assertions
var (
	_ clientLogModeProvider = &LoadOptions{}
)

// credentialsCacheOptionsProvider implementor assertions
var (
	_ credentialsCacheOptionsProvider = &LoadOptions{}
)

// credentialsProviderProvider implementor assertions
var (
	_ credentialsProviderProvider = &LoadOptions{}
)

// customCABundleProvider implementor assertions
var (
	_ customCABundleProvider = &EnvConfig{}
	_ customCABundleProvider = &LoadOptions{}
)

// defaultRegionProvider implementor assertions
var (
	_ defaultRegionProvider = &LoadOptions{}
)

// ec2IMDSRegionProvider implementor assertions
var (
	_ ec2IMDSRegionProvider = &LoadOptions{}
)

// ec2RoleCredentialOptionsProvider implementor assertions
var (
	_ ec2RoleCredentialOptionsProvider = &LoadOptions{}
)

// endpointCredentialOptionsProvider implementor assertions
var (
	_ endpointCredentialOptionsProvider = &LoadOptions{}
)

// endpointResolverProvider implementor assertions
var (
	_ endpointResolverProvider = &LoadOptions{}
)

// httpClientProvider implementor assertions
var (
	_ httpClientProvider = &LoadOptions{}
)

// logConfigurationWarningsProvider implementor assertions
var (
	_ logConfigurationWarningsProvider = &LoadOptions{}
)

// loggerProvider implementor assertions
var (
	_ loggerProvider = &LoadOptions{}
)

// processCredentialOptions implementor assertions
var (
	_ processCredentialOptions = &LoadOptions{}
)

// regionProvider implementor assertions
var (
	_ regionProvider = &EnvConfig{}
	_ regionProvider = &SharedConfig{}
	_ regionProvider = &LoadOptions{}
	_ regionProvider = &UseEC2IMDSRegion{}
)

// retryProvider implementor assertions
var (
	_ retryProvider = &LoadOptions{}
)

// sharedConfigFilesProvider implementor assertions
var (
	_ sharedConfigFilesProvider = &EnvConfig{}
	_ sharedConfigFilesProvider = &LoadOptions{}
)

// sharedConfigProfileProvider implementor assertions
var (
	_ sharedConfigProfileProvider = &EnvConfig{}
	_ sharedConfigProfileProvider = &LoadOptions{}
)

// webIdentityRoleCredentialOptionsProvider implementor assertions
var (
	_ webIdentityRoleCredentialOptionsProvider = &LoadOptions{}
)
