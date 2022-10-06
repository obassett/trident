// Code generated by go-swagger; DO NOT EDIT.

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

// StoragePoolDisk storage pool disk
//
// swagger:model storage_pool_disk
type StoragePoolDisk struct {

	// disk
	Disk *StoragePoolDiskDisk `json:"disk,omitempty"`

	// Raw capacity of the disk, in bytes.
	// Read Only: true
	TotalSize int64 `json:"total_size,omitempty"`

	// Usable capacity of this disk, in bytes.
	// Read Only: true
	UsableSize int64 `json:"usable_size,omitempty"`
}

// Validate validates this storage pool disk
func (m *StoragePoolDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePoolDisk) validateDisk(formats strfmt.Registry) error {
	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if m.Disk != nil {
		if err := m.Disk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this storage pool disk based on the context it is used
func (m *StoragePoolDisk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsableSize(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePoolDisk) contextValidateDisk(ctx context.Context, formats strfmt.Registry) error {

	if m.Disk != nil {
		if err := m.Disk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *StoragePoolDisk) contextValidateTotalSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "total_size", "body", int64(m.TotalSize)); err != nil {
		return err
	}

	return nil
}

func (m *StoragePoolDisk) contextValidateUsableSize(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "usable_size", "body", int64(m.UsableSize)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoragePoolDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePoolDisk) UnmarshalBinary(b []byte) error {
	var res StoragePoolDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StoragePoolDiskDisk Reference to the constituent disk object.
//
// swagger:model StoragePoolDiskDisk
type StoragePoolDiskDisk struct {

	// links
	Links *StoragePoolDiskDiskLinks `json:"_links,omitempty"`

	// name
	// Example: 1.0.1
	Name string `json:"name,omitempty"`
}

// Validate validates this storage pool disk disk
func (m *StoragePoolDiskDisk) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePoolDiskDisk) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this storage pool disk disk based on the context it is used
func (m *StoragePoolDiskDisk) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePoolDiskDisk) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoragePoolDiskDisk) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePoolDiskDisk) UnmarshalBinary(b []byte) error {
	var res StoragePoolDiskDisk
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StoragePoolDiskDiskLinks storage pool disk disk links
//
// swagger:model StoragePoolDiskDiskLinks
type StoragePoolDiskDiskLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this storage pool disk disk links
func (m *StoragePoolDiskDiskLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePoolDiskDiskLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this storage pool disk disk links based on the context it is used
func (m *StoragePoolDiskDiskLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StoragePoolDiskDiskLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StoragePoolDiskDiskLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StoragePoolDiskDiskLinks) UnmarshalBinary(b []byte) error {
	var res StoragePoolDiskDiskLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}