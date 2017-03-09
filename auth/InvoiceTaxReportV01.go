package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03400101 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:auth.034.001.01 Document"`
	Message *InvoiceTaxReportV01 `xml:"InvcTaxRpt"`
}

func (d *Document03400101) AddMessage() *InvoiceTaxReportV01 {
	d.Message = new(InvoiceTaxReportV01)
	return d.Message
}

// The InvoiceTaxReport message is sent by tax responsible to tax authority. Tax authorities require corporates to report their sales based value added tax (VAT). This message is targeted to this reporting based on information in sales invoices and card transactions.
type InvoiceTaxReportV01 struct {

	// Defines message level identification, number of individual tax reports and tax authority.
	InvoiceTaxReportHeader *iso20022.TaxReportHeader1 `xml:"InvcTaxRptHdr"`

	// Contains all needed party details for tax agency (sender of the TaxReport) and tax authority (receiver of the TaxReport) and the details of the reported sales transaction and calculated tax related that specific business transaction.
	TaxReport []*iso20022.TaxReport1 `xml:"TaxRpt"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InvoiceTaxReportV01) AddInvoiceTaxReportHeader() *iso20022.TaxReportHeader1 {
	i.InvoiceTaxReportHeader = new(iso20022.TaxReportHeader1)
	return i.InvoiceTaxReportHeader
}

func (i *InvoiceTaxReportV01) AddTaxReport() *iso20022.TaxReport1 {
	newValue := new(iso20022.TaxReport1)
	i.TaxReport = append(i.TaxReport, newValue)
	return newValue
}

func (i *InvoiceTaxReportV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
