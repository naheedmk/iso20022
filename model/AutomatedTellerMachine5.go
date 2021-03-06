package model

// ATM information.
type AutomatedTellerMachine5 struct {

	// ATM terminal device identification for the acquirer and the issuer.
	Identification *Max35Text `xml:"Id"`

	// ATM terminal device identification for the ATM manager.
	AdditionalIdentification *Max35Text `xml:"AddtlId,omitempty"`

	// ATM terminal device identification for the branch.
	SequenceNumber *Max35Text `xml:"SeqNb,omitempty"`

	// Reference currency of the ATM.
	BaseCurrency *ActiveCurrencyCode `xml:"BaseCcy"`

	// Location of the ATM.
	Location *PostalAddress17 `xml:"Lctn,omitempty"`

	// Indicates the environment of the transaction.
	LocationCategory *TransactionEnvironment2Code `xml:"LctnCtgy,omitempty"`

	// ATM terminal equipment.
	Equipment *ATMEquipment1 `xml:"Eqpmnt,omitempty"`

	// List of ATM devices out of service.
	OutOfServiceDevice []*ATMDevice2Code `xml:"OutOfSvcDvc,omitempty"`

	// Mechanism used to protect the message of the ATM protocol.
	MessageProtection *MessageProtection1Code `xml:"MsgPrtcn,omitempty"`
}

func (a *AutomatedTellerMachine5) SetIdentification(value string) {
	a.Identification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine5) SetAdditionalIdentification(value string) {
	a.AdditionalIdentification = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine5) SetSequenceNumber(value string) {
	a.SequenceNumber = (*Max35Text)(&value)
}

func (a *AutomatedTellerMachine5) SetBaseCurrency(value string) {
	a.BaseCurrency = (*ActiveCurrencyCode)(&value)
}

func (a *AutomatedTellerMachine5) AddLocation() *PostalAddress17 {
	a.Location = new(PostalAddress17)
	return a.Location
}

func (a *AutomatedTellerMachine5) SetLocationCategory(value string) {
	a.LocationCategory = (*TransactionEnvironment2Code)(&value)
}

func (a *AutomatedTellerMachine5) AddEquipment() *ATMEquipment1 {
	a.Equipment = new(ATMEquipment1)
	return a.Equipment
}

func (a *AutomatedTellerMachine5) AddOutOfServiceDevice(value string) {
	a.OutOfServiceDevice = append(a.OutOfServiceDevice, (*ATMDevice2Code)(&value))
}

func (a *AutomatedTellerMachine5) SetMessageProtection(value string) {
	a.MessageProtection = (*MessageProtection1Code)(&value)
}
