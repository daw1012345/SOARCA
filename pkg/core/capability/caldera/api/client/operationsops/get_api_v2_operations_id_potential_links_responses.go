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

	"soarca/internal/capability/caldera/api/models"
)

// GetAPIV2OperationsIDPotentialLinksReader is a Reader for the GetAPIV2OperationsIDPotentialLinks structure.
type GetAPIV2OperationsIDPotentialLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2OperationsIDPotentialLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2OperationsIDPotentialLinksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/operations/{id}/potential-links] GetAPIV2OperationsIDPotentialLinks", response, response.Code())
	}
}

// NewGetAPIV2OperationsIDPotentialLinksOK creates a GetAPIV2OperationsIDPotentialLinksOK with default headers values
func NewGetAPIV2OperationsIDPotentialLinksOK() *GetAPIV2OperationsIDPotentialLinksOK {
	return &GetAPIV2OperationsIDPotentialLinksOK{}
}

/*
GetAPIV2OperationsIDPotentialLinksOK describes a response with status code 200, with default header values.

Response contains a list of link objects for the requested id.
*/
type GetAPIV2OperationsIDPotentialLinksOK struct {
	Payload []*models.PartialLink
}

// IsSuccess returns true when this get Api v2 operations Id potential links o k response has a 2xx status code
func (o *GetAPIV2OperationsIDPotentialLinksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 operations Id potential links o k response has a 3xx status code
func (o *GetAPIV2OperationsIDPotentialLinksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 operations Id potential links o k response has a 4xx status code
func (o *GetAPIV2OperationsIDPotentialLinksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 operations Id potential links o k response has a 5xx status code
func (o *GetAPIV2OperationsIDPotentialLinksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 operations Id potential links o k response a status code equal to that given
func (o *GetAPIV2OperationsIDPotentialLinksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v2 operations Id potential links o k response
func (o *GetAPIV2OperationsIDPotentialLinksOK) Code() int {
	return 200
}

func (o *GetAPIV2OperationsIDPotentialLinksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/operations/{id}/potential-links][%d] getApiV2OperationsIdPotentialLinksOK %s", 200, payload)
}

func (o *GetAPIV2OperationsIDPotentialLinksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/operations/{id}/potential-links][%d] getApiV2OperationsIdPotentialLinksOK %s", 200, payload)
}

func (o *GetAPIV2OperationsIDPotentialLinksOK) GetPayload() []*models.PartialLink {
	return o.Payload
}

func (o *GetAPIV2OperationsIDPotentialLinksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
