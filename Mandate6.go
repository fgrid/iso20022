package iso20022

// Information that serves as a basis to debit an account.
type Mandate6 struct {

	// Unique identification, as assigned by the creditor, to unambiguously identify the mandate.
	MandateIdentification *Max35Text `xml:"MndtId"`

	// Identification for the mandate request, as assigned by the initiating party.
	MandateRequestIdentification *Max35Text `xml:"MndtReqId,omitempty"`

	// Specifies the type of mandate, such as paper, electronic or scheme.
	Type *MandateTypeInformation1 `xml:"Tp,omitempty"`

	// Provides details of the duration of the mandate and occurrence of the underlying transactions.
	Occurrences *MandateOccurrences3 `xml:"Ocrncs,omitempty"`

	// Fixed amount to be collected from the debtor's account.
	CollectionAmount *ActiveCurrencyAndAmount `xml:"ColltnAmt,omitempty"`

	// Maximum amount that may be collected from the debtor's account, per instruction.
	MaximumAmount *ActiveCurrencyAndAmount `xml:"MaxAmt,omitempty"`

	// Provides the reason for the setup of the mandate.
	Reason *MandateSetupReason1Choice `xml:"Rsn,omitempty"`

	// Credit party that signs the mandate.
	CreditorSchemeIdentification *PartyIdentification43 `xml:"CdtrSchmeId,omitempty"`

	// Party that signs the mandate and to whom an amount of money is due.
	Creditor *PartyIdentification43 `xml:"Cdtr,omitempty"`

	// Unambiguous identification of the account of the creditor to which a credit entry will be posted as a result of the payment transaction.
	CreditorAccount *CashAccount24 `xml:"CdtrAcct,omitempty"`

	// Financial institution servicing an account for the creditor.
	CreditorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"CdtrAgt,omitempty"`

	// Ultimate party to which an amount of money is due.
	UltimateCreditor *PartyIdentification43 `xml:"UltmtCdtr,omitempty"`

	// Party that signs the mandate and owes an amount of money to the (ultimate) creditor.
	Debtor *PartyIdentification43 `xml:"Dbtr,omitempty"`

	// Unambiguous identification of the account of the debtor, to which a debit entry will be made as a result of the transaction.
	DebtorAccount *CashAccount24 `xml:"DbtrAcct,omitempty"`

	// Financial institution servicing an account for the debtor.
	DebtorAgent *BranchAndFinancialInstitutionIdentification5 `xml:"DbtrAgt,omitempty"`

	// Ultimate party that owes an amount of money to the (ultimate) creditor.
	UltimateDebtor *PartyIdentification43 `xml:"UltmtDbtr,omitempty"`

	// Provides information to identify the underlying documents associated with the mandate.
	ReferredDocument []*ReferredDocumentInformation6 `xml:"RfrdDoc,omitempty"`
}

func (m *Mandate6) SetMandateIdentification(value string) {
	m.MandateIdentification = (*Max35Text)(&value)
}

func (m *Mandate6) SetMandateRequestIdentification(value string) {
	m.MandateRequestIdentification = (*Max35Text)(&value)
}

func (m *Mandate6) AddType() *MandateTypeInformation1 {
	m.Type = new(MandateTypeInformation1)
	return m.Type
}

func (m *Mandate6) AddOccurrences() *MandateOccurrences3 {
	m.Occurrences = new(MandateOccurrences3)
	return m.Occurrences
}

func (m *Mandate6) SetCollectionAmount(value, currency string) {
	m.CollectionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate6) SetMaximumAmount(value, currency string) {
	m.MaximumAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (m *Mandate6) AddReason() *MandateSetupReason1Choice {
	m.Reason = new(MandateSetupReason1Choice)
	return m.Reason
}

func (m *Mandate6) AddCreditorSchemeIdentification() *PartyIdentification43 {
	m.CreditorSchemeIdentification = new(PartyIdentification43)
	return m.CreditorSchemeIdentification
}

func (m *Mandate6) AddCreditor() *PartyIdentification43 {
	m.Creditor = new(PartyIdentification43)
	return m.Creditor
}

func (m *Mandate6) AddCreditorAccount() *CashAccount24 {
	m.CreditorAccount = new(CashAccount24)
	return m.CreditorAccount
}

func (m *Mandate6) AddCreditorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.CreditorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.CreditorAgent
}

func (m *Mandate6) AddUltimateCreditor() *PartyIdentification43 {
	m.UltimateCreditor = new(PartyIdentification43)
	return m.UltimateCreditor
}

func (m *Mandate6) AddDebtor() *PartyIdentification43 {
	m.Debtor = new(PartyIdentification43)
	return m.Debtor
}

func (m *Mandate6) AddDebtorAccount() *CashAccount24 {
	m.DebtorAccount = new(CashAccount24)
	return m.DebtorAccount
}

func (m *Mandate6) AddDebtorAgent() *BranchAndFinancialInstitutionIdentification5 {
	m.DebtorAgent = new(BranchAndFinancialInstitutionIdentification5)
	return m.DebtorAgent
}

func (m *Mandate6) AddUltimateDebtor() *PartyIdentification43 {
	m.UltimateDebtor = new(PartyIdentification43)
	return m.UltimateDebtor
}

func (m *Mandate6) AddReferredDocument() *ReferredDocumentInformation6 {
	newValue := new(ReferredDocumentInformation6)
	m.ReferredDocument = append(m.ReferredDocument, newValue)
	return newValue
}
