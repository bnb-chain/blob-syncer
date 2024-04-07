// Code generated by go-swagger; DO NOT EDIT.

package blob

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetBlobSidecarsByBlockNumParams creates a new GetBlobSidecarsByBlockNumParams object
//
// There are no default values defined in the spec.
func NewGetBlobSidecarsByBlockNumParams() GetBlobSidecarsByBlockNumParams {

	return GetBlobSidecarsByBlockNumParams{}
}

// GetBlobSidecarsByBlockNumParams contains all the bound params for the get blob sidecars by block num operation
// typically these are obtained from a http.Request
//
// swagger:parameters getBlobSidecarsByBlockNum
type GetBlobSidecarsByBlockNumParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*blockNum
	  Required: true
	  In: path
	*/
	BlockNum int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetBlobSidecarsByBlockNumParams() beforehand.
func (o *GetBlobSidecarsByBlockNumParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rBlockNum, rhkBlockNum, _ := route.Params.GetOK("blockNum")
	if err := o.bindBlockNum(rBlockNum, rhkBlockNum, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindBlockNum binds and validates parameter BlockNum from path.
func (o *GetBlobSidecarsByBlockNumParams) bindBlockNum(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("blockNum", "path", "int64", raw)
	}
	o.BlockNum = value

	return nil
}