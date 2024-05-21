//go:generate mockery
package gandi

import (
	"github.com/lcd1232/go-gandi/certificate"
	"github.com/lcd1232/go-gandi/config"
	"github.com/lcd1232/go-gandi/domain"
	"github.com/lcd1232/go-gandi/email"
	"github.com/lcd1232/go-gandi/livedns"
	"github.com/lcd1232/go-gandi/simplehosting"
)

// NewDomainClient returns a client to the Gandi Domains API
// It expects an API key, available from https://account.gandi.net/en/
func NewDomainClient(config config.Config) *domain.Domain {
	return domain.New(config)
}

// NewEmailClient returns a client to the Gandi Email API
// It expects an API key, available from https://account.gandi.net/en/
func NewEmailClient(config config.Config) *email.Email {
	return email.New(config)
}

// NewLiveDNSClient returns a client to the Gandi Domains API
// It expects an API key, available from https://account.gandi.net/en/
func NewLiveDNSClient(config config.Config) *livedns.LiveDNS {
	return livedns.New(config)
}

// NewSimpleHostingClient returns a client to the Gandi Simple Hosting API
// It expects an API key, available from https://account.gandi.net/en/
func NewSimpleHostingClient(config config.Config) *simplehosting.SimpleHosting {
	return simplehosting.New(config)
}

// NewCertificateClient returns a client to the Gandi Certificate API
// It expects an API key, available from https://account.gandi.net/en/
func NewCertificateClient(config config.Config) *certificate.Certificate {
	return certificate.New(config)
}
