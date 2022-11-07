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

// IpamVlanGroupsReadReader is a Reader for the IpamVlanGroupsRead structure.
type IpamVlanGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVlanGroupsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsReadOK creates a IpamVlanGroupsReadOK with default headers values
func NewIpamVlanGroupsReadOK() *IpamVlanGroupsReadOK {
	return &IpamVlanGroupsReadOK{}
}

/*
IpamVlanGroupsReadOK describes a response with status code 200, with default header values.

IpamVlanGroupsReadOK ipam vlan groups read o k
*/
type IpamVlanGroupsReadOK struct {
	Payload *models.VLANGroup
}

// IsSuccess returns true when this ipam vlan groups read o k response has a 2xx status code
func (o *IpamVlanGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups read o k response has a 3xx status code
func (o *IpamVlanGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups read o k response has a 4xx status code
func (o *IpamVlanGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups read o k response has a 5xx status code
func (o *IpamVlanGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups read o k response a status code equal to that given
func (o *IpamVlanGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamVlanGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsReadOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsReadOK) String() string {
	return fmt.Sprintf("[GET /ipam/vlan-groups/{id}/][%d] ipamVlanGroupsReadOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsReadOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsReadDefault creates a IpamVlanGroupsReadDefault with default headers values
func NewIpamVlanGroupsReadDefault(code int) *IpamVlanGroupsReadDefault {
	return &IpamVlanGroupsReadDefault{
		_statusCode: code,
	}
}

/*
IpamVlanGroupsReadDefault describes a response with status code -1, with default header values.

IpamVlanGroupsReadDefault ipam vlan groups read default
*/
type IpamVlanGroupsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlan groups read default response
func (o *IpamVlanGroupsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam vlan groups read default response has a 2xx status code
func (o *IpamVlanGroupsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlan groups read default response has a 3xx status code
func (o *IpamVlanGroupsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlan groups read default response has a 4xx status code
func (o *IpamVlanGroupsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlan groups read default response has a 5xx status code
func (o *IpamVlanGroupsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlan groups read default response a status code equal to that given
func (o *IpamVlanGroupsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamVlanGroupsReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/vlan-groups/{id}/][%d] ipam_vlan-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsReadDefault) String() string {
	return fmt.Sprintf("[GET /ipam/vlan-groups/{id}/][%d] ipam_vlan-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
