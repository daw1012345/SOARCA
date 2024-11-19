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

// GetAPIV2OperationsIDLinksLinkIDResultReader is a Reader for the GetAPIV2OperationsIDLinksLinkIDResult structure.
type GetAPIV2OperationsIDLinksLinkIDResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2OperationsIDLinksLinkIDResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2OperationsIDLinksLinkIDResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/operations/{id}/links/{link_id}/result] GetAPIV2OperationsIDLinksLinkIDResult", response, response.Code())
	}
}

// NewGetAPIV2OperationsIDLinksLinkIDResultOK creates a GetAPIV2OperationsIDLinksLinkIDResultOK with default headers values
func NewGetAPIV2OperationsIDLinksLinkIDResultOK() *GetAPIV2OperationsIDLinksLinkIDResultOK {
	return &GetAPIV2OperationsIDLinksLinkIDResultOK{}
}

/*
GetAPIV2OperationsIDLinksLinkIDResultOK describes a response with status code 200, with default header values.

Contains a dictionary with the requested link and its results dictionary.
*/
type GetAPIV2OperationsIDLinksLinkIDResultOK struct {
	Payload *models.LinkResult
}

// IsSuccess returns true when this get Api v2 operations Id links link Id result o k response has a 2xx status code
func (o *GetAPIV2OperationsIDLinksLinkIDResultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 operations Id links link Id result o k response has a 3xx status code
func (o *GetAPIV2OperationsIDLinksLinkIDResultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 operations Id links link Id result o k response has a 4xx status code
func (o *GetAPIV2OperationsIDLinksLinkIDResultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 operations Id links link Id result o k response has a 5xx status code
func (o *GetAPIV2OperationsIDLinksLinkIDResultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 operations Id links link Id result o k response a status code equal to that given
func (o *GetAPIV2OperationsIDLinksLinkIDResultOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v2 operations Id links link Id result o k response
func (o *GetAPIV2OperationsIDLinksLinkIDResultOK) Code() int {
	return 200
}

func (o *GetAPIV2OperationsIDLinksLinkIDResultOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/operations/{id}/links/{link_id}/result][%d] getApiV2OperationsIdLinksLinkIdResultOK %s", 200, payload)
}

func (o *GetAPIV2OperationsIDLinksLinkIDResultOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/operations/{id}/links/{link_id}/result][%d] getApiV2OperationsIdLinksLinkIdResultOK %s", 200, payload)
}

func (o *GetAPIV2OperationsIDLinksLinkIDResultOK) GetPayload() *models.LinkResult {
	return o.Payload
}

func (o *GetAPIV2OperationsIDLinksLinkIDResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LinkResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
