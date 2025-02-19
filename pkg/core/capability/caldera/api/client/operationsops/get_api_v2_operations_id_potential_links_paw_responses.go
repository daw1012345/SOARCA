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

// GetAPIV2OperationsIDPotentialLinksPawReader is a Reader for the GetAPIV2OperationsIDPotentialLinksPaw structure.
type GetAPIV2OperationsIDPotentialLinksPawReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2OperationsIDPotentialLinksPawReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2OperationsIDPotentialLinksPawOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/operations/{id}/potential-links/{paw}] GetAPIV2OperationsIDPotentialLinksPaw", response, response.Code())
	}
}

// NewGetAPIV2OperationsIDPotentialLinksPawOK creates a GetAPIV2OperationsIDPotentialLinksPawOK with default headers values
func NewGetAPIV2OperationsIDPotentialLinksPawOK() *GetAPIV2OperationsIDPotentialLinksPawOK {
	return &GetAPIV2OperationsIDPotentialLinksPawOK{}
}

/*
GetAPIV2OperationsIDPotentialLinksPawOK describes a response with status code 200, with default header values.

All potential links for operation and the specified agent paw.
*/
type GetAPIV2OperationsIDPotentialLinksPawOK struct {
	Payload *models.PartialLink
}

// IsSuccess returns true when this get Api v2 operations Id potential links paw o k response has a 2xx status code
func (o *GetAPIV2OperationsIDPotentialLinksPawOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 operations Id potential links paw o k response has a 3xx status code
func (o *GetAPIV2OperationsIDPotentialLinksPawOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 operations Id potential links paw o k response has a 4xx status code
func (o *GetAPIV2OperationsIDPotentialLinksPawOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 operations Id potential links paw o k response has a 5xx status code
func (o *GetAPIV2OperationsIDPotentialLinksPawOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 operations Id potential links paw o k response a status code equal to that given
func (o *GetAPIV2OperationsIDPotentialLinksPawOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v2 operations Id potential links paw o k response
func (o *GetAPIV2OperationsIDPotentialLinksPawOK) Code() int {
	return 200
}

func (o *GetAPIV2OperationsIDPotentialLinksPawOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/operations/{id}/potential-links/{paw}][%d] getApiV2OperationsIdPotentialLinksPawOK %s", 200, payload)
}

func (o *GetAPIV2OperationsIDPotentialLinksPawOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/operations/{id}/potential-links/{paw}][%d] getApiV2OperationsIdPotentialLinksPawOK %s", 200, payload)
}

func (o *GetAPIV2OperationsIDPotentialLinksPawOK) GetPayload() *models.PartialLink {
	return o.Payload
}

func (o *GetAPIV2OperationsIDPotentialLinksPawOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartialLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
