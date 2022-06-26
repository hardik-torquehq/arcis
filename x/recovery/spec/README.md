<!--
order: 0
title: "Recovery Overview"
parent:
  title: "recovery"
-->

# `recovery`

Recover tokens that are stuck on unsupported Arcis accounts.

## Abstract

This document specifies the  `x/recovery` module of the Arcis Hub.

The `x/recovery` module enables users on Arcis to recover locked funds that were transferred to accounts whose keys are not supported on Arcis. This happened in particular after the initial Arcis launch (`v1.1.2`), where users transferred tokens to a `secp256k1` Arcis address via IBC in order to [claim their airdrop](https://arcis.dev/modules/claims/). To be EVM compatible, [keys on Arcis](https://arcis.dev/technical_concepts/accounts.html#arcis-accounts) are generated using the `eth_secp256k1` key type which results in a different address derivation than e.g. the `secp256k1` key type used by other Cosmos chains.

At the time of Arcis’ relaunch, the value of locked tokens on unsupported accounts sits at $36,291.28 worth of OSMO and $268.86 worth of ATOM tokens according to the [Mintscan](https://www.mintscan.io/arcis/assets) block explorer. With the `x/recovery` module, users can recover these tokens back to their own addresses in the originating chains by performing IBC transfers from authorized IBC channels (i.e Osmosis for OSMO, Cosmos Hub for ATOM).

## Contents

1. **[Concepts](01_concepts.md)**
2. **[State](02_state.md)**
3. **[Hooks](03_hooks.md)**
4. **[Events](04_events.md)**
5. **[Parameters](05_parameters.md)**
6. **[Clients](06_clients.md)**
