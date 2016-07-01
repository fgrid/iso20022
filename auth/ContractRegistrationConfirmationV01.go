package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document01900101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:auth.019.001.01 Document"`
	Message *ContractRegistrationConfirmationV01 `xml:"CtrctRegnConf"`
}

func (d *Document01900101) AddMessage() *ContractRegistrationConfirmationV01 {
	d.Message = new(ContractRegistrationConfirmationV01)
	return d.Message
}

// The ContractRegistrationConfirmation message is sent by the registration agent to the reporting party to register the contract subject to currency control.
type ContractRegistrationConfirmationV01 struct {

	// Characteristics shared by all individual items included in the message.
	GroupHeader *iso20022.CurrencyControlHeader2 `xml:"GrpHdr"`

	// Identifies the contract details which is registered for currency control.
	RegisteredContract []*iso20022.RegisteredContract4 `xml:"RegdCtrct"`

	// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`

}


func (c *ContractRegistrationConfirmationV01) AddGroupHeader() *iso20022.CurrencyControlHeader2 {
	c.GroupHeader = new(iso20022.CurrencyControlHeader2)
	return c.GroupHeader
}

func (c *ContractRegistrationConfirmationV01) AddRegisteredContract() *iso20022.RegisteredContract4 {
	newValue := new (iso20022.RegisteredContract4)
	c.RegisteredContract = append(c.RegisteredContract, newValue)
	return newValue
}

func (c *ContractRegistrationConfirmationV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new (iso20022.SupplementaryData1)
	c.SupplementaryData = append(c.SupplementaryData, newValue)
	return newValue
}

