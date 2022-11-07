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
)

// IpamPrefixesBulkDeleteReader is a Reader for the IpamPrefixesBulkDelete structure.
type IpamPrefixesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamPrefixesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamPrefixesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesBulkDeleteNoContent creates a IpamPrefixesBulkDeleteNoContent with default headers values
func NewIpamPrefixesBulkDeleteNoContent() *IpamPrefixesBulkDeleteNoContent {
	return &IpamPrefixesBulkDeleteNoContent{}
}

/*
IpamPrefixesBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamPrefixesBulkDeleteNoContent ipam prefixes bulk delete no content
*/
type IpamPrefixesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this ipam prefixes bulk delete no content response has a 2xx status code
func (o *IpamPrefixesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes bulk delete no content response has a 3xx status code
func (o *IpamPrefixesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes bulk delete no content response has a 4xx status code
func (o *IpamPrefixesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes bulk delete no content response has a 5xx status code
func (o *IpamPrefixesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes bulk delete no content response a status code equal to that given
func (o *IpamPrefixesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamPrefixesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/prefixes/][%d] ipamPrefixesBulkDeleteNoContent ", 204)
}

func (o *IpamPrefixesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/prefixes/][%d] ipamPrefixesBulkDeleteNoContent ", 204)
}

func (o *IpamPrefixesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamPrefixesBulkDeleteDefault creates a IpamPrefixesBulkDeleteDefault with default headers values
func NewIpamPrefixesBulkDeleteDefault(code int) *IpamPrefixesBulkDeleteDefault {
	return &IpamPrefixesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamPrefixesBulkDeleteDefault describes a response with status code -1, with default header values.

IpamPrefixesBulkDeleteDefault ipam prefixes bulk delete default
*/
type IpamPrefixesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam prefixes bulk delete default response
func (o *IpamPrefixesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam prefixes bulk delete default response has a 2xx status code
func (o *IpamPrefixesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam prefixes bulk delete default response has a 3xx status code
func (o *IpamPrefixesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam prefixes bulk delete default response has a 4xx status code
func (o *IpamPrefixesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam prefixes bulk delete default response has a 5xx status code
func (o *IpamPrefixesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam prefixes bulk delete default response a status code equal to that given
func (o *IpamPrefixesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamPrefixesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/prefixes/][%d] ipam_prefixes_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/prefixes/][%d] ipam_prefixes_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
