<!--
order: 8
-->

# Clients

A user can query the `x/incentives` module using the CLI, JSON-RPC, gRPC or
REST.

## CLI

Find below a list of `arcisd` commands added with the `x/inflation` module. You
can obtain the full list by using the `arcisd -h` command.

### Queries

The `query` commands allow users to query `inflation` state.

**`period`**

Allows users to query the current inflation period.

```go
arcisd query inflation period [flags]
```

**`epoch-mint-provision`**

Allows users to query the current inflation epoch provisions value.

```go
arcisd query inflation epoch-mint-provision [flags]
```

**`skipped-epochs`**

Allows users to query the current number of skipped epochs.

```go
arcisd query inflation skipped-epochs [flags]
```

**`total-supply`**

Allows users to query the total supply of tokens in circulation.

```go
arcisd query inflation total-supply [flags]
```

**`inflation-rate`**

Allows users to query the inflation rate of the current period.

```go
arcisd query inflation inflation-rate [flags]
```

**`params`**

Allows users to query the current inflation parameters.

```go
arcisd query inflation params [flags]
```

### Proposals

The `tx gov submit-proposal` commands allow users to query create a proposal
using the governance module CLI:

**`param-change`**

Allows users to submit a `ParameterChangeProposal`.

```bash
arcisd tx gov submit-proposal param-change [proposal-file] [flags]
```

## gRPC

### Queries

| Verb   | Method                                        | Description                                   |
| ------ | --------------------------------------------- | --------------------------------------------- |
| `gRPC` | `arcis.inflation.v1.Query/Period`             | Gets current inflation period                 |
| `gRPC` | `arcis.inflation.v1.Query/EpochMintProvision` | Gets current inflation epoch provisions value |
| `gRPC` | `arcis.inflation.v1.Query/Params`             | Gets current inflation parameters             |
| `gRPC` | `arcis.inflation.v1.Query/SkippedEpochs`      | Gets current number of skipped epochs         |
| `gRPC` | `arcis.inflation.v1.Query/TotalSupply`        | Gets current total supply                     |
| `gRPC` | `arcis.inflation.v1.Query/InflationRate`      | Gets current inflation rate                   |
| `GET`  | `/arcis/inflation/v1/period`                  | Gets current inflation period                 |
| `GET`  | `/arcis/inflation/v1/epoch_mint_provision`    | Gets current inflation epoch provisions value |
| `GET`  | `/arcis/inflation/v1/skipped_epochs`          | Gets current number of skipped epochs         |
| `GET`  | `/arcis/inflation/v1/total_supply`          | Gets current total supply                     |
| `GET`  | `/arcis/inflation/v1/inflation_rate`          | Gets current inflation rate                   |
| `GET`  | `/arcis/inflation/v1/params`                  | Gets current inflation parameters             |
