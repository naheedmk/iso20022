package model

// Specifies corporate action dates.
type CorporateActionDate17 struct {

	// Date on which the movement is due to take place (cash and/or securities).
	PaymentDate *DateFormat19Choice `xml:"PmtDt"`

	// Date at which assets become available to the account owner (in a credit entry), or cease to be available to the account owner (in a debit entry).
	ValueDate *DateFormat11Choice `xml:"ValDt,omitempty"`

	// Date/time at which a foreign exchange rate will be determined.
	ForeignExchangeRateFixingDate *DateFormat19Choice `xml:"FXRateFxgDt,omitempty"`

	// Date on which a payment can be made, for example, if payment date is a non-business day or to indicate the first payment date of an offer.
	EarliestPaymentDate *DateFormat19Choice `xml:"EarlstPmtDt,omitempty"`
}

func (c *CorporateActionDate17) AddPaymentDate() *DateFormat19Choice {
	c.PaymentDate = new(DateFormat19Choice)
	return c.PaymentDate
}

func (c *CorporateActionDate17) AddValueDate() *DateFormat11Choice {
	c.ValueDate = new(DateFormat11Choice)
	return c.ValueDate
}

func (c *CorporateActionDate17) AddForeignExchangeRateFixingDate() *DateFormat19Choice {
	c.ForeignExchangeRateFixingDate = new(DateFormat19Choice)
	return c.ForeignExchangeRateFixingDate
}

func (c *CorporateActionDate17) AddEarliestPaymentDate() *DateFormat19Choice {
	c.EarliestPaymentDate = new(DateFormat19Choice)
	return c.EarliestPaymentDate
}
