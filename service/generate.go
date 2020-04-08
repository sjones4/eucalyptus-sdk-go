// Package service contains automatically generated AWS clients.
package service

//go:generate go run -tags codegen ../private/model/cli/gen-api/main.go -path=../service ../models/apis/*/*/api-2.json
//go:generate sed --in-place s/github.com\/aws\/aws-sdk-go-v2\/internal\/awsutil/github.com\/sjones4\/eucalyptus-sdk-go\/internal\/awsutil/ ../service/euprop/api_types.go ../service/euprop/api_op_DescribeProperties.go ../service/euprop/api_op_ModifyPropertyValue.go
//go:generate gofmt -s -w ../service
