package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01500103 struct {
	XMLName xml.Name        `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.015.001.03 Document"`
	Message *DeltaReportV03 `xml:"DltaRpt"`
}

func (d *Document01500103) AddMessage() *DeltaReportV03 {
	d.Message = new(DeltaReportV03)
	return d.Message
}

// Scope
// The DeltaReport message is sent by the matching application to the parties involved in the request of a baseline amendment.
// The message is used to list the differences between the established and the newly proposed baseline.
// Usage
// The DeltaReport message can be sent by the matching application to
// - the parties involved in the amendment of a baseline that has been established in the push-through mode. In the outlined scenario the message is sent to the requester of the amendment to acknowledge the receipt of the request and to list the differences between the established and the newly proposed baseline and to the counterparty to list the differences between the established and the newly proposed baseline and to request the acceptance or rejection of the amendment request,
// or
// - the party that has requested the amendment of a baseline established in the lodge mode. In the outlined scenario the message is used to confirm the changes to the baseline and to list the differences between the amended baseline and the baseline established earlier.
type DeltaReportV03 struct {

	// Identifies the report.
	ReportIdentification *iso20022.MessageIdentification1 `xml:"RptId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *iso20022.TransactionStatus4 `xml:"TxSts"`

	// Sequence number of the proposed baseline amendment.
	AmendmentNumber *iso20022.Count1 `xml:"AmdmntNb"`

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

	// Reference to the identification of the baseline included in the amendment request.
	SubmitterProposedBaselineReference *iso20022.DocumentIdentification1 `xml:"SubmitrPropsdBaselnRef"`

	// Detailed comparison between the currently established baseline elements and the proposed ones.
	UpdatedElement []*iso20022.ComparisonResult2 `xml:"UpdtdElmt"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (d *DeltaReportV03) AddReportIdentification() *iso20022.MessageIdentification1 {
	d.ReportIdentification = new(iso20022.MessageIdentification1)
	return d.ReportIdentification
}

func (d *DeltaReportV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	d.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return d.TransactionIdentification
}

func (d *DeltaReportV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	d.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return d.EstablishedBaselineIdentification
}

func (d *DeltaReportV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	d.TransactionStatus = new(iso20022.TransactionStatus4)
	return d.TransactionStatus
}

func (d *DeltaReportV03) AddAmendmentNumber() *iso20022.Count1 {
	d.AmendmentNumber = new(iso20022.Count1)
	return d.AmendmentNumber
}

func (d *DeltaReportV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new(iso20022.DocumentIdentification5)
	d.UserTransactionReference = append(d.UserTransactionReference, newValue)
	return newValue
}

func (d *DeltaReportV03) AddBuyer() *iso20022.PartyIdentification26 {
	d.Buyer = new(iso20022.PartyIdentification26)
	return d.Buyer
}

func (d *DeltaReportV03) AddSeller() *iso20022.PartyIdentification26 {
	d.Seller = new(iso20022.PartyIdentification26)
	return d.Seller
}

func (d *DeltaReportV03) AddBuyerBank() *iso20022.BICIdentification1 {
	d.BuyerBank = new(iso20022.BICIdentification1)
	return d.BuyerBank
}

func (d *DeltaReportV03) AddSellerBank() *iso20022.BICIdentification1 {
	d.SellerBank = new(iso20022.BICIdentification1)
	return d.SellerBank
}

func (d *DeltaReportV03) AddSubmitterProposedBaselineReference() *iso20022.DocumentIdentification1 {
	d.SubmitterProposedBaselineReference = new(iso20022.DocumentIdentification1)
	return d.SubmitterProposedBaselineReference
}

func (d *DeltaReportV03) AddUpdatedElement() *iso20022.ComparisonResult2 {
	newValue := new(iso20022.ComparisonResult2)
	d.UpdatedElement = append(d.UpdatedElement, newValue)
	return newValue
}

func (d *DeltaReportV03) AddRequestForAction() *iso20022.PendingActivity2 {
	d.RequestForAction = new(iso20022.PendingActivity2)
	return d.RequestForAction
}
