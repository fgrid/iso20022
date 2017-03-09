package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01600103 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.016.001.03 Document"`
	Message *ErrorReportV03 `xml:"ErrRpt"`
}

func (d *Document01600103) AddMessage() *ErrorReportV03 {
	d.Message = new(ErrorReportV03)
	return d.Message
}

// Scope
// The ErrorReport message is sent by the matching application to the party from which it received a message.
// This message is used to inform about the inability of the matching application to process a received message.
// Usage
// The ErrorReport message can be sent to a party from which the matching application received a message to inform about its inability to process the received message because
// - the syntax of the message is incorrect,or
// - the message content is inconsistent,or
// - according to the workflow implemented in the matching application, it did not expect the received message.
type ErrorReportV03 struct {

	// Identifies the report.
	ReportIdentification *iso20022.MessageIdentification1 `xml:"RptId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId,omitempty"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *iso20022.TransactionStatus4 `xml:"TxSts,omitempty"`

	// Reference to the transaction for the financial institution which is the sender of the rejected message.
	UserTransactionReference *iso20022.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Reference to the identification of the rejected message.
	RejectedMessageReference *iso20022.MessageIdentification1 `xml:"RjctdMsgRef,omitempty"`

	// Specifies the total number of errors identified in the rejected message.
	NumberOfErrors *iso20022.Count1 `xml:"NbOfErrs"`

	// Describes the error that is the cause of the rejection.
	ErrorDescription []*iso20022.ValidationResult3 `xml:"ErrDesc"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (e *ErrorReportV03) AddReportIdentification() *iso20022.MessageIdentification1 {
	e.ReportIdentification = new(iso20022.MessageIdentification1)
	return e.ReportIdentification
}

func (e *ErrorReportV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	e.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return e.TransactionIdentification
}

func (e *ErrorReportV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	e.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return e.EstablishedBaselineIdentification
}

func (e *ErrorReportV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	e.TransactionStatus = new(iso20022.TransactionStatus4)
	return e.TransactionStatus
}

func (e *ErrorReportV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	e.UserTransactionReference = new(iso20022.DocumentIdentification5)
	return e.UserTransactionReference
}

func (e *ErrorReportV03) AddRejectedMessageReference() *iso20022.MessageIdentification1 {
	e.RejectedMessageReference = new(iso20022.MessageIdentification1)
	return e.RejectedMessageReference
}

func (e *ErrorReportV03) AddNumberOfErrors() *iso20022.Count1 {
	e.NumberOfErrors = new(iso20022.Count1)
	return e.NumberOfErrors
}

func (e *ErrorReportV03) AddErrorDescription() *iso20022.ValidationResult3 {
	newValue := new(iso20022.ValidationResult3)
	e.ErrorDescription = append(e.ErrorDescription, newValue)
	return newValue
}

func (e *ErrorReportV03) AddRequestForAction() *iso20022.PendingActivity2 {
	e.RequestForAction = new(iso20022.PendingActivity2)
	return e.RequestForAction
}
