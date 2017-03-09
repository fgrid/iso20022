package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01800101 struct {
	XMLName xml.Name                        `xml:"urn:iso:std:iso:20022:tech:xsd:auth.018.001.01 Document"`
	Message *ContractRegistrationRequestV01 `xml:"CtrctRegnReq"`
}

func (d *Document01800101) AddMessage() *ContractRegistrationRequestV01 {
	d.Message = new(ContractRegistrationRequestV01)
	return d.Message
}

// The ContractRegistrationRequest message is sent by the reporting party to the registration agent to initiate the registration of a new contract subject to currency control.
type ContractRegistrationRequestV01 struct {

	// Characteristics shared by all individual items included in the message.
	GroupHeader *iso20022.CurrencyControlHeader1 `xml:"GrpHdr"`

	// Identifies the currency control contract details for which the registration is requested.
	ContractRegistration []*iso20022.ContractRegistration1 `xml:"CtrctRegn"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (c *ContractRegistrationRequestV01) AddGroupHeader() *iso20022.CurrencyControlHeader1 {
	c.GroupHeader = new(iso20022.CurrencyControlHeader1)
	return c.GroupHeader
}

func (c *ContractRegistrationRequestV01) AddContractRegistration() *iso20022.ContractRegistration1 {
	newValue := new(iso20022.ContractRegistration1)
	c.ContractRegistration = append(c.ContractRegistration, newValue)
	return newValue
}

func (c *ContractRegistrationRequestV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}
