package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01300103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.013.001.03 Document"`
	Message *DataSetMatchReportV03 `xml:"DataSetMtchRpt"`
}

func (d *Document01300103) AddMessage() *DataSetMatchReportV03 {
	d.Message = new(DataSetMatchReportV03)
	return d.Message
}

// Scope
// The DataSetMatchReport message is sent by the matching application to the parties involved in a data set match.
// This message is used to either
// - inform about the successful match of data sets submitted with the instruction match or pre-match (DataSetSubmission message) and the related baseline,or
// - inform about mis-matches found between data sets submitted with the instruction match or pre-match (DataSetSubmission message) and the related baseline.
// Usage
// The DataSetMatchReport message can be sent by the matching application to the party requesting a data set pre-match for a transaction established in the push-through mode. In the outlined scenario, the DataSetMatchReport message will either inform about the successful pre-match or list the mis-matches between the data set(s) conveyed with the DataSetSubmission message and the related baseline.
// or
// The DataSetMatchReport message can be sent by the matching application to the parties involved in a data set match for a transaction established in the push-through mode. In the outlined scenario, the DataSetMatchReport message will either inform about the successful match or list the mis-matches between the data set(s) conveyed with the DataSetSubmission message and the related baseline.
// or
// The DataSetMatchReport message can be sent by the matching application to the party requesting a data set match or pre-match for a transaction established in the lodge mode. In the outlined scenario, the DataSetMatchReport will either inform about the successful match or list the mis-matches between the data set(s) conveyed with the DataSetSubmission message and the related baseline.
type DataSetMatchReportV03 struct {

	// Identifies the report. 
	ReportIdentification *iso20022.MessageIdentification1 `xml:"RptId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established. 
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId"`

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

	// Identifies the documents compared in this report.
	ComparedDocumentReference []*iso20022.DocumentIdentification10 `xml:"CmpardDocRef"`

	// Specifies whether the data set was submitted for match or pre-match.
	SubmissionType *iso20022.ReportType3 `xml:"SubmissnTp"`

	// Description of the differences between the data set(s) and the baseline.
	Report *iso20022.MisMatchReport3 `xml:"Rpt"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`

}


func (d *DataSetMatchReportV03) AddReportIdentification() *iso20022.MessageIdentification1 {
	d.ReportIdentification = new(iso20022.MessageIdentification1)
	return d.ReportIdentification
}

func (d *DataSetMatchReportV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	d.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return d.TransactionIdentification
}

func (d *DataSetMatchReportV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	d.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return d.EstablishedBaselineIdentification
}

func (d *DataSetMatchReportV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	d.TransactionStatus = new(iso20022.TransactionStatus4)
	return d.TransactionStatus
}

func (d *DataSetMatchReportV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new (iso20022.DocumentIdentification5)
	d.UserTransactionReference = append(d.UserTransactionReference, newValue)
	return newValue
}

func (d *DataSetMatchReportV03) AddBuyer() *iso20022.PartyIdentification26 {
	d.Buyer = new(iso20022.PartyIdentification26)
	return d.Buyer
}

func (d *DataSetMatchReportV03) AddSeller() *iso20022.PartyIdentification26 {
	d.Seller = new(iso20022.PartyIdentification26)
	return d.Seller
}

func (d *DataSetMatchReportV03) AddBuyerBank() *iso20022.BICIdentification1 {
	d.BuyerBank = new(iso20022.BICIdentification1)
	return d.BuyerBank
}

func (d *DataSetMatchReportV03) AddSellerBank() *iso20022.BICIdentification1 {
	d.SellerBank = new(iso20022.BICIdentification1)
	return d.SellerBank
}

func (d *DataSetMatchReportV03) AddComparedDocumentReference() *iso20022.DocumentIdentification10 {
	newValue := new (iso20022.DocumentIdentification10)
	d.ComparedDocumentReference = append(d.ComparedDocumentReference, newValue)
	return newValue
}

func (d *DataSetMatchReportV03) AddSubmissionType() *iso20022.ReportType3 {
	d.SubmissionType = new(iso20022.ReportType3)
	return d.SubmissionType
}

func (d *DataSetMatchReportV03) AddReport() *iso20022.MisMatchReport3 {
	d.Report = new(iso20022.MisMatchReport3)
	return d.Report
}

func (d *DataSetMatchReportV03) AddRequestForAction() *iso20022.PendingActivity2 {
	d.RequestForAction = new(iso20022.PendingActivity2)
	return d.RequestForAction
}

