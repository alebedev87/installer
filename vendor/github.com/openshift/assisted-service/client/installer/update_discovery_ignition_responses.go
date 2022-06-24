// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// UpdateDiscoveryIgnitionReader is a Reader for the UpdateDiscoveryIgnition structure.
type UpdateDiscoveryIgnitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDiscoveryIgnitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateDiscoveryIgnitionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateDiscoveryIgnitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateDiscoveryIgnitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateDiscoveryIgnitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateDiscoveryIgnitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdateDiscoveryIgnitionMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateDiscoveryIgnitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateDiscoveryIgnitionCreated creates a UpdateDiscoveryIgnitionCreated with default headers values
func NewUpdateDiscoveryIgnitionCreated() *UpdateDiscoveryIgnitionCreated {
	return &UpdateDiscoveryIgnitionCreated{}
}

/* UpdateDiscoveryIgnitionCreated describes a response with status code 201, with default header values.

Success.
*/
type UpdateDiscoveryIgnitionCreated struct {
}

func (o *UpdateDiscoveryIgnitionCreated) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}/discovery-ignition][%d] updateDiscoveryIgnitionCreated ", 201)
}

func (o *UpdateDiscoveryIgnitionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateDiscoveryIgnitionBadRequest creates a UpdateDiscoveryIgnitionBadRequest with default headers values
func NewUpdateDiscoveryIgnitionBadRequest() *UpdateDiscoveryIgnitionBadRequest {
	return &UpdateDiscoveryIgnitionBadRequest{}
}

/* UpdateDiscoveryIgnitionBadRequest describes a response with status code 400, with default header values.

Error.
*/
type UpdateDiscoveryIgnitionBadRequest struct {
	Payload *models.Error
}

func (o *UpdateDiscoveryIgnitionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}/discovery-ignition][%d] updateDiscoveryIgnitionBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateDiscoveryIgnitionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDiscoveryIgnitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDiscoveryIgnitionUnauthorized creates a UpdateDiscoveryIgnitionUnauthorized with default headers values
func NewUpdateDiscoveryIgnitionUnauthorized() *UpdateDiscoveryIgnitionUnauthorized {
	return &UpdateDiscoveryIgnitionUnauthorized{}
}

/* UpdateDiscoveryIgnitionUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type UpdateDiscoveryIgnitionUnauthorized struct {
	Payload *models.InfraError
}

func (o *UpdateDiscoveryIgnitionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}/discovery-ignition][%d] updateDiscoveryIgnitionUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateDiscoveryIgnitionUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateDiscoveryIgnitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDiscoveryIgnitionForbidden creates a UpdateDiscoveryIgnitionForbidden with default headers values
func NewUpdateDiscoveryIgnitionForbidden() *UpdateDiscoveryIgnitionForbidden {
	return &UpdateDiscoveryIgnitionForbidden{}
}

/* UpdateDiscoveryIgnitionForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type UpdateDiscoveryIgnitionForbidden struct {
	Payload *models.InfraError
}

func (o *UpdateDiscoveryIgnitionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}/discovery-ignition][%d] updateDiscoveryIgnitionForbidden  %+v", 403, o.Payload)
}
func (o *UpdateDiscoveryIgnitionForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateDiscoveryIgnitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDiscoveryIgnitionNotFound creates a UpdateDiscoveryIgnitionNotFound with default headers values
func NewUpdateDiscoveryIgnitionNotFound() *UpdateDiscoveryIgnitionNotFound {
	return &UpdateDiscoveryIgnitionNotFound{}
}

/* UpdateDiscoveryIgnitionNotFound describes a response with status code 404, with default header values.

Error.
*/
type UpdateDiscoveryIgnitionNotFound struct {
	Payload *models.Error
}

func (o *UpdateDiscoveryIgnitionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}/discovery-ignition][%d] updateDiscoveryIgnitionNotFound  %+v", 404, o.Payload)
}
func (o *UpdateDiscoveryIgnitionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDiscoveryIgnitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDiscoveryIgnitionMethodNotAllowed creates a UpdateDiscoveryIgnitionMethodNotAllowed with default headers values
func NewUpdateDiscoveryIgnitionMethodNotAllowed() *UpdateDiscoveryIgnitionMethodNotAllowed {
	return &UpdateDiscoveryIgnitionMethodNotAllowed{}
}

/* UpdateDiscoveryIgnitionMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type UpdateDiscoveryIgnitionMethodNotAllowed struct {
	Payload *models.Error
}

func (o *UpdateDiscoveryIgnitionMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}/discovery-ignition][%d] updateDiscoveryIgnitionMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *UpdateDiscoveryIgnitionMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDiscoveryIgnitionMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDiscoveryIgnitionInternalServerError creates a UpdateDiscoveryIgnitionInternalServerError with default headers values
func NewUpdateDiscoveryIgnitionInternalServerError() *UpdateDiscoveryIgnitionInternalServerError {
	return &UpdateDiscoveryIgnitionInternalServerError{}
}

/* UpdateDiscoveryIgnitionInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type UpdateDiscoveryIgnitionInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateDiscoveryIgnitionInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}/discovery-ignition][%d] updateDiscoveryIgnitionInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateDiscoveryIgnitionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateDiscoveryIgnitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
