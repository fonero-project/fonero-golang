---
title: Overview
---

Horizon is an API server for the Fonero ecosystem.  It acts as the interface between [fonero-core](https://github.com/fonero-project/fonero-core) and applications that want to access the Fonero network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Fonero ecosystem](https://www.fonero.org/developers/guides/) for details of where Horizon fits in. You can also watch a [talk on Horizon](https://www.youtube.com/watch?v=AtJ-f6Ih4A4) by Fonero.org developer Scott Fleckenstein:

[![Horizon: API webserver for the Fonero network](https://img.youtube.com/vi/AtJ-f6Ih4A4/sddefault.jpg "Horizon: API webserver for the Fonero network")](https://www.youtube.com/watch?v=AtJ-f6Ih4A4)

Horizon provides a RESTful API to allow client applications to interact with the Fonero network. You can communicate with Horizon using cURL or just your web browser. However, if you're building a client application, you'll likely want to use a Fonero SDK in the language of your client.
SDF provides a [JavaScript SDK](https://www.fonero.org/developers/js-fonero-sdk/learn/index.html) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net: [https://horizon-testnet.fonero.org/](https://horizon-testnet.fonero.org/) and one that is connected to the public Fonero network:
[https://horizon.fonero.org/](https://horizon.fonero.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/fonero-project/js-fonero-sdk)
- [Java](https://github.com/fonero-project/java-fonero-sdk)
- [Go](https://github.com/fonero-project/fonero-golang)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/fonero-project/ruby-fonero-sdk)
- [Python](https://github.com/FoneroCN/py-fonero-base)
- [C#](https://github.com/QuantozTechnology/csharp-fonero-base)
