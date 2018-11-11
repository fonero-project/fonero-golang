# Fonero Sign

This folder contains `fonero-sign` a simple utility to make it easy to add your signature to a transaction envelope.  When run on the terminal it:

1.  Prompts your for a base64-encoded envelope:
2.  Asks for your private seed.
3.  Outputs a new envelope with your signature added.

## Installing

```bash
$ go get -u github.com/fonero-project/fonero-golang/tools/fonero-sign
```

## Running

```bash
$ fonero-sign
```
