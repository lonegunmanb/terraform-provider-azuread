package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId{}

// IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId is a struct representing the Resource ID for a Identity Governance Entitlement Management Access Package Assignment Id Access Package Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Scope
type IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId struct {
	AccessPackageAssignmentId        string
	AccessPackageResourceRoleScopeId string
	AccessPackageResourceScopeId     string
}

// NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeID returns a new IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId struct
func NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeID(accessPackageAssignmentId string, accessPackageResourceRoleScopeId string, accessPackageResourceScopeId string) IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId {
	return IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId{
		AccessPackageAssignmentId:        accessPackageAssignmentId,
		AccessPackageResourceRoleScopeId: accessPackageResourceRoleScopeId,
		AccessPackageResourceScopeId:     accessPackageResourceScopeId,
	}
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeID parses 'input' into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeID(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeIDInsensitively(input string) (*IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessPackageAssignmentId, ok = input.Parsed["accessPackageAssignmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageAssignmentId", input)
	}

	if id.AccessPackageResourceRoleScopeId, ok = input.Parsed["accessPackageResourceRoleScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceRoleScopeId", input)
	}

	if id.AccessPackageResourceScopeId, ok = input.Parsed["accessPackageResourceScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessPackageResourceScopeId", input)
	}

	return nil
}

// ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeID checks that 'input' can be parsed as a Identity Governance Entitlement Management Access Package Assignment Id Access Package Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Scope ID
func ValidateIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Entitlement Management Access Package Assignment Id Access Package Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId) ID() string {
	fmtString := "/identityGovernance/entitlementManagement/accessPackageAssignments/%s/accessPackage/accessPackageResourceRoleScopes/%s/accessPackageResourceScope/accessPackageResource/accessPackageResourceScopes/%s"
	return fmt.Sprintf(fmtString, id.AccessPackageAssignmentId, id.AccessPackageResourceRoleScopeId, id.AccessPackageResourceScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Entitlement Management Access Package Assignment Id Access Package Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("entitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("accessPackageAssignments", "accessPackageAssignments", "accessPackageAssignments"),
		resourceids.UserSpecifiedSegment("accessPackageAssignmentId", "accessPackageAssignmentId"),
		resourceids.StaticSegment("accessPackage", "accessPackage", "accessPackage"),
		resourceids.StaticSegment("accessPackageResourceRoleScopes", "accessPackageResourceRoleScopes", "accessPackageResourceRoleScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceRoleScopeId", "accessPackageResourceRoleScopeId"),
		resourceids.StaticSegment("accessPackageResourceScope", "accessPackageResourceScope", "accessPackageResourceScope"),
		resourceids.StaticSegment("accessPackageResource", "accessPackageResource", "accessPackageResource"),
		resourceids.StaticSegment("accessPackageResourceScopes", "accessPackageResourceScopes", "accessPackageResourceScopes"),
		resourceids.UserSpecifiedSegment("accessPackageResourceScopeId", "accessPackageResourceScopeId"),
	}
}

// String returns a human-readable description of this Identity Governance Entitlement Management Access Package Assignment Id Access Package Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Scope ID
func (id IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAccessPackageResourceRoleScopeIdAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceScopeId) String() string {
	components := []string{
		fmt.Sprintf("Access Package Assignment: %q", id.AccessPackageAssignmentId),
		fmt.Sprintf("Access Package Resource Role Scope: %q", id.AccessPackageResourceRoleScopeId),
		fmt.Sprintf("Access Package Resource Scope: %q", id.AccessPackageResourceScopeId),
	}
	return fmt.Sprintf("Identity Governance Entitlement Management Access Package Assignment Id Access Package Access Package Resource Role Scope Id Access Package Resource Scope Access Package Resource Access Package Resource Scope (%s)", strings.Join(components, "\n"))
}