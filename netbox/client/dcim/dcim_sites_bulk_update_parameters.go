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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// NewDcimSitesBulkUpdateParams creates a new DcimSitesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSitesBulkUpdateParams() *DcimSitesBulkUpdateParams {
	return &DcimSitesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSitesBulkUpdateParamsWithTimeout creates a new DcimSitesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimSitesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimSitesBulkUpdateParams {
	return &DcimSitesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimSitesBulkUpdateParamsWithContext creates a new DcimSitesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimSitesBulkUpdateParamsWithContext(ctx context.Context) *DcimSitesBulkUpdateParams {
	return &DcimSitesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimSitesBulkUpdateParamsWithHTTPClient creates a new DcimSitesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSitesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimSitesBulkUpdateParams {
	return &DcimSitesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimSitesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim sites bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimSitesBulkUpdateParams struct {

	// Data.
	Data *models.WritableSite

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim sites bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesBulkUpdateParams) WithDefaults() *DcimSitesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim sites bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim sites bulk update params
func (o *DcimSitesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimSitesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim sites bulk update params
func (o *DcimSitesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim sites bulk update params
func (o *DcimSitesBulkUpdateParams) WithContext(ctx context.Context) *DcimSitesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim sites bulk update params
func (o *DcimSitesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim sites bulk update params
func (o *DcimSitesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimSitesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim sites bulk update params
func (o *DcimSitesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim sites bulk update params
func (o *DcimSitesBulkUpdateParams) WithData(data *models.WritableSite) *DcimSitesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim sites bulk update params
func (o *DcimSitesBulkUpdateParams) SetData(data *models.WritableSite) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSitesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
