package auth

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document00100101 struct {
	XMLName xml.Name                      `xml:"urn:iso:std:iso:20022:tech:xsd:auth.001.001.01 Document"`
	Message *InformationRequestOpeningV01 `xml:"InfReqOpng"`
}

func (d *Document00100101) AddMessage() *InformationRequestOpeningV01 {
	d.Message = new(InformationRequestOpeningV01)
	return d.Message
}

// This message is sent by the authorities (police, customs, tax authorities, enforcement authorities) to a financial institution to request account and other banking and financial information. Requested information can relate to accounts, their signatories and beneficiaries and co-owners as well as movements plus positions on these accounts.
//
// Requests are underpinned by specific legal texts.
type InformationRequestOpeningV01 struct {

	// Unique identification for the specific investigation as known by the requesting party.
	InvestigationIdentification *iso20022.Max35Text `xml:"InvstgtnId"`

	// Provides details on the legal basis of the request.
	LegalMandateBasis *iso20022.LegalMandate1 `xml:"LglMndtBsis"`

	// Specifies the confidentiality status of the investigation.
	ConfidentialityStatus *iso20022.YesNoIndicator `xml:"CnfdtltySts"`

	// Specifies the date by when the financial institutiion needs to provide a response.
	DueDate *iso20022.DueDate1 `xml:"DueDt,omitempty"`

	// Specifies the dates between which period the authority requests that the financial institution provides a response to the information request.
	InvestigationPeriod *iso20022.DateOrDateTimePeriodChoice `xml:"InvstgtnPrd"`

	// Specifies the the search criteria for the financial institution to perform the search on. The search criteria can be an account, a customer identification or a payment instrument type.
	SearchCriteria *iso20022.SearchCriteria1Choice `xml:"SchCrit"`

	// Additional information that can not be captured in the structured fields and/or any other specific block.
	SupplementaryData []*iso20022.SupplementaryData1 `xml:"SplmtryData,omitempty"`
}

func (i *InformationRequestOpeningV01) SetInvestigationIdentification(value string) {
	i.InvestigationIdentification = (*iso20022.Max35Text)(&value)
}

func (i *InformationRequestOpeningV01) AddLegalMandateBasis() *iso20022.LegalMandate1 {
	i.LegalMandateBasis = new(iso20022.LegalMandate1)
	return i.LegalMandateBasis
}

func (i *InformationRequestOpeningV01) SetConfidentialityStatus(value string) {
	i.ConfidentialityStatus = (*iso20022.YesNoIndicator)(&value)
}

func (i *InformationRequestOpeningV01) AddDueDate() *iso20022.DueDate1 {
	i.DueDate = new(iso20022.DueDate1)
	return i.DueDate
}

func (i *InformationRequestOpeningV01) AddInvestigationPeriod() *iso20022.DateOrDateTimePeriodChoice {
	i.InvestigationPeriod = new(iso20022.DateOrDateTimePeriodChoice)
	return i.InvestigationPeriod
}

func (i *InformationRequestOpeningV01) AddSearchCriteria() *iso20022.SearchCriteria1Choice {
	i.SearchCriteria = new(iso20022.SearchCriteria1Choice)
	return i.SearchCriteria
}

func (i *InformationRequestOpeningV01) AddSupplementaryData() *iso20022.SupplementaryData1 {
	newValue := new(iso20022.SupplementaryData1)
	i.SupplementaryData = append(i.SupplementaryData, newValue)
	return newValue
}
