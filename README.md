## Eucalyptus SDK for GO

A Eucalyptus SDK for Go using the [AWS SDK for Go](https://github.com/aws/aws-sdk-go-v2)

The SDK is used by the [Eucalyptus Cloud CLI](https://github.com/sjones4/euca)

### Generating Service APIs

On update of the JSON models service code should be regenerated:

    go generate ./service/

### Usage

To update to the latest SDK:

    go get -u github.com/sjones4/eucalyptus-sdk-go

To use the SDK first load configuration:

    import (
        "github.com/aws/aws-sdk-go-v2/aws"
        "github.com/aws/aws-sdk-go-v2/aws/external"
    )

    const (
        endpoint = "..."
        profile = "..."
        region = "eucalyptus"
    )

    cfg, _ := external.LoadDefaultAWSConfig(
        external.WithEndpointResolverFunc(func(aws.EndpointResolver) aws.EndpointResolver {
            return aws.ResolveWithEndpointURL(endpoint)
        }),
        external.WithSharedConfigProfile(profile),
        external.WithRegion(region),
    )

To enable debug logging:

    cfg.LogLevel = aws.LogDebugWithHTTPBody

With the configuration you can create a service client.

#### Properties

Example of creating a properties service client:

    import "github.com/sjones4/eucalyptus-sdk-go/service/euprop"
    
    const (
        DefaultEndpoint = "http://127.0.0.1:8773/services/Properties"
        MyEndpoint = "http://properties.cloud-10-10-10-101.euca.me:8773"
    )
    
    cfg := ...
    
    svc := euprop.New(cfg)

The example above does not include error handling.

#### Service Management

Example of creating a service management service client:

    import "github.com/sjones4/eucalyptus-sdk-go/service/euserv"
    
    const (
        DefaultEndpoint = "http://127.0.0.1:8773/services/Empyrean"
        MyEndpoint = "http://bootstrap.cloud-10-10-10-101.euca.me:8773"
    )
    
    cfg := ...
    
    svc := euserv.New(cfg)

The example above does not include error handling.

### Troubleshooting

If you encounter signature validation errors then you may need to remove a trailing / from the endpoint url.