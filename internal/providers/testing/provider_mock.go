// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package testing

import (
	"fmt"
	"sync"

	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
	"github.com/zclconf/go-cty/cty/msgpack"

	"github.com/hashicorp/terraform/internal/configs/hcl2shim"
	"github.com/hashicorp/terraform/internal/providers"
)

var _ providers.Interface = (*MockProvider)(nil)

// MockProvider implements providers.Interface but mocks out all the
// calls for testing purposes.
//
// This is distinct from providers.Mock which is actually available to Terraform
// configuration and test authors. This type is only for use in internal testing
// of Terraform itself.
type MockProvider struct {
	sync.Mutex

	// Anything you want, in case you need to store extra data with the mock.
	Meta interface{}

	GetProviderSchemaCalled   bool
	GetProviderSchemaResponse *providers.GetProviderSchemaResponse

	GetResourceIdentitySchemasCalled   bool
	GetResourceIdentitySchemasResponse *providers.GetResourceIdentitySchemasResponse

	ValidateProviderConfigCalled   bool
	ValidateProviderConfigResponse *providers.ValidateProviderConfigResponse
	ValidateProviderConfigRequest  providers.ValidateProviderConfigRequest
	ValidateProviderConfigFn       func(providers.ValidateProviderConfigRequest) providers.ValidateProviderConfigResponse

	ValidateResourceConfigCalled   bool
	ValidateResourceConfigResponse *providers.ValidateResourceConfigResponse
	ValidateResourceConfigRequest  providers.ValidateResourceConfigRequest
	ValidateResourceConfigFn       func(providers.ValidateResourceConfigRequest) providers.ValidateResourceConfigResponse

	ValidateDataResourceConfigCalled   bool
	ValidateDataResourceConfigResponse *providers.ValidateDataResourceConfigResponse
	ValidateDataResourceConfigRequest  providers.ValidateDataResourceConfigRequest
	ValidateDataResourceConfigFn       func(providers.ValidateDataResourceConfigRequest) providers.ValidateDataResourceConfigResponse

	ValidateListResourceConfigCalled   bool
	ValidateListResourceConfigResponse *providers.ValidateListResourceConfigResponse
	ValidateListResourceConfigRequest  providers.ValidateListResourceConfigRequest
	ValidateListResourceConfigFn       func(providers.ValidateListResourceConfigRequest) providers.ValidateListResourceConfigResponse

	UpgradeResourceStateCalled   bool
	UpgradeResourceStateResponse *providers.UpgradeResourceStateResponse
	UpgradeResourceStateRequest  providers.UpgradeResourceStateRequest
	UpgradeResourceStateFn       func(providers.UpgradeResourceStateRequest) providers.UpgradeResourceStateResponse

	UpgradeResourceIdentityCalled   bool
	UpgradeResourceIdentityResponse *providers.UpgradeResourceIdentityResponse
	UpgradeResourceIdentityRequest  providers.UpgradeResourceIdentityRequest
	UpgradeResourceIdentityFn       func(providers.UpgradeResourceIdentityRequest) providers.UpgradeResourceIdentityResponse

	ConfigureProviderCalled   bool
	ConfigureProviderResponse *providers.ConfigureProviderResponse
	ConfigureProviderRequest  providers.ConfigureProviderRequest
	ConfigureProviderFn       func(providers.ConfigureProviderRequest) providers.ConfigureProviderResponse

	StopCalled   bool
	StopFn       func() error
	StopResponse error

	ReadResourceCalled   bool
	ReadResourceResponse *providers.ReadResourceResponse
	ReadResourceRequest  providers.ReadResourceRequest
	ReadResourceFn       func(providers.ReadResourceRequest) providers.ReadResourceResponse

	PlanResourceChangeCalled   bool
	PlanResourceChangeResponse *providers.PlanResourceChangeResponse
	PlanResourceChangeRequest  providers.PlanResourceChangeRequest
	PlanResourceChangeFn       func(providers.PlanResourceChangeRequest) providers.PlanResourceChangeResponse

	ApplyResourceChangeCalled   bool
	ApplyResourceChangeResponse *providers.ApplyResourceChangeResponse
	ApplyResourceChangeRequest  providers.ApplyResourceChangeRequest
	ApplyResourceChangeFn       func(providers.ApplyResourceChangeRequest) providers.ApplyResourceChangeResponse

	ImportResourceStateCalled   bool
	ImportResourceStateResponse *providers.ImportResourceStateResponse
	ImportResourceStateRequest  providers.ImportResourceStateRequest
	ImportResourceStateFn       func(providers.ImportResourceStateRequest) providers.ImportResourceStateResponse

	MoveResourceStateCalled   bool
	MoveResourceStateResponse *providers.MoveResourceStateResponse
	MoveResourceStateRequest  providers.MoveResourceStateRequest
	MoveResourceStateFn       func(providers.MoveResourceStateRequest) providers.MoveResourceStateResponse

	ReadDataSourceCalled   bool
	ReadDataSourceResponse *providers.ReadDataSourceResponse
	ReadDataSourceRequest  providers.ReadDataSourceRequest
	ReadDataSourceFn       func(providers.ReadDataSourceRequest) providers.ReadDataSourceResponse

	ValidateEphemeralResourceConfigCalled   bool
	ValidateEphemeralResourceConfigResponse *providers.ValidateEphemeralResourceConfigResponse
	ValidateEphemeralResourceConfigRequest  providers.ValidateEphemeralResourceConfigRequest
	ValidateEphemeralResourceConfigFn       func(providers.ValidateEphemeralResourceConfigRequest) providers.ValidateEphemeralResourceConfigResponse
	OpenEphemeralResourceCalled             bool
	OpenEphemeralResourceResponse           *providers.OpenEphemeralResourceResponse
	OpenEphemeralResourceRequest            providers.OpenEphemeralResourceRequest
	OpenEphemeralResourceFn                 func(providers.OpenEphemeralResourceRequest) providers.OpenEphemeralResourceResponse
	RenewEphemeralResourceCalled            bool
	RenewEphemeralResourceResponse          *providers.RenewEphemeralResourceResponse
	RenewEphemeralResourceRequest           providers.RenewEphemeralResourceRequest
	RenewEphemeralResourceFn                func(providers.RenewEphemeralResourceRequest) providers.RenewEphemeralResourceResponse
	CloseEphemeralResourceCalled            bool
	CloseEphemeralResourceResponse          *providers.CloseEphemeralResourceResponse
	CloseEphemeralResourceRequest           providers.CloseEphemeralResourceRequest
	CloseEphemeralResourceFn                func(providers.CloseEphemeralResourceRequest) providers.CloseEphemeralResourceResponse

	CallFunctionCalled   bool
	CallFunctionResponse providers.CallFunctionResponse
	CallFunctionRequest  providers.CallFunctionRequest
	CallFunctionFn       func(providers.CallFunctionRequest) providers.CallFunctionResponse

	ListResourceCalled   bool
	ListResourceResponse providers.ListResourceResponse
	ListResourceRequest  providers.ListResourceRequest
	ListResourceFn       func(providers.ListResourceRequest) providers.ListResourceResponse

	ValidateStateStoreConfigCalled   bool
	ValidateStateStoreConfigResponse *providers.ValidateStateStoreConfigResponse
	ValidateStateStoreConfigRequest  providers.ValidateStateStoreConfigRequest
	ValidateStateStoreConfigFn       func(providers.ValidateStateStoreConfigRequest) providers.ValidateStateStoreConfigResponse

	ConfigureStateStoreCalled   bool
	ConfigureStateStoreResponse *providers.ConfigureStateStoreResponse
	ConfigureStateStoreRequest  providers.ConfigureStateStoreRequest
	ConfigureStateStoreFn       func(providers.ConfigureStateStoreRequest) providers.ConfigureStateStoreResponse

	GetStatesCalled   bool
	GetStatesResponse *providers.GetStatesResponse
	GetStatesRequest  providers.GetStatesRequest
	GetStatesFn       func(providers.GetStatesRequest) providers.GetStatesResponse

	DeleteStateCalled   bool
	DeleteStateResponse *providers.DeleteStateResponse
	DeleteStateRequest  providers.DeleteStateRequest
	DeleteStateFn       func(providers.DeleteStateRequest) providers.DeleteStateResponse

	PlanActionCalled   bool
	PlanActionResponse providers.PlanActionResponse
	PlanActionRequest  providers.PlanActionRequest
	PlanActionFn       func(providers.PlanActionRequest) providers.PlanActionResponse

	InvokeActionCalled   bool
	InvokeActionResponse providers.InvokeActionResponse
	InvokeActionRequest  providers.InvokeActionRequest
	InvokeActionFn       func(providers.InvokeActionRequest) providers.InvokeActionResponse

	ValidateActionCalled         bool
	ValidateActionConfigRequest  providers.ValidateActionConfigRequest
	ValidateActionConfigResponse *providers.ValidateActionConfigResponse
	ValidateActionConfigFn       func(providers.ValidateActionConfigRequest) providers.ValidateActionConfigResponse

	CloseCalled bool
	CloseError  error
}

func (p *MockProvider) GetProviderSchema() providers.GetProviderSchemaResponse {
	defer p.beginWrite()()
	p.GetProviderSchemaCalled = true
	return p.getProviderSchema()
}

func (p *MockProvider) getProviderSchema() providers.GetProviderSchemaResponse {
	// This version of getProviderSchema doesn't do any locking, so it's suitable to
	// call from other methods of this mock as long as they are already
	// holding the lock.
	if p.GetProviderSchemaResponse != nil {
		return *p.GetProviderSchemaResponse
	}

	return providers.GetProviderSchemaResponse{
		Provider:          providers.Schema{},
		DataSources:       map[string]providers.Schema{},
		ResourceTypes:     map[string]providers.Schema{},
		ListResourceTypes: map[string]providers.Schema{},
		StateStores:       map[string]providers.Schema{},
	}
}

func (p *MockProvider) GetResourceIdentitySchemas() providers.GetResourceIdentitySchemasResponse {
	defer p.beginWrite()()
	p.GetResourceIdentitySchemasCalled = true

	return p.getResourceIdentitySchemas()
}

func (p *MockProvider) getResourceIdentitySchemas() providers.GetResourceIdentitySchemasResponse {
	if p.GetResourceIdentitySchemasResponse != nil {
		return *p.GetResourceIdentitySchemasResponse
	}

	resp := providers.GetResourceIdentitySchemasResponse{IdentityTypes: make(map[string]providers.IdentitySchema)}
	if p.GetProviderSchemaResponse != nil {

		for typeName, schema := range p.GetProviderSchemaResponse.ResourceTypes {
			if schema.Identity != nil {
				resp.IdentityTypes[typeName] = providers.IdentitySchema{
					Version: schema.IdentityVersion,
					Body:    schema.Identity,
				}
			}
		}

	}

	return resp
}

func (p *MockProvider) ValidateProviderConfig(r providers.ValidateProviderConfigRequest) (resp providers.ValidateProviderConfigResponse) {
	defer p.beginWrite()()

	p.ValidateProviderConfigCalled = true
	p.ValidateProviderConfigRequest = r
	if p.ValidateProviderConfigFn != nil {
		return p.ValidateProviderConfigFn(r)
	}

	if p.ValidateProviderConfigResponse != nil {
		return *p.ValidateProviderConfigResponse
	}

	resp.PreparedConfig = r.Config
	return resp
}

func (p *MockProvider) ValidateResourceConfig(r providers.ValidateResourceConfigRequest) (resp providers.ValidateResourceConfigResponse) {
	defer p.beginWrite()()

	p.ValidateResourceConfigCalled = true
	p.ValidateResourceConfigRequest = r

	// Marshall the value to replicate behavior by the GRPC protocol,
	// and return any relevant errors
	resourceSchema, ok := p.getProviderSchema().ResourceTypes[r.TypeName]
	if !ok {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no schema found for %q", r.TypeName))
		return resp
	}

	_, err := msgpack.Marshal(r.Config, resourceSchema.Body.ImpliedType())
	if err != nil {
		resp.Diagnostics = resp.Diagnostics.Append(err)
		return resp
	}

	if p.ValidateResourceConfigFn != nil {
		return p.ValidateResourceConfigFn(r)
	}

	if p.ValidateResourceConfigResponse != nil {
		return *p.ValidateResourceConfigResponse
	}

	return resp
}

func (p *MockProvider) ValidateDataResourceConfig(r providers.ValidateDataResourceConfigRequest) (resp providers.ValidateDataResourceConfigResponse) {
	defer p.beginWrite()()

	p.ValidateDataResourceConfigCalled = true
	p.ValidateDataResourceConfigRequest = r

	// Marshall the value to replicate behavior by the GRPC protocol
	dataSchema, ok := p.getProviderSchema().DataSources[r.TypeName]
	if !ok {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no schema found for %q", r.TypeName))
		return resp
	}
	_, err := msgpack.Marshal(r.Config, dataSchema.Body.ImpliedType())
	if err != nil {
		resp.Diagnostics = resp.Diagnostics.Append(err)
		return resp
	}

	if p.ValidateDataResourceConfigFn != nil {
		return p.ValidateDataResourceConfigFn(r)
	}

	if p.ValidateDataResourceConfigResponse != nil {
		return *p.ValidateDataResourceConfigResponse
	}

	return resp
}

func (p *MockProvider) ValidateEphemeralResourceConfig(r providers.ValidateEphemeralResourceConfigRequest) (resp providers.ValidateEphemeralResourceConfigResponse) {
	defer p.beginWrite()()

	p.ValidateEphemeralResourceConfigCalled = true
	p.ValidateEphemeralResourceConfigRequest = r

	// Marshall the value to replicate behavior by the GRPC protocol
	ephemeralSchema, ok := p.getProviderSchema().EphemeralResourceTypes[r.TypeName]
	if !ok {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no schema found for %q", r.TypeName))
		return resp
	}
	_, err := msgpack.Marshal(r.Config, ephemeralSchema.Body.ImpliedType())
	if err != nil {
		resp.Diagnostics = resp.Diagnostics.Append(err)
		return resp
	}

	if p.ValidateEphemeralResourceConfigFn != nil {
		return p.ValidateEphemeralResourceConfigFn(r)
	}

	if p.ValidateEphemeralResourceConfigResponse != nil {
		return *p.ValidateEphemeralResourceConfigResponse
	}

	return resp
}

func (p *MockProvider) ValidateListResourceConfig(r providers.ValidateListResourceConfigRequest) (resp providers.ValidateListResourceConfigResponse) {
	defer p.beginWrite()()

	p.ValidateListResourceConfigCalled = true
	p.ValidateListResourceConfigRequest = r

	// Marshall the value to replicate behavior by the GRPC protocol
	listSchema, ok := p.getProviderSchema().ListResourceTypes[r.TypeName]
	if !ok {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no schema found for %q", r.TypeName))
		return resp
	}
	_, err := msgpack.Marshal(r.Config, listSchema.Body.ImpliedType())
	if err != nil {
		resp.Diagnostics = resp.Diagnostics.Append(err)
		return resp
	}

	if p.ValidateListResourceConfigFn != nil {
		return p.ValidateListResourceConfigFn(r)
	}

	if p.ValidateListResourceConfigResponse != nil {
		return *p.ValidateListResourceConfigResponse
	}

	return resp
}

// UpgradeResourceState mocks out the response from the provider during an UpgradeResourceState RPC
// The default logic will return the resource's state unchanged, unless other logic is defined on the mock (e.g. UpgradeResourceStateFn)
//
// When using this mock you may need to provide custom logic if the plugin-framework alters values in state,
// e.g. when handling write-only attributes.
func (p *MockProvider) UpgradeResourceState(r providers.UpgradeResourceStateRequest) (resp providers.UpgradeResourceStateResponse) {
	defer p.beginWrite()()

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before UpgradeResourceState %q", r.TypeName))
		return resp
	}

	schema, ok := p.getProviderSchema().ResourceTypes[r.TypeName]
	if !ok {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no schema found for %q", r.TypeName))
		return resp
	}

	schemaType := schema.Body.ImpliedType()

	p.UpgradeResourceStateCalled = true
	p.UpgradeResourceStateRequest = r

	if p.UpgradeResourceStateFn != nil {
		return p.UpgradeResourceStateFn(r)
	}

	if p.UpgradeResourceStateResponse != nil {
		return *p.UpgradeResourceStateResponse
	}

	switch {
	case r.RawStateFlatmap != nil:
		v, err := hcl2shim.HCL2ValueFromFlatmap(r.RawStateFlatmap, schemaType)
		if err != nil {
			resp.Diagnostics = resp.Diagnostics.Append(err)
			return resp
		}
		resp.UpgradedState = v
	case len(r.RawStateJSON) > 0:
		v, err := ctyjson.Unmarshal(r.RawStateJSON, schemaType)

		if err != nil {
			resp.Diagnostics = resp.Diagnostics.Append(err)
			return resp
		}
		resp.UpgradedState = v
	}

	return resp
}

func (p *MockProvider) UpgradeResourceIdentity(r providers.UpgradeResourceIdentityRequest) (resp providers.UpgradeResourceIdentityResponse) {
	defer p.beginWrite()()

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before UpgradeResourceIdentity %q", r.TypeName))
		return resp
	}
	p.UpgradeResourceIdentityCalled = true
	p.UpgradeResourceIdentityRequest = r

	if p.UpgradeResourceIdentityFn != nil {
		return p.UpgradeResourceIdentityFn(r)
	}

	if p.UpgradeResourceIdentityResponse != nil {
		return *p.UpgradeResourceIdentityResponse
	}

	schema, ok := p.getProviderSchema().ResourceTypes[r.TypeName]

	if !ok || schema.Identity == nil {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no identity schema found for %q", r.TypeName))
		return resp
	}

	identityType := schema.Identity.ImpliedType()

	v, err := ctyjson.Unmarshal(r.RawIdentityJSON, identityType)

	if err != nil {
		resp.Diagnostics = resp.Diagnostics.Append(err)
		return resp
	}
	resp.UpgradedIdentity = v

	return resp
}

func (p *MockProvider) ConfigureProvider(r providers.ConfigureProviderRequest) (resp providers.ConfigureProviderResponse) {
	defer p.beginWrite()()

	p.ConfigureProviderCalled = true
	p.ConfigureProviderRequest = r

	if p.ConfigureProviderFn != nil {
		return p.ConfigureProviderFn(r)
	}

	if p.ConfigureProviderResponse != nil {
		return *p.ConfigureProviderResponse
	}

	return resp
}

func (p *MockProvider) Stop() error {
	// We intentionally don't lock in this one because the whole point of this
	// method is to be called concurrently with another operation that can
	// be cancelled.  The provider itself is responsible for handling
	// any concurrency concerns in this case.

	p.StopCalled = true
	if p.StopFn != nil {
		return p.StopFn()
	}

	return p.StopResponse
}

func (p *MockProvider) ReadResource(r providers.ReadResourceRequest) (resp providers.ReadResourceResponse) {
	defer p.beginWrite()()

	p.ReadResourceCalled = true
	p.ReadResourceRequest = r

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before ReadResource %q", r.TypeName))
		return resp
	}

	if p.ReadResourceFn != nil {
		return p.ReadResourceFn(r)
	}

	if p.ReadResourceResponse != nil {
		resp = *p.ReadResourceResponse

		// Make sure the NewState conforms to the schema.
		// This isn't always the case for the existing tests.
		schema, ok := p.getProviderSchema().ResourceTypes[r.TypeName]
		if !ok {
			resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no schema found for %q", r.TypeName))
			return resp
		}

		newState, err := schema.Body.CoerceValue(resp.NewState)
		if err != nil {
			resp.Diagnostics = resp.Diagnostics.Append(err)
		}
		resp.NewState = newState
		if resp.Identity.IsNull() {
			resp.Identity = r.CurrentIdentity
		}

		return resp
	}

	// otherwise just return the same state we received without the write-only attributes
	// since there are old tests without a schema we default to the prior state
	if schema, ok := p.getProviderSchema().ResourceTypes[r.TypeName]; ok {

		newVal, err := cty.Transform(r.PriorState, func(path cty.Path, v cty.Value) (cty.Value, error) {
			// We're only concerned with known null values, which can be computed
			// by the provider.
			if !v.IsKnown() {
				return v, nil
			}

			attrSchema := schema.Body.AttributeByPath(path)
			if attrSchema == nil {
				// this is an intermediate path which does not represent an attribute
				return v, nil
			}

			// Write-only attributes always return null
			if attrSchema.WriteOnly {
				return cty.NullVal(v.Type()), nil
			}

			return v, nil
		})
		if err != nil {
			resp.Diagnostics = resp.Diagnostics.Append(err)
		}
		resp.NewState = newVal
	} else {
		resp.NewState = r.PriorState
	}

	resp.Identity = r.CurrentIdentity
	resp.Private = r.Private
	return resp
}

func (p *MockProvider) PlanResourceChange(r providers.PlanResourceChangeRequest) (resp providers.PlanResourceChangeResponse) {
	defer p.beginWrite()()

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before PlanResourceChange %q", r.TypeName))
		return resp
	}

	p.PlanResourceChangeCalled = true
	p.PlanResourceChangeRequest = r

	if p.PlanResourceChangeFn != nil {
		return p.PlanResourceChangeFn(r)
	}

	if p.PlanResourceChangeResponse != nil {
		return *p.PlanResourceChangeResponse
	}

	// this is a destroy plan,
	if r.ProposedNewState.IsNull() {
		resp.PlannedState = r.ProposedNewState
		resp.PlannedPrivate = r.PriorPrivate
		return resp
	}

	schema, ok := p.getProviderSchema().ResourceTypes[r.TypeName]
	if !ok {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no schema found for %q", r.TypeName))
		return resp
	}

	// The default plan behavior is to accept the proposed value, and mark all
	// nil computed attributes as unknown.
	val, err := cty.Transform(r.ProposedNewState, func(path cty.Path, v cty.Value) (cty.Value, error) {
		// We're only concerned with known null values, which can be computed
		// by the provider.
		if !v.IsKnown() {
			return v, nil
		}

		attrSchema := schema.Body.AttributeByPath(path)
		if attrSchema == nil {
			// this is an intermediate path which does not represent an attribute
			return v, nil
		}

		// Write-only attributes always return null
		if attrSchema.WriteOnly {
			return cty.NullVal(v.Type()), nil
		}

		// get the current configuration value, to detect when a
		// computed+optional attributes has become unset
		configVal, err := path.Apply(r.Config)
		if err != nil {
			// cty can't currently apply some paths, so don't try to guess
			// what's needed here and return the proposed part of the value.
			// This is only a helper to create a default plan value, any tests
			// relying on specific plan behavior will create their own
			// PlanResourceChange responses.
			return v, nil
		}

		switch {
		case attrSchema.Computed && !attrSchema.Optional && v.IsNull():
			// this is the easy path, this value is not yet set, and _must_ be computed
			return cty.UnknownVal(v.Type()), nil

		case attrSchema.Computed && attrSchema.Optional && !v.IsNull() && configVal.IsNull():
			// If an optional+computed value has gone from set to unset, it
			// becomes computed. (this was not possible to do with legacy
			// providers)
			return cty.UnknownVal(v.Type()), nil
		}

		return v, nil
	})
	if err != nil {
		resp.Diagnostics = resp.Diagnostics.Append(err)
		return resp
	}

	resp.PlannedPrivate = r.PriorPrivate
	resp.PlannedState = val

	return resp
}

func (p *MockProvider) ApplyResourceChange(r providers.ApplyResourceChangeRequest) (resp providers.ApplyResourceChangeResponse) {
	defer p.beginWrite()()

	p.ApplyResourceChangeCalled = true
	p.ApplyResourceChangeRequest = r

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before ApplyResourceChange %q", r.TypeName))
		return resp
	}

	if p.ApplyResourceChangeFn != nil {
		return p.ApplyResourceChangeFn(r)
	}

	if p.ApplyResourceChangeResponse != nil {
		return *p.ApplyResourceChangeResponse
	}

	// if the value is nil, we return that directly to correspond to a delete
	if r.PlannedState.IsNull() {
		resp.NewState = r.PlannedState
		resp.NewIdentity = r.PlannedIdentity
		return resp
	}

	// the default behavior will be to create the minimal valid apply value by
	// setting unknowns (which correspond to computed attributes) to a zero
	// value.
	val, _ := cty.Transform(r.PlannedState, func(path cty.Path, v cty.Value) (cty.Value, error) {
		if !v.IsKnown() {
			ty := v.Type()
			switch {
			case ty == cty.String:
				return cty.StringVal(""), nil
			case ty == cty.Number:
				return cty.NumberIntVal(0), nil
			case ty == cty.Bool:
				return cty.False, nil
			case ty.IsMapType():
				return cty.MapValEmpty(ty.ElementType()), nil
			case ty.IsListType():
				return cty.ListValEmpty(ty.ElementType()), nil
			default:
				return cty.NullVal(ty), nil
			}
		}
		return v, nil
	})

	resp.NewState = val
	resp.Private = r.PlannedPrivate
	resp.NewIdentity = r.PlannedIdentity

	return resp
}

func (p *MockProvider) ImportResourceState(r providers.ImportResourceStateRequest) (resp providers.ImportResourceStateResponse) {
	defer p.beginWrite()()

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before ImportResourceState %q", r.TypeName))
		return resp
	}

	p.ImportResourceStateCalled = true
	p.ImportResourceStateRequest = r
	if p.ImportResourceStateFn != nil {
		return p.ImportResourceStateFn(r)
	}

	if p.ImportResourceStateResponse != nil {
		resp = *p.ImportResourceStateResponse

		// take a copy of the slice, because it is read by the resource instance
		importedResources := make([]providers.ImportedResource, len(resp.ImportedResources))
		copy(importedResources, resp.ImportedResources)

		// fixup the cty value to match the schema
		for i, res := range importedResources {
			schema, ok := p.getProviderSchema().ResourceTypes[res.TypeName]
			if !ok {
				resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no schema found for %q", res.TypeName))
				return resp
			}

			var err error
			res.State, err = schema.Body.CoerceValue(res.State)
			if err != nil {
				resp.Diagnostics = resp.Diagnostics.Append(err)
				return resp
			}

			importedResources[i] = res
		}
		resp.ImportedResources = importedResources
	}

	return resp
}

func (p *MockProvider) MoveResourceState(r providers.MoveResourceStateRequest) (resp providers.MoveResourceStateResponse) {
	defer p.beginWrite()()

	p.MoveResourceStateCalled = true
	p.MoveResourceStateRequest = r
	if p.MoveResourceStateFn != nil {
		return p.MoveResourceStateFn(r)
	}

	if p.MoveResourceStateResponse != nil {
		resp = *p.MoveResourceStateResponse
	}

	return resp
}

func (p *MockProvider) ReadDataSource(r providers.ReadDataSourceRequest) (resp providers.ReadDataSourceResponse) {
	defer p.beginWrite()()

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before ReadDataSource %q", r.TypeName))
		return resp
	}

	p.ReadDataSourceCalled = true
	p.ReadDataSourceRequest = r

	if p.ReadDataSourceFn != nil {
		return p.ReadDataSourceFn(r)
	}

	if p.ReadDataSourceResponse != nil {
		resp = *p.ReadDataSourceResponse
	}

	return resp
}

func (p *MockProvider) OpenEphemeralResource(r providers.OpenEphemeralResourceRequest) (resp providers.OpenEphemeralResourceResponse) {
	defer p.beginWrite()()

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before OpenEphemeralResource %q", r.TypeName))
		return resp
	}

	p.OpenEphemeralResourceCalled = true
	p.OpenEphemeralResourceRequest = r

	if p.OpenEphemeralResourceFn != nil {
		return p.OpenEphemeralResourceFn(r)
	}

	if p.OpenEphemeralResourceResponse != nil {
		resp = *p.OpenEphemeralResourceResponse
	}

	return resp
}

func (p *MockProvider) RenewEphemeralResource(r providers.RenewEphemeralResourceRequest) (resp providers.RenewEphemeralResourceResponse) {
	defer p.beginWrite()()

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before RenewEphemeralResource %q", r.TypeName))
		return resp
	}

	p.RenewEphemeralResourceCalled = true
	p.RenewEphemeralResourceRequest = r

	if p.RenewEphemeralResourceFn != nil {
		return p.RenewEphemeralResourceFn(r)
	}

	if p.RenewEphemeralResourceResponse != nil {
		resp = *p.RenewEphemeralResourceResponse
	}

	return resp
}

func (p *MockProvider) CloseEphemeralResource(r providers.CloseEphemeralResourceRequest) (resp providers.CloseEphemeralResourceResponse) {
	defer p.beginWrite()()

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before CloseEphemeralResource %q", r.TypeName))
		return resp
	}

	p.CloseEphemeralResourceCalled = true
	p.CloseEphemeralResourceRequest = r

	if p.CloseEphemeralResourceFn != nil {
		return p.CloseEphemeralResourceFn(r)
	}

	if p.CloseEphemeralResourceResponse != nil {
		resp = *p.CloseEphemeralResourceResponse
	}

	return resp
}

func (p *MockProvider) CallFunction(r providers.CallFunctionRequest) providers.CallFunctionResponse {
	defer p.beginWrite()()

	p.CallFunctionCalled = true
	p.CallFunctionRequest = r

	if p.CallFunctionFn != nil {
		return p.CallFunctionFn(r)
	}

	return p.CallFunctionResponse
}

func (p *MockProvider) ListResource(r providers.ListResourceRequest) providers.ListResourceResponse {
	p.Lock()
	defer p.Unlock()
	p.ListResourceCalled = true
	p.ListResourceRequest = r

	if p.ListResourceFn != nil {
		return p.ListResourceFn(r)
	}

	return p.ListResourceResponse
}

func (p *MockProvider) ValidateStateStoreConfig(r providers.ValidateStateStoreConfigRequest) (resp providers.ValidateStateStoreConfigResponse) {
	p.Lock()
	defer p.Unlock()

	p.ValidateStateStoreConfigCalled = true
	p.ValidateStateStoreConfigRequest = r

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before ValidateStateStoreConfig %q", r.TypeName))
		return resp
	}

	if p.ValidateStateStoreConfigResponse != nil {
		return *p.ValidateStateStoreConfigResponse
	}

	if p.ValidateStateStoreConfigFn != nil {
		return p.ValidateStateStoreConfigFn(r)
	}

	// In the absence of any custom logic, we do basic validation of the received config against the schema.
	//
	// Marshall the value to replicate behavior by the GRPC protocol,
	// and return any relevant errors
	storeSchema, ok := p.getProviderSchema().StateStores[r.TypeName]
	if !ok {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no schema found for state store %q", r.TypeName))
		return resp
	}

	_, err := msgpack.Marshal(r.Config, storeSchema.Body.ImpliedType())
	if err != nil {
		resp.Diagnostics = resp.Diagnostics.Append(err)
		return resp
	}

	return resp
}

func (p *MockProvider) ConfigureStateStore(r providers.ConfigureStateStoreRequest) (resp providers.ConfigureStateStoreResponse) {
	p.Lock()
	defer p.Unlock()

	p.ConfigureStateStoreCalled = true
	p.ConfigureStateStoreRequest = r

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("Configure not called before ConfigureStateStore %q", r.TypeName))
		return resp
	}

	if p.ConfigureStateStoreFn != nil {
		return p.ConfigureStateStoreFn(r)
	}

	// In the absence of any custom logic, we do the logic below.
	//
	// Marshall the value to replicate behavior by the GRPC protocol,
	// and return any relevant errors
	storeSchema, ok := p.getProviderSchema().StateStores[r.TypeName]
	if !ok {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("no schema found for state store %q", r.TypeName))
		return resp
	}

	if p.ConfigureStateStoreResponse != nil {
		return *p.ConfigureStateStoreResponse
	}

	_, err := msgpack.Marshal(r.Config, storeSchema.Body.ImpliedType())
	if err != nil {
		resp.Diagnostics = resp.Diagnostics.Append(err)
		return resp
	}

	return resp
}

func (p *MockProvider) GetStates(r providers.GetStatesRequest) (resp providers.GetStatesResponse) {
	p.Lock()
	defer p.Unlock()

	p.GetStatesCalled = true
	p.GetStatesRequest = r

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("ConfigureProvider not called before GetStates %q", r.TypeName))
	}
	if !p.ConfigureStateStoreCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("ConfigureStateStore not called before GetStates %q", r.TypeName))
	}
	if resp.Diagnostics.HasErrors() {
		return resp
	}

	if p.GetStatesResponse != nil {
		return *p.GetStatesResponse
	}

	if p.GetStatesFn != nil {
		return p.GetStatesFn(r)
	}

	// If the mock has no further inputs, return an empty list.
	// The state store should be reporting a minimum of the default workspace usually,
	// but this should be achieved by querying data storage and identifying the artifact
	// for that workspace, and reporting that the workspace exists.
	resp.States = []string{}

	return resp
}

func (p *MockProvider) DeleteState(r providers.DeleteStateRequest) (resp providers.DeleteStateResponse) {
	p.Lock()
	defer p.Unlock()

	p.DeleteStateCalled = true
	p.DeleteStateRequest = r

	if !p.ConfigureProviderCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("ConfigureProvider not called before DeleteState %q", r.TypeName))
	}
	if !p.ConfigureStateStoreCalled {
		resp.Diagnostics = resp.Diagnostics.Append(fmt.Errorf("ConfigureStateStore not called before DeleteState %q", r.TypeName))
	}

	if p.DeleteStateResponse != nil {
		return *p.DeleteStateResponse
	}

	if p.DeleteStateFn != nil {
		return p.DeleteStateFn(r)
	}

	// There's no logic we can include here in the absence of other fields on the mock.

	// If the response contains no diagnostics then the deletion is assumed to be successful.
	return resp
}

func (p *MockProvider) PlanAction(r providers.PlanActionRequest) (resp providers.PlanActionResponse) {
	p.Lock()
	defer p.Unlock()

	p.PlanActionCalled = true
	p.PlanActionRequest = r

	if p.PlanActionFn != nil {
		return p.PlanActionFn(r)
	}

	return p.PlanActionResponse
}

func (p *MockProvider) InvokeAction(r providers.InvokeActionRequest) (resp providers.InvokeActionResponse) {
	p.Lock()
	defer p.Unlock()

	p.InvokeActionCalled = true
	p.InvokeActionRequest = r

	if p.InvokeActionFn != nil {
		return p.InvokeActionFn(r)
	}

	return p.InvokeActionResponse
}

func (p *MockProvider) Close() error {
	defer p.beginWrite()()

	p.CloseCalled = true
	return p.CloseError
}

func (p *MockProvider) beginWrite() func() {
	p.Lock()
	return p.Unlock
}

func (p *MockProvider) ValidateActionConfig(r providers.ValidateActionConfigRequest) (resp providers.ValidateActionConfigResponse) {
	defer p.beginWrite()

	p.ValidateActionCalled = true
	p.ValidateActionConfigRequest = r

	if p.ValidateActionConfigFn != nil {
		return p.ValidateActionConfigFn(r)
	}

	if p.ValidateActionConfigResponse != nil {
		return *p.ValidateActionConfigResponse
	}

	return resp
}
