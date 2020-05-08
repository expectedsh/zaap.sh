package deployer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/satori/go.uuid"
)

func (s *Server) createDnsEntry(ctx context.Context, applicationId uuid.UUID) error {
	application, err := s.applicationStore.FindWithRunner(ctx, applicationId)
	if err != nil {
		return err
	} else if application == nil {
		return ErrApplicationNotFound
	}

	for _, ip := range application.Runner.ExternalIps {
		_, err = s.cloudflareClient.CreateDNSRecord(s.config.CloudflareZoneId, cloudflare.DNSRecord{
			Type:    "A",
			Name:    application.DefaultDomain,
			Content: ip,
			TTL:     1,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) deleteDnsEntry(domain string) error {
	records, err := s.cloudflareClient.DNSRecords(s.config.CloudflareZoneId, cloudflare.DNSRecord{
		Name: domain,
	})
	if err != nil {
		return err
	}

	for _, record := range records {
		if err = s.cloudflareClient.DeleteDNSRecord(s.config.CloudflareZoneId, record.ID); err != nil {
			return err
		}
	}

	return nil
}
