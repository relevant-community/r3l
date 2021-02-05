# r3l

**r3l** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

This app demonstrates the usage of the [Relevant Reputation Protocol](https://github.com/relevant-community/reputation) and Cosmos off-chain workers.

How it works:

- The app is pre-initialzed with some votes (pagerank links) and personalization node in the genesis state
- Once the offchain worker is started, it queries the votes from the r3l chain, computes reputation scores, and submits the results as an 'oracle' claim
- The Oracle module stores the all claims and associated validator votes for the claims
- The every EndBlock, the r3l module queries Oracle for the pending rounds, tallies the votes and commits the results on-chain if the concensus threshold is reached

Oracle TODOs:
 - [ ] Add option to delegate oracle worker (so validators can use a different account)
 - [ ] Expose querying pendingVotes and RoundVotes via cli and rpc endpoints (currently only used internally)
 - [ ] Implement a commit-reveal pattern to prevent free-rider oracle validators
 - [ ] Documentation
 
 Example TODO:
 - [ ] Exchange price query example

## Run the r3l chain and the off-chain worker

You will need to make sure you have cloned and built the Cosmos Relayer: https://github.com/cosmos/relayer

```
starport serve
```

in a new terminal window, run

```
r3ld start-worker --from user1 --chain-id r3l --keyring-backend test
```

to check results, in a new terminal window, run

```
r3ld query r3l list-score
```

and you should see some reputation scores

----
starport docs:

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
