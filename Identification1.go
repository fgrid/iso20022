package iso20022

// Unique identifier of a document, message or transaction.
type Identification1 struct {

	// Unique identifier of a document, message or transaction.
	Identification *Max35Text `xml:"Id"`
}

func (i *Identification1) SetIdentification(value string) {
	i.Identification = (*Max35Text)(&value)
}
