package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesContentApplicabilitySettings struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Offer if the update is recommended by a vendor in the list, otherwise withhold the offer.
	OfferWhileRecommendedBy *[]string `json:"offerWhileRecommendedBy,omitempty"`

	// Settings for governing safeguard-holds on offering content.
	Safeguard *WindowsUpdatesSafeguardSettings `json:"safeguard,omitempty"`
}