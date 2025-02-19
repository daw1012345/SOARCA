// Code generated by go-swagger; DO NOT EDIT.

package planners

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

// HeadAPIV2PlannersPlannerIDReader is a Reader for the HeadAPIV2PlannersPlannerID structure.
type HeadAPIV2PlannersPlannerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadAPIV2PlannersPlannerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadAPIV2PlannersPlannerIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[HEAD /api/v2/planners/{planner_id}] HeadAPIV2PlannersPlannerID", response, response.Code())
	}
}

// NewHeadAPIV2PlannersPlannerIDOK creates a HeadAPIV2PlannersPlannerIDOK with default headers values
func NewHeadAPIV2PlannersPlannerIDOK() *HeadAPIV2PlannersPlannerIDOK {
	return &HeadAPIV2PlannersPlannerIDOK{}
}

/*
HeadAPIV2PlannersPlannerIDOK describes a response with status code 200, with default header values.

Returns a planner with the specified id in `PlannerSchema` format.
*/
type HeadAPIV2PlannersPlannerIDOK struct {
	Payload *models.PartialPlanner
}

// IsSuccess returns true when this head Api v2 planners planner Id o k response has a 2xx status code
func (o *HeadAPIV2PlannersPlannerIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this head Api v2 planners planner Id o k response has a 3xx status code
func (o *HeadAPIV2PlannersPlannerIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this head Api v2 planners planner Id o k response has a 4xx status code
func (o *HeadAPIV2PlannersPlannerIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this head Api v2 planners planner Id o k response has a 5xx status code
func (o *HeadAPIV2PlannersPlannerIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this head Api v2 planners planner Id o k response a status code equal to that given
func (o *HeadAPIV2PlannersPlannerIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the head Api v2 planners planner Id o k response
func (o *HeadAPIV2PlannersPlannerIDOK) Code() int {
	return 200
}

func (o *HeadAPIV2PlannersPlannerIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/planners/{planner_id}][%d] headApiV2PlannersPlannerIdOK %s", 200, payload)
}

func (o *HeadAPIV2PlannersPlannerIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/planners/{planner_id}][%d] headApiV2PlannersPlannerIdOK %s", 200, payload)
}

func (o *HeadAPIV2PlannersPlannerIDOK) GetPayload() *models.PartialPlanner {
	return o.Payload
}

func (o *HeadAPIV2PlannersPlannerIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartialPlanner)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
