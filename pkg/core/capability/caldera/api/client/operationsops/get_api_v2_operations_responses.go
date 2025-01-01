// Code generated by go-swagger; DO NOT EDIT.

package operationsops

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

// GetAPIV2OperationsReader is a Reader for the GetAPIV2Operations structure.
type GetAPIV2OperationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2OperationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2OperationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/operations] GetAPIV2Operations", response, response.Code())
	}
}

// NewGetAPIV2OperationsOK creates a GetAPIV2OperationsOK with default headers values
func NewGetAPIV2OperationsOK() *GetAPIV2OperationsOK {
	return &GetAPIV2OperationsOK{}
}

/*
GetAPIV2OperationsOK describes a response with status code 200, with default header values.

The response is a list of all operations.
*/
type GetAPIV2OperationsOK struct {
	Payload []*models.PartialOperation
}

// IsSuccess returns true when this get Api v2 operations o k response has a 2xx status code
func (o *GetAPIV2OperationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 operations o k response has a 3xx status code
func (o *GetAPIV2OperationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 operations o k response has a 4xx status code
func (o *GetAPIV2OperationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 operations o k response has a 5xx status code
func (o *GetAPIV2OperationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 operations o k response a status code equal to that given
func (o *GetAPIV2OperationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v2 operations o k response
func (o *GetAPIV2OperationsOK) Code() int {
	return 200
}

func (o *GetAPIV2OperationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/operations][%d] getApiV2OperationsOK %s", 200, payload)
}

func (o *GetAPIV2OperationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/operations][%d] getApiV2OperationsOK %s", 200, payload)
}

func (o *GetAPIV2OperationsOK) GetPayload() []*models.PartialOperation {
	return o.Payload
}

func (o *GetAPIV2OperationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
