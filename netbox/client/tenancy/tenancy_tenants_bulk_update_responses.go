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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// TenancyTenantsBulkUpdateReader is a Reader for the TenancyTenantsBulkUpdate structure.
type TenancyTenantsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantsBulkUpdateOK creates a TenancyTenantsBulkUpdateOK with default headers values
func NewTenancyTenantsBulkUpdateOK() *TenancyTenantsBulkUpdateOK {
	return &TenancyTenantsBulkUpdateOK{}
}

/* TenancyTenantsBulkUpdateOK describes a response with status code 200, with default header values.

TenancyTenantsBulkUpdateOK tenancy tenants bulk update o k
*/
type TenancyTenantsBulkUpdateOK struct {
	Payload *models.Tenant
}

func (o *TenancyTenantsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenants/][%d] tenancyTenantsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *TenancyTenantsBulkUpdateOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *TenancyTenantsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantsBulkUpdateDefault creates a TenancyTenantsBulkUpdateDefault with default headers values
func NewTenancyTenantsBulkUpdateDefault(code int) *TenancyTenantsBulkUpdateDefault {
	return &TenancyTenantsBulkUpdateDefault{
		_statusCode: code,
	}
}

/* TenancyTenantsBulkUpdateDefault describes a response with status code -1, with default header values.

TenancyTenantsBulkUpdateDefault tenancy tenants bulk update default
*/
type TenancyTenantsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy tenants bulk update default response
func (o *TenancyTenantsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyTenantsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenants/][%d] tenancy_tenants_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *TenancyTenantsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
