// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Operator
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateTenantRequest update tenant request
//
// swagger:model updateTenantRequest
type UpdateTenantRequest struct {

	// image
	// Pattern: ^((.*?)/(.*?):(.+))$
	Image string `json:"image,omitempty"`

	// image pull secret
	ImagePullSecret string `json:"image_pull_secret,omitempty"`

	// image registry
	ImageRegistry *ImageRegistry `json:"image_registry,omitempty"`
}

// Validate validates this update tenant request
func (m *UpdateTenantRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageRegistry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTenantRequest) validateImage(formats strfmt.Registry) error {
	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if err := validate.Pattern("image", "body", m.Image, `^((.*?)/(.*?):(.+))$`); err != nil {
		return err
	}

	return nil
}

func (m *UpdateTenantRequest) validateImageRegistry(formats strfmt.Registry) error {
	if swag.IsZero(m.ImageRegistry) { // not required
		return nil
	}

	if m.ImageRegistry != nil {
		if err := m.ImageRegistry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image_registry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image_registry")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update tenant request based on the context it is used
func (m *UpdateTenantRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateImageRegistry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateTenantRequest) contextValidateImageRegistry(ctx context.Context, formats strfmt.Registry) error {

	if m.ImageRegistry != nil {
		if err := m.ImageRegistry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image_registry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("image_registry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateTenantRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateTenantRequest) UnmarshalBinary(b []byte) error {
	var res UpdateTenantRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}