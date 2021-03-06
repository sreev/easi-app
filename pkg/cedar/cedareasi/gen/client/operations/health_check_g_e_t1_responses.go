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

// HealthCheckGET1Reader is a Reader for the HealthCheckGET1 structure.
type HealthCheckGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthCheckGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHealthCheckGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHealthCheckGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewHealthCheckGET1OK creates a HealthCheckGET1OK with default headers values
func NewHealthCheckGET1OK() *HealthCheckGET1OK {
	return &HealthCheckGET1OK{}
}

/*HealthCheckGET1OK handles this case with default header values.

OK
*/
type HealthCheckGET1OK struct {
	Payload *models.HealthCheckGETResponse
}

func (o *HealthCheckGET1OK) Error() string {
	return fmt.Sprintf("[GET /healthCheck][%d] healthCheckGET1OK  %+v", 200, o.Payload)
}

func (o *HealthCheckGET1OK) GetPayload() *models.HealthCheckGETResponse {
	return o.Payload
}

func (o *HealthCheckGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HealthCheckGETResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHealthCheckGET1Unauthorized creates a HealthCheckGET1Unauthorized with default headers values
func NewHealthCheckGET1Unauthorized() *HealthCheckGET1Unauthorized {
	return &HealthCheckGET1Unauthorized{}
}

/*HealthCheckGET1Unauthorized handles this case with default header values.

Access Denied
*/
type HealthCheckGET1Unauthorized struct {
}

func (o *HealthCheckGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /healthCheck][%d] healthCheckGET1Unauthorized ", 401)
}

func (o *HealthCheckGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
