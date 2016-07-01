package semt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:semt.014.001.01 Document"`
	Message *IntraPositionMovementStatusAdviceV01 `xml:"IntraPosMvmntStsAdvc"`
}

func (d *Document01400101) AddMessage() *IntraPositionMovementStatusAdviceV01 {
	d.Message = new(IntraPositionMovementStatusAdviceV01)
	return d.Message
}

// Scope
// An account servicer sends a IntraPositionMovementStatusAdvice to an account owner to advise the status of a intra-position movement instruction previously sent by the account owner.
// The account servicer/owner relationship may be:
// - a central securities depository or another settlement market infrastructure acting on behalf of their participants
// - an agent (sub-custodian) acting on behalf of their global custodian customer, or
// - a custodian acting on behalf of an investment management institution or a broker/dealer.
// Usage
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type IntraPositionMovementStatusAdviceV01 struct {

	// Information that unambiguously identifies an IntraPositionMovementStatusAdvice message as known by the account servicer.
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// Unambiguous identification of a transaction as per the account owner (or the instructing party managing the account).
	TransactionIdentification *iso20022.TransactionIdentifications3 `xml:"TxId"`

	// Provides details on the processing status of the transaction.
	ProcessingStatus *iso20022.IntraPositionProcessingStatus1Choice `xml:"PrcgSts,omitempty"`

	// Provides the status of settlement of a transaction.
	SettlementStatus *iso20022.SettlementStatus2Choice `xml:"SttlmSts,omitempty"`

	// Identifies the details of the transaction.
	TransactionDetails *iso20022.IntraPositionDetails4 `xml:"TxDtls,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	Extension []*iso20022.Extension2 `xml:"Xtnsn,omitempty"`

}


func (i *IntraPositionMovementStatusAdviceV01) AddIdentification() *iso20022.DocumentIdentification11 {
	i.Identification = new(iso20022.DocumentIdentification11)
	return i.Identification
}

func (i *IntraPositionMovementStatusAdviceV01) AddTransactionIdentification() *iso20022.TransactionIdentifications3 {
	i.TransactionIdentification = new(iso20022.TransactionIdentifications3)
	return i.TransactionIdentification
}

func (i *IntraPositionMovementStatusAdviceV01) AddProcessingStatus() *iso20022.IntraPositionProcessingStatus1Choice {
	i.ProcessingStatus = new(iso20022.IntraPositionProcessingStatus1Choice)
	return i.ProcessingStatus
}

func (i *IntraPositionMovementStatusAdviceV01) AddSettlementStatus() *iso20022.SettlementStatus2Choice {
	i.SettlementStatus = new(iso20022.SettlementStatus2Choice)
	return i.SettlementStatus
}

func (i *IntraPositionMovementStatusAdviceV01) AddTransactionDetails() *iso20022.IntraPositionDetails4 {
	i.TransactionDetails = new(iso20022.IntraPositionDetails4)
	return i.TransactionDetails
}

func (i *IntraPositionMovementStatusAdviceV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	i.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return i.MessageOriginator
}

func (i *IntraPositionMovementStatusAdviceV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	i.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return i.MessageRecipient
}

func (i *IntraPositionMovementStatusAdviceV01) AddExtension() *iso20022.Extension2 {
	newValue := new (iso20022.Extension2)
	i.Extension = append(i.Extension, newValue)
	return newValue
}

