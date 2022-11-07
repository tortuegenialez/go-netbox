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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// UsersPermissionsCreateReader is a Reader for the UsersPermissionsCreate structure.
type UsersPermissionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPermissionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersPermissionsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersPermissionsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersPermissionsCreateCreated creates a UsersPermissionsCreateCreated with default headers values
func NewUsersPermissionsCreateCreated() *UsersPermissionsCreateCreated {
	return &UsersPermissionsCreateCreated{}
}

/*
UsersPermissionsCreateCreated describes a response with status code 201, with default header values.

UsersPermissionsCreateCreated users permissions create created
*/
type UsersPermissionsCreateCreated struct {
	Payload *models.ObjectPermission
}

// IsSuccess returns true when this users permissions create created response has a 2xx status code
func (o *UsersPermissionsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users permissions create created response has a 3xx status code
func (o *UsersPermissionsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users permissions create created response has a 4xx status code
func (o *UsersPermissionsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users permissions create created response has a 5xx status code
func (o *UsersPermissionsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users permissions create created response a status code equal to that given
func (o *UsersPermissionsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UsersPermissionsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /users/permissions/][%d] usersPermissionsCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersPermissionsCreateCreated) String() string {
	return fmt.Sprintf("[POST /users/permissions/][%d] usersPermissionsCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersPermissionsCreateCreated) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *UsersPermissionsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersPermissionsCreateDefault creates a UsersPermissionsCreateDefault with default headers values
func NewUsersPermissionsCreateDefault(code int) *UsersPermissionsCreateDefault {
	return &UsersPermissionsCreateDefault{
		_statusCode: code,
	}
}

/*
UsersPermissionsCreateDefault describes a response with status code -1, with default header values.

UsersPermissionsCreateDefault users permissions create default
*/
type UsersPermissionsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users permissions create default response
func (o *UsersPermissionsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users permissions create default response has a 2xx status code
func (o *UsersPermissionsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users permissions create default response has a 3xx status code
func (o *UsersPermissionsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users permissions create default response has a 4xx status code
func (o *UsersPermissionsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users permissions create default response has a 5xx status code
func (o *UsersPermissionsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users permissions create default response a status code equal to that given
func (o *UsersPermissionsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersPermissionsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /users/permissions/][%d] users_permissions_create default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsCreateDefault) String() string {
	return fmt.Sprintf("[POST /users/permissions/][%d] users_permissions_create default  %+v", o._statusCode, o.Payload)
}

func (o *UsersPermissionsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersPermissionsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
