package iso20022

// Chain of parties involved in the settlement of a transaction, including receipts and deliveries, book transfers, treasury deals, or other activities, resulting in the movement of a security or amount of money from one account to another.
type ReceivingPartiesAndAccount14 struct {

	// Party that acts on behalf of the buyer of securities when the buyer does not have a direct relationship with the receiving agent.
	ReceiversCustodianDetails *PartyIdentificationAndAccount124 `xml:"RcvrsCtdnDtls,omitempty"`

	// Party that the receiver's custodian uses to effect the receipt of a security, when the receiver's custodian does not have a direct relationship with the receiving agent.
	ReceiversIntermediary1Details *PartyIdentificationAndAccount124 `xml:"RcvrsIntrmy1Dtls,omitempty"`

	// Party that interacts with the receiver’s intermediary.
	ReceiversIntermediary2Details *PartyIdentificationAndAccount124 `xml:"RcvrsIntrmy2Dtls,omitempty"`

	// Party that receives securities from the delivering agent via the place of settlement, for example, securities central depository.
	ReceivingAgentDetails *PartyIdentificationAndAccount123 `xml:"RcvgAgtDtls"`

	// Identifies the securities settlement system to be used.
	SecuritiesSettlementSystem *Max35Text `xml:"SctiesSttlmSys,omitempty"`

	// Place where settlement of the securities takes place.
	PlaceOfSettlementDetails *PartyIdentification97 `xml:"PlcOfSttlmDtls,omitempty"`
}

func (r *ReceivingPartiesAndAccount14) AddReceiversCustodianDetails() *PartyIdentificationAndAccount124 {
	r.ReceiversCustodianDetails = new(PartyIdentificationAndAccount124)
	return r.ReceiversCustodianDetails
}

func (r *ReceivingPartiesAndAccount14) AddReceiversIntermediary1Details() *PartyIdentificationAndAccount124 {
	r.ReceiversIntermediary1Details = new(PartyIdentificationAndAccount124)
	return r.ReceiversIntermediary1Details
}

func (r *ReceivingPartiesAndAccount14) AddReceiversIntermediary2Details() *PartyIdentificationAndAccount124 {
	r.ReceiversIntermediary2Details = new(PartyIdentificationAndAccount124)
	return r.ReceiversIntermediary2Details
}

func (r *ReceivingPartiesAndAccount14) AddReceivingAgentDetails() *PartyIdentificationAndAccount123 {
	r.ReceivingAgentDetails = new(PartyIdentificationAndAccount123)
	return r.ReceivingAgentDetails
}

func (r *ReceivingPartiesAndAccount14) SetSecuritiesSettlementSystem(value string) {
	r.SecuritiesSettlementSystem = (*Max35Text)(&value)
}

func (r *ReceivingPartiesAndAccount14) AddPlaceOfSettlementDetails() *PartyIdentification97 {
	r.PlaceOfSettlementDetails = new(PartyIdentification97)
	return r.PlaceOfSettlementDetails
}
