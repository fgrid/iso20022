package iso20022

// Unique identifier of a document, message or transaction.
type Identification16 struct {

	// Unique identifier of a document, message or transaction.
	Identification *RestrictedFINXMax16Text `xml:"Id"`
}

func (i *Identification16) SetIdentification(value string) {
	i.Identification = (*RestrictedFINXMax16Text)(&value)
}
