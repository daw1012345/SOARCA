// Code generated by go-swagger; DO NOT EDIT.

package facts

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

// GetAPIV2FactsOperationIDReader is a Reader for the GetAPIV2FactsOperationID structure.
type GetAPIV2FactsOperationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2FactsOperationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2FactsOperationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/facts/{operation_id}] GetAPIV2FactsOperationID", response, response.Code())
	}
}

// NewGetAPIV2FactsOperationIDOK creates a GetAPIV2FactsOperationIDOK with default headers values
func NewGetAPIV2FactsOperationIDOK() *GetAPIV2FactsOperationIDOK {
	return &GetAPIV2FactsOperationIDOK{}
}

/*
GetAPIV2FactsOperationIDOK describes a response with status code 200, with default header values.

Returns a list of facts associated with operation_id, dumped in `FactSchema` format.
*/
type GetAPIV2FactsOperationIDOK struct {
	Payload []*models.PartialFact
}

// IsSuccess returns true when this get Api v2 facts operation Id o k response has a 2xx status code
func (o *GetAPIV2FactsOperationIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 facts operation Id o k response has a 3xx status code
func (o *GetAPIV2FactsOperationIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 facts operation Id o k response has a 4xx status code
func (o *GetAPIV2FactsOperationIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 facts operation Id o k response has a 5xx status code
func (o *GetAPIV2FactsOperationIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 facts operation Id o k response a status code equal to that given
func (o *GetAPIV2FactsOperationIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v2 facts operation Id o k response
func (o *GetAPIV2FactsOperationIDOK) Code() int {
	return 200
}

func (o *GetAPIV2FactsOperationIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/facts/{operation_id}][%d] getApiV2FactsOperationIdOK %s", 200, payload)
}

func (o *GetAPIV2FactsOperationIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/facts/{operation_id}][%d] getApiV2FactsOperationIdOK %s", 200, payload)
}

func (o *GetAPIV2FactsOperationIDOK) GetPayload() []*models.PartialFact {
	return o.Payload
}

func (o *GetAPIV2FactsOperationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
