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

	"soarca/pkg/core/capability/caldera/api/models"
)

// GetAPIV2RelationshipsOperationIDReader is a Reader for the GetAPIV2RelationshipsOperationID structure.
type GetAPIV2RelationshipsOperationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2RelationshipsOperationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2RelationshipsOperationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/relationships/{operation_id}] GetAPIV2RelationshipsOperationID", response, response.Code())
	}
}

// NewGetAPIV2RelationshipsOperationIDOK creates a GetAPIV2RelationshipsOperationIDOK with default headers values
func NewGetAPIV2RelationshipsOperationIDOK() *GetAPIV2RelationshipsOperationIDOK {
	return &GetAPIV2RelationshipsOperationIDOK{}
}

/*
GetAPIV2RelationshipsOperationIDOK describes a response with status code 200, with default header values.

Returns a list of matching relationships, dumped in `RelationshipSchema` format.
*/
type GetAPIV2RelationshipsOperationIDOK struct {
	Payload *models.Relationship
}

// IsSuccess returns true when this get Api v2 relationships operation Id o k response has a 2xx status code
func (o *GetAPIV2RelationshipsOperationIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 relationships operation Id o k response has a 3xx status code
func (o *GetAPIV2RelationshipsOperationIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 relationships operation Id o k response has a 4xx status code
func (o *GetAPIV2RelationshipsOperationIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 relationships operation Id o k response has a 5xx status code
func (o *GetAPIV2RelationshipsOperationIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 relationships operation Id o k response a status code equal to that given
func (o *GetAPIV2RelationshipsOperationIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v2 relationships operation Id o k response
func (o *GetAPIV2RelationshipsOperationIDOK) Code() int {
	return 200
}

func (o *GetAPIV2RelationshipsOperationIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/relationships/{operation_id}][%d] getApiV2RelationshipsOperationIdOK %s", 200, payload)
}

func (o *GetAPIV2RelationshipsOperationIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/relationships/{operation_id}][%d] getApiV2RelationshipsOperationIdOK %s", 200, payload)
}

func (o *GetAPIV2RelationshipsOperationIDOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *GetAPIV2RelationshipsOperationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
