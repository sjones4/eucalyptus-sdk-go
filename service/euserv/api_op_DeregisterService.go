// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package euserv

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sjones4/eucalyptus-sdk-go/internal/awsutil"
)

type DeregisterServiceInput struct {
	_ struct{} `type:"structure"`

	// Name is a required field
	Name *string `locationName:"Name" type:"string" required:"true"`

	Type *string `locationName:"Type" type:"string"`
}

// String returns the string representation
func (s DeregisterServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeregisterServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeregisterServiceInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeregisterServiceOutput struct {
	_ struct{} `locationName:"DeregisterServiceResponseType" type:"structure"`

	DeregisteredServices []ServiceId `locationName:"deregisteredServices" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DeregisterServiceOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeregisterService = "DeregisterService"

// DeregisterServiceRequest returns a request value for making API operation for
// Eucalyptus Cloud Service Management Service.
//
//    // Example sending a request using DeregisterServiceRequest.
//    req := client.DeregisterServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/euserv-2010-01-01/DeregisterService
func (c *Client) DeregisterServiceRequest(input *DeregisterServiceInput) DeregisterServiceRequest {
	op := &aws.Operation{
		Name:       opDeregisterService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeregisterServiceInput{}
	}

	req := c.newRequest(op, input, &DeregisterServiceOutput{})
	return DeregisterServiceRequest{Request: req, Input: input, Copy: c.DeregisterServiceRequest}
}

// DeregisterServiceRequest is the request type for the
// DeregisterService API operation.
type DeregisterServiceRequest struct {
	*aws.Request
	Input *DeregisterServiceInput
	Copy  func(*DeregisterServiceInput) DeregisterServiceRequest
}

// Send marshals and sends the DeregisterService API request.
func (r DeregisterServiceRequest) Send(ctx context.Context) (*DeregisterServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterServiceResponse{
		DeregisterServiceOutput: r.Request.Data.(*DeregisterServiceOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterServiceResponse is the response type for the
// DeregisterService API operation.
type DeregisterServiceResponse struct {
	*DeregisterServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterService request.
func (r *DeregisterServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
