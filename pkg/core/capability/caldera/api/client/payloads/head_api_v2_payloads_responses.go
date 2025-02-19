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

// HeadAPIV2PayloadsReader is a Reader for the HeadAPIV2Payloads structure.
type HeadAPIV2PayloadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadAPIV2PayloadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadAPIV2PayloadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[HEAD /api/v2/payloads] HeadAPIV2Payloads", response, response.Code())
	}
}

// NewHeadAPIV2PayloadsOK creates a HeadAPIV2PayloadsOK with default headers values
func NewHeadAPIV2PayloadsOK() *HeadAPIV2PayloadsOK {
	return &HeadAPIV2PayloadsOK{}
}

/*
HeadAPIV2PayloadsOK describes a response with status code 200, with default header values.

Returns a list of all payloads in PayloadSchema format.
*/
type HeadAPIV2PayloadsOK struct {
	Payload *models.Payload
}

// IsSuccess returns true when this head Api v2 payloads o k response has a 2xx status code
func (o *HeadAPIV2PayloadsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this head Api v2 payloads o k response has a 3xx status code
func (o *HeadAPIV2PayloadsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this head Api v2 payloads o k response has a 4xx status code
func (o *HeadAPIV2PayloadsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this head Api v2 payloads o k response has a 5xx status code
func (o *HeadAPIV2PayloadsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this head Api v2 payloads o k response a status code equal to that given
func (o *HeadAPIV2PayloadsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the head Api v2 payloads o k response
func (o *HeadAPIV2PayloadsOK) Code() int {
	return 200
}

func (o *HeadAPIV2PayloadsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/payloads][%d] headApiV2PayloadsOK %s", 200, payload)
}

func (o *HeadAPIV2PayloadsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/payloads][%d] headApiV2PayloadsOK %s", 200, payload)
}

func (o *HeadAPIV2PayloadsOK) GetPayload() *models.Payload {
	return o.Payload
}

func (o *HeadAPIV2PayloadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Payload)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
