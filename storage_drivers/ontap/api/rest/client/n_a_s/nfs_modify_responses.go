// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NfsModifyReader is a Reader for the NfsModify structure.
type NfsModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NfsModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNfsModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNfsModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNfsModifyOK creates a NfsModifyOK with default headers values
func NewNfsModifyOK() *NfsModifyOK {
	return &NfsModifyOK{}
}

/* NfsModifyOK describes a response with status code 200, with default header values.

OK
*/
type NfsModifyOK struct {
}

func (o *NfsModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/services/{svm.uuid}][%d] nfsModifyOK ", 200)
}

func (o *NfsModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNfsModifyDefault creates a NfsModifyDefault with default headers values
func NewNfsModifyDefault(code int) *NfsModifyDefault {
	return &NfsModifyDefault{
		_statusCode: code,
	}
}

/* NfsModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 3276916    | Vserver is not running |
| 3277069    | Cannot disable TCP because the SnapDiff RPC server is in the \\\"on\\\" state |
| 3277087    | Attempting to reduce the number of bits used for NFSv3 FSIDs and File IDs from 64 to 32 on Vserver. This could result in collisions between different File IDs and is not recommended |
| 3277088    | Attempting to increase the number of bits used for NFSv3 FSIDs and File IDs from 32 to 64 on Vserver. This could result in older client software no longer working with the volumes owned by Vserver  |
| 3277090    | Attempting to disallow multiple FSIDs per mount point on Vserver. Since this Vserver currently uses 32-bit NFSv3 FSIDs and File IDs, this could result in collisions between different File IDs and is not recommended |
| 3277099    | Domain name contains invalid characters or its too short. Allowed characters are: alphabetical characters (A-Za-z), numeric characters (0-9), minus sign (-), and the period (.). The first character must be alphabetical or numeric, last character must not be a minus sign or a period. Minimum supported length: 2 characters, maximum of 256 characters |

*/
type NfsModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nfs modify default response
func (o *NfsModifyDefault) Code() int {
	return o._statusCode
}

func (o *NfsModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/nfs/services/{svm.uuid}][%d] nfs_modify default  %+v", o._statusCode, o.Payload)
}
func (o *NfsModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NfsModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}