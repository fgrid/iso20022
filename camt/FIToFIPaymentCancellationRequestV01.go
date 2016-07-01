package camt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document05600101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:camt.056.001.01 Document"`
	Message *FIToFIPaymentCancellationRequestV01 `xml:"FIToFIPmtCxlReq"`
}

func (d *Document05600101) AddMessage() *FIToFIPaymentCancellationRequestV01 {
	d.Message = new(FIToFIPaymentCancellationRequestV01)
	return d.Message
}

// Scope
// The FIToFI Payment Cancellation Request message is sent by a case creator/case assigner to a case assignee.
// This message is used to request the cancellation of an original payment instruction. The FIToFI Payment Cancellation Request message is exchanged between the instructing agent and the instructed agent to request the cancellation of a interbank payment message previously sent (such as FIToFICustomerCreditTransfer, FIToFICustomerDirectDebit or FinancialInstitutionCreditTransfer).
// Usage
// The FIToFI Payment Cancellation Request message must be answered with a:
// - Resolution Of Investigation message with a positive final outcome when the case assignee can perform the requested cancellation
// - Resolution Of Investigation message with a negative final outcome when the case assignee may perform the requested cancellation but fails to do so (too late, irrevocable instruction, ...)
// - Reject Investigation message when the case assignee is unable or not authorised to perform the requested cancellation
// - Notification Of Case Assignment message to indicate whether the case assignee will take on the case himself or reassign the case to a subsequent party in the payment processing chain.
// A FIToFI Payment Cancellation Request message concerns one and only one original payment instruction at a time.
// When a case assignee successfully performs a cancellation, it must return the corresponding funds to the case assigner. It may provide some details about the return in the Resolution Of Investigation message.
// The processing of a FIToFI Payment Cancellation Request message case may lead to a Debit Authorisation Request message sent to the creditor by its account servicing institution.
// The FIToFI Payment Cancellation Request message may be used to escalate a case after an unsuccessful request to modify the payment. In this scenario, the case identification remains the same as in the original FIToFI Payment Cancellation Request message and the element ReopenCaseIndication is set to 'Yes' or 'true'.
type FIToFIPaymentCancellationRequestV01 struct {

	// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The Assigner must be the sender of this confirmation and the Assignee must be the receiver.
	Assignment *iso20022.CaseAssignment2 `xml:"Assgnmt"`

	// Identifies the investigation case.
	Case *iso20022.Case2 `xml:"Case,omitempty"`

	// Provides details on the number of transactions and the control sum of the message.
	ControlData *iso20022.ControlData1 `xml:"CtrlData,omitempty"`

	// Identifies the payment instruction to be cancelled.
	Underlying []*iso20022.UnderlyingTransaction2 `xml:"Undrlyg"`

}


func (f *FIToFIPaymentCancellationRequestV01) AddAssignment() *iso20022.CaseAssignment2 {
	f.Assignment = new(iso20022.CaseAssignment2)
	return f.Assignment
}

func (f *FIToFIPaymentCancellationRequestV01) AddCase() *iso20022.Case2 {
	f.Case = new(iso20022.Case2)
	return f.Case
}

func (f *FIToFIPaymentCancellationRequestV01) AddControlData() *iso20022.ControlData1 {
	f.ControlData = new(iso20022.ControlData1)
	return f.ControlData
}

func (f *FIToFIPaymentCancellationRequestV01) AddUnderlying() *iso20022.UnderlyingTransaction2 {
	newValue := new (iso20022.UnderlyingTransaction2)
	f.Underlying = append(f.Underlying, newValue)
	return newValue
}

