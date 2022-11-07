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

// DcimPowerFeedsBulkPartialUpdateReader is a Reader for the DcimPowerFeedsBulkPartialUpdate structure.
type DcimPowerFeedsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerFeedsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsBulkPartialUpdateOK creates a DcimPowerFeedsBulkPartialUpdateOK with default headers values
func NewDcimPowerFeedsBulkPartialUpdateOK() *DcimPowerFeedsBulkPartialUpdateOK {
	return &DcimPowerFeedsBulkPartialUpdateOK{}
}

/*
DcimPowerFeedsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerFeedsBulkPartialUpdateOK dcim power feeds bulk partial update o k
*/
type DcimPowerFeedsBulkPartialUpdateOK struct {
	Payload *models.PowerFeed
}

// IsSuccess returns true when this dcim power feeds bulk partial update o k response has a 2xx status code
func (o *DcimPowerFeedsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power feeds bulk partial update o k response has a 3xx status code
func (o *DcimPowerFeedsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power feeds bulk partial update o k response has a 4xx status code
func (o *DcimPowerFeedsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power feeds bulk partial update o k response has a 5xx status code
func (o *DcimPowerFeedsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power feeds bulk partial update o k response a status code equal to that given
func (o *DcimPowerFeedsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerFeedsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-feeds/][%d] dcimPowerFeedsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-feeds/][%d] dcimPowerFeedsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsBulkPartialUpdateOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsBulkPartialUpdateDefault creates a DcimPowerFeedsBulkPartialUpdateDefault with default headers values
func NewDcimPowerFeedsBulkPartialUpdateDefault(code int) *DcimPowerFeedsBulkPartialUpdateDefault {
	return &DcimPowerFeedsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerFeedsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerFeedsBulkPartialUpdateDefault dcim power feeds bulk partial update default
*/
type DcimPowerFeedsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds bulk partial update default response
func (o *DcimPowerFeedsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power feeds bulk partial update default response has a 2xx status code
func (o *DcimPowerFeedsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power feeds bulk partial update default response has a 3xx status code
func (o *DcimPowerFeedsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power feeds bulk partial update default response has a 4xx status code
func (o *DcimPowerFeedsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power feeds bulk partial update default response has a 5xx status code
func (o *DcimPowerFeedsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power feeds bulk partial update default response a status code equal to that given
func (o *DcimPowerFeedsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerFeedsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-feeds/][%d] dcim_power-feeds_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-feeds/][%d] dcim_power-feeds_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
