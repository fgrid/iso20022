package remt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100103 struct {
	XMLName xml.Name             `xml:"urn:iso:std:iso:20022:tech:xsd:remt.001.001.03 Document"`
	Message *RemittanceAdviceV03 `xml:"RmtAdvc"`
}

func (d *Document00100103) AddMessage() *RemittanceAdviceV03 {
	d.Message = new(RemittanceAdviceV03)
	return d.Message
}

// The RemittanceAdvice message allows the originator to provide remittance details that can be associated with a payment.
type RemittanceAdviceV03 struct {

	// Set of characteristics shared by all remittance information included in the message.
	GroupHeader *iso20022.GroupHeader62 `xml:"GrpHdr"`

	// Provides information to enable the matching of an entry with the items that the associated payment is intended to settle, such as commercial invoices in an accounts' receivable system, tax obligations, or garnishment orders.
	RemittanceInformation []*iso20022.RemittanceInformation13 `xml:"RmtInf"`

	// Additional information that cannot be  captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (r *RemittanceAdviceV03) AddGroupHeader() *iso20022.GroupHeader62 {
	r.GroupHeader = new(iso20022.GroupHeader62)
	return r.GroupHeader
}

func (r *RemittanceAdviceV03) AddRemittanceInformation() *iso20022.RemittanceInformation13 {
	newValue := new(iso20022.RemittanceInformation13)
	r.RemittanceInformation = append(r.RemittanceInformation, newValue)
	return newValue
}

func (r *RemittanceAdviceV03) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}
