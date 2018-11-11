---
title: Horizon
---

Horizon is the server for the client facing API for the Fonero ecosystem.  It acts as the interface between [fonero-core](https://www.fonero.org/developers/learn/fonero-core) and applications that want to access the Fonero network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the Fonero ecosystem](https://www.fonero.org/developers/guides/) for more details.

You can interact directly with horizon via curl or a web browser but SDF provides a [JavaScript SDK](https://www.fonero.org/developers/js-fonero-sdk/learn/) for clients to use to interact with Horizon.

SDF runs a instance of Horizon that is connected to the test net [https://horizon-testnet.fonero.org/](https://horizon-testnet.fonero.org/).

## Libraries

SDF maintained libraries:<br />
- [JavaScript](https://github.com/fonero-project/js-fonero-sdk)
- [Java](https://github.com/fonero-project/java-fonero-sdk)
- [Go](https://github.com/fonero-project/fonero-golang)

Community maintained libraries (in various states of completeness) for interacting with Horizon in other languages:<br>
- [Ruby](https://github.com/fonero-project/ruby-fonero-sdk)
- [Python](https://github.com/FoneroCN/py-fonero-base)
- [C# .NET 2.0](https://github.com/QuantozTechnology/csharp-fonero-base)
- [C# .NET Core 2.x](https://github.com/elucidsoft/dotnetcore-fonero-sdk)
- [C++](https://bitbucket.org/bnogal/foneroqore/wiki/Home)
