package domain

import (
	"testing"

	"github.com/go-gandi/go-gandi/internal/client"
	"github.com/go-gandi/go-gandi/internal/mocks"
)

func TestDomainCheck(t *testing.T) {
	mockDoer := mocks.NewMockDoer(t)
	c := client.New("apikey123", "", "", "", false, false, 0, mockDoer)
	d := NewFromClient(*c)
	_ = d
	panic("not implemented")
}
