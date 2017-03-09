package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01700103 struct {
	XMLName xml.Name                           `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.017.001.03 Document"`
	Message *ForwardDataSetSubmissionReportV03 `xml:"FwdDataSetSubmissnRpt"`
}

func (d *Document01700103) AddMessage() *ForwardDataSetSubmissionReportV03 {
	d.Message = new(ForwardDataSetSubmissionReportV03)
	return d.Message
}

// Scope
// The ForwardDataSetSubmissionReport message is sent by the matching application to the counterparty(ies) of the submitter of data sets.
// This message is used to pass on information related to the purchasing agreement(s) covered by the transaction(s) referred to in the message.
// Usage
// The ForwardDataSetSubmission message can be sent by the matching application to forward the details of a DataSetSubmission message that it has obtained.
type ForwardDataSetSubmissionReportV03 struct {

	// Identifies the report.
	ReportIdentification *iso20022.MessageIdentification1 `xml:"RptId"`

	// Identifies the transactions that this submission relates to and provides associated information.
	RelatedTransactionReferences []*iso20022.DataSetSubmissionReferences4 `xml:"RltdTxRefs"`

	// This reference must be used for all data sets belonging to the same submission group.
	CommonSubmissionReference *iso20022.SimpleIdentificationInformation `xml:"CmonSubmissnRef"`

	// The financial institution that has submitted the data sets to the matching engine.
	Submitter *iso20022.BICIdentification1 `xml:"Submitr"`

	// The financial institution of the buyer, uniquely identified by its BIC.
	BuyerBank *iso20022.BICIdentification1 `xml:"BuyrBk"`

	// The financial institution of the seller, uniquely identified by its BIC.
	SellerBank *iso20022.BICIdentification1 `xml:"SellrBk"`

	// Commercial information that has been submitted to the matching application by the other party.
	CommercialDataSet *iso20022.CommercialDataSet3 `xml:"ComrclDataSet,omitempty"`

	// Transport information that has been submitted to the matching application by the other party.
	TransportDataSet *iso20022.TransportDataSet3 `xml:"TrnsprtDataSet,omitempty"`

	// Insurance information that has been submitted to the matching application by the other party.
	InsuranceDataSet *iso20022.InsuranceDataSet1 `xml:"InsrncDataSet,omitempty"`

	// Certificate information that has been submitted to the matching application by the other party.
	CertificateDataSet []*iso20022.CertificateDataSet1 `xml:"CertDataSet,omitempty"`

	// Other certificate information that has been submitted to the matching application by the other party.
	OtherCertificateDataSet []*iso20022.OtherCertificateDataSet1 `xml:"OthrCertDataSet,omitempty"`

	// Next processing step required.
	RequestForAction *iso20022.PendingActivity2 `xml:"ReqForActn,omitempty"`
}

func (f *ForwardDataSetSubmissionReportV03) AddReportIdentification() *iso20022.MessageIdentification1 {
	f.ReportIdentification = new(iso20022.MessageIdentification1)
	return f.ReportIdentification
}

func (f *ForwardDataSetSubmissionReportV03) AddRelatedTransactionReferences() *iso20022.DataSetSubmissionReferences4 {
	newValue := new(iso20022.DataSetSubmissionReferences4)
	f.RelatedTransactionReferences = append(f.RelatedTransactionReferences, newValue)
	return newValue
}

func (f *ForwardDataSetSubmissionReportV03) AddCommonSubmissionReference() *iso20022.SimpleIdentificationInformation {
	f.CommonSubmissionReference = new(iso20022.SimpleIdentificationInformation)
	return f.CommonSubmissionReference
}

func (f *ForwardDataSetSubmissionReportV03) AddSubmitter() *iso20022.BICIdentification1 {
	f.Submitter = new(iso20022.BICIdentification1)
	return f.Submitter
}

func (f *ForwardDataSetSubmissionReportV03) AddBuyerBank() *iso20022.BICIdentification1 {
	f.BuyerBank = new(iso20022.BICIdentification1)
	return f.BuyerBank
}

func (f *ForwardDataSetSubmissionReportV03) AddSellerBank() *iso20022.BICIdentification1 {
	f.SellerBank = new(iso20022.BICIdentification1)
	return f.SellerBank
}

func (f *ForwardDataSetSubmissionReportV03) AddCommercialDataSet() *iso20022.CommercialDataSet3 {
	f.CommercialDataSet = new(iso20022.CommercialDataSet3)
	return f.CommercialDataSet
}

func (f *ForwardDataSetSubmissionReportV03) AddTransportDataSet() *iso20022.TransportDataSet3 {
	f.TransportDataSet = new(iso20022.TransportDataSet3)
	return f.TransportDataSet
}

func (f *ForwardDataSetSubmissionReportV03) AddInsuranceDataSet() *iso20022.InsuranceDataSet1 {
	f.InsuranceDataSet = new(iso20022.InsuranceDataSet1)
	return f.InsuranceDataSet
}

func (f *ForwardDataSetSubmissionReportV03) AddCertificateDataSet() *iso20022.CertificateDataSet1 {
	newValue := new(iso20022.CertificateDataSet1)
	f.CertificateDataSet = append(f.CertificateDataSet, newValue)
	return newValue
}

func (f *ForwardDataSetSubmissionReportV03) AddOtherCertificateDataSet() *iso20022.OtherCertificateDataSet1 {
	newValue := new(iso20022.OtherCertificateDataSet1)
	f.OtherCertificateDataSet = append(f.OtherCertificateDataSet, newValue)
	return newValue
}

func (f *ForwardDataSetSubmissionReportV03) AddRequestForAction() *iso20022.PendingActivity2 {
	f.RequestForAction = new(iso20022.PendingActivity2)
	return f.RequestForAction
}
