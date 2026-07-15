# ELA Worker Service

[![JavaScript](https://img.shields.io/badge/JavaScript-ES2023-F7DF1E?style=flat&logo=javascript&logoColor=black)](https://developer.mozilla.org/en-US/docs/Web/JavaScript)
[![Cloudflare Workers](https://img.shields.io/badge/Cloudflare-Workers-F38020?style=flat&logo=cloudflare&logoColor=white)](https://workers.cloudflare.com/)
[![License: GPL v2](https://img.shields.io/badge/License-GPL%20v2-blue.svg)](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)

> A fast Linux launcher built with JavaScript and Cloudflare Workers, designed for managing the state, instances, and runtime of multiple Linux distributions.

---

## Overview

ELA Worker Service is a lightweight API service that provides remote logic support for ELA.exe.

- Designed for users with low memory who cannot run the full ELA desktop version
- Core logic runs on Cloudflare Workers; local client only handles rendering
- All users share the same Worker instance

---

## Architecture

ELA.exe (Windows) -> hardcoded Worker URL -> shared Worker -> HL sequences / registry data -> ELA.exe executes locally

Worker only sends instructions. Local machine does all the real work.

---

## Deployment

Only needed if you are maintaining the Worker. End users do not need to deploy.

npm install -g wrangler
wrangler login
wrangler deploy

---

## API Endpoints

GET /api/instances
Returns list of supported Linux distributions.

GET /api/registry
Returns list of registered instances.

POST /api/registry
Creates a new instance.

DELETE /api/registry/:id
Deletes an instance.

POST /api/hl/parse
Parses an HL sequence.

POST /api/ls/exec
Executes an LS script.

GET /api/version
Returns ELA version information.

---

## Configuration

Worker URL is hardcoded in ELA.exe. End users do not need to configure anything.

Open ELA.exe and it will automatically connect to the official Worker.

No setup required.

---

## Limitations

Cloudflare Workers free tier limit:
- 100,000 requests per day
- 100 Worker scripts per account
- 128 MB memory per Worker

If the limit is exceeded, the Worker may return HTTP 404 or 503.

All users share the same Worker to avoid hitting the limit.

---

## Performance

This version is slower than the full desktop version due to network latency.

Memory usage is very low, suitable for low-end devices.

Internet connection is required. Offline mode is not supported.

---

## Maintenance Status

This Worker version may be discontinued in the future.

It is provided as a transitional solution for users with low memory.

If the Worker stops, old ELA.exe versions will lose network functionality.

Users are encouraged to use the full desktop version for long-term use.

---

## Funding and Updates

ELA updates depend on maintainer time and resources.

There is no automatic update system at this time.

Users must manually download new versions from GitHub.

If the project receives funding, update speed may increase.

---

## Source Code

https://github.com/allow-make/Efficient-Linux-Launch-ELA-

Fork and self-host if needed.

---

## License

GPL-2.0
