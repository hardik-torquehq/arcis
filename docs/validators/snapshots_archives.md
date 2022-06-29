<!--
order: 6
-->

# Snapshots & Archive Nodes

Quickly sync your node with Arcis using a snapshot or serve queries for prev versions using archive nodes {synopsis}

## List of Snapshots and Archives

Below is a list of publicly available snapshots that you can use to sync with the Arcis mainnet and
archived [9001-1 mainnet](https://github.com/Ambiplatforms-TORQUE/mainnet/tree/main/arcis_9001-1):

<!-- markdown-link-check-disable -->
:::: tabs
::: tab Snapshots

| Name        | URL                                                                     |
| -------------|------------------------------------------------------------------------ |
| `Staketab`   | [github.com/staketab/nginx-cosmos-snap](https://github.com/staketab/nginx-cosmos-snap/blob/main/docs/arcis.md) |
| `Polkachu`   | [polkachu.com](https://www.polkachu.com/tendermint_snapshots/arcis)                   |
| `Nodes Guru` | [snapshots.nodes.guru/arcis_9001-2/](snapshots.nodes.guru/arcis_9001-2/)                   |
:::
::: tab Archives
<!-- markdown-link-check-disable -->

| Name           | URL                                                                             |
| ---------------|---------------------------------------------------------------------------------|
| `Nodes Guru`   | [snapshots.nodes.guru/arcis_9001-1](https://snapshots.nodes.guru/arcis_9001-1/)                                    |
| `Polkachu`     | [polkachu.com/tendermint_snapshots/arcis](https://www.polkachu.com/tendermint_snapshots/arcis)                           |
| `Forbole`      | [bigdipper.live/arcis_9001-1](https://s3.bigdipper.live.eu-central-1.linodeobjects.com/arcis_9001-1.tar.lz4) |
:::
::::

To access snapshots and archives, follow the process below (this code snippet is to access a snapshot of the current network, `arcis_9001-2`, from Nodes Guru):

```bash
cd $HOME/.arcisd/data
wget https://snapshots.nodes.guru/arcis_9001-2/arcis_9001-2-410819.tar
tar xf arcis_9001-2-410819.tar
```
