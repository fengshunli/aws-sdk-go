// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribestreamingservice

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// One or more arguments to the StartStreamTranscription or StartMedicalStreamTranscription
	// operation was not valid. For example, MediaEncoding or LanguageCode used
	// not valid values. Check the specified parameters and try your request again.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// A new stream started with the same session ID. The current stream has been
	// terminated.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	//
	// A problem occurred while processing the audio. Amazon Transcribe terminated
	// processing.
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Your client has exceeded one of the Amazon Transcribe limits. This is typically
	// the audio length limit. Break your audio stream into smaller chunks and try
	// your request again.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The service is currently unavailable. Try your request later.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":         newErrorBadRequestException,
	"ConflictException":           newErrorConflictException,
	"InternalFailureException":    newErrorInternalFailureException,
	"LimitExceededException":      newErrorLimitExceededException,
	"ServiceUnavailableException": newErrorServiceUnavailableException,
}
