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

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimModuleBayTemplatesUpdateReader is a Reader for the DcimModuleBayTemplatesUpdate structure.
type DcimModuleBayTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBayTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBayTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBayTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBayTemplatesUpdateOK creates a DcimModuleBayTemplatesUpdateOK with default headers values
func NewDcimModuleBayTemplatesUpdateOK() *DcimModuleBayTemplatesUpdateOK {
	return &DcimModuleBayTemplatesUpdateOK{}
}

/*
DcimModuleBayTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimModuleBayTemplatesUpdateOK dcim module bay templates update o k
*/
type DcimModuleBayTemplatesUpdateOK struct {
	Payload *models.ModuleBayTemplate
}

// IsSuccess returns true when this dcim module bay templates update o k response has a 2xx status code
func (o *DcimModuleBayTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bay templates update o k response has a 3xx status code
func (o *DcimModuleBayTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bay templates update o k response has a 4xx status code
func (o *DcimModuleBayTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bay templates update o k response has a 5xx status code
func (o *DcimModuleBayTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bay templates update o k response a status code equal to that given
func (o *DcimModuleBayTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModuleBayTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBayTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBayTemplatesUpdateOK) GetPayload() *models.ModuleBayTemplate {
	return o.Payload
}

func (o *DcimModuleBayTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBayTemplatesUpdateDefault creates a DcimModuleBayTemplatesUpdateDefault with default headers values
func NewDcimModuleBayTemplatesUpdateDefault(code int) *DcimModuleBayTemplatesUpdateDefault {
	return &DcimModuleBayTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBayTemplatesUpdateDefault describes a response with status code -1, with default header values.

DcimModuleBayTemplatesUpdateDefault dcim module bay templates update default
*/
type DcimModuleBayTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bay templates update default response
func (o *DcimModuleBayTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bay templates update default response has a 2xx status code
func (o *DcimModuleBayTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bay templates update default response has a 3xx status code
func (o *DcimModuleBayTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bay templates update default response has a 4xx status code
func (o *DcimModuleBayTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bay templates update default response has a 5xx status code
func (o *DcimModuleBayTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bay templates update default response a status code equal to that given
func (o *DcimModuleBayTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBayTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bay-templates/{id}/][%d] dcim_module-bay-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/module-bay-templates/{id}/][%d] dcim_module-bay-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBayTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
