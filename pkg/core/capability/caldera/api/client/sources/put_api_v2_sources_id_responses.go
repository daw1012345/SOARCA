// Code generated by go-swagger; DO NOT EDIT.

package sources

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

// PutAPIV2SourcesIDReader is a Reader for the PutAPIV2SourcesID structure.
type PutAPIV2SourcesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV2SourcesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPIV2SourcesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /api/v2/sources/{id}] PutAPIV2SourcesID", response, response.Code())
	}
}

// NewPutAPIV2SourcesIDOK creates a PutAPIV2SourcesIDOK with default headers values
func NewPutAPIV2SourcesIDOK() *PutAPIV2SourcesIDOK {
	return &PutAPIV2SourcesIDOK{}
}

/*
PutAPIV2SourcesIDOK describes a response with status code 200, with default header values.

Returns a single Source dumped in SourceSchema format.
*/
type PutAPIV2SourcesIDOK struct {
	Payload *models.Source
}

// IsSuccess returns true when this put Api v2 sources Id o k response has a 2xx status code
func (o *PutAPIV2SourcesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put Api v2 sources Id o k response has a 3xx status code
func (o *PutAPIV2SourcesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put Api v2 sources Id o k response has a 4xx status code
func (o *PutAPIV2SourcesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put Api v2 sources Id o k response has a 5xx status code
func (o *PutAPIV2SourcesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put Api v2 sources Id o k response a status code equal to that given
func (o *PutAPIV2SourcesIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put Api v2 sources Id o k response
func (o *PutAPIV2SourcesIDOK) Code() int {
	return 200
}

func (o *PutAPIV2SourcesIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/sources/{id}][%d] putApiV2SourcesIdOK %s", 200, payload)
}

func (o *PutAPIV2SourcesIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/sources/{id}][%d] putApiV2SourcesIdOK %s", 200, payload)
}

func (o *PutAPIV2SourcesIDOK) GetPayload() *models.Source {
	return o.Payload
}

func (o *PutAPIV2SourcesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Source)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
