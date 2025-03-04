# v1.21.8 (2022-09-02)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.7 (2022-08-31)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.6 (2022-08-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.5 (2022-08-24)

* **Documentation**: Documentation updates for AWS Transfer Family

# v1.21.4 (2022-08-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.3 (2022-08-09)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.2 (2022-08-08)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.1 (2022-08-01)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.21.0 (2022-07-26)

* **Feature**: AWS Transfer Family now supports Applicability Statement 2 (AS2), a network protocol used for the secure and reliable transfer of critical Business-to-Business (B2B) data over the public internet using HTTP/HTTPS as the transport mechanism.

# v1.20.2 (2022-07-05)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.1 (2022-06-29)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.20.0 (2022-06-22)

* **Feature**: Until today, the service supported only RSA host keys and user keys. Now with this launch, Transfer Family has expanded the support for ECDSA and ED25519 host keys and user keys, enabling customers to support a broader set of clients by choosing RSA, ECDSA, and ED25519 host and user keys.

# v1.19.1 (2022-06-07)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.19.0 (2022-05-18)

* **Feature**: AWS Transfer Family now supports SetStat server configuration option, which provides the ability to ignore SetStat command issued by file transfer clients, enabling customers to upload files without any errors.

# v1.18.7 (2022-05-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.6 (2022-05-12)

* **Documentation**: AWS Transfer Family now accepts ECDSA keys for server host keys

# v1.18.5 (2022-04-25)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.4 (2022-04-19)

* **Documentation**: This release contains corrected HomeDirectoryMappings examples for several API functions: CreateAccess, UpdateAccess, CreateUser, and UpdateUser,.

# v1.18.3 (2022-03-30)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.2 (2022-03-24)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.1 (2022-03-23)

* **Documentation**: Documentation updates for AWS Transfer Family to describe how to remove an associated workflow from a server.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.18.0 (2022-03-10)

* **Feature**: Adding more descriptive error types for managed workflows

# v1.17.0 (2022-03-08)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Feature**: Updated service client model to latest release.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.16.0 (2022-02-24)

* **Feature**: API client updated
* **Feature**: Adds RetryMaxAttempts and RetryMod to API client Options. This allows the API clients' default Retryer to be configured from the shared configuration files or environment variables. Adding a new Retry mode of `Adaptive`. `Adaptive` retry mode is an experimental mode, adding client rate limiting when throttles reponses are received from an API. See [retry.AdaptiveMode](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/aws/retry#AdaptiveMode) for more details, and configuration options.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.15.0 (2022-01-14)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.14.0 (2022-01-07)

* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.13.0 (2021-12-21)

* **Feature**: API Paginators now support specifying the initial starting token, and support stopping on empty string tokens.
* **Feature**: API client updated

# v1.12.1 (2021-12-02)

* **Bug Fix**: Fixes a bug that prevented aws.EndpointResolverWithOptions from being used by the service client. ([#1514](https://github.com/aws/aws-sdk-go-v2/pull/1514))
* **Dependency Update**: Updated to the latest SDK module versions

# v1.12.0 (2021-11-19)

* **Feature**: API client updated
* **Dependency Update**: Updated to the latest SDK module versions

# v1.11.0 (2021-11-12)

* **Feature**: Service clients now support custom endpoints that have an initial URI path defined.

# v1.10.0 (2021-11-06)

* **Feature**: The SDK now supports configuration of FIPS and DualStack endpoints using environment variables, shared configuration, or programmatically.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.9.0 (2021-10-21)

* **Feature**: Updated  to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.1 (2021-10-11)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.8.0 (2021-09-30)

* **Feature**: API client updated

# v1.7.1 (2021-09-17)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.7.0 (2021-09-02)

* **Feature**: API client updated

# v1.6.0 (2021-08-27)

* **Feature**: Updated API model to latest revision.
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.3 (2021-08-19)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.2 (2021-08-04)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version.
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.1 (2021-07-15)

* **Dependency Update**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.5.0 (2021-06-25)

* **Feature**: API client updated
* **Feature**: Updated `github.com/aws/smithy-go` to latest version
* **Dependency Update**: Updated to the latest SDK module versions

# v1.4.1 (2021-06-11)

* **Documentation**: Updated to latest API model.

# v1.4.0 (2021-05-25)

* **Feature**: API client updated

# v1.3.1 (2021-05-20)

* **Dependency Update**: Updated to the latest SDK module versions

# v1.3.0 (2021-05-14)

* **Feature**: Constant has been added to modules to enable runtime version inspection for reporting.
* **Dependency Update**: Updated to the latest SDK module versions

