# traefik-plugin-localipaddr

[![Build Status](https://github.com/qiudesong/traefik-plugin-localipaddr/workflows/Main/badge.svg?branch=master)](https://github.com/tommoulard/traefik-plugin-shellexec/actions)

This is a plugin for [Traefik](https://traefik.io) which terminates a connection by outputing local address ( ipv4 / ipv6 ).

## Usage

### Configuration

For now, there is no configuration required.

Here is an example of a file provider dynamic configuration (given here in YAML), where the interesting part is the `http.middlewares` section:

```yaml
# Dynamic configuration

http:
  routers:
    my-router:
      rule: host(`demo.localhost`)
      service: service-foo
      entryPoints:
        - web
      middlewares:
        - traefik-plugin-localipaddr

  services:
   service-foo:
      loadBalancer:
        servers:
          - url: http://127.0.0.1:5000

  middlewares:
    traefik-plugin-localipaddr:
      plugin:
        traefik-plugin-localipaddr:
          enabled: true
```
