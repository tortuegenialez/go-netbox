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

// IpamRirsBulkDeleteReader is a Reader for the IpamRirsBulkDelete structure.
type IpamRirsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRirsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRirsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRirsBulkDeleteNoContent creates a IpamRirsBulkDeleteNoContent with default headers values
func NewIpamRirsBulkDeleteNoContent() *IpamRirsBulkDeleteNoContent {
	return &IpamRirsBulkDeleteNoContent{}
}

/*
IpamRirsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamRirsBulkDeleteNoContent ipam rirs bulk delete no content
*/
type IpamRirsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this ipam rirs bulk delete no content response has a 2xx status code
func (o *IpamRirsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam rirs bulk delete no content response has a 3xx status code
func (o *IpamRirsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs bulk delete no content response has a 4xx status code
func (o *IpamRirsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam rirs bulk delete no content response has a 5xx status code
func (o *IpamRirsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs bulk delete no content response a status code equal to that given
func (o *IpamRirsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamRirsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipamRirsBulkDeleteNoContent ", 204)
}

func (o *IpamRirsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipamRirsBulkDeleteNoContent ", 204)
}

func (o *IpamRirsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamRirsBulkDeleteDefault creates a IpamRirsBulkDeleteDefault with default headers values
func NewIpamRirsBulkDeleteDefault(code int) *IpamRirsBulkDeleteDefault {
	return &IpamRirsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamRirsBulkDeleteDefault describes a response with status code -1, with default header values.

IpamRirsBulkDeleteDefault ipam rirs bulk delete default
*/
type IpamRirsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam rirs bulk delete default response
func (o *IpamRirsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam rirs bulk delete default response has a 2xx status code
func (o *IpamRirsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam rirs bulk delete default response has a 3xx status code
func (o *IpamRirsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam rirs bulk delete default response has a 4xx status code
func (o *IpamRirsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam rirs bulk delete default response has a 5xx status code
func (o *IpamRirsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam rirs bulk delete default response a status code equal to that given
func (o *IpamRirsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamRirsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipam_rirs_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipam_rirs_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRirsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
