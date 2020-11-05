# r3l

**r3l** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

This is a very early WIP. At the moment the app demonstrates the usage of the [Relevant Reputation Protocol](https://github.com/relevant-community/reputation) via a naive implementation of off-chain workers.

How it works:

- The app is pre-initialzed with some votes (pagerank links) and personalization node in the genesis state
- We query the current votes and other params in the r3l module's BeginBlock method
- We then begin a go routine that computes the reputation scores.
- Once results are returned (in some future block) we issue a TX to write them on-chain.

TODO: all of the obove should happen in a separate client process.
TODO: implement oracle reporting - results need to reach consensus before being written onchain

## Test the app

You will need to make sure you have cloned and built the Cosmos Relayer: https://github.com/cosmos/relayer

```
starport serve -v // so we can see when the TXs are issued
```

in a separate window, run

```
r3ld query r3l list-score
```

and you should see some reputation scores

## Get started

```
starport serve
```

`serve` command installs dependencies, initializes and runs the application.

## Configure

Initialization parameters of your app are stored in `config.yml`.

### `accounts`

A list of user accounts created during genesis of your application.

| Key   | Required | Type            | Description                                       |
| ----- | -------- | --------------- | ------------------------------------------------- |
| name  | Y        | String          | Local name of the key pair                        |
| coins | Y        | List of Strings | Initial coins with denominations (e.g. "100coin") |

## Learn more

- [Starport](https://github.com/tendermint/starport)
- [Cosmos SDK documentation](https://docs.cosmos.network)
- [Cosmos Tutorials](https://tutorials.cosmos.network)
- [Channel on Discord](https://discord.gg/W8trcGV)
