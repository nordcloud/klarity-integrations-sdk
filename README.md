# Klarity Integrations SDK Clients

This repository contains SDK for common languages to use with Klarity Integrations API
SDK Clients are generated automatically by the [OpenAPI Generator](https://openapi-generator.tech) project
every time the new version of API is being released.


## Golang
[go/integrations](./go/integrations) directory contains SDK for the GO programming language that allows to
connect with Klarity Integrations API.

you can import it with:

```bash
go get github.com/nordcloud/klarity-integrations-sdk/go/integrations
```

## Python
[python/integrations](./python/integrations) directory contains SDK for the Python programming language that allows to connect with Klarity Integrations API

you can import it using:

- pip:

```bash
pip install git+https://github.com/nordcloud/klarity-integrations-sdk.git#subdirectory=python/integrations
```

- pipenv

```bash
pipenv install "git+https://github.com/nordcloud/klarity-integrations-sdk.git#subdirectory=python/integrations&egg=integrations"
```
