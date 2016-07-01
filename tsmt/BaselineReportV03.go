package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01100103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.011.001.03 Document"`
	Message *BaselineReportV03 `xml:"BaselnRpt"`
}

func (d *Document01100103) AddMessage() *BaselineReportV03 {
	d.Message = new(BaselineReportV03)
	return d.Message
}

// Scope
// The BaselineReport message is sent by the matching application to the parties involved in an amendment request or to the parties involved in a data set match.
// The message is used to report either a pre-calculation or final calculation of the dynamic part of an established baseline.
// Usage
// The BaselineReport message can be sent by the matching application to the parties involved in an amendment request for a transaction established in the push-through mode. In the outlined scenario, the message is sent
// - to the party requested to accept or reject an amendment request after the matching application has received a BaselineAmendmentRequest message. The message informs about the provisional status of the dynamic part of the baseline.
// - to the requester and the accepter of an amendment request after the matching application has received an AmendmentAcceptance message conveying the acceptance of the amendment request. The message informs about the actual status of the dynamic part of the baseline.
// or
// The BaselineReport message can be sent by the matching application to the party which has sent an amendment request for a transaction established in the lodge mode. In the outlined scenario the message is used to inform about the actual status of the dynamic part of the baseline.
// or
// The BaselineReport message can be sent by the matching application to the parties involved in a data set match. In the outlined scenario, the message is sent
// - to the submitter of the data set(s) in the case of a data set match for a transaction established in the lodge mode.
// - to the submitter of the data set(s) and to the counterparty in case of a data set match for a transaction established in the push-through mode.The message can be sent after a successful data-set match or after the acceptance of mis-matched data sets to inform about the actual status of the dynamic part of the baseline.
type BaselineReportV03 struct {

	// Identifies the report. 
	ReportIdentification *iso20022.MessageIdentification1 `xml:"RptId"`

	// Reference to the related message at the origin of the report or sent at the same time than the report.
	RelatedMessageReference *iso20022.MessageIdentification1 `xml:"RltdMsgRef,omitempty"`

	// Type of baseline report.
	ReportType *iso20022.ReportType2 `xml:"RptTp"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	// 
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.  
	EstablishedBaselineIdentification *iso20022.DocumentIdentification6 `xml:"EstblishdBaselnId"`

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

	// Information on the goods
	ReportedLineItem *iso20022.LineItem8 `xml:"RptdLineItm"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`

}


func (b *BaselineReportV03) AddReportIdentification() *iso20022.MessageIdentification1 {
	b.ReportIdentification = new(iso20022.MessageIdentification1)
	return b.ReportIdentification
}

func (b *BaselineReportV03) AddRelatedMessageReference() *iso20022.MessageIdentification1 {
	b.RelatedMessageReference = new(iso20022.MessageIdentification1)
	return b.RelatedMessageReference
}

func (b *BaselineReportV03) AddReportType() *iso20022.ReportType2 {
	b.ReportType = new(iso20022.ReportType2)
	return b.ReportType
}

func (b *BaselineReportV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	b.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return b.TransactionIdentification
}

func (b *BaselineReportV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification6 {
	b.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification6)
	return b.EstablishedBaselineIdentification
}

func (b *BaselineReportV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	b.TransactionStatus = new(iso20022.TransactionStatus4)
	return b.TransactionStatus
}

func (b *BaselineReportV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new (iso20022.DocumentIdentification5)
	b.UserTransactionReference = append(b.UserTransactionReference, newValue)
	return newValue
}

func (b *BaselineReportV03) AddBuyer() *iso20022.PartyIdentification26 {
	b.Buyer = new(iso20022.PartyIdentification26)
	return b.Buyer
}

func (b *BaselineReportV03) AddSeller() *iso20022.PartyIdentification26 {
	b.Seller = new(iso20022.PartyIdentification26)
	return b.Seller
}

func (b *BaselineReportV03) AddBuyerBank() *iso20022.BICIdentification1 {
	b.BuyerBank = new(iso20022.BICIdentification1)
	return b.BuyerBank
}

func (b *BaselineReportV03) AddSellerBank() *iso20022.BICIdentification1 {
	b.SellerBank = new(iso20022.BICIdentification1)
	return b.SellerBank
}

func (b *BaselineReportV03) AddReportedLineItem() *iso20022.LineItem8 {
	b.ReportedLineItem = new(iso20022.LineItem8)
	return b.ReportedLineItem
}

func (b *BaselineReportV03) AddRequestForAction() *iso20022.PendingActivity2 {
	b.RequestForAction = new(iso20022.PendingActivity2)
	return b.RequestForAction
}

