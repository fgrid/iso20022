package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01000103 struct {
	XMLName xml.Name                `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.010.001.03 Document"`
	Message *BaselineMatchReportV03 `xml:"BaselnMtchRpt"`
}

func (d *Document01000103) AddMessage() *BaselineMatchReportV03 {
	d.Message = new(BaselineMatchReportV03)
	return d.Message
}

// Scope
// The BaselineMatchReport message is sent by the matching application to the parties involved in the establishment of a transaction.
// The message is used to inform about either the successful establishment of a transaction (baseline) or the mis-matches found between two baseline initiation messages.
// Usage
// The BaselineMatchReport message can be sent by the matching application to
// - the parties involved in the establishment of a transaction in the push-through mode, that is the senders of InitialBaselineSubmission and BaselineReSubmission messages including the instruction push-through. In the outlined scenario the message is used to inform either about the successful establishment of a transaction in the matching application or about mis-matches found between two baseline initiation messages,or
// - the party establishing a transaction in the lodge mode, that is the sender of an InitialBaselineSubmission message including the instruction lodge. In the outlined scenario the message is used to inform about the successful establishment of a transaction in the matching application.
type BaselineMatchReportV03 struct {

	// Identifies the report.
	ReportIdentification *iso20022.MessageIdentification1 `xml:"RptId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *iso20022.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for each financial institution which is a party to the transaction.
	UserTransactionReference []*iso20022.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Party that buys goods or services, or a financial instrument.
	Buyer *iso20022.PartyIdentification26 `xml:"Buyr"`

	// Party that sells goods or services, or a financial instrument.
	Seller *iso20022.PartyIdentification26 `xml:"Sellr"`

	// The financial institution of the buyer, uniquely identified by its BIC.
	BuyerBank *iso20022.BICIdentification1 `xml:"BuyrBk"`

	// The financial institution of the seller, uniquely identified by its BIC.
	SellerBank *iso20022.BICIdentification1 `xml:"SellrBk"`

	// Specifies the number of matching trials for a baseline.
	BaselineEstablishmentTrials *iso20022.Limit1 `xml:"BaselnEstblishmtTrils"`

	// Identifies the two baselines compared in this report.
	ComparedDocumentReference []*iso20022.DocumentIdentification4 `xml:"CmpardDocRef"`

	// Description of the differences between the two proposed baselines
	Report *iso20022.MisMatchReport3 `xml:"Rpt"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (b *BaselineMatchReportV03) AddReportIdentification() *iso20022.MessageIdentification1 {
	b.ReportIdentification = new(iso20022.MessageIdentification1)
	return b.ReportIdentification
}

func (b *BaselineMatchReportV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	b.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return b.TransactionIdentification
}

func (b *BaselineMatchReportV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	b.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return b.EstablishedBaselineIdentification
}

func (b *BaselineMatchReportV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	b.TransactionStatus = new(iso20022.TransactionStatus4)
	return b.TransactionStatus
}

func (b *BaselineMatchReportV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new(iso20022.DocumentIdentification5)
	b.UserTransactionReference = append(b.UserTransactionReference, newValue)
	return newValue
}

func (b *BaselineMatchReportV03) AddBuyer() *iso20022.PartyIdentification26 {
	b.Buyer = new(iso20022.PartyIdentification26)
	return b.Buyer
}

func (b *BaselineMatchReportV03) AddSeller() *iso20022.PartyIdentification26 {
	b.Seller = new(iso20022.PartyIdentification26)
	return b.Seller
}

func (b *BaselineMatchReportV03) AddBuyerBank() *iso20022.BICIdentification1 {
	b.BuyerBank = new(iso20022.BICIdentification1)
	return b.BuyerBank
}

func (b *BaselineMatchReportV03) AddSellerBank() *iso20022.BICIdentification1 {
	b.SellerBank = new(iso20022.BICIdentification1)
	return b.SellerBank
}

func (b *BaselineMatchReportV03) AddBaselineEstablishmentTrials() *iso20022.Limit1 {
	b.BaselineEstablishmentTrials = new(iso20022.Limit1)
	return b.BaselineEstablishmentTrials
}

func (b *BaselineMatchReportV03) AddComparedDocumentReference() *iso20022.DocumentIdentification4 {
	newValue := new(iso20022.DocumentIdentification4)
	b.ComparedDocumentReference = append(b.ComparedDocumentReference, newValue)
	return newValue
}

func (b *BaselineMatchReportV03) AddReport() *iso20022.MisMatchReport3 {
	b.Report = new(iso20022.MisMatchReport3)
	return b.Report
}

func (b *BaselineMatchReportV03) AddRequestForAction() *iso20022.PendingActivity2 {
	b.RequestForAction = new(iso20022.PendingActivity2)
	return b.RequestForAction
}
