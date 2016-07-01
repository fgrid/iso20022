package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03700204 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.002.04 Document"`
	Message *PortfolioTransferNotification002V04 `xml:"PrtflTrfNtfctn"`
}

func (d *Document03700204) AddMessage() *PortfolioTransferNotification002V04 {
	d.Message = new(PortfolioTransferNotification002V04)
	return d.Message
}

// Scope
// An account servicer sends a PortfolioTransferNotification to another account servicer to exchange transfer settlement details information during a retail or institutional client portfolio transfer.
// The account servicers will typically be local agents or global custodians acting on behalf of an investment management institution, a broker/dealer or a retail client.
// 
// Usage
// By exchange of transfer settlement details, it is understood the providing, by the delivering account servicer to the receiving account servicer, of the settlement details (trade date, settlement date, delivering settlement chain, quantities, etc.) of the individual transfers that will take place during a full or partial portfolio transfer. This delivering account servicer message may also include, for validation, the receiving settlement chain as provided by the client. In case the receiving settlement chain is not available to the delivering account servicer, the receiving account servicer may in return provide to the delivering account servicer the receiving settlement chain using the same message.
// 
// The message may also be used to:
// - re-send a message previously sent,
// - provide a third party with a copy of a message for information,
// - re-send to a third party a copy of a message for information
// using the relevant elements in the Business Application Header.
type PortfolioTransferNotification002V04 struct {

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Provides general information on the notification
	StatementGeneralDetails *iso20022.Statement50 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification109 `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount30 `xml:"SfkpgAcct"`

	// Details of transfer.
	TransferNotificationDetails []*iso20022.SecuritiesTradeDetails57 `xml:"TrfNtfctnDtls,omitempty"`

}


func (p *PortfolioTransferNotification002V04) AddPagination() *iso20022.Pagination {
	p.Pagination = new(iso20022.Pagination)
	return p.Pagination
}

func (p *PortfolioTransferNotification002V04) AddStatementGeneralDetails() *iso20022.Statement50 {
	p.StatementGeneralDetails = new(iso20022.Statement50)
	return p.StatementGeneralDetails
}

func (p *PortfolioTransferNotification002V04) AddAccountOwner() *iso20022.PartyIdentification109 {
	p.AccountOwner = new(iso20022.PartyIdentification109)
	return p.AccountOwner
}

func (p *PortfolioTransferNotification002V04) AddSafekeepingAccount() *iso20022.SecuritiesAccount30 {
	p.SafekeepingAccount = new(iso20022.SecuritiesAccount30)
	return p.SafekeepingAccount
}

func (p *PortfolioTransferNotification002V04) AddTransferNotificationDetails() *iso20022.SecuritiesTradeDetails57 {
	newValue := new (iso20022.SecuritiesTradeDetails57)
	p.TransferNotificationDetails = append(p.TransferNotificationDetails, newValue)
	return newValue
}

