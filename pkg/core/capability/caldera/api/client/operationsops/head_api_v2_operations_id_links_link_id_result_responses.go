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

	"soarca/pkg/core/capability/caldera/api/models"
)

// HeadAPIV2OperationsIDLinksLinkIDResultReader is a Reader for the HeadAPIV2OperationsIDLinksLinkIDResult structure.
type HeadAPIV2OperationsIDLinksLinkIDResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HeadAPIV2OperationsIDLinksLinkIDResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHeadAPIV2OperationsIDLinksLinkIDResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[HEAD /api/v2/operations/{id}/links/{link_id}/result] HeadAPIV2OperationsIDLinksLinkIDResult", response, response.Code())
	}
}

// NewHeadAPIV2OperationsIDLinksLinkIDResultOK creates a HeadAPIV2OperationsIDLinksLinkIDResultOK with default headers values
func NewHeadAPIV2OperationsIDLinksLinkIDResultOK() *HeadAPIV2OperationsIDLinksLinkIDResultOK {
	return &HeadAPIV2OperationsIDLinksLinkIDResultOK{}
}

/*
HeadAPIV2OperationsIDLinksLinkIDResultOK describes a response with status code 200, with default header values.

Contains a dictionary with the requested link and its results dictionary.
*/
type HeadAPIV2OperationsIDLinksLinkIDResultOK struct {
	Payload *models.LinkResult
}

// IsSuccess returns true when this head Api v2 operations Id links link Id result o k response has a 2xx status code
func (o *HeadAPIV2OperationsIDLinksLinkIDResultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this head Api v2 operations Id links link Id result o k response has a 3xx status code
func (o *HeadAPIV2OperationsIDLinksLinkIDResultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this head Api v2 operations Id links link Id result o k response has a 4xx status code
func (o *HeadAPIV2OperationsIDLinksLinkIDResultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this head Api v2 operations Id links link Id result o k response has a 5xx status code
func (o *HeadAPIV2OperationsIDLinksLinkIDResultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this head Api v2 operations Id links link Id result o k response a status code equal to that given
func (o *HeadAPIV2OperationsIDLinksLinkIDResultOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the head Api v2 operations Id links link Id result o k response
func (o *HeadAPIV2OperationsIDLinksLinkIDResultOK) Code() int {
	return 200
}

func (o *HeadAPIV2OperationsIDLinksLinkIDResultOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/operations/{id}/links/{link_id}/result][%d] headApiV2OperationsIdLinksLinkIdResultOK %s", 200, payload)
}

func (o *HeadAPIV2OperationsIDLinksLinkIDResultOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[HEAD /api/v2/operations/{id}/links/{link_id}/result][%d] headApiV2OperationsIdLinksLinkIdResultOK %s", 200, payload)
}

func (o *HeadAPIV2OperationsIDLinksLinkIDResultOK) GetPayload() *models.LinkResult {
	return o.Payload
}

func (o *HeadAPIV2OperationsIDLinksLinkIDResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LinkResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
