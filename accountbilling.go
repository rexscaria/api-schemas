// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cfrex

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/cf-rex-go/internal/apijson"
	"github.com/stainless-sdks/cf-rex-go/internal/requestconfig"
	"github.com/stainless-sdks/cf-rex-go/option"
)

// AccountBillingService contains methods and other services that help with
// interacting with the cf-rex API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountBillingService] method instead.
type AccountBillingService struct {
	Options []option.RequestOption
}

// NewAccountBillingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountBillingService(opts ...option.RequestOption) (r *AccountBillingService) {
	r = &AccountBillingService{}
	r.Options = opts
	return
}

// Gets the current billing profile for the account.
func (r *AccountBillingService) GetProfile(ctx context.Context, accountID string, opts ...option.RequestOption) (res *BillingResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/billing/profile", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BillingResponseSingle struct {
	Result BillingResponseSingleResult `json:"result"`
	JSON   billingResponseSingleJSON   `json:"-"`
	APIResponseSingleBilling
}

// billingResponseSingleJSON contains the JSON metadata for the struct
// [BillingResponseSingle]
type billingResponseSingleJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingResponseSingleJSON) RawJSON() string {
	return r.raw
}

type BillingResponseSingleResult struct {
	// Billing item identifier tag.
	ID                     string                          `json:"id"`
	AccountType            string                          `json:"account_type"`
	Address                string                          `json:"address"`
	Address2               string                          `json:"address2"`
	Balance                string                          `json:"balance"`
	CardExpiryMonth        int64                           `json:"card_expiry_month"`
	CardExpiryYear         int64                           `json:"card_expiry_year"`
	CardNumber             string                          `json:"card_number"`
	City                   string                          `json:"city"`
	Company                string                          `json:"company"`
	Country                string                          `json:"country"`
	CreatedOn              time.Time                       `json:"created_on" format:"date-time"`
	DeviceData             string                          `json:"device_data"`
	EditedOn               time.Time                       `json:"edited_on" format:"date-time"`
	EnterpriseBillingEmail string                          `json:"enterprise_billing_email"`
	EnterprisePrimaryEmail string                          `json:"enterprise_primary_email"`
	FirstName              string                          `json:"first_name"`
	IsPartner              bool                            `json:"is_partner"`
	LastName               string                          `json:"last_name"`
	NextBillDate           time.Time                       `json:"next_bill_date" format:"date-time"`
	PaymentAddress         string                          `json:"payment_address"`
	PaymentAddress2        string                          `json:"payment_address2"`
	PaymentCity            string                          `json:"payment_city"`
	PaymentCountry         string                          `json:"payment_country"`
	PaymentEmail           string                          `json:"payment_email"`
	PaymentFirstName       string                          `json:"payment_first_name"`
	PaymentGateway         string                          `json:"payment_gateway"`
	PaymentLastName        string                          `json:"payment_last_name"`
	PaymentNonce           string                          `json:"payment_nonce"`
	PaymentState           string                          `json:"payment_state"`
	PaymentZipcode         string                          `json:"payment_zipcode"`
	PrimaryEmail           string                          `json:"primary_email"`
	State                  string                          `json:"state"`
	TaxIDType              string                          `json:"tax_id_type"`
	Telephone              string                          `json:"telephone"`
	UseLegacy              bool                            `json:"use_legacy"`
	ValidationCode         string                          `json:"validation_code"`
	Vat                    string                          `json:"vat"`
	Zipcode                string                          `json:"zipcode"`
	JSON                   billingResponseSingleResultJSON `json:"-"`
}

// billingResponseSingleResultJSON contains the JSON metadata for the struct
// [BillingResponseSingleResult]
type billingResponseSingleResultJSON struct {
	ID                     apijson.Field
	AccountType            apijson.Field
	Address                apijson.Field
	Address2               apijson.Field
	Balance                apijson.Field
	CardExpiryMonth        apijson.Field
	CardExpiryYear         apijson.Field
	CardNumber             apijson.Field
	City                   apijson.Field
	Company                apijson.Field
	Country                apijson.Field
	CreatedOn              apijson.Field
	DeviceData             apijson.Field
	EditedOn               apijson.Field
	EnterpriseBillingEmail apijson.Field
	EnterprisePrimaryEmail apijson.Field
	FirstName              apijson.Field
	IsPartner              apijson.Field
	LastName               apijson.Field
	NextBillDate           apijson.Field
	PaymentAddress         apijson.Field
	PaymentAddress2        apijson.Field
	PaymentCity            apijson.Field
	PaymentCountry         apijson.Field
	PaymentEmail           apijson.Field
	PaymentFirstName       apijson.Field
	PaymentGateway         apijson.Field
	PaymentLastName        apijson.Field
	PaymentNonce           apijson.Field
	PaymentState           apijson.Field
	PaymentZipcode         apijson.Field
	PrimaryEmail           apijson.Field
	State                  apijson.Field
	TaxIDType              apijson.Field
	Telephone              apijson.Field
	UseLegacy              apijson.Field
	ValidationCode         apijson.Field
	Vat                    apijson.Field
	Zipcode                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BillingResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingResponseSingleResultJSON) RawJSON() string {
	return r.raw
}
