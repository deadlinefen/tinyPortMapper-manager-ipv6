package main

import (
	"flag"
	"fmt"

	"github.com/deadlinefen/tinyPortMapper-manager-ipv6/pkg/config"
	"github.com/deadlinefen/tinyPortMapper-manager-ipv6/pkg/ddns"
	"github.com/pkg/errors"
)

var domain string

func init() {
	flag.StringVar(&domain, "d", "tencent.com", "resolute domain")

	flag.Parse()
}

func main() {
	r := &config.Resolution{
		Domain: domain,
		Dns:    []string{"8.8.8.8"},
		Ttl:    1,
	}
	fmt.Printf("resolution: %+v\n", r)
	checker := ddns.NewResoluterFactory(r).NewResoluter()

	ip, err := checker.ResoluteOneIp()
	if err != nil {
		panic(errors.Wrapf(err, "resolute failed"))
	}

	fmt.Printf("ipv6: %s\n", ip)
}
