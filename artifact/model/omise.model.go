package model

import "encoding/json"

func UnmarshalOmiseCharge(data []byte) (OmiseCharge, error) {
	var r OmiseCharge
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *OmiseCharge) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type OmiseCharge struct {
	Object                   string          `json:"object"`
	ID                       string          `json:"id"`
	Location                 string          `json:"location"`
	Amount                   int64           `json:"amount"`
	AcquirerReferenceNumber  interface{}     `json:"acquirer_reference_number"`
	Net                      int64           `json:"net"`
	Fee                      int64           `json:"fee"`
	FeeVat                   int64           `json:"fee_vat"`
	Interest                 int64           `json:"interest"`
	InterestVat              int64           `json:"interest_vat"`
	FundingAmount            int64           `json:"funding_amount"`
	RefundedAmount           int64           `json:"refunded_amount"`
	TransactionFees          TransactionFees `json:"transaction_fees"`
	PlatformFee              PlatformFee     `json:"platform_fee"`
	Currency                 string          `json:"currency"`
	FundingCurrency          string          `json:"funding_currency"`
	IP                       interface{}     `json:"ip"`
	Refunds                  Refunds         `json:"refunds"`
	Link                     interface{}     `json:"link"`
	Description              interface{}     `json:"description"`
	Metadata                 Metadata        `json:"metadata"`
	Card                     interface{}     `json:"card"`
	Source                   Source          `json:"source"`
	Schedule                 interface{}     `json:"schedule"`
	Customer                 interface{}     `json:"customer"`
	Dispute                  interface{}     `json:"dispute"`
	Transaction              interface{}     `json:"transaction"`
	FailureCode              interface{}     `json:"failure_code"`
	FailureMessage           interface{}     `json:"failure_message"`
	Status                   string          `json:"status"`
	AuthorizeURI             interface{}     `json:"authorize_uri"`
	ReturnURI                interface{}     `json:"return_uri"`
	CreatedAt                string          `json:"created_at"`
	PaidAt                   interface{}     `json:"paid_at"`
	ExpiresAt                string          `json:"expires_at"`
	ExpiredAt                interface{}     `json:"expired_at"`
	ReversedAt               interface{}     `json:"reversed_at"`
	ZeroInterestInstallments bool            `json:"zero_interest_installments"`
	Branch                   interface{}     `json:"branch"`
	Terminal                 interface{}     `json:"terminal"`
	Device                   interface{}     `json:"device"`
	Authorized               bool            `json:"authorized"`
	Capturable               bool            `json:"capturable"`
	Capture                  bool            `json:"capture"`
	Disputable               bool            `json:"disputable"`
	Livemode                 bool            `json:"livemode"`
	Refundable               bool            `json:"refundable"`
	PartiallyRefundable      bool            `json:"partially_refundable"`
	Reversed                 bool            `json:"reversed"`
	Reversible               bool            `json:"reversible"`
	Voided                   bool            `json:"voided"`
	Paid                     bool            `json:"paid"`
	Expired                  bool            `json:"expired"`
	ApprovalCode             interface{}     `json:"approval_code"`
}

type Metadata struct {
}

type PlatformFee struct {
	Fixed      interface{} `json:"fixed"`
	Amount     interface{} `json:"amount"`
	Percentage interface{} `json:"percentage"`
}

type Refunds struct {
	Object   string        `json:"object"`
	Data     []interface{} `json:"data"`
	Limit    int64         `json:"limit"`
	Offset   int64         `json:"offset"`
	Total    int64         `json:"total"`
	Location string        `json:"location"`
	Order    string        `json:"order"`
	From     string        `json:"from"`
	To       string        `json:"to"`
}

type Source struct {
	Object                   string             `json:"object"`
	ID                       string             `json:"id"`
	Livemode                 bool               `json:"livemode"`
	Location                 string             `json:"location"`
	Amount                   int64              `json:"amount"`
	Barcode                  interface{}        `json:"barcode"`
	Bank                     interface{}        `json:"bank"`
	CreatedAt                string             `json:"created_at"`
	Currency                 string             `json:"currency"`
	Email                    interface{}        `json:"email"`
	Flow                     string             `json:"flow"`
	InstallmentTerm          interface{}        `json:"installment_term"`
	IP                       interface{}        `json:"ip"`
	AbsorptionType           interface{}        `json:"absorption_type"`
	Name                     interface{}        `json:"name"`
	MobileNumber             interface{}        `json:"mobile_number"`
	PhoneNumber              interface{}        `json:"phone_number"`
	PlatformType             interface{}        `json:"platform_type"`
	ScannableCode            ScannableCode      `json:"scannable_code"`
	Billing                  interface{}        `json:"billing"`
	Shipping                 interface{}        `json:"shipping"`
	Items                    []interface{}      `json:"items"`
	References               interface{}        `json:"references"`
	ProviderReferences       ProviderReferences `json:"provider_references"`
	StoreID                  interface{}        `json:"store_id"`
	StoreName                interface{}        `json:"store_name"`
	TerminalID               interface{}        `json:"terminal_id"`
	Type                     string             `json:"type"`
	ZeroInterestInstallments interface{}        `json:"zero_interest_installments"`
	ChargeStatus             string             `json:"charge_status"`
	ReceiptAmount            interface{}        `json:"receipt_amount"`
	Discounts                []interface{}      `json:"discounts"`
}

type ProviderReferences struct {
	ReferenceNumber1 string      `json:"reference_number_1"`
	ReferenceNumber2 interface{} `json:"reference_number_2"`
}

type ScannableCode struct {
	Object string `json:"object"`
	Type   string `json:"type"`
	Image  Image  `json:"image"`
}

type Image struct {
	Object      string `json:"object"`
	Livemode    bool   `json:"livemode"`
	ID          string `json:"id"`
	Deleted     bool   `json:"deleted"`
	Filename    string `json:"filename"`
	Location    string `json:"location"`
	Kind        string `json:"kind"`
	DownloadURI string `json:"download_uri"`
	CreatedAt   string `json:"created_at"`
}

type TransactionFees struct {
	FeeFlat string `json:"fee_flat"`
	FeeRate string `json:"fee_rate"`
	VatRate string `json:"vat_rate"`
}
