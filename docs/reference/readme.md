---
title: Overview
---

The Go SDK contains packages for interacting with most aspects of the fonero ecosystem.  In addition to generally useful, low-level packages such as [`keypair`](https://godoc.org/github.com/fonero-project/fonero-golang/keypair) (used for creating fonero-compliant public/secret key pairs), the Go SDK also contains code for the server applications and client tools written in go.

## Godoc reference

The most accurate and up-to-date reference information on the Go SDK is found within godoc.  The godoc.org service automatically updates the documentation for the Go SDK everytime github is updated.  The godoc for all of our packages can be found at (https://godoc.org/github.com/fonero-project/fonero-golang).

## Client Packages

The Go SDK contains packages for interacting with the various fonero services:

- [`horizon`](https://godoc.org/github.com/fonero-project/fonero-golang/clients/horizon) provides client access to a horizon server, allowing you to load account information, stream payments, post transactions and more.
- [`fonerotoml`](https://godoc.org/github.com/fonero-project/fonero-golang/clients/fonerotoml) provides the ability to resolve Fonero.toml files from the internet.  You can read about [Fonero.toml concepts here](../../guides/concepts/fonero-toml.md).
- [`federation`](https://godoc.org/github.com/fonero-project/fonero-golang/clients/federation) makes it easy to resolve a fonero addresses (e.g. `scott*fonero.org`) into a fonero account ID suitable for use within a transaction.

