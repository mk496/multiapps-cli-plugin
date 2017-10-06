package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetBreakpointsParams creates a new GetBreakpointsParams object
// with the default values initialized.
func NewGetBreakpointsParams() *GetBreakpointsParams {

	return &GetBreakpointsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetBreakpointsParamsWithTimeout creates a new GetBreakpointsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetBreakpointsParamsWithTimeout(timeout time.Duration) *GetBreakpointsParams {

	return &GetBreakpointsParams{

		timeout: timeout,
	}
}

// NewGetBreakpointsParamsWithContext creates a new GetBreakpointsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetBreakpointsParamsWithContext(ctx context.Context) *GetBreakpointsParams {

	return &GetBreakpointsParams{

		Context: ctx,
	}
}

/*GetBreakpointsParams contains all the parameters to send to the API endpoint
for the get breakpoints operation typically these are written to a http.Request
*/
type GetBreakpointsParams struct {
	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the get breakpoints params
func (o *GetBreakpointsParams) WithTimeout(timeout time.Duration) *GetBreakpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get breakpoints params
func (o *GetBreakpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get breakpoints params
func (o *GetBreakpointsParams) WithContext(ctx context.Context) *GetBreakpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get breakpoints params
func (o *GetBreakpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *GetBreakpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}