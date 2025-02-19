// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// GetAPIV2AgentsPawReader is a Reader for the GetAPIV2AgentsPaw structure.
type GetAPIV2AgentsPawReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2AgentsPawReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2AgentsPawOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/agents/{paw}] GetAPIV2AgentsPaw", response, response.Code())
	}
}

// NewGetAPIV2AgentsPawOK creates a GetAPIV2AgentsPawOK with default headers values
func NewGetAPIV2AgentsPawOK() *GetAPIV2AgentsPawOK {
	return &GetAPIV2AgentsPawOK{}
}

/*
GetAPIV2AgentsPawOK describes a response with status code 200, with default header values.

Returns JSON response with specified Agent
*/
type GetAPIV2AgentsPawOK struct {
	Payload *models.PartialAgent
}

// IsSuccess returns true when this get Api v2 agents paw o k response has a 2xx status code
func (o *GetAPIV2AgentsPawOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 agents paw o k response has a 3xx status code
func (o *GetAPIV2AgentsPawOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 agents paw o k response has a 4xx status code
func (o *GetAPIV2AgentsPawOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 agents paw o k response has a 5xx status code
func (o *GetAPIV2AgentsPawOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 agents paw o k response a status code equal to that given
func (o *GetAPIV2AgentsPawOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v2 agents paw o k response
func (o *GetAPIV2AgentsPawOK) Code() int {
	return 200
}

func (o *GetAPIV2AgentsPawOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/agents/{paw}][%d] getApiV2AgentsPawOK %s", 200, payload)
}

func (o *GetAPIV2AgentsPawOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/agents/{paw}][%d] getApiV2AgentsPawOK %s", 200, payload)
}

func (o *GetAPIV2AgentsPawOK) GetPayload() *models.PartialAgent {
	return o.Payload
}

func (o *GetAPIV2AgentsPawOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartialAgent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
