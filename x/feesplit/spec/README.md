<!--
order: 0
title: "Fees Overview"
parent:
  title: "fees"
-->

# `feesplit`

## Abstract

This document specifies the internal `x/feesplit` module of the Arcis Hub.

The `x/feesplit` module enables the Arcis Hub to support splitting transaction fees between block proposer and smart contract deployers. As a part of the [Arcis Token Model](https://arcis.blog/the-arcis-token-model-edc07014978b), this mechanism aims to increase the adoption of the Arcis Hub by offering a new stable source of income for smart contract deployers. Developers can register their smart contracts and everytime someone interacts with a registered smart contract, the contract deployer or their assigned withdrawal account receives a part of the transaction fees.

Together, all registered smart contracts make up the Arcis dApp Store: paying developers and network operators for their services via built-in shared fee revenue model.

## Contents

1. **[Concepts](01_concepts.md)**
2. **[State](02_state.md)**
3. **[State Transitions](03_state_transitions.md)**
4. **[Transactions](04_transactions.md)**
5. **[Hooks](05_hooks.md)**
6. **[Events](06_events.md)**
7. **[Parameters](07_parameters.md)**
8. **[Clients](08_clients.md)**
9. **[Future Improvements](09_improvements.md)**
