package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00700102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.007.001.02 Document"`
	Message *AmendmentRejectionV02 `xml:"AmdmntRjctn"`
}

func (d *Document00700102) AddMessage() *AmendmentRejectionV02 {
	d.Message = new(AmendmentRejectionV02)
	return d.Message
}

// Scope
// The AmendmentRejection message is sent by the party requested to accept or reject an amendment to the matching application.
// This message is used to reject an amendment request.
// Usage
// The AmendmentRejection message can be sent by the party requested to accept or reject an amendment to inform that it rejects the requested amendment.
// The message can be sent in response to a FullPushThroughReport and DeltaReport message conveying the details of a BaselineAmendmentRequest message.
// The acceptance of an amendment request can be achieved by sending an AmendmentAcceptance message.
type AmendmentRejectionV02 struct {

	// Identifies the rejection message.
	RejectionIdentification *iso20022.MessageIdentification1 `xml:"RjctnId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Reference to the identification of the delta report that contained the amendment.
	DeltaReportReference *iso20022.MessageIdentification1 `xml:"DltaRptRef"`

	// Sequence number of the rejected baseline amendment.
	RejectedAmendmentNumber *iso20022.Count1 `xml:"RjctdAmdmntNb"`

	// Specifies the reaons for rejecting the amendment.                                                    
	RejectionReason *iso20022.RejectionReason1Choice `xml:"RjctnRsn"`

}


func (a *AmendmentRejectionV02) AddRejectionIdentification() *iso20022.MessageIdentification1 {
	a.RejectionIdentification = new(iso20022.MessageIdentification1)
	return a.RejectionIdentification
}

func (a *AmendmentRejectionV02) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	a.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return a.TransactionIdentification
}

func (a *AmendmentRejectionV02) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	a.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return a.SubmitterTransactionReference
}

func (a *AmendmentRejectionV02) AddDeltaReportReference() *iso20022.MessageIdentification1 {
	a.DeltaReportReference = new(iso20022.MessageIdentification1)
	return a.DeltaReportReference
}

func (a *AmendmentRejectionV02) AddRejectedAmendmentNumber() *iso20022.Count1 {
	a.RejectedAmendmentNumber = new(iso20022.Count1)
	return a.RejectedAmendmentNumber
}

func (a *AmendmentRejectionV02) AddRejectionReason() *iso20022.RejectionReason1Choice {
	a.RejectionReason = new(iso20022.RejectionReason1Choice)
	return a.RejectionReason
}

