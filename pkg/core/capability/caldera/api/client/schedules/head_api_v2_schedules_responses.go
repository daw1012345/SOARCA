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

// HeadAPIV2SchedulesReader is a Reader for the HeadAPIV2Schedules structure.
type HeadAPIV2SchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadAPIV2SchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadAPIV2SchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[HEAD /api/v2/schedules] HeadAPIV2Schedules", response, response.Code())
	}
}

// NewHeadAPIV2SchedulesOK creates a HeadAPIV2SchedulesOK with default headers values
func NewHeadAPIV2SchedulesOK() *HeadAPIV2SchedulesOK {
	return &HeadAPIV2SchedulesOK{}
}

/*
HeadAPIV2SchedulesOK describes a response with status code 200, with default header values.

The response is a list of all scheduled operations.
*/
type HeadAPIV2SchedulesOK struct {
	Payload []*models.PartialSchedule
}

// IsSuccess returns true when this head Api v2 schedules o k response has a 2xx status code
func (o *HeadAPIV2SchedulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this head Api v2 schedules o k response has a 3xx status code
func (o *HeadAPIV2SchedulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this head Api v2 schedules o k response has a 4xx status code
func (o *HeadAPIV2SchedulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this head Api v2 schedules o k response has a 5xx status code
func (o *HeadAPIV2SchedulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this head Api v2 schedules o k response a status code equal to that given
func (o *HeadAPIV2SchedulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the head Api v2 schedules o k response
func (o *HeadAPIV2SchedulesOK) Code() int {
	return 200
}

func (o *HeadAPIV2SchedulesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/schedules][%d] headApiV2SchedulesOK %s", 200, payload)
}

func (o *HeadAPIV2SchedulesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/schedules][%d] headApiV2SchedulesOK %s", 200, payload)
}

func (o *HeadAPIV2SchedulesOK) GetPayload() []*models.PartialSchedule {
	return o.Payload
}

func (o *HeadAPIV2SchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
