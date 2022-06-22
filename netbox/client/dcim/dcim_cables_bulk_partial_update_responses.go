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

// DcimCablesBulkPartialUpdateReader is a Reader for the DcimCablesBulkPartialUpdate structure.
type DcimCablesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimCablesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCablesBulkPartialUpdateOK creates a DcimCablesBulkPartialUpdateOK with default headers values
func NewDcimCablesBulkPartialUpdateOK() *DcimCablesBulkPartialUpdateOK {
	return &DcimCablesBulkPartialUpdateOK{}
}

/* DcimCablesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimCablesBulkPartialUpdateOK dcim cables bulk partial update o k
*/
type DcimCablesBulkPartialUpdateOK struct {
	Payload *models.Cable
}

func (o *DcimCablesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/cables/][%d] dcimCablesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimCablesBulkPartialUpdateOK) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCablesBulkPartialUpdateDefault creates a DcimCablesBulkPartialUpdateDefault with default headers values
func NewDcimCablesBulkPartialUpdateDefault(code int) *DcimCablesBulkPartialUpdateDefault {
	return &DcimCablesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimCablesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimCablesBulkPartialUpdateDefault dcim cables bulk partial update default
*/
type DcimCablesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cables bulk partial update default response
func (o *DcimCablesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimCablesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/cables/][%d] dcim_cables_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimCablesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
