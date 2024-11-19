// Code generated by go-swagger; DO NOT EDIT.

package relationships

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

// PatchAPIV2RelationshipsReader is a Reader for the PatchAPIV2Relationships structure.
type PatchAPIV2RelationshipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPIV2RelationshipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAPIV2RelationshipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PATCH /api/v2/relationships] PatchAPIV2Relationships", response, response.Code())
	}
}

// NewPatchAPIV2RelationshipsOK creates a PatchAPIV2RelationshipsOK with default headers values
func NewPatchAPIV2RelationshipsOK() *PatchAPIV2RelationshipsOK {
	return &PatchAPIV2RelationshipsOK{}
}

/*
PatchAPIV2RelationshipsOK describes a response with status code 200, with default header values.

Returns the updated relationship(s), dumped in `RelationshipSchema` format.
*/
type PatchAPIV2RelationshipsOK struct {
	Payload *models.Relationship
}

// IsSuccess returns true when this patch Api v2 relationships o k response has a 2xx status code
func (o *PatchAPIV2RelationshipsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch Api v2 relationships o k response has a 3xx status code
func (o *PatchAPIV2RelationshipsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch Api v2 relationships o k response has a 4xx status code
func (o *PatchAPIV2RelationshipsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch Api v2 relationships o k response has a 5xx status code
func (o *PatchAPIV2RelationshipsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch Api v2 relationships o k response a status code equal to that given
func (o *PatchAPIV2RelationshipsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch Api v2 relationships o k response
func (o *PatchAPIV2RelationshipsOK) Code() int {
	return 200
}

func (o *PatchAPIV2RelationshipsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/relationships][%d] patchApiV2RelationshipsOK %s", 200, payload)
}

func (o *PatchAPIV2RelationshipsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /api/v2/relationships][%d] patchApiV2RelationshipsOK %s", 200, payload)
}

func (o *PatchAPIV2RelationshipsOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *PatchAPIV2RelationshipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
