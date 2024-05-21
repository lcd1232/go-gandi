package domain

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-gandi/go-gandi/internal/client"
	"github.com/go-gandi/go-gandi/internal/mocks"
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
						Options: PricingOptions{
							Period: "golive",
						},
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
