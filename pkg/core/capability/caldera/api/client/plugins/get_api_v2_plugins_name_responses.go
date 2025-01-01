// Code generated by go-swagger; DO NOT EDIT.

package plugins

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

// GetAPIV2PluginsNameReader is a Reader for the GetAPIV2PluginsName structure.
type GetAPIV2PluginsNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2PluginsNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2PluginsNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /api/v2/plugins/{name}] GetAPIV2PluginsName", response, response.Code())
	}
}

// NewGetAPIV2PluginsNameOK creates a GetAPIV2PluginsNameOK with default headers values
func NewGetAPIV2PluginsNameOK() *GetAPIV2PluginsNameOK {
	return &GetAPIV2PluginsNameOK{}
}

/*
GetAPIV2PluginsNameOK describes a response with status code 200, with default header values.

Returns a plugin in `PluginSchema` format with the requested name, if it exists.
*/
type GetAPIV2PluginsNameOK struct {
	Payload *models.PartialPlugin
}

// IsSuccess returns true when this get Api v2 plugins name o k response has a 2xx status code
func (o *GetAPIV2PluginsNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Api v2 plugins name o k response has a 3xx status code
func (o *GetAPIV2PluginsNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Api v2 plugins name o k response has a 4xx status code
func (o *GetAPIV2PluginsNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Api v2 plugins name o k response has a 5xx status code
func (o *GetAPIV2PluginsNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Api v2 plugins name o k response a status code equal to that given
func (o *GetAPIV2PluginsNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Api v2 plugins name o k response
func (o *GetAPIV2PluginsNameOK) Code() int {
	return 200
}

func (o *GetAPIV2PluginsNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/plugins/{name}][%d] getApiV2PluginsNameOK %s", 200, payload)
}

func (o *GetAPIV2PluginsNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v2/plugins/{name}][%d] getApiV2PluginsNameOK %s", 200, payload)
}

func (o *GetAPIV2PluginsNameOK) GetPayload() *models.PartialPlugin {
	return o.Payload
}

func (o *GetAPIV2PluginsNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartialPlugin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
