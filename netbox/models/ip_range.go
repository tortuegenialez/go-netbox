// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IPRange IP range
//
// swagger:model IPRange
type IPRange struct {

	// Children
	// Read Only: true
	Children int64 `json:"children,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// End address
	//
	// IPv4 or IPv6 address (with mask)
	// Required: true
	EndAddress *string `json:"end_address"`

	// family
	Family *IPRangeFamily `json:"family,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// role
	Role *NestedRole `json:"role,omitempty"`

	// Size
	// Read Only: true
	Size int64 `json:"size,omitempty"`

	// Start address
	//
	// IPv4 or IPv6 address (with mask)
	// Required: true
	StartAddress *string `json:"start_address"`

	// status
	Status *IPRangeStatus `json:"status,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags,omitempty"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// vrf
	Vrf *NestedVRF `json:"vrf,omitempty"`
}

// Validate validates this IP range
func (m *IPRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFamily(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVrf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPRange) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) validateEndAddress(formats strfmt.Registry) error {

	if err := validate.Required("end_address", "body", m.EndAddress); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) validateFamily(formats strfmt.Registry) error {
	if swag.IsZero(m.Family) { // not required
		return nil
	}

	if m.Family != nil {
		if err := m.Family.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("family")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("family")
			}
			return err
		}
	}

	return nil
}

func (m *IPRange) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *IPRange) validateStartAddress(formats strfmt.Registry) error {

	if err := validate.Required("start_address", "body", m.StartAddress); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *IPRange) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IPRange) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *IPRange) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) validateVrf(formats strfmt.Registry) error {
	if swag.IsZero(m.Vrf) { // not required
		return nil
	}

	if m.Vrf != nil {
		if err := m.Vrf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vrf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vrf")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this IP range based on the context it is used
func (m *IPRange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChildren(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFamily(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRole(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVrf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPRange) contextValidateChildren(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "children", "body", int64(m.Children)); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) contextValidateFamily(ctx context.Context, formats strfmt.Registry) error {

	if m.Family != nil {
		if err := m.Family.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("family")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("family")
			}
			return err
		}
	}

	return nil
}

func (m *IPRange) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) contextValidateRole(ctx context.Context, formats strfmt.Registry) error {

	if m.Role != nil {
		if err := m.Role.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

func (m *IPRange) contextValidateSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "size", "body", int64(m.Size)); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *IPRange) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IPRange) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenant != nil {
		if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *IPRange) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *IPRange) contextValidateVrf(ctx context.Context, formats strfmt.Registry) error {

	if m.Vrf != nil {
		if err := m.Vrf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vrf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vrf")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPRange) UnmarshalBinary(b []byte) error {
	var res IPRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPRangeFamily Family
//
// swagger:model IPRangeFamily
type IPRangeFamily struct {

	// label
	// Required: true
	// Enum: [IPv4 IPv6]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [4 6]
	Value *int64 `json:"value"`
}

// Validate validates this IP range family
func (m *IPRangeFamily) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ipRangeFamilyTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IPv4","IPv6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipRangeFamilyTypeLabelPropEnum = append(ipRangeFamilyTypeLabelPropEnum, v)
	}
}

const (

	// IPRangeFamilyLabelIPV4 captures enum value "IPv4"
	IPRangeFamilyLabelIPV4 string = "IPv4"

	// IPRangeFamilyLabelIPV6 captures enum value "IPv6"
	IPRangeFamilyLabelIPV6 string = "IPv6"
)

// prop value enum
func (m *IPRangeFamily) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ipRangeFamilyTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPRangeFamily) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("family"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("family"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var ipRangeFamilyTypeValuePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[4,6]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipRangeFamilyTypeValuePropEnum = append(ipRangeFamilyTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *IPRangeFamily) validateValueEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, ipRangeFamilyTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPRangeFamily) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("family"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("family"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this IP range family based on the context it is used
func (m *IPRangeFamily) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *IPRangeFamily) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPRangeFamily) UnmarshalBinary(b []byte) error {
	var res IPRangeFamily
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IPRangeStatus Status
//
// swagger:model IPRangeStatus
type IPRangeStatus struct {

	// label
	// Required: true
	// Enum: [Active Reserved Deprecated]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [active reserved deprecated]
	Value *string `json:"value"`
}

// Validate validates this IP range status
func (m *IPRangeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ipRangeStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Reserved","Deprecated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipRangeStatusTypeLabelPropEnum = append(ipRangeStatusTypeLabelPropEnum, v)
	}
}

const (

	// IPRangeStatusLabelActive captures enum value "Active"
	IPRangeStatusLabelActive string = "Active"

	// IPRangeStatusLabelReserved captures enum value "Reserved"
	IPRangeStatusLabelReserved string = "Reserved"

	// IPRangeStatusLabelDeprecated captures enum value "Deprecated"
	IPRangeStatusLabelDeprecated string = "Deprecated"
)

// prop value enum
func (m *IPRangeStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ipRangeStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPRangeStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var ipRangeStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","reserved","deprecated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ipRangeStatusTypeValuePropEnum = append(ipRangeStatusTypeValuePropEnum, v)
	}
}

const (

	// IPRangeStatusValueActive captures enum value "active"
	IPRangeStatusValueActive string = "active"

	// IPRangeStatusValueReserved captures enum value "reserved"
	IPRangeStatusValueReserved string = "reserved"

	// IPRangeStatusValueDeprecated captures enum value "deprecated"
	IPRangeStatusValueDeprecated string = "deprecated"
)

// prop value enum
func (m *IPRangeStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ipRangeStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IPRangeStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this IP range status based on context it is used
func (m *IPRangeStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IPRangeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPRangeStatus) UnmarshalBinary(b []byte) error {
	var res IPRangeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
