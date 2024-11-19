// Code generated by go-swagger; DO NOT EDIT.

package objectives

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

// PatchAPIV2ObjectivesIDReader is a Reader for the PatchAPIV2ObjectivesID structure.
type PatchAPIV2ObjectivesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIV2ObjectivesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPIV2ObjectivesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /api/v2/objectives/{id}] PatchAPIV2ObjectivesID", response, response.Code())
	}
}

// NewPatchAPIV2ObjectivesIDOK creates a PatchAPIV2ObjectivesIDOK with default headers values
func NewPatchAPIV2ObjectivesIDOK() *PatchAPIV2ObjectivesIDOK {
	return &PatchAPIV2ObjectivesIDOK{}
}

/*
PatchAPIV2ObjectivesIDOK describes a response with status code 200, with default header values.

The updated Objective in ObjectiveSchema format.
*/
type PatchAPIV2ObjectivesIDOK struct {
	Payload *models.Objective
}

// IsSuccess returns true when this patch Api v2 objectives Id o k response has a 2xx status code
func (o *PatchAPIV2ObjectivesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api v2 objectives Id o k response has a 3xx status code
func (o *PatchAPIV2ObjectivesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api v2 objectives Id o k response has a 4xx status code
func (o *PatchAPIV2ObjectivesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api v2 objectives Id o k response has a 5xx status code
func (o *PatchAPIV2ObjectivesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api v2 objectives Id o k response a status code equal to that given
func (o *PatchAPIV2ObjectivesIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch Api v2 objectives Id o k response
func (o *PatchAPIV2ObjectivesIDOK) Code() int {
	return 200
}

func (o *PatchAPIV2ObjectivesIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/objectives/{id}][%d] patchApiV2ObjectivesIdOK %s", 200, payload)
}

func (o *PatchAPIV2ObjectivesIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/objectives/{id}][%d] patchApiV2ObjectivesIdOK %s", 200, payload)
}

func (o *PatchAPIV2ObjectivesIDOK) GetPayload() *models.Objective {
	return o.Payload
}

func (o *PatchAPIV2ObjectivesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Objective)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
