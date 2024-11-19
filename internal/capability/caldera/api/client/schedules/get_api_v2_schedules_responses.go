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

// GetAPIV2SchedulesReader is a Reader for the GetAPIV2Schedules structure.
type GetAPIV2SchedulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2SchedulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2SchedulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/schedules] GetAPIV2Schedules", response, response.Code())
	}
}

// NewGetAPIV2SchedulesOK creates a GetAPIV2SchedulesOK with default headers values
func NewGetAPIV2SchedulesOK() *GetAPIV2SchedulesOK {
	return &GetAPIV2SchedulesOK{}
}

/*
GetAPIV2SchedulesOK describes a response with status code 200, with default header values.

The response is a list of all scheduled operations.
*/
type GetAPIV2SchedulesOK struct {
	Payload []*models.PartialSchedule
}

// IsSuccess returns true when this get Api v2 schedules o k response has a 2xx status code
func (o *GetAPIV2SchedulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 schedules o k response has a 3xx status code
func (o *GetAPIV2SchedulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 schedules o k response has a 4xx status code
func (o *GetAPIV2SchedulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 schedules o k response has a 5xx status code
func (o *GetAPIV2SchedulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 schedules o k response a status code equal to that given
func (o *GetAPIV2SchedulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v2 schedules o k response
func (o *GetAPIV2SchedulesOK) Code() int {
	return 200
}

func (o *GetAPIV2SchedulesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/schedules][%d] getApiV2SchedulesOK %s", 200, payload)
}

func (o *GetAPIV2SchedulesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/schedules][%d] getApiV2SchedulesOK %s", 200, payload)
}

func (o *GetAPIV2SchedulesOK) GetPayload() []*models.PartialSchedule {
	return o.Payload
}

func (o *GetAPIV2SchedulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
