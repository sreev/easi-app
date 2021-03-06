// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cmsgov/easi-app/pkg/cedar/cedareasi/gen/models"
)

// SystemsGET2Reader is a Reader for the SystemsGET2 structure.
type SystemsGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SystemsGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSystemsGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSystemsGET2BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSystemsGET2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSystemsGET2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSystemsGET2OK creates a SystemsGET2OK with default headers values
func NewSystemsGET2OK() *SystemsGET2OK {
	return &SystemsGET2OK{}
}

/*SystemsGET2OK handles this case with default header values.

OK
*/
type SystemsGET2OK struct {
	Payload *models.SystemsGETResponse
}

func (o *SystemsGET2OK) Error() string {
	return fmt.Sprintf("[GET /systems][%d] systemsGET2OK  %+v", 200, o.Payload)
}

func (o *SystemsGET2OK) GetPayload() *models.SystemsGETResponse {
	return o.Payload
}

func (o *SystemsGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsGETResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemsGET2BadRequest creates a SystemsGET2BadRequest with default headers values
func NewSystemsGET2BadRequest() *SystemsGET2BadRequest {
	return &SystemsGET2BadRequest{}
}

/*SystemsGET2BadRequest handles this case with default header values.

Bad Request
*/
type SystemsGET2BadRequest struct {
	Payload *models.SystemsGETResponse
}

func (o *SystemsGET2BadRequest) Error() string {
	return fmt.Sprintf("[GET /systems][%d] systemsGET2BadRequest  %+v", 400, o.Payload)
}

func (o *SystemsGET2BadRequest) GetPayload() *models.SystemsGETResponse {
	return o.Payload
}

func (o *SystemsGET2BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsGETResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemsGET2Unauthorized creates a SystemsGET2Unauthorized with default headers values
func NewSystemsGET2Unauthorized() *SystemsGET2Unauthorized {
	return &SystemsGET2Unauthorized{}
}

/*SystemsGET2Unauthorized handles this case with default header values.

Access Denied
*/
type SystemsGET2Unauthorized struct {
	Payload *models.SystemsGETResponse
}

func (o *SystemsGET2Unauthorized) Error() string {
	return fmt.Sprintf("[GET /systems][%d] systemsGET2Unauthorized  %+v", 401, o.Payload)
}

func (o *SystemsGET2Unauthorized) GetPayload() *models.SystemsGETResponse {
	return o.Payload
}

func (o *SystemsGET2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsGETResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSystemsGET2InternalServerError creates a SystemsGET2InternalServerError with default headers values
func NewSystemsGET2InternalServerError() *SystemsGET2InternalServerError {
	return &SystemsGET2InternalServerError{}
}

/*SystemsGET2InternalServerError handles this case with default header values.

Internal Server Error
*/
type SystemsGET2InternalServerError struct {
	Payload *models.SystemsGETResponse
}

func (o *SystemsGET2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /systems][%d] systemsGET2InternalServerError  %+v", 500, o.Payload)
}

func (o *SystemsGET2InternalServerError) GetPayload() *models.SystemsGETResponse {
	return o.Payload
}

func (o *SystemsGET2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemsGETResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
