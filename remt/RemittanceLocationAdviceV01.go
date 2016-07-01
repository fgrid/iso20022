package remt

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00200101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:remt.002.001.01 Document"`
	Message *RemittanceLocationAdviceV01 `xml:"RmtLctnAdvc"`
}

func (d *Document00200101) AddMessage() *RemittanceLocationAdviceV01 {
	d.Message = new(RemittanceLocationAdviceV01)
	return d.Message
}

// The RemittanceLocationAdvice message allows the originator of the message to identify where the remittance advice is located for a related payment.
type RemittanceLocationAdviceV01 struct {

	// Set of characteristics shared by all remittance location information included in the message.
	GroupHeader *iso20022.GroupHeader62 `xml:"GrpHdr"`

	// Provides information related to location and/or delivery of the remittance information.  This information is used to enable the matching of an entry with the items that the associated payment is intended to settle.
	RemittanceLocation []*iso20022.RemittanceLocation3 `xml:"RmtLctn"`

	// Additional information that cannot be  captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (r *RemittanceLocationAdviceV01) AddGroupHeader() *iso20022.GroupHeader62 {
	r.GroupHeader = new(iso20022.GroupHeader62)
	return r.GroupHeader
}

func (r *RemittanceLocationAdviceV01) AddRemittanceLocation() *iso20022.RemittanceLocation3 {
	newValue := new (iso20022.RemittanceLocation3)
	r.RemittanceLocation = append(r.RemittanceLocation, newValue)
	return newValue
}

func (r *RemittanceLocationAdviceV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	r.SupplementaryData = append(r.SupplementaryData, newValue)
	return newValue
}

