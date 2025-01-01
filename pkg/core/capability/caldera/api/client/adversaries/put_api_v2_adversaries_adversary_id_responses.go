// Code generated by go-swagger; DO NOT EDIT.

package adversaries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"soarca/internal/capability/caldera/api/models"
)

// PutAPIV2AdversariesAdversaryIDReader is a Reader for the PutAPIV2AdversariesAdversaryID structure.
type PutAPIV2AdversariesAdversaryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV2AdversariesAdversaryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPIV2AdversariesAdversaryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /api/v2/adversaries/{adversary_id}] PutAPIV2AdversariesAdversaryID", response, response.Code())
	}
}

// NewPutAPIV2AdversariesAdversaryIDOK creates a PutAPIV2AdversariesAdversaryIDOK with default headers values
func NewPutAPIV2AdversariesAdversaryIDOK() *PutAPIV2AdversariesAdversaryIDOK {
	return &PutAPIV2AdversariesAdversaryIDOK{}
}

/*
PutAPIV2AdversariesAdversaryIDOK describes a response with status code 200, with default header values.

A single adversary, either newly created or updated, in AdversarySchema format.
*/
type PutAPIV2AdversariesAdversaryIDOK struct {
	Payload *models.Adversary
}

// IsSuccess returns true when this put Api v2 adversaries adversary Id o k response has a 2xx status code
func (o *PutAPIV2AdversariesAdversaryIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put Api v2 adversaries adversary Id o k response has a 3xx status code
func (o *PutAPIV2AdversariesAdversaryIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put Api v2 adversaries adversary Id o k response has a 4xx status code
func (o *PutAPIV2AdversariesAdversaryIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put Api v2 adversaries adversary Id o k response has a 5xx status code
func (o *PutAPIV2AdversariesAdversaryIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put Api v2 adversaries adversary Id o k response a status code equal to that given
func (o *PutAPIV2AdversariesAdversaryIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put Api v2 adversaries adversary Id o k response
func (o *PutAPIV2AdversariesAdversaryIDOK) Code() int {
	return 200
}

func (o *PutAPIV2AdversariesAdversaryIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/adversaries/{adversary_id}][%d] putApiV2AdversariesAdversaryIdOK %s", 200, payload)
}

func (o *PutAPIV2AdversariesAdversaryIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/adversaries/{adversary_id}][%d] putApiV2AdversariesAdversaryIdOK %s", 200, payload)
}

func (o *PutAPIV2AdversariesAdversaryIDOK) GetPayload() *models.Adversary {
	return o.Payload
}

func (o *PutAPIV2AdversariesAdversaryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Adversary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
