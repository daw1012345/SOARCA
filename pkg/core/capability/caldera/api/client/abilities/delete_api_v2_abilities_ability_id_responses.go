// Code generated by go-swagger; DO NOT EDIT.

package abilities

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

// DeleteAPIV2AbilitiesAbilityIDReader is a Reader for the DeleteAPIV2AbilitiesAbilityID structure.
type DeleteAPIV2AbilitiesAbilityIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV2AbilitiesAbilityIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAPIV2AbilitiesAbilityIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /api/v2/abilities/{ability_id}] DeleteAPIV2AbilitiesAbilityID", response, response.Code())
	}
}

// NewDeleteAPIV2AbilitiesAbilityIDNoContent creates a DeleteAPIV2AbilitiesAbilityIDNoContent with default headers values
func NewDeleteAPIV2AbilitiesAbilityIDNoContent() *DeleteAPIV2AbilitiesAbilityIDNoContent {
	return &DeleteAPIV2AbilitiesAbilityIDNoContent{}
}

/*
DeleteAPIV2AbilitiesAbilityIDNoContent describes a response with status code 204, with default header values.

HTTP 204 Status Code (No Content)
*/
type DeleteAPIV2AbilitiesAbilityIDNoContent struct {
	Payload *models.Ability
}

// IsSuccess returns true when this delete Api v2 abilities ability Id no content response has a 2xx status code
func (o *DeleteAPIV2AbilitiesAbilityIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete Api v2 abilities ability Id no content response has a 3xx status code
func (o *DeleteAPIV2AbilitiesAbilityIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete Api v2 abilities ability Id no content response has a 4xx status code
func (o *DeleteAPIV2AbilitiesAbilityIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete Api v2 abilities ability Id no content response has a 5xx status code
func (o *DeleteAPIV2AbilitiesAbilityIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete Api v2 abilities ability Id no content response a status code equal to that given
func (o *DeleteAPIV2AbilitiesAbilityIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete Api v2 abilities ability Id no content response
func (o *DeleteAPIV2AbilitiesAbilityIDNoContent) Code() int {
	return 204
}

func (o *DeleteAPIV2AbilitiesAbilityIDNoContent) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/abilities/{ability_id}][%d] deleteApiV2AbilitiesAbilityIdNoContent %s", 204, payload)
}

func (o *DeleteAPIV2AbilitiesAbilityIDNoContent) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /api/v2/abilities/{ability_id}][%d] deleteApiV2AbilitiesAbilityIdNoContent %s", 204, payload)
}

func (o *DeleteAPIV2AbilitiesAbilityIDNoContent) GetPayload() *models.Ability {
	return o.Payload
}

func (o *DeleteAPIV2AbilitiesAbilityIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ability)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
