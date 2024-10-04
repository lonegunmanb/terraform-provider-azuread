package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskProfile struct {
	// This is the count of human identities that have been assigned to this riskScoreBracket,
	HumanCount *int64 `json:"humanCount,omitempty"`

	// This is the count of nonhuman identities that have been assigned to this riskScoreBracket
	NonHumanCount *int64 `json:"nonHumanCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}