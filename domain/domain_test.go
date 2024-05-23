package domain

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/lcd1232/go-gandi/internal/client"
	"github.com/lcd1232/go-gandi/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func readFile(t *testing.T, path string) []byte {
	t.Helper()
	b, err := os.ReadFile("testdata/" + path)
	require.NoError(t, err)
	return b
}

func toPointerTime(t time.Time) *time.Time {
	return &t
}

func toBool(b bool) *bool {
	return &b
}

func TestDomainCheck(t *testing.T) {
	mockDoer := mocks.NewMockDoer(t)
	mockDoer.EXPECT().Do(mock.Anything).RunAndReturn(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "https://api.gandi.net/v5/domain/check?country=US&currency=USD&name=example.com", r.URL.String())
		return &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(readFile(t, "domain_check_success.json"))),
		}, nil
	})
	c := client.New("apikey123", "", "", "", false, false, 0, mockDoer)
	d := NewFromClient(*c)
	periodEndsAt := time.Date(2020, 3, 28, 15, 59, 59, 0, time.UTC)
	want := DomainAvailability{
		Currency: "EUR",
		Grid:     "A",
		Products: []Product{
			{
				Status: "available",
				Periods: []Period{
					{
						Name:     "eap5",
						StartsAt: time.Date(2019, 2, 25, 16, 0, 0, 0, time.UTC),
						EndsAt:   &periodEndsAt,
					},
					{
						Name:     "golive",
						StartsAt: time.Date(2020, 3, 28, 16, 0, 0, 0, time.UTC),
					},
				},
				Name:    "example.com",
				Process: "create",
				Taxes: []Tax{
					{
						Type: "service",
						Rate: 20,
						Name: "vat",
					},
				},
				Prices: []Price{
					{
						MaxDuration:      1,
						DurationUnit:     "y",
						MinDuration:      1,
						Discount:         false,
						PriceAfterTaxes:  878.44,
						PriceBeforeTaxes: 732.03,
						Type:             "premium",
					},
					{
						MaxDuration:      1,
						DurationUnit:     "y",
						MinDuration:      1,
						Discount:         false,
						PriceAfterTaxes:  775.12,
						PriceBeforeTaxes: 645.93,
						Type:             "premium",
					},
				},
			},
		},
		Taxes: []Tax{
			{
				Type: "service",
				Rate: 20,
				Name: "vat",
			},
		},
	}
	got, err := d.GetDomainAvailability(DomainAvailabilityRequest{
		Domain:   "example.com",
		Country:  "US",
		Currency: "USD",
	})
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestDomainCheckWithSharingID(t *testing.T) {
	mockDoer := mocks.NewMockDoer(t)
	mockDoer.EXPECT().Do(mock.Anything).RunAndReturn(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "https://api.gandi.net/v5/domain/check?country=US&currency=USD&name=example.com&sharing_id=123", r.URL.String())
		return &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(readFile(t, "domain_check_success.json"))),
		}, nil
	})
	c := client.New("apikey123", "", "", "123", false, false, 0, mockDoer)
	d := NewFromClient(*c)
	periodEndsAt := time.Date(2020, 3, 28, 15, 59, 59, 0, time.UTC)
	want := DomainAvailability{
		Currency: "EUR",
		Grid:     "A",
		Products: []Product{
			{
				Status: "available",
				Periods: []Period{
					{
						Name:     "eap5",
						StartsAt: time.Date(2019, 2, 25, 16, 0, 0, 0, time.UTC),
						EndsAt:   &periodEndsAt,
					},
					{
						Name:     "golive",
						StartsAt: time.Date(2020, 3, 28, 16, 0, 0, 0, time.UTC),
					},
				},
				Name:    "example.com",
				Process: "create",
				Taxes: []Tax{
					{
						Type: "service",
						Rate: 20,
						Name: "vat",
					},
				},
				Prices: []Price{
					{
						MaxDuration:      1,
						DurationUnit:     "y",
						MinDuration:      1,
						Discount:         false,
						PriceAfterTaxes:  878.44,
						PriceBeforeTaxes: 732.03,
						Type:             "premium",
					},
					{
						MaxDuration:      1,
						DurationUnit:     "y",
						MinDuration:      1,
						Discount:         false,
						PriceAfterTaxes:  775.12,
						PriceBeforeTaxes: 645.93,
						Type:             "premium",
					},
				},
			},
		},
		Taxes: []Tax{
			{
				Type: "service",
				Rate: 20,
				Name: "vat",
			},
		},
	}
	got, err := d.GetDomainAvailability(DomainAvailabilityRequest{
		Domain:   "example.com",
		Country:  "US",
		Currency: "USD",
	})
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestDomainGetDomain(t *testing.T) {
	mockDoer := mocks.NewMockDoer(t)
	mockDoer.EXPECT().Do(mock.Anything).RunAndReturn(func(r *http.Request) (*http.Response, error) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "https://api.gandi.net/v5/domain/domains/example.com", r.URL.String())
		return &http.Response{
			Status:     http.StatusText(http.StatusOK),
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(bytes.NewReader(readFile(t, "domain_details_success.json"))),
		}, nil
	})
	c := client.New("apikey123", "", "", "", false, false, 0, mockDoer)
	d := NewFromClient(*c)
	want := Details{
		AutoRenew: &AutoRenew{
			Duration: 1,
			Dates: []*time.Time{
				toPointerTime(time.Date(2026, 3, 15, 21, 30, 47, 0, time.UTC)),
				toPointerTime(time.Date(2026, 3, 31, 22, 30, 47, 0, time.UTC)),
				toPointerTime(time.Date(2026, 4, 14, 22, 30, 47, 0, time.UTC)),
			},
			Href: "https://api.gandi.net/v5/domain/example.com/auto-renew",
		},
		CanTLDLock: toBool(true),
		Contacts: &Contacts{
			Owner: &Contact{
				OrgName:        "Company Inc",
				FamilyName:     "Doe",
				GivenName:      "John",
				ContactType:    "company",
				StreetAddr:     "123 Main St",
				Zip:            "90001",
				City:           "Los Angeles",
				State:          "US-CA",
				Country:        "US",
				Email:          "support@example.com",
				MailObfuscated: toBool(true),
				Phone:          "+1.2135551212",
				DataObfuscated: toBool(true),
				Validation:     "none",
			},
			Admin: &Contact{
				FamilyName:     "Doe",
				GivenName:      "John",
				ContactType:    "person",
				StreetAddr:     "123 Main St",
				Zip:            "90001",
				City:           "Los Angeles",
				State:          "US-CA",
				Country:        "US",
				Email:          "support@example.com",
				MailObfuscated: toBool(true),
				Phone:          "+1.2135551212",
				DataObfuscated: toBool(true),
				Validation:     "none",
			},
			Tech: &Contact{
				FamilyName:     "Doe",
				GivenName:      "John",
				ContactType:    "person",
				StreetAddr:     "123 Main St",
				Zip:            "90001",
				City:           "Los Angeles",
				State:          "US-CA",
				Country:        "US",
				Email:          "support@example.com",
				MailObfuscated: toBool(true),
				Phone:          "+1.2135551212",
				DataObfuscated: toBool(true),
				Validation:     "none",
			},
			Billing: &Contact{
				FamilyName:     "Doe",
				GivenName:      "John",
				ContactType:    "person",
				StreetAddr:     "123 Main St",
				Zip:            "90001",
				City:           "Los Angeles",
				State:          "US-CA",
				Country:        "US",
				Email:          "support@example.com",
				MailObfuscated: toBool(true),
				Phone:          "+1.2135551212",
				DataObfuscated: toBool(true),
				Validation:     "none",
			},
		},
		Dates: &ResponseDates{
			CreatedAt:           toPointerTime(time.Date(2018, 10, 7, 12, 1, 12, 0, time.UTC)),
			DeletesAt:           toPointerTime(time.Date(2026, 5, 30, 12, 30, 47, 0, time.UTC)),
			HoldBeginsAt:        toPointerTime(time.Date(2026, 4, 15, 22, 30, 47, 0, time.UTC)),
			HoldEndsAt:          toPointerTime(time.Date(2026, 5, 30, 22, 30, 47, 0, time.UTC)),
			PendingDeleteEndsAt: toPointerTime(time.Date(2026, 7, 14, 12, 30, 47, 0, time.UTC)),
			RegistryCreatedAt:   toPointerTime(time.Date(2018, 4, 15, 22, 30, 47, 0, time.UTC)),
			RegistryEndsAt:      toPointerTime(time.Date(2026, 4, 15, 22, 30, 47, 0, time.UTC)),
			RenewBeginsAt:       toPointerTime(time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)),
			UpdatedAt:           toPointerTime(time.Date(2024, 5, 22, 12, 32, 11, 0, time.UTC)),
			AuthInfoExpiresAt:   toPointerTime(time.Date(2026, 5, 22, 12, 32, 11, 0, time.UTC)),
		},
		FQDN:        "example.com",
		FQDNUnicode: "example.com",
		Href:        "https://api.gandi.net/v5/domain/domains/example.com",
		Nameservers: []string{"ns1.gandi.net", "ns2.gandi.net", "ns3.gandi.net"},
		Services:    []string{"gandilivedns"},
		SharingSpace: &SharingSpace{
			ID:   "d828bdcb-934a-4d1b-ae1d-d663b948e51a",
			Name: "Company",
		},
		Status:   []string{"clientTransferProhibited"},
		TLD:      "com",
		AuthInfo: "zjzxhgjrsdf!asd",
		ID:       "f0996c41-12d1-458b-964f-04b045b45e2d",
	}
	got, err := d.GetDomain("example.com")
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestGetContactType(t *testing.T) {
	for _, tc := range []struct {
		name  string
		value string
		want  string
	}{
		{
			name:  "individual as a string",
			value: "individual",
			want:  "individual",
		},
		{
			name:  "individual as a number",
			value: "0",
			want:  "individual",
		},
		{
			name:  "company as a string",
			value: "company",
			want:  "company",
		},
		{
			name:  "company as a number",
			value: "1",
			want:  "company",
		},
		{
			name:  "association as a string",
			value: "association",
			want:  "association",
		},
		{
			name:  "association as a number",
			value: "2",
			want:  "association",
		},
		{
			name:  "public body as a string",
			value: "publicbody",
			want:  "publicbody",
		},
		{
			name:  "public body as a number",
			value: "3",
			want:  "publicbody",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c := contact{
				ContactType: json.Number(tc.value),
			}
			assert.Equal(t, tc.want, c.GetContactType())
		})
	}
}
