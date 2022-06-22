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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasJournalEntriesDeleteReader is a Reader for the ExtrasJournalEntriesDelete structure.
type ExtrasJournalEntriesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasJournalEntriesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasJournalEntriesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasJournalEntriesDeleteNoContent creates a ExtrasJournalEntriesDeleteNoContent with default headers values
func NewExtrasJournalEntriesDeleteNoContent() *ExtrasJournalEntriesDeleteNoContent {
	return &ExtrasJournalEntriesDeleteNoContent{}
}

/* ExtrasJournalEntriesDeleteNoContent describes a response with status code 204, with default header values.

ExtrasJournalEntriesDeleteNoContent extras journal entries delete no content
*/
type ExtrasJournalEntriesDeleteNoContent struct {
}

func (o *ExtrasJournalEntriesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/journal-entries/{id}/][%d] extrasJournalEntriesDeleteNoContent ", 204)
}

func (o *ExtrasJournalEntriesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasJournalEntriesDeleteDefault creates a ExtrasJournalEntriesDeleteDefault with default headers values
func NewExtrasJournalEntriesDeleteDefault(code int) *ExtrasJournalEntriesDeleteDefault {
	return &ExtrasJournalEntriesDeleteDefault{
		_statusCode: code,
	}
}

/* ExtrasJournalEntriesDeleteDefault describes a response with status code -1, with default header values.

ExtrasJournalEntriesDeleteDefault extras journal entries delete default
*/
type ExtrasJournalEntriesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras journal entries delete default response
func (o *ExtrasJournalEntriesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasJournalEntriesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/journal-entries/{id}/][%d] extras_journal-entries_delete default  %+v", o._statusCode, o.Payload)
}
func (o *ExtrasJournalEntriesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasJournalEntriesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
