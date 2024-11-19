// Code generated by go-swagger; DO NOT EDIT.

package operationsops

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

// HeadAPIV2OperationsIDLinksLinkIDReader is a Reader for the HeadAPIV2OperationsIDLinksLinkID structure.
type HeadAPIV2OperationsIDLinksLinkIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadAPIV2OperationsIDLinksLinkIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadAPIV2OperationsIDLinksLinkIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[HEAD /api/v2/operations/{id}/links/{link_id}] HeadAPIV2OperationsIDLinksLinkID", response, response.Code())
	}
}

// NewHeadAPIV2OperationsIDLinksLinkIDOK creates a HeadAPIV2OperationsIDLinksLinkIDOK with default headers values
func NewHeadAPIV2OperationsIDLinksLinkIDOK() *HeadAPIV2OperationsIDLinksLinkIDOK {
	return &HeadAPIV2OperationsIDLinksLinkIDOK{}
}

/*
HeadAPIV2OperationsIDLinksLinkIDOK describes a response with status code 200, with default header values.

The link matching the provided `link_id` within the operation matching `id`. Use fields from the `BaseGetOneQuerySchema` in the request body to add `include` and `exclude` filters.
*/
type HeadAPIV2OperationsIDLinksLinkIDOK struct {
	Payload *models.PartialLink
}

// IsSuccess returns true when this head Api v2 operations Id links link Id o k response has a 2xx status code
func (o *HeadAPIV2OperationsIDLinksLinkIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this head Api v2 operations Id links link Id o k response has a 3xx status code
func (o *HeadAPIV2OperationsIDLinksLinkIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this head Api v2 operations Id links link Id o k response has a 4xx status code
func (o *HeadAPIV2OperationsIDLinksLinkIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this head Api v2 operations Id links link Id o k response has a 5xx status code
func (o *HeadAPIV2OperationsIDLinksLinkIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this head Api v2 operations Id links link Id o k response a status code equal to that given
func (o *HeadAPIV2OperationsIDLinksLinkIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the head Api v2 operations Id links link Id o k response
func (o *HeadAPIV2OperationsIDLinksLinkIDOK) Code() int {
	return 200
}

func (o *HeadAPIV2OperationsIDLinksLinkIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/operations/{id}/links/{link_id}][%d] headApiV2OperationsIdLinksLinkIdOK %s", 200, payload)
}

func (o *HeadAPIV2OperationsIDLinksLinkIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/operations/{id}/links/{link_id}][%d] headApiV2OperationsIdLinksLinkIdOK %s", 200, payload)
}

func (o *HeadAPIV2OperationsIDLinksLinkIDOK) GetPayload() *models.PartialLink {
	return o.Payload
}

func (o *HeadAPIV2OperationsIDLinksLinkIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartialLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
