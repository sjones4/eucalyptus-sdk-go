// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

// Eucalyptus query binding inconsistencies
package euquery

import (
	"encoding/xml"
	request "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"io/ioutil"
)

// Eucalyptus error response has non-standard document element and Errors wrapper
type xmlErrorResponse struct {
	XMLName   xml.Name `xml:"Response"`
	Code      string   `xml:"Errors>Error>Code"`
	Message   string   `xml:"Errors>Error>Message"`
	RequestID string   `xml:"RequestId"`
}

var UnmarshalErrorHandler = request.NamedHandler{Name: "euca.query.UnmarshalError", Fn: UnmarshalError}

// UnmarshalError unmarshals an error response for an non-standard Eucalyptus Query service.
func UnmarshalError(r *request.Request) {
	defer r.HTTPResponse.Body.Close()

	bodyBytes, err := ioutil.ReadAll(r.HTTPResponse.Body)
	if err != nil {
		r.Error = awserr.New("SerializationError", "failed to read from query HTTP response body", err)
		return
	}

	// First check for specific error
	resp := xmlErrorResponse{}
	decodeErr := xml.Unmarshal(bodyBytes, &resp)
	if decodeErr == nil {
		reqID := resp.RequestID
		if reqID == "" {
			reqID = r.RequestID
		}
		r.Error = awserr.NewRequestFailure(
			awserr.New(resp.Code, resp.Message, nil),
			r.HTTPResponse.StatusCode,
			reqID,
		)
		return
	}

	// Check for unhandled error
	if r.HTTPResponse.StatusCode == 503 {
		r.Error = awserr.NewRequestFailure(
			awserr.New("ServiceUnavailableException", "service is unavailable", nil),
			r.HTTPResponse.StatusCode,
			r.RequestID,
		)
		return
	}

	// Failed to retrieve any error message from the response body
	r.Error = awserr.New("SerializationError",
		"failed to decode query XML error response", decodeErr)
}
