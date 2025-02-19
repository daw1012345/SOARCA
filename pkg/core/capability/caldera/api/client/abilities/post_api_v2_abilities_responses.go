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

// PostAPIV2AbilitiesReader is a Reader for the PostAPIV2Abilities structure.
type PostAPIV2AbilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV2AbilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAPIV2AbilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /api/v2/abilities] PostAPIV2Abilities", response, response.Code())
	}
}

// NewPostAPIV2AbilitiesOK creates a PostAPIV2AbilitiesOK with default headers values
func NewPostAPIV2AbilitiesOK() *PostAPIV2AbilitiesOK {
	return &PostAPIV2AbilitiesOK{}
}

/*
PostAPIV2AbilitiesOK describes a response with status code 200, with default header values.

JSON dictionary representation of the created Ability.
*/
type PostAPIV2AbilitiesOK struct {
	Payload *models.Ability
}

// IsSuccess returns true when this post Api v2 abilities o k response has a 2xx status code
func (o *PostAPIV2AbilitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post Api v2 abilities o k response has a 3xx status code
func (o *PostAPIV2AbilitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post Api v2 abilities o k response has a 4xx status code
func (o *PostAPIV2AbilitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post Api v2 abilities o k response has a 5xx status code
func (o *PostAPIV2AbilitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post Api v2 abilities o k response a status code equal to that given
func (o *PostAPIV2AbilitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post Api v2 abilities o k response
func (o *PostAPIV2AbilitiesOK) Code() int {
	return 200
}

func (o *PostAPIV2AbilitiesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/abilities][%d] postApiV2AbilitiesOK %s", 200, payload)
}

func (o *PostAPIV2AbilitiesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /api/v2/abilities][%d] postApiV2AbilitiesOK %s", 200, payload)
}

func (o *PostAPIV2AbilitiesOK) GetPayload() *models.Ability {
	return o.Payload
}

func (o *PostAPIV2AbilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Ability)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
