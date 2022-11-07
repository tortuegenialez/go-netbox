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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimPowerPanelsDeleteReader is a Reader for the DcimPowerPanelsDelete structure.
type DcimPowerPanelsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerPanelsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPanelsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPanelsDeleteNoContent creates a DcimPowerPanelsDeleteNoContent with default headers values
func NewDcimPowerPanelsDeleteNoContent() *DcimPowerPanelsDeleteNoContent {
	return &DcimPowerPanelsDeleteNoContent{}
}

/*
DcimPowerPanelsDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerPanelsDeleteNoContent dcim power panels delete no content
*/
type DcimPowerPanelsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim power panels delete no content response has a 2xx status code
func (o *DcimPowerPanelsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power panels delete no content response has a 3xx status code
func (o *DcimPowerPanelsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power panels delete no content response has a 4xx status code
func (o *DcimPowerPanelsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power panels delete no content response has a 5xx status code
func (o *DcimPowerPanelsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power panels delete no content response a status code equal to that given
func (o *DcimPowerPanelsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimPowerPanelsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-panels/{id}/][%d] dcimPowerPanelsDeleteNoContent ", 204)
}

func (o *DcimPowerPanelsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-panels/{id}/][%d] dcimPowerPanelsDeleteNoContent ", 204)
}

func (o *DcimPowerPanelsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimPowerPanelsDeleteDefault creates a DcimPowerPanelsDeleteDefault with default headers values
func NewDcimPowerPanelsDeleteDefault(code int) *DcimPowerPanelsDeleteDefault {
	return &DcimPowerPanelsDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPanelsDeleteDefault describes a response with status code -1, with default header values.

DcimPowerPanelsDeleteDefault dcim power panels delete default
*/
type DcimPowerPanelsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power panels delete default response
func (o *DcimPowerPanelsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power panels delete default response has a 2xx status code
func (o *DcimPowerPanelsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power panels delete default response has a 3xx status code
func (o *DcimPowerPanelsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power panels delete default response has a 4xx status code
func (o *DcimPowerPanelsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power panels delete default response has a 5xx status code
func (o *DcimPowerPanelsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power panels delete default response a status code equal to that given
func (o *DcimPowerPanelsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerPanelsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-panels/{id}/][%d] dcim_power-panels_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-panels/{id}/][%d] dcim_power-panels_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
