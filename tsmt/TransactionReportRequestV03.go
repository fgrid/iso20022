package tsmt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document04200103 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:tsmt.042.001.03 Document"`
	Message *TransactionReportRequestV03 `xml:"TxRptReq"`
}

func (d *Document04200103) AddMessage() *TransactionReportRequestV03 {
	d.Message = new(TransactionReportRequestV03)
	return d.Message
}

// Scope
// The TransactionReportRequest message is sent by a party involved in a transaction to the matching application.
// This message is used to request a report on details of transactions registered in the matching application.
// Usage
// The TransactionReportRequest message can be sent by either party involved in a transaction to request a report on a variety of details of all transactions that the requester is involved in. For example, the message can be used to request a report on all transactions that the requester is involved in with a certain customer.
type TransactionReportRequestV03 struct {

	// Identifies the request message.
	RequestIdentification *iso20022.MessageIdentification1 `xml:"ReqId"`

	// Parameters to be used as criteria for the content of the transaction report.
	ReportSpecification *iso20022.ReportSpecification4 `xml:"RptSpcfctn"`

}


func (t *TransactionReportRequestV03) AddRequestIdentification() *iso20022.MessageIdentification1 {
	t.RequestIdentification = new(iso20022.MessageIdentification1)
	return t.RequestIdentification
}

func (t *TransactionReportRequestV03) AddReportSpecification() *iso20022.ReportSpecification4 {
	t.ReportSpecification = new(iso20022.ReportSpecification4)
	return t.ReportSpecification
}

