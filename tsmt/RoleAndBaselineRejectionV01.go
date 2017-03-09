package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document05000101 struct {
	XMLName xml.Name                     `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.050.001.01 Document"`
	Message *RoleAndBaselineRejectionV01 `xml:"RoleAndBaselnRjctn"`
}

func (d *Document05000101) AddMessage() *RoleAndBaselineRejectionV01 {
	d.Message = new(RoleAndBaselineRejectionV01)
	return d.Message
}

// Scope
// The RoleAndBaselineRejection message is sent by a secondary bank to the matching application if it rejects to join the transaction based on the baseline and the role that it is expected to play.
// Usage
// The RoleAndBaselineRejection message is sent in response to a message that is a direct request to join a transaction.
type RoleAndBaselineRejectionV01 struct {

	// Identifies the rejection message.
	RejectionIdentification *iso20022.MessageIdentification1 `xml:"RjctnId"`

	// Reference to the message that contained the baseline and is rejected.
	RelatedMessageReference *iso20022.MessageIdentification1 `xml:"RltdMsgRef"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	//
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Reason why the user cannot accept the request.
	RejectionReason *iso20022.Reason2 `xml:"RjctnRsn,omitempty"`
}

func (r *RoleAndBaselineRejectionV01) AddRejectionIdentification() *iso20022.MessageIdentification1 {
	r.RejectionIdentification = new(iso20022.MessageIdentification1)
	return r.RejectionIdentification
}

func (r *RoleAndBaselineRejectionV01) AddRelatedMessageReference() *iso20022.MessageIdentification1 {
	r.RelatedMessageReference = new(iso20022.MessageIdentification1)
	return r.RelatedMessageReference
}

func (r *RoleAndBaselineRejectionV01) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	r.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return r.TransactionIdentification
}

func (r *RoleAndBaselineRejectionV01) AddRejectionReason() *iso20022.Reason2 {
	r.RejectionReason = new(iso20022.Reason2)
	return r.RejectionReason
}
