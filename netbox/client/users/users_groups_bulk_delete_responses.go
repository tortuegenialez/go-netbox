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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersGroupsBulkDeleteReader is a Reader for the UsersGroupsBulkDelete structure.
type UsersGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUsersGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersGroupsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersGroupsBulkDeleteNoContent creates a UsersGroupsBulkDeleteNoContent with default headers values
func NewUsersGroupsBulkDeleteNoContent() *UsersGroupsBulkDeleteNoContent {
	return &UsersGroupsBulkDeleteNoContent{}
}

/*
UsersGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

UsersGroupsBulkDeleteNoContent users groups bulk delete no content
*/
type UsersGroupsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this users groups bulk delete no content response has a 2xx status code
func (o *UsersGroupsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users groups bulk delete no content response has a 3xx status code
func (o *UsersGroupsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users groups bulk delete no content response has a 4xx status code
func (o *UsersGroupsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this users groups bulk delete no content response has a 5xx status code
func (o *UsersGroupsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this users groups bulk delete no content response a status code equal to that given
func (o *UsersGroupsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *UsersGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/groups/][%d] usersGroupsBulkDeleteNoContent ", 204)
}

func (o *UsersGroupsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /users/groups/][%d] usersGroupsBulkDeleteNoContent ", 204)
}

func (o *UsersGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersGroupsBulkDeleteDefault creates a UsersGroupsBulkDeleteDefault with default headers values
func NewUsersGroupsBulkDeleteDefault(code int) *UsersGroupsBulkDeleteDefault {
	return &UsersGroupsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
UsersGroupsBulkDeleteDefault describes a response with status code -1, with default header values.

UsersGroupsBulkDeleteDefault users groups bulk delete default
*/
type UsersGroupsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users groups bulk delete default response
func (o *UsersGroupsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users groups bulk delete default response has a 2xx status code
func (o *UsersGroupsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users groups bulk delete default response has a 3xx status code
func (o *UsersGroupsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users groups bulk delete default response has a 4xx status code
func (o *UsersGroupsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users groups bulk delete default response has a 5xx status code
func (o *UsersGroupsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users groups bulk delete default response a status code equal to that given
func (o *UsersGroupsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersGroupsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/groups/][%d] users_groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *UsersGroupsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /users/groups/][%d] users_groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *UsersGroupsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersGroupsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
