package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document02100101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:auth.021.001.01 Document"`
	Message *ContractRegistrationAmendmentRequestV01 `xml:"CtrctRegnAmdmntReq"`
}

func (d *Document02100101) AddMessage() *ContractRegistrationAmendmentRequestV01 {
	d.Message = new(ContractRegistrationAmendmentRequestV01)
	return d.Message
}

// The ContractRegistrationAmendmentRequest message is sent by the reporting party to the registration agent to amend the registered contract subject to currency control.
type ContractRegistrationAmendmentRequestV01 struct {

	// Characteristics shared by all individual items included in the message.
	GroupHeader *iso20022.CurrencyControlHeader1 `xml:"GrpHdr"`

	// Details on the amendment of the registered contract.
	ContractRegistrationAmendment []*iso20022.RegisteredContract1 `xml:"CtrctRegnAmdmnt"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (c *ContractRegistrationAmendmentRequestV01) AddGroupHeader() *iso20022.CurrencyControlHeader1 {
	c.GroupHeader = new(iso20022.CurrencyControlHeader1)
	return c.GroupHeader
}

func (c *ContractRegistrationAmendmentRequestV01) AddContractRegistrationAmendment() *iso20022.RegisteredContract1 {
	newValue := new (iso20022.RegisteredContract1)
	c.ContractRegistrationAmendment = append(c.ContractRegistrationAmendment, newValue)
	return newValue
}

func (c *ContractRegistrationAmendmentRequestV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

