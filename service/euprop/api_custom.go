// Copyright (c) 2020 Steve Jones
// SPDX-License-Identifier: BSD-2-Clause

// Generated API client customizations
package euprop

import (
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/sjones4/eucalyptus-sdk-go/private/protocol/euquery"
)

func init() {
	initClient = clientInitializer
}

func clientInitializer(client *Client) {
	client.Handlers.UnmarshalError.Remove(query.UnmarshalErrorHandler)
	client.Handlers.UnmarshalError.PushBackNamed(euquery.UnmarshalErrorHandler)
}
