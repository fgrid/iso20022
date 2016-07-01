package iso20022

// Date of last revision.
type UpdatedDate struct {

	// Date of last revision.
	Date *ISODate `xml:"Dt"`

}


func (u *UpdatedDate) SetDate(value string) {
	u.Date = (*ISODate)(&value)
}

