package sese

import (
	"encoding/xml"

	"github.com/fgrid/iso20022"
)

type Document03700101 struct {
	XMLName xml.Name `xml:"urn:iso:std:iso:20022:tech:xsd:sese.037.001.01 Document"`
	Message *PortfolioTransferNotificationV01 `xml:"PrtflTrfNtfctn"`
}

func (d *Document03700101) AddMessage() *PortfolioTransferNotificationV01 {
	d.Message = new(PortfolioTransferNotificationV01)
	return d.Message
}

// Scope
// An account servicer sends a PortfolioTransferNotification to another account servicer to exchange transfer settlement details information during a retail or institutional client portfolio transfers.
// The account servicers will typically be local agents or global custodians acting on behalf of an investment management institution, a broker/dealer or a retail client.
// Usage
// By exchange of transfer settlement details, it is understood the providing, by the delivering account servicer to the receiving account servicer, of the settlement details (trade date, settlement date, delivering settlement chain, quantities, etc.) of the individual transfers that will take place during a full or partial portfolio transfer. This delivering account servicer message may also include, for validation, the receiving settlement chain as provided by the client. In case the receiving settlement chain is not available to the delivering account servicer, the receiving account servicer may in return provide to the delivering account servicer the receiving settlement chain using the same message.
// The message may also be used to:
// - re-send a message previously sent (the sub-function of the message is Duplicate),
// - provide a third party with a copy of a message for information (the sub-function of the message is Copy),
// - re-send to a third party a copy of a message for information (the sub-function of the message is Copy Duplicate).
// ISO 15022 - 20022 Coexistence
// This ISO 20022 message is reversed engineered from ISO 15022. Both standards will coexist for a certain number of years. Until this coexistence period ends, the usage of certain data types is restricted to ensure interoperability between ISO 15022 and 20022 users. Compliance to these rules is mandatory in a coexistence environment.  The coexistence restrictions are described in a Textual Rule linked to the Message Items they concern. These coexistence textual rules are clearly identified as follows:  “CoexistenceXxxxRule”.
type PortfolioTransferNotificationV01 struct {

	// Information that unambiguously identifies a PortfolioTransferNotification message as known by the sender.
	Identification *iso20022.DocumentIdentification11 `xml:"Id"`

	// Page number of the message (within a statement) and continuation indicator to indicate that the statement is to continue or that the message is the last page of the statement.
	Pagination *iso20022.Pagination `xml:"Pgntn"`

	// Provides general information on the notification
	StatementGeneralDetails *iso20022.Statement19 `xml:"StmtGnlDtls"`

	// Party that legally owns the account.
	AccountOwner *iso20022.PartyIdentification13Choice `xml:"AcctOwnr,omitempty"`

	// Account to or from which a securities entry is made.
	SafekeepingAccount *iso20022.SecuritiesAccount13 `xml:"SfkpgAcct"`

	// Details of transfer.
	TransferNotificationDetails []*iso20022.SecuritiesTradeDetails7 `xml:"TrfNtfctnDtls,omitempty"`

	// Party that originated the message, if other than the sender.
	MessageOriginator *iso20022.PartyIdentification10Choice `xml:"MsgOrgtr,omitempty"`

	// Party that is the final destination of the message, if other than the receiver.
	MessageRecipient *iso20022.PartyIdentification10Choice `xml:"MsgRcpt,omitempty"`

}


func (p *PortfolioTransferNotificationV01) AddIdentification() *iso20022.DocumentIdentification11 {
	p.Identification = new(iso20022.DocumentIdentification11)
	return p.Identification
}

func (p *PortfolioTransferNotificationV01) AddPagination() *iso20022.Pagination {
	p.Pagination = new(iso20022.Pagination)
	return p.Pagination
}

func (p *PortfolioTransferNotificationV01) AddStatementGeneralDetails() *iso20022.Statement19 {
	p.StatementGeneralDetails = new(iso20022.Statement19)
	return p.StatementGeneralDetails
}

func (p *PortfolioTransferNotificationV01) AddAccountOwner() *iso20022.PartyIdentification13Choice {
	p.AccountOwner = new(iso20022.PartyIdentification13Choice)
	return p.AccountOwner
}

func (p *PortfolioTransferNotificationV01) AddSafekeepingAccount() *iso20022.SecuritiesAccount13 {
	p.SafekeepingAccount = new(iso20022.SecuritiesAccount13)
	return p.SafekeepingAccount
}

func (p *PortfolioTransferNotificationV01) AddTransferNotificationDetails() *iso20022.SecuritiesTradeDetails7 {
	newValue := new (iso20022.SecuritiesTradeDetails7)
	p.TransferNotificationDetails = append(p.TransferNotificationDetails, newValue)
	return newValue
}

func (p *PortfolioTransferNotificationV01) AddMessageOriginator() *iso20022.PartyIdentification10Choice {
	p.MessageOriginator = new(iso20022.PartyIdentification10Choice)
	return p.MessageOriginator
}

func (p *PortfolioTransferNotificationV01) AddMessageRecipient() *iso20022.PartyIdentification10Choice {
	p.MessageRecipient = new(iso20022.PartyIdentification10Choice)
	return p.MessageRecipient
}

