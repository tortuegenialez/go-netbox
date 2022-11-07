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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// IpamFhrpGroupsBulkUpdateReader is a Reader for the IpamFhrpGroupsBulkUpdate structure.
type IpamFhrpGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamFhrpGroupsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupsBulkUpdateOK creates a IpamFhrpGroupsBulkUpdateOK with default headers values
func NewIpamFhrpGroupsBulkUpdateOK() *IpamFhrpGroupsBulkUpdateOK {
	return &IpamFhrpGroupsBulkUpdateOK{}
}

/*
IpamFhrpGroupsBulkUpdateOK describes a response with status code 200, with default header values.

IpamFhrpGroupsBulkUpdateOK ipam fhrp groups bulk update o k
*/
type IpamFhrpGroupsBulkUpdateOK struct {
	Payload *models.FHRPGroup
}

// IsSuccess returns true when this ipam fhrp groups bulk update o k response has a 2xx status code
func (o *IpamFhrpGroupsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp groups bulk update o k response has a 3xx status code
func (o *IpamFhrpGroupsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp groups bulk update o k response has a 4xx status code
func (o *IpamFhrpGroupsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp groups bulk update o k response has a 5xx status code
func (o *IpamFhrpGroupsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp groups bulk update o k response a status code equal to that given
func (o *IpamFhrpGroupsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamFhrpGroupsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/fhrp-groups/][%d] ipamFhrpGroupsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/fhrp-groups/][%d] ipamFhrpGroupsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupsBulkUpdateOK) GetPayload() *models.FHRPGroup {
	return o.Payload
}

func (o *IpamFhrpGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupsBulkUpdateDefault creates a IpamFhrpGroupsBulkUpdateDefault with default headers values
func NewIpamFhrpGroupsBulkUpdateDefault(code int) *IpamFhrpGroupsBulkUpdateDefault {
	return &IpamFhrpGroupsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamFhrpGroupsBulkUpdateDefault describes a response with status code -1, with default header values.

IpamFhrpGroupsBulkUpdateDefault ipam fhrp groups bulk update default
*/
type IpamFhrpGroupsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam fhrp groups bulk update default response
func (o *IpamFhrpGroupsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam fhrp groups bulk update default response has a 2xx status code
func (o *IpamFhrpGroupsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam fhrp groups bulk update default response has a 3xx status code
func (o *IpamFhrpGroupsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam fhrp groups bulk update default response has a 4xx status code
func (o *IpamFhrpGroupsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam fhrp groups bulk update default response has a 5xx status code
func (o *IpamFhrpGroupsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam fhrp groups bulk update default response a status code equal to that given
func (o *IpamFhrpGroupsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamFhrpGroupsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/fhrp-groups/][%d] ipam_fhrp-groups_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/fhrp-groups/][%d] ipam_fhrp-groups_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
