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

	"soarca/pkg/core/capability/caldera/api/models"
)

// HeadAPIV2SchedulesIDReader is a Reader for the HeadAPIV2SchedulesID structure.
type HeadAPIV2SchedulesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadAPIV2SchedulesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadAPIV2SchedulesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[HEAD /api/v2/schedules/{id}] HeadAPIV2SchedulesID", response, response.Code())
	}
}

// NewHeadAPIV2SchedulesIDOK creates a HeadAPIV2SchedulesIDOK with default headers values
func NewHeadAPIV2SchedulesIDOK() *HeadAPIV2SchedulesIDOK {
	return &HeadAPIV2SchedulesIDOK{}
}

/*
HeadAPIV2SchedulesIDOK describes a response with status code 200, with default header values.

The response is a single dumped Scheduled object.
*/
type HeadAPIV2SchedulesIDOK struct {
	Payload *models.PartialSchedule
}

// IsSuccess returns true when this head Api v2 schedules Id o k response has a 2xx status code
func (o *HeadAPIV2SchedulesIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this head Api v2 schedules Id o k response has a 3xx status code
func (o *HeadAPIV2SchedulesIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this head Api v2 schedules Id o k response has a 4xx status code
func (o *HeadAPIV2SchedulesIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this head Api v2 schedules Id o k response has a 5xx status code
func (o *HeadAPIV2SchedulesIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this head Api v2 schedules Id o k response a status code equal to that given
func (o *HeadAPIV2SchedulesIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the head Api v2 schedules Id o k response
func (o *HeadAPIV2SchedulesIDOK) Code() int {
	return 200
}

func (o *HeadAPIV2SchedulesIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/schedules/{id}][%d] headApiV2SchedulesIdOK %s", 200, payload)
}

func (o *HeadAPIV2SchedulesIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/schedules/{id}][%d] headApiV2SchedulesIdOK %s", 200, payload)
}

func (o *HeadAPIV2SchedulesIDOK) GetPayload() *models.PartialSchedule {
	return o.Payload
}

func (o *HeadAPIV2SchedulesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartialSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
