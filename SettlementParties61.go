package iso20022

// Specifies settlement parties (delivering/receiving).
type SettlementParties61 struct {

	// First receiving party in the settlement chain. In a plain vanilla settlement, it is the central securities depository where the receiving side of the transaction requests to receive the financial instrument.
	Depository *PartyIdentification120 `xml:"Dpstry,omitempty"`

	// Party that interacts with the depository.
	Party1 *PartyIdentificationAndAccount128 `xml:"Pty1,omitempty"`

	// Party that interacts with the party1.
	Party2 *PartyIdentificationAndAccount128 `xml:"Pty2,omitempty"`

	// Party that interacts with the party2.
	Party3 *PartyIdentificationAndAccount128 `xml:"Pty3,omitempty"`
}

func (s *SettlementParties61) AddDepository() *PartyIdentification120 {
	s.Depository = new(PartyIdentification120)
	return s.Depository
}

func (s *SettlementParties61) AddParty1() *PartyIdentificationAndAccount128 {
	s.Party1 = new(PartyIdentificationAndAccount128)
	return s.Party1
}

func (s *SettlementParties61) AddParty2() *PartyIdentificationAndAccount128 {
	s.Party2 = new(PartyIdentificationAndAccount128)
	return s.Party2
}

func (s *SettlementParties61) AddParty3() *PartyIdentificationAndAccount128 {
	s.Party3 = new(PartyIdentificationAndAccount128)
	return s.Party3
}
