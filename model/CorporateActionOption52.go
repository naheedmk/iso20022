package model

// Provides information about the corporate action option.
type CorporateActionOption52 struct {

	// Number identifying the available corporate action options.
	OptionNumber *Exact3NumericText `xml:"OptnNb"`

	// Specifies the corporate action options available to the account owner.
	OptionType *CorporateActionOption10Choice `xml:"OptnTp"`

	// Specifies how fractions resulting from derived securities will be processed or how prorated decisions will be rounding, if provided with a pro ration rate.
	FractionDisposition *FractionDispositionType19Choice `xml:"FrctnDspstn,omitempty"`

	// Specifies the conditions that apply to the offer.
	OfferType []*OfferTypeFormat3Choice `xml:"OfferTp,omitempty"`

	// Specifies the features that may apply to a corporate action option.
	OptionFeatures []*OptionFeaturesFormat9Choice `xml:"OptnFeatrs,omitempty"`

	// Specifies the type of intermediates securities distribution.
	IntermediateSecuritiesDistributionType *IntermediateSecuritiesDistributionTypeFormat9Choice `xml:"IntrmdtSctiesDstrbtnTp,omitempty"`

	// Specifies the status of the option.
	OptionAvailabilityStatus *OptionAvailabilityStatus1Choice `xml:"OptnAvlbtySts,omitempty"`

	// Indicates the type of certification/breakdown.
	CertificationBreakdownType []*BeneficiaryCertificationType5Choice `xml:"CertfctnBrkdwnTp,omitempty"`

	// Holder of the security has to certify, in line with the terms of the corporate action, that it is not domiciled in the country indicated.
	NonDomicileCountry []*CountryCode `xml:"NonDmclCtry,omitempty"`

	// Country of domicile in which the Corporate Action option is valid. The holder of the security has to certify that it is domiciled in the country indicated.
	ValidDomicileCountry []*CountryCode `xml:"VldDmclCtry,omitempty"`

	// Currency in which the cash disbursed from an interest or dividend payment is offered.
	CurrencyOption *ActiveCurrencyCode `xml:"CcyOptn,omitempty"`

	// Indicates whether the corporate action movement is a default processing or a standing instruction.
	DefaultProcessingOrStandingInstruction *DefaultProcessingOrStandingInstruction1Choice `xml:"DfltPrcgOrStgInstr"`

	// Indicates whether redemption charges apply.
	ChargesAppliedIndicator *YesNoIndicator `xml:"ChrgsApldInd,omitempty"`

	// Indicates whether or not certification/breakdown is required from the account owner.
	// Yes: certification required
	// No: no certification required
	CertificationBreakdownIndicator *YesNoIndicator `xml:"CertfctnBrkdwnInd,omitempty"`

	// Indicates whether withdrawal of instruction is allowed.
	WithdrawalAllowedIndicator *YesNoIndicator `xml:"WdrwlAllwdInd,omitempty"`

	// Indicates whether change of instruction is allowed.
	ChangeAllowedIndicator *YesNoIndicator `xml:"ChngAllwdInd,omitempty"`

	// Identifies the financial instrument.
	SecurityIdentification *SecurityIdentification14 `xml:"SctyId,omitempty"`

	// Provides information about the dates related to a corporate action option.
	DateDetails *CorporateActionDate15 `xml:"DtDtls,omitempty"`

	// Provides information about the periods related to a corporate action option.
	PeriodDetails *CorporateActionPeriod7 `xml:"PrdDtls,omitempty"`

	// Provides information about rates and amounts related to a corporate action option.
	RateAndAmountDetails *CorporateActionRate37 `xml:"RateAndAmtDtls,omitempty"`

	// Provides information about the prices related to a corporate action option.
	PriceDetails *CorporateActionPrice28 `xml:"PricDtls,omitempty"`

	// Provides information about securities quantity related to a corporate action option.
	SecuritiesQuantity *SecuritiesOption23 `xml:"SctiesQty,omitempty"`

	// Provides information about securities movement related to a corporate action option.
	SecuritiesMovementDetails []*SecuritiesOption38 `xml:"SctiesMvmntDtls,omitempty"`

	// Provides information about the cash movement linked to the corporate action option.
	CashMovementDetails []*CashOption25 `xml:"CshMvmntDtls,omitempty"`

	// Provides additional information about the corporate action movement.
	AdditionalInformation *CorporateActionNarrative20 `xml:"AddtlInf,omitempty"`
}

func (c *CorporateActionOption52) SetOptionNumber(value string) {
	c.OptionNumber = (*Exact3NumericText)(&value)
}

func (c *CorporateActionOption52) AddOptionType() *CorporateActionOption10Choice {
	c.OptionType = new(CorporateActionOption10Choice)
	return c.OptionType
}

func (c *CorporateActionOption52) AddFractionDisposition() *FractionDispositionType19Choice {
	c.FractionDisposition = new(FractionDispositionType19Choice)
	return c.FractionDisposition
}

func (c *CorporateActionOption52) AddOfferType() *OfferTypeFormat3Choice {
	newValue := new(OfferTypeFormat3Choice)
	c.OfferType = append(c.OfferType, newValue)
	return newValue
}

func (c *CorporateActionOption52) AddOptionFeatures() *OptionFeaturesFormat9Choice {
	newValue := new(OptionFeaturesFormat9Choice)
	c.OptionFeatures = append(c.OptionFeatures, newValue)
	return newValue
}

func (c *CorporateActionOption52) AddIntermediateSecuritiesDistributionType() *IntermediateSecuritiesDistributionTypeFormat9Choice {
	c.IntermediateSecuritiesDistributionType = new(IntermediateSecuritiesDistributionTypeFormat9Choice)
	return c.IntermediateSecuritiesDistributionType
}

func (c *CorporateActionOption52) AddOptionAvailabilityStatus() *OptionAvailabilityStatus1Choice {
	c.OptionAvailabilityStatus = new(OptionAvailabilityStatus1Choice)
	return c.OptionAvailabilityStatus
}

func (c *CorporateActionOption52) AddCertificationBreakdownType() *BeneficiaryCertificationType5Choice {
	newValue := new(BeneficiaryCertificationType5Choice)
	c.CertificationBreakdownType = append(c.CertificationBreakdownType, newValue)
	return newValue
}

func (c *CorporateActionOption52) AddNonDomicileCountry(value string) {
	c.NonDomicileCountry = append(c.NonDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption52) AddValidDomicileCountry(value string) {
	c.ValidDomicileCountry = append(c.ValidDomicileCountry, (*CountryCode)(&value))
}

func (c *CorporateActionOption52) SetCurrencyOption(value string) {
	c.CurrencyOption = (*ActiveCurrencyCode)(&value)
}

func (c *CorporateActionOption52) AddDefaultProcessingOrStandingInstruction() *DefaultProcessingOrStandingInstruction1Choice {
	c.DefaultProcessingOrStandingInstruction = new(DefaultProcessingOrStandingInstruction1Choice)
	return c.DefaultProcessingOrStandingInstruction
}

func (c *CorporateActionOption52) SetChargesAppliedIndicator(value string) {
	c.ChargesAppliedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption52) SetCertificationBreakdownIndicator(value string) {
	c.CertificationBreakdownIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption52) SetWithdrawalAllowedIndicator(value string) {
	c.WithdrawalAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption52) SetChangeAllowedIndicator(value string) {
	c.ChangeAllowedIndicator = (*YesNoIndicator)(&value)
}

func (c *CorporateActionOption52) AddSecurityIdentification() *SecurityIdentification14 {
	c.SecurityIdentification = new(SecurityIdentification14)
	return c.SecurityIdentification
}

func (c *CorporateActionOption52) AddDateDetails() *CorporateActionDate15 {
	c.DateDetails = new(CorporateActionDate15)
	return c.DateDetails
}

func (c *CorporateActionOption52) AddPeriodDetails() *CorporateActionPeriod7 {
	c.PeriodDetails = new(CorporateActionPeriod7)
	return c.PeriodDetails
}

func (c *CorporateActionOption52) AddRateAndAmountDetails() *CorporateActionRate37 {
	c.RateAndAmountDetails = new(CorporateActionRate37)
	return c.RateAndAmountDetails
}

func (c *CorporateActionOption52) AddPriceDetails() *CorporateActionPrice28 {
	c.PriceDetails = new(CorporateActionPrice28)
	return c.PriceDetails
}

func (c *CorporateActionOption52) AddSecuritiesQuantity() *SecuritiesOption23 {
	c.SecuritiesQuantity = new(SecuritiesOption23)
	return c.SecuritiesQuantity
}

func (c *CorporateActionOption52) AddSecuritiesMovementDetails() *SecuritiesOption38 {
	newValue := new(SecuritiesOption38)
	c.SecuritiesMovementDetails = append(c.SecuritiesMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption52) AddCashMovementDetails() *CashOption25 {
	newValue := new(CashOption25)
	c.CashMovementDetails = append(c.CashMovementDetails, newValue)
	return newValue
}

func (c *CorporateActionOption52) AddAdditionalInformation() *CorporateActionNarrative20 {
	c.AdditionalInformation = new(CorporateActionNarrative20)
	return c.AdditionalInformation
}
