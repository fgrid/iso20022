package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01400105 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.014.001.05 Document"`
	Message *DataSetSubmissionV05 `xml:"DataSetSubmissn"`
}

func (d *Document01400105) AddMessage() *DataSetSubmissionV05 {
	d.Message = new(DataSetSubmissionV05)
	return d.Message
}

// Scope
// The DataSetSubmission message is sent by a party involved in the transaction to the matching application.
// This message is used to trigger either a match or a pre-match of the information submitted with the message.
// Usage
// The DataSetSubmission message can be sent by either party with the instruction pre-match. In the outlined scenario, the matching application will compare the data set(s) conveyed by the DataSetSubmission message with the established baseline and report the matching result to the requester of the data set pre-match by sending a DataSetMatchReport message.
// or
// The DataSetSubmission message can be sent by the party specified in the baseline as data set submitter with the instruction match. In the outlined scenario, the matching application will compare the data set(s) conveyed by the DataSetSubmission message with the established baseline and report the matching result to
// - the parties involved in a transaction established in the push-through mode, or
// - the initiator of a transaction established in the lodge mode.
// The DataSetSubmission message can be used to submit multiple data sets for multiple transactions (baselines) at the same time. However, all transactions (baselines) covered by the message must be for the same parties, that is transaction initiator and counterparty, as well as for the same buyer and seller.
// The DataSetSubmission message consists of data reflecting trade information related to the purchasing agreement covered by the transaction(s), for example shipment date, invoice amount.
type DataSetSubmissionV05 struct {

	// Identifies the submitted information.
	SubmissionIdentification *iso20022.MessageIdentification1 `xml:"SubmissnId"`

	// Identifies the transactions that this submission relates to and provides associated information.
	RelatedTransactionReferences []*iso20022.DataSetSubmissionReferences3 `xml:"RltdTxRefs"`

	// This reference must be used for all data sets belonging to the same submission group.
	CommonSubmissionReference *iso20022.SimpleIdentificationInformation `xml:"CmonSubmissnRef"`

	// Specifies the instruction given by the submitter. 
	Instruction *iso20022.InstructionType3 `xml:"Instr"`

	// The financial institution of the buyer, uniquely identified by its BIC. 
	BuyerBank *iso20022.BICIdentification1 `xml:"BuyrBk"`

	// The financial institution of the seller, uniquely identified by its BIC. 
	SellerBank *iso20022.BICIdentification1 `xml:"SellrBk"`

	// Commercial information that is submitted to the matching application for processing.
	CommercialDataSet *iso20022.CommercialDataSet5 `xml:"ComrclDataSet,omitempty"`

	// Transport information that is submitted to the matching application for processing.
	TransportDataSet *iso20022.TransportDataSet5 `xml:"TrnsprtDataSet,omitempty"`

	// Insurance information that is submitted to the matching application for processing.
	InsuranceDataSet *iso20022.InsuranceDataSet1 `xml:"InsrncDataSet,omitempty"`

	// Certificate information that is submitted to the matching application for processing.
	CertificateDataSet []*iso20022.CertificateDataSet2 `xml:"CertDataSet,omitempty"`

	// Other certificate information that is submitted to the matching application for processing.
	OtherCertificateDataSet []*iso20022.OtherCertificateDataSet2 `xml:"OthrCertDataSet,omitempty"`

}


func (d *DataSetSubmissionV05) AddSubmissionIdentification() *iso20022.MessageIdentification1 {
	d.SubmissionIdentification = new(iso20022.MessageIdentification1)
	return d.SubmissionIdentification
}

func (d *DataSetSubmissionV05) AddRelatedTransactionReferences() *iso20022.DataSetSubmissionReferences3 {
	newValue := new (iso20022.DataSetSubmissionReferences3)
	d.RelatedTransactionReferences = append(d.RelatedTransactionReferences, newValue)
	return newValue
}

func (d *DataSetSubmissionV05) AddCommonSubmissionReference() *iso20022.SimpleIdentificationInformation {
	d.CommonSubmissionReference = new(iso20022.SimpleIdentificationInformation)
	return d.CommonSubmissionReference
}

func (d *DataSetSubmissionV05) AddInstruction() *iso20022.InstructionType3 {
	d.Instruction = new(iso20022.InstructionType3)
	return d.Instruction
}

func (d *DataSetSubmissionV05) AddBuyerBank() *iso20022.BICIdentification1 {
	d.BuyerBank = new(iso20022.BICIdentification1)
	return d.BuyerBank
}

func (d *DataSetSubmissionV05) AddSellerBank() *iso20022.BICIdentification1 {
	d.SellerBank = new(iso20022.BICIdentification1)
	return d.SellerBank
}

func (d *DataSetSubmissionV05) AddCommercialDataSet() *iso20022.CommercialDataSet5 {
	d.CommercialDataSet = new(iso20022.CommercialDataSet5)
	return d.CommercialDataSet
}

func (d *DataSetSubmissionV05) AddTransportDataSet() *iso20022.TransportDataSet5 {
	d.TransportDataSet = new(iso20022.TransportDataSet5)
	return d.TransportDataSet
}

func (d *DataSetSubmissionV05) AddInsuranceDataSet() *iso20022.InsuranceDataSet1 {
	d.InsuranceDataSet = new(iso20022.InsuranceDataSet1)
	return d.InsuranceDataSet
}

func (d *DataSetSubmissionV05) AddCertificateDataSet() *iso20022.CertificateDataSet2 {
	newValue := new (iso20022.CertificateDataSet2)
	d.CertificateDataSet = append(d.CertificateDataSet, newValue)
	return newValue
}

func (d *DataSetSubmissionV05) AddOtherCertificateDataSet() *iso20022.OtherCertificateDataSet2 {
	newValue := new (iso20022.OtherCertificateDataSet2)
	d.OtherCertificateDataSet = append(d.OtherCertificateDataSet, newValue)
	return newValue
}

