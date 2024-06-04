package caddy_dns_rfc2136

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/certmagic"
	"github.com/libdns/rfc2136"
)

type Provider struct {
	*rfc2136.Provider
}

func init() {
	caddy.RegisterModule(Provider{})
}

func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.rfc2136",
		New: func() caddy.Module { return &Provider{new(rfc2136.Provider)} },
	}
}

func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			var dst *string
			switch d.Val() {
			case "key_name":
				dst = &p.KeyName
			case "key":
				dst = &p.Key
			case "key_alg":
				dst = &p.KeyAlg
			case "server":
				dst = &p.Server
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}

			if d.NextArg() {
				*dst = d.Val()
			}
			if d.NextArg() {
				return d.ArgErr()
			}
		}
	}

	if p.Key == "" {
		return d.Err("rfc2136: missing key")
	}
	if p.KeyName == "" {
		return d.Err("rfc2136: missing key_name")
	}
	if p.KeyAlg == "" {
		return d.Err("rfc2136: missing key_alg")
	}
	if p.Server == "" {
		return d.Err("rfc2136: missing server")
	}

	return nil
}

var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ certmagic.DNSProvider = (*Provider)(nil)
)
