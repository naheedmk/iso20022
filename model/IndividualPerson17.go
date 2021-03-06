package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson17 struct {

	// Name received at birth, eg, maiden name.
	BirthName *Max35Text `xml:"BirthNm"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Unique and unambiguous identification of a person, eg, passport.
	Identification *PersonIdentification6 `xml:"Id,omitempty"`

	// Postal address of a party.
	Address *LongPostalAddress2Choice `xml:"Adr,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingParty *PartyIdentification9Choice `xml:"EmplngPty,omitempty"`

	// Specifies details related to the attendance card.
	AttendanceCardDetails *AttendanceCard2 `xml:"AttndncCardDtls"`
}

func (i *IndividualPerson17) SetBirthName(value string) {
	i.BirthName = (*Max35Text)(&value)
}

func (i *IndividualPerson17) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson17) AddIdentification() *PersonIdentification6 {
	i.Identification = new(PersonIdentification6)
	return i.Identification
}

func (i *IndividualPerson17) AddAddress() *LongPostalAddress2Choice {
	i.Address = new(LongPostalAddress2Choice)
	return i.Address
}

func (i *IndividualPerson17) AddEmployingParty() *PartyIdentification9Choice {
	i.EmployingParty = new(PartyIdentification9Choice)
	return i.EmployingParty
}

func (i *IndividualPerson17) AddAttendanceCardDetails() *AttendanceCard2 {
	i.AttendanceCardDetails = new(AttendanceCard2)
	return i.AttendanceCardDetails
}
