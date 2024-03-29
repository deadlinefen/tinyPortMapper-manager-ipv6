package ddns

import (
	"github.com/deadlinefen/tinyPortMapper-manager-ipv6/pkg/config"
	"github.com/miekg/dns"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type IResoluter interface {
	ResoluteOneIp() (string, error)
}

type Resoluter struct {
	resolution *config.Resolution
	dnsIndex   int
}

func (c *Resoluter) nextDnsIndex() {
	c.dnsIndex = (c.dnsIndex + 1) % len(c.resolution.Dns)
}

func (c *Resoluter) GetDns() string {
	defer c.nextDnsIndex()
	return c.resolution.Dns[c.dnsIndex]
}

func (c *Resoluter) ResoluteOneIp() (string, error) {
	dnsServer := c.GetDns()
	log.Debugf("resolute one ip from %s...", dnsServer)

	client := new(dns.Client)
	msg := new(dns.Msg)
	msg.SetQuestion(c.resolution.Domain+".", dns.TypeAAAA)

	r, _, err := client.Exchange(msg, dnsServer+":53")
	if err != nil {
		return "", err
	}
	if r.Rcode != dns.RcodeSuccess {
		return "", errors.Errorf("DNS lookup failed with code: %d", r.Rcode)
	}
	if len(r.Answer) < 1 {
		return "", errors.Errorf("DNS answer size is %d", len(r.Answer))
	}

	ip, ok := r.Answer[0].(*dns.AAAA)
	if !ok {
		return "", errors.Errorf("DNS answer ip is not ipv6")
	}

	return ip.AAAA.String(), nil
}

type IResoluterFactory interface {
	NewResoluter() IResoluter
}

type ResoluterFactory struct {
	resolution *config.Resolution
}

func (cf *ResoluterFactory) NewResoluter() IResoluter {
	return &Resoluter{resolution: cf.resolution, dnsIndex: 0}
}

func NewResoluterFactory(resolution *config.Resolution) IResoluterFactory {
	return &ResoluterFactory{resolution: resolution}
}
