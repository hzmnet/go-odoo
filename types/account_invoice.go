package types

import (
	"time"
)

type AccountInvoice struct {
	AccessToken                    string      `xmlrpc:"access_token"`
	AccountId                      Many2One    `xmlrpc:"account_id"`
	ActivityDateDeadline           time.Time   `xmlrpc:"activity_date_deadline"`
	ActivityIds                    []int64     `xmlrpc:"activity_ids"`
	ActivityState                  interface{} `xmlrpc:"activity_state"`
	ActivitySummary                string      `xmlrpc:"activity_summary"`
	ActivityTypeId                 Many2One    `xmlrpc:"activity_type_id"`
	ActivityUserId                 Many2One    `xmlrpc:"activity_user_id"`
	AmountTax                      float64     `xmlrpc:"amount_tax"`
	AmountTotal                    float64     `xmlrpc:"amount_total"`
	AmountTotalCompanySigned       float64     `xmlrpc:"amount_total_company_signed"`
	AmountTotalSigned              float64     `xmlrpc:"amount_total_signed"`
	AmountUntaxed                  float64     `xmlrpc:"amount_untaxed"`
	AmountUntaxedSigned            float64     `xmlrpc:"amount_untaxed_signed"`
	CampaignId                     Many2One    `xmlrpc:"campaign_id"`
	CashRoundingId                 Many2One    `xmlrpc:"cash_rounding_id"`
	Comment                        string      `xmlrpc:"comment"`
	CommercialPartnerId            Many2One    `xmlrpc:"commercial_partner_id"`
	CompanyCurrencyId              Many2One    `xmlrpc:"company_currency_id"`
	CompanyId                      Many2One    `xmlrpc:"company_id"`
	CreateDate                     time.Time   `xmlrpc:"create_date"`
	CreateUid                      Many2One    `xmlrpc:"create_uid"`
	CurrencyId                     Many2One    `xmlrpc:"currency_id"`
	Date                           time.Time   `xmlrpc:"date"`
	DateDue                        time.Time   `xmlrpc:"date_due"`
	DateInvoice                    time.Time   `xmlrpc:"date_invoice"`
	DisplayName                    string      `xmlrpc:"display_name"`
	FiscalPositionId               Many2One    `xmlrpc:"fiscal_position_id"`
	HasOutstanding                 bool        `xmlrpc:"has_outstanding"`
	Id                             int64       `xmlrpc:"id"`
	IncotermsId                    Many2One    `xmlrpc:"incoterms_id"`
	InvoiceLineIds                 []int64     `xmlrpc:"invoice_line_ids"`
	JournalId                      Many2One    `xmlrpc:"journal_id"`
	LastUpdate                     time.Time   `xmlrpc:"__last_update"`
	MediumId                       Many2One    `xmlrpc:"medium_id"`
	MessageChannelIds              []int64     `xmlrpc:"message_channel_ids"`
	MessageFollowerIds             []int64     `xmlrpc:"message_follower_ids"`
	MessageIds                     []int64     `xmlrpc:"message_ids"`
	MessageIsFollower              bool        `xmlrpc:"message_is_follower"`
	MessageLastPost                time.Time   `xmlrpc:"message_last_post"`
	MessageNeedaction              bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter       int64       `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds              []int64     `xmlrpc:"message_partner_ids"`
	MessageUnread                  bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter           int64       `xmlrpc:"message_unread_counter"`
	MoveId                         Many2One    `xmlrpc:"move_id"`
	MoveName                       string      `xmlrpc:"move_name"`
	Name                           string      `xmlrpc:"name"`
	Number                         string      `xmlrpc:"number"`
	Origin                         string      `xmlrpc:"origin"`
	OutstandingCreditsDebitsWidget string      `xmlrpc:"outstanding_credits_debits_widget"`
	PartnerBankId                  Many2One    `xmlrpc:"partner_bank_id"`
	PartnerId                      Many2One    `xmlrpc:"partner_id"`
	PartnerShippingId              Many2One    `xmlrpc:"partner_shipping_id"`
	PaymentIds                     []int64     `xmlrpc:"payment_ids"`
	PaymentMoveLineIds             []int64     `xmlrpc:"payment_move_line_ids"`
	PaymentsWidget                 string      `xmlrpc:"payments_widget"`
	PaymentTermId                  Many2One    `xmlrpc:"payment_term_id"`
	PortalUrl                      string      `xmlrpc:"portal_url"`
	PurchaseId                     Many2One    `xmlrpc:"purchase_id"`
	Reconciled                     bool        `xmlrpc:"reconciled"`
	Reference                      string      `xmlrpc:"reference"`
	ReferenceType                  interface{} `xmlrpc:"reference_type"`
	RefundInvoiceId                Many2One    `xmlrpc:"refund_invoice_id"`
	RefundInvoiceIds               []int64     `xmlrpc:"refund_invoice_ids"`
	Residual                       float64     `xmlrpc:"residual"`
	ResidualCompanySigned          float64     `xmlrpc:"residual_company_signed"`
	ResidualSigned                 float64     `xmlrpc:"residual_signed"`
	Sent                           bool        `xmlrpc:"sent"`
	SequenceNumberNext             string      `xmlrpc:"sequence_number_next"`
	SequenceNumberNextPrefix       string      `xmlrpc:"sequence_number_next_prefix"`
	SourceId                       Many2One    `xmlrpc:"source_id"`
	State                          interface{} `xmlrpc:"state"`
	TaxLineIds                     []int64     `xmlrpc:"tax_line_ids"`
	TeamId                         Many2One    `xmlrpc:"team_id"`
	TimesheetCount                 int64       `xmlrpc:"timesheet_count"`
	TimesheetIds                   []int64     `xmlrpc:"timesheet_ids"`
	Type                           interface{} `xmlrpc:"type"`
	UserId                         Many2One    `xmlrpc:"user_id"`
	WebsiteMessageIds              []int64     `xmlrpc:"website_message_ids"`
	WriteDate                      time.Time   `xmlrpc:"write_date"`
	WriteUid                       Many2One    `xmlrpc:"write_uid"`
}

type AccountInvoiceNil struct {
	AccessToken                    interface{} `xmlrpc:"access_token"`
	AccountId                      interface{} `xmlrpc:"account_id"`
	ActivityDateDeadline           interface{} `xmlrpc:"activity_date_deadline"`
	ActivityIds                    interface{} `xmlrpc:"activity_ids"`
	ActivityState                  interface{} `xmlrpc:"activity_state"`
	ActivitySummary                interface{} `xmlrpc:"activity_summary"`
	ActivityTypeId                 interface{} `xmlrpc:"activity_type_id"`
	ActivityUserId                 interface{} `xmlrpc:"activity_user_id"`
	AmountTax                      interface{} `xmlrpc:"amount_tax"`
	AmountTotal                    interface{} `xmlrpc:"amount_total"`
	AmountTotalCompanySigned       interface{} `xmlrpc:"amount_total_company_signed"`
	AmountTotalSigned              interface{} `xmlrpc:"amount_total_signed"`
	AmountUntaxed                  interface{} `xmlrpc:"amount_untaxed"`
	AmountUntaxedSigned            interface{} `xmlrpc:"amount_untaxed_signed"`
	CampaignId                     interface{} `xmlrpc:"campaign_id"`
	CashRoundingId                 interface{} `xmlrpc:"cash_rounding_id"`
	Comment                        interface{} `xmlrpc:"comment"`
	CommercialPartnerId            interface{} `xmlrpc:"commercial_partner_id"`
	CompanyCurrencyId              interface{} `xmlrpc:"company_currency_id"`
	CompanyId                      interface{} `xmlrpc:"company_id"`
	CreateDate                     interface{} `xmlrpc:"create_date"`
	CreateUid                      interface{} `xmlrpc:"create_uid"`
	CurrencyId                     interface{} `xmlrpc:"currency_id"`
	Date                           interface{} `xmlrpc:"date"`
	DateDue                        interface{} `xmlrpc:"date_due"`
	DateInvoice                    interface{} `xmlrpc:"date_invoice"`
	DisplayName                    interface{} `xmlrpc:"display_name"`
	FiscalPositionId               interface{} `xmlrpc:"fiscal_position_id"`
	HasOutstanding                 bool        `xmlrpc:"has_outstanding"`
	Id                             interface{} `xmlrpc:"id"`
	IncotermsId                    interface{} `xmlrpc:"incoterms_id"`
	InvoiceLineIds                 interface{} `xmlrpc:"invoice_line_ids"`
	JournalId                      interface{} `xmlrpc:"journal_id"`
	LastUpdate                     interface{} `xmlrpc:"__last_update"`
	MediumId                       interface{} `xmlrpc:"medium_id"`
	MessageChannelIds              interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds             interface{} `xmlrpc:"message_follower_ids"`
	MessageIds                     interface{} `xmlrpc:"message_ids"`
	MessageIsFollower              bool        `xmlrpc:"message_is_follower"`
	MessageLastPost                interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction              bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter       interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds              interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread                  bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter           interface{} `xmlrpc:"message_unread_counter"`
	MoveId                         interface{} `xmlrpc:"move_id"`
	MoveName                       interface{} `xmlrpc:"move_name"`
	Name                           interface{} `xmlrpc:"name"`
	Number                         interface{} `xmlrpc:"number"`
	Origin                         interface{} `xmlrpc:"origin"`
	OutstandingCreditsDebitsWidget interface{} `xmlrpc:"outstanding_credits_debits_widget"`
	PartnerBankId                  interface{} `xmlrpc:"partner_bank_id"`
	PartnerId                      interface{} `xmlrpc:"partner_id"`
	PartnerShippingId              interface{} `xmlrpc:"partner_shipping_id"`
	PaymentIds                     interface{} `xmlrpc:"payment_ids"`
	PaymentMoveLineIds             interface{} `xmlrpc:"payment_move_line_ids"`
	PaymentsWidget                 interface{} `xmlrpc:"payments_widget"`
	PaymentTermId                  interface{} `xmlrpc:"payment_term_id"`
	PortalUrl                      interface{} `xmlrpc:"portal_url"`
	PurchaseId                     interface{} `xmlrpc:"purchase_id"`
	Reconciled                     bool        `xmlrpc:"reconciled"`
	Reference                      interface{} `xmlrpc:"reference"`
	ReferenceType                  interface{} `xmlrpc:"reference_type"`
	RefundInvoiceId                interface{} `xmlrpc:"refund_invoice_id"`
	RefundInvoiceIds               interface{} `xmlrpc:"refund_invoice_ids"`
	Residual                       interface{} `xmlrpc:"residual"`
	ResidualCompanySigned          interface{} `xmlrpc:"residual_company_signed"`
	ResidualSigned                 interface{} `xmlrpc:"residual_signed"`
	Sent                           bool        `xmlrpc:"sent"`
	SequenceNumberNext             interface{} `xmlrpc:"sequence_number_next"`
	SequenceNumberNextPrefix       interface{} `xmlrpc:"sequence_number_next_prefix"`
	SourceId                       interface{} `xmlrpc:"source_id"`
	State                          interface{} `xmlrpc:"state"`
	TaxLineIds                     interface{} `xmlrpc:"tax_line_ids"`
	TeamId                         interface{} `xmlrpc:"team_id"`
	TimesheetCount                 interface{} `xmlrpc:"timesheet_count"`
	TimesheetIds                   interface{} `xmlrpc:"timesheet_ids"`
	Type                           interface{} `xmlrpc:"type"`
	UserId                         interface{} `xmlrpc:"user_id"`
	WebsiteMessageIds              interface{} `xmlrpc:"website_message_ids"`
	WriteDate                      interface{} `xmlrpc:"write_date"`
	WriteUid                       interface{} `xmlrpc:"write_uid"`
}

var AccountInvoiceModel string = "account.invoice"

type AccountInvoices []AccountInvoice

type AccountInvoicesNil []AccountInvoiceNil

func (s *AccountInvoice) NilableType_() interface{} {
	return &AccountInvoiceNil{}
}

func (ns *AccountInvoiceNil) Type_() interface{} {
	s := &AccountInvoice{}
	return load(ns, s)
}

func (s *AccountInvoices) NilableType_() interface{} {
	return &AccountInvoicesNil{}
}

func (ns *AccountInvoicesNil) Type_() interface{} {
	s := &AccountInvoices{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*AccountInvoice))
	}
	return s
}
