package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02500101 struct {
	XMLName xml.Name                                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.025.001.01 Document"`
	Message *CurrencyControlSupportingDocumentDeliveryV01 `xml:"CcyCtrlSpprtgDocDlvry"`
}

func (d *Document02500101) AddMessage() *CurrencyControlSupportingDocumentDeliveryV01 {
	d.Message = new(CurrencyControlSupportingDocumentDeliveryV01)
	return d.Message
}

// The CurrencyControlSupportingDocumentDelivery message is sent by either the reporting party (respectively the registration agent or the registration agent (respectively the reporting party) in response to the supporting document request.
type CurrencyControlSupportingDocumentDeliveryV01 struct {

	// Characteristics shared by all individual items included in the message.
	GroupHeader *iso20022.CurrencyControlHeader3 `xml:"GrpHdr"`

	// Details of the supporting document provided for the registered contract.
	SupportingDocument []*iso20022.SupportingDocument1 `xml:"SpprtgDoc"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *CurrencyControlSupportingDocumentDeliveryV01) AddGroupHeader() *iso20022.CurrencyControlHeader3 {
	c.GroupHeader = new(iso20022.CurrencyControlHeader3)
	return c.GroupHeader
}

func (c *CurrencyControlSupportingDocumentDeliveryV01) AddSupportingDocument() *iso20022.SupportingDocument1 {
	newValue := new(iso20022.SupportingDocument1)
	c.SupportingDocument = append(c.SupportingDocument, newValue)
	return newValue
}

func (c *CurrencyControlSupportingDocumentDeliveryV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
