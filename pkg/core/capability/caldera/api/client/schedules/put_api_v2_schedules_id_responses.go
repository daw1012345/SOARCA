// Code generated by go-swagger; DO NOT EDIT.

package schedules

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

// PutAPIV2SchedulesIDReader is a Reader for the PutAPIV2SchedulesID structure.
type PutAPIV2SchedulesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV2SchedulesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAPIV2SchedulesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /api/v2/schedules/{id}] PutAPIV2SchedulesID", response, response.Code())
	}
}

// NewPutAPIV2SchedulesIDOK creates a PutAPIV2SchedulesIDOK with default headers values
func NewPutAPIV2SchedulesIDOK() *PutAPIV2SchedulesIDOK {
	return &PutAPIV2SchedulesIDOK{}
}

/*
PutAPIV2SchedulesIDOK describes a response with status code 200, with default header values.

The response is a dump of the newly replaced Schedule object.
*/
type PutAPIV2SchedulesIDOK struct {
	Payload *models.Schedule
}

// IsSuccess returns true when this put Api v2 schedules Id o k response has a 2xx status code
func (o *PutAPIV2SchedulesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put Api v2 schedules Id o k response has a 3xx status code
func (o *PutAPIV2SchedulesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put Api v2 schedules Id o k response has a 4xx status code
func (o *PutAPIV2SchedulesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put Api v2 schedules Id o k response has a 5xx status code
func (o *PutAPIV2SchedulesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put Api v2 schedules Id o k response a status code equal to that given
func (o *PutAPIV2SchedulesIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the put Api v2 schedules Id o k response
func (o *PutAPIV2SchedulesIDOK) Code() int {
	return 200
}

func (o *PutAPIV2SchedulesIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/schedules/{id}][%d] putApiV2SchedulesIdOK %s", 200, payload)
}

func (o *PutAPIV2SchedulesIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /api/v2/schedules/{id}][%d] putApiV2SchedulesIdOK %s", 200, payload)
}

func (o *PutAPIV2SchedulesIDOK) GetPayload() *models.Schedule {
	return o.Payload
}

func (o *PutAPIV2SchedulesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Schedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
