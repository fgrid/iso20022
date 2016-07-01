package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00500102 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.005.001.02 Document"`
	Message *AmendmentAcceptanceV02 `xml:"AmdmntAccptnc"`
}

func (d *Document00500102) AddMessage() *AmendmentAcceptanceV02 {
	d.Message = new(AmendmentAcceptanceV02)
	return d.Message
}

// Scope
// The AmendmentAcceptance message is sent by the party requested to accept or reject an amendment to the matching application.
// The message is used to accept an amendment request.
// Usage
// The AmendmentAcceptance message can be sent by the party requested to accept or reject an amendment to inform that it accepts the requested amendment.
// The message can be sent in response to a FullPushThroughReport and DeltaReport message conveying the details of a BaselineAmendmentRequest message.
// The rejection of an amendment request can be achieved by sending an AmendmentRejection message.
type AmendmentAcceptanceV02 struct {

	// Identifies the acceptance message.
	AcceptanceIdentification *iso20022.MessageIdentification1 `xml:"AccptncId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reference to the identification of the transaction for the requesting financial institution.
	SubmitterTransactionReference *iso20022.SimpleIdentificationInformation `xml:"SubmitrTxRef,omitempty"`

	// Reference to the identification of the delta report that contained the amendment.
	DeltaReportReference *iso20022.MessageIdentification1 `xml:"DltaRptRef"`

	// Sequence number of the accepted baseline amendment.
	AcceptedAmendmentNumber *iso20022.Count1 `xml:"AccptdAmdmntNb"`

}


func (a *AmendmentAcceptanceV02) AddAcceptanceIdentification() *iso20022.MessageIdentification1 {
	a.AcceptanceIdentification = new(iso20022.MessageIdentification1)
	return a.AcceptanceIdentification
}

func (a *AmendmentAcceptanceV02) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	a.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return a.TransactionIdentification
}

func (a *AmendmentAcceptanceV02) AddSubmitterTransactionReference() *iso20022.SimpleIdentificationInformation {
	a.SubmitterTransactionReference = new(iso20022.SimpleIdentificationInformation)
	return a.SubmitterTransactionReference
}

func (a *AmendmentAcceptanceV02) AddDeltaReportReference() *iso20022.MessageIdentification1 {
	a.DeltaReportReference = new(iso20022.MessageIdentification1)
	return a.DeltaReportReference
}

func (a *AmendmentAcceptanceV02) AddAcceptedAmendmentNumber() *iso20022.Count1 {
	a.AcceptedAmendmentNumber = new(iso20022.Count1)
	return a.AcceptedAmendmentNumber
}

