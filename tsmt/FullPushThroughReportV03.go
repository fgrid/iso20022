package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01800103 struct {
	XMLName xml.Name                  `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.018.001.03 Document"`
	Message *FullPushThroughReportV03 `xml:"FullPushThrghRpt"`
}

func (d *Document01800103) AddMessage() *FullPushThroughReportV03 {
	d.Message = new(FullPushThroughReportV03)
	return d.Message
}

// Scope
// The FullPushThroughReport message is sent by the matching application to a party involved in a transaction.
// This message is used to pass on information that the matching application has received from the submitter. The forwarded information can originate from an InitialBaselineSubmission or BaselineReSubmission or BaselineAmendmentRequest message.
// Usage
// The FullPushThroughReport message can be sent by the matching application to a party to convey
// - the details of an InitialBaselineSubmission message that it has obtained,or
// - the details of a BaselineResubmission message that it has obtained,or
// - the details of a BaselineAmendmentRequest message that it has obtained.
type FullPushThroughReportV03 struct {

	// Identifies the report.
	ReportIdentification *iso20022.MessageIdentification1 `xml:"RptId"`

	// Unique identification assigned by the matching application to the transaction.
	// This identification is to be used in any communication between the parties.
	TransactionIdentification *iso20022.SimpleIdentificationInformation `xml:"TxId"`

	// Unique identification assigned by the matching application to the baseline when it is established.
	EstablishedBaselineIdentification *iso20022.DocumentIdentification3 `xml:"EstblishdBaselnId,omitempty"`

	// Identifies the status of the transaction by means of a code.
	TransactionStatus *iso20022.TransactionStatus4 `xml:"TxSts"`

	// Reference to the transaction for the financial institution which submitted the baseline.
	UserTransactionReference []*iso20022.DocumentIdentification5 `xml:"UsrTxRef,omitempty"`

	// Specifies the type of report.
	ReportPurpose *iso20022.ReportType1 `xml:"RptPurp"`

	// Specifies the commercial details of the underlying transaction.
	PushedThroughBaseline *iso20022.Baseline3 `xml:"PushdThrghBaseln"`

	// Person to be contacted in the organisation of the buyer.
	BuyerContactPerson []*iso20022.ContactIdentification1 `xml:"BuyrCtctPrsn,omitempty"`

	// Person to be contacted in the organisation of the seller.
	SellerContactPerson []*iso20022.ContactIdentification1 `xml:"SellrCtctPrsn,omitempty"`

	// Person to be contacted in the buyer's bank.
	BuyerBankContactPerson []*iso20022.ContactIdentification1 `xml:"BuyrBkCtctPrsn,omitempty"`

	// Person to be contacted in the seller's bank.
	SellerBankContactPerson []*iso20022.ContactIdentification1 `xml:"SellrBkCtctPrsn,omitempty"`

	// Person to be contacted in another bank than the seller or buyer's bank.
	OtherBankContactPerson []*iso20022.ContactIdentification3 `xml:"OthrBkCtctPrsn,omitempty"`

	// Information on the next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (f *FullPushThroughReportV03) AddReportIdentification() *iso20022.MessageIdentification1 {
	f.ReportIdentification = new(iso20022.MessageIdentification1)
	return f.ReportIdentification
}

func (f *FullPushThroughReportV03) AddTransactionIdentification() *iso20022.SimpleIdentificationInformation {
	f.TransactionIdentification = new(iso20022.SimpleIdentificationInformation)
	return f.TransactionIdentification
}

func (f *FullPushThroughReportV03) AddEstablishedBaselineIdentification() *iso20022.DocumentIdentification3 {
	f.EstablishedBaselineIdentification = new(iso20022.DocumentIdentification3)
	return f.EstablishedBaselineIdentification
}

func (f *FullPushThroughReportV03) AddTransactionStatus() *iso20022.TransactionStatus4 {
	f.TransactionStatus = new(iso20022.TransactionStatus4)
	return f.TransactionStatus
}

func (f *FullPushThroughReportV03) AddUserTransactionReference() *iso20022.DocumentIdentification5 {
	newValue := new(iso20022.DocumentIdentification5)
	f.UserTransactionReference = append(f.UserTransactionReference, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddReportPurpose() *iso20022.ReportType1 {
	f.ReportPurpose = new(iso20022.ReportType1)
	return f.ReportPurpose
}

func (f *FullPushThroughReportV03) AddPushedThroughBaseline() *iso20022.Baseline3 {
	f.PushedThroughBaseline = new(iso20022.Baseline3)
	return f.PushedThroughBaseline
}

func (f *FullPushThroughReportV03) AddBuyerContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	f.BuyerContactPerson = append(f.BuyerContactPerson, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddSellerContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	f.SellerContactPerson = append(f.SellerContactPerson, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddBuyerBankContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	f.BuyerBankContactPerson = append(f.BuyerBankContactPerson, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddSellerBankContactPerson() *iso20022.ContactIdentification1 {
	newValue := new(iso20022.ContactIdentification1)
	f.SellerBankContactPerson = append(f.SellerBankContactPerson, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddOtherBankContactPerson() *iso20022.ContactIdentification3 {
	newValue := new(iso20022.ContactIdentification3)
	f.OtherBankContactPerson = append(f.OtherBankContactPerson, newValue)
	return newValue
}

func (f *FullPushThroughReportV03) AddRequestForAction() *iso20022.PendingActivity2 {
	f.RequestForAction = new(iso20022.PendingActivity2)
	return f.RequestForAction
}
