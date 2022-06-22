RFC2136 DNS provider module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records using RFC2136 Dynamic Updates.

## Caddy module name

```
dns.providers.rfc2136
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
    "module": "acme",
    "challenges": {
        "dns": {
            "provider": {
                "name": "rfc2136",
                "key": "cWnu6Ju9zOki4f7Q+da2KKGo0KOXbCf6Pej6hW3geC4=",
                "key_name": "test",
                "key_alg": "hmac-sha256",
                "server": "1.2.3.4:53"
            }
        }
    }
}
```

or with the Caddyfile:

```
# globally
{
    acme_dns rfc2136 {
        key_name "test"
        key_alg "hmac-sha256"
        key "cWnu6Ju9zOki4f7Q+da2KKGo0KOXbCf6Pej6hW3geC4="
        server "1.2.3.4:53"
    }
}
```

```
# one site
tls {
    dns rfc2136 {
        key_name "test"
        key_alg "hmac-sha256"
        key "cWnu6Ju9zOki4f7Q+da2KKGo0KOXbCf6Pej6hW3geC4="
        server "1.2.3.4:53"
    }
}
```
