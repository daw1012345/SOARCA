// Code generated by go-swagger; DO NOT EDIT.

package payloads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"soarca/pkg/core/capability/caldera/api/models"
)

// GetAPIV2PayloadsReader is a Reader for the GetAPIV2Payloads structure.
type GetAPIV2PayloadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2PayloadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2PayloadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/payloads] GetAPIV2Payloads", response, response.Code())
	}
}

// NewGetAPIV2PayloadsOK creates a GetAPIV2PayloadsOK with default headers values
func NewGetAPIV2PayloadsOK() *GetAPIV2PayloadsOK {
	return &GetAPIV2PayloadsOK{}
}

/*
GetAPIV2PayloadsOK describes a response with status code 200, with default header values.

Returns a list of all payloads in PayloadSchema format.
*/
type GetAPIV2PayloadsOK struct {
	Payload *models.Payload
}

// IsSuccess returns true when this get Api v2 payloads o k response has a 2xx status code
func (o *GetAPIV2PayloadsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 payloads o k response has a 3xx status code
func (o *GetAPIV2PayloadsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 payloads o k response has a 4xx status code
func (o *GetAPIV2PayloadsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 payloads o k response has a 5xx status code
func (o *GetAPIV2PayloadsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 payloads o k response a status code equal to that given
func (o *GetAPIV2PayloadsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v2 payloads o k response
func (o *GetAPIV2PayloadsOK) Code() int {
	return 200
}

func (o *GetAPIV2PayloadsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/payloads][%d] getApiV2PayloadsOK %s", 200, payload)
}

func (o *GetAPIV2PayloadsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/payloads][%d] getApiV2PayloadsOK %s", 200, payload)
}

func (o *GetAPIV2PayloadsOK) GetPayload() *models.Payload {
	return o.Payload
}

func (o *GetAPIV2PayloadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Payload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
