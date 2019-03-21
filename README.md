# Etherus [![Build Status](https://travis-ci.org/Etherus-org/etherus.svg?branch=dukei-develop-f1.8.6)](https://travis-ci.org/Etherus-org/etherus)

Ethereum powered by Tendermint consensus with dynamic validators

* You can try demo at https://etherus.org
* Blockchain explorer: https://explorer.etherus.org
* Wallet: https://wallet.etherus.org
* RPC endpoint: https://rpc-alpha.etherus.org
* Tendermint P2P persistent peer: 26593a301734740cdf503ea8f4f8b218446360fc@127.0.0.1:26656
* Alpha genesis.json: https://gist.github.com/dukei/1e8889b33430495cfbb9818ab8d09761

## Features

Etherus is fully compatible with the standard Ethereum tooling such as [geth](https://github.com/ethereum/go-ethereum), [mist](https://github.com/ethereum/mist), [truffle](https://github.com/trufflesuite/truffle) and [remix](http://remix.ethereum.org). Please
install whichever tooling suits you best.

## Installation

Get binary package for your OS. https://github.com/Etherus-org/etherus/releases/tag/v1.8.6-alpha

Run Etherus with Tendermint. On the first run Etherus will create a wallet for you and Tendermint will show you your validator public key. If you want your node to be a validator you should register it on Validators smart contract. 

Go to https://wallet.etherus.org/#contracts and choose Validators contract from list.
Execute addDeposit method. Specify your validator public key as vPub parameter and your Etherus address as nodeAddr. Send 2500 ETR with this call (otherwise the call will fail). If you don't have ETR you can request them at info@etherus.org.

Validators will became effective after the next block. You will be able to see how your node proposes blocks and gets rewarded at https://explorer.etherus.org

## Building from source
For Linux and MacOS:
```
export GOPATH=$(pwd)
export PATH=$GOPATH/bin:$PATH
```
For Windows (cmd):
```
SET "GOPATH=%CD%"
SET "PATH=%GOPATH%/bin;%PATH%"
```
And for any platform afterall:
```
git clone -b dukei-develop-f1.8.6 "https://github.com/etherus-org/etherus.git" "src/github.com/ya-enot/etherus"
git clone -b dukei-develop "https://github.com/etherus-org/tendermint.git" "src/github.com/tendermint/tendermint"
pushd .
cd src/github.com/ya-enot/etherus
make check install
cd ../../tendermint/tendermint
make check install
popd
```
### Running Etherus

#### Initialisation
To get started, you need to initialise the genesis block for tendermint core and go-ethereum. We provide initialisation
files with reasonable defaults and money allocated into a predefined account. If you installed from binary or docker
please download these default files [here](https://github.com/tendermint/ethermint/tree/develop/setup).

You can choose where to store the ethermint files with `--datadir`. For this guide, we will use `~/.ethermint`, which is a reasonable default in most cases.

Before you can run ethermint you need to initialise tendermint and ethermint with their respective genesis states.
Please switch into the folder where you have the initialisation files. If you installed from source you can just follow
these instructions.

```bash
etherus --datadir ~/.etherus --with-tendermint init
```

which will also invoke `tendermint init --home ~/.etherus/tendermint`. You can prevent Tendermint from
being starting by excluding the flag `--with-tendermint` for example:

```bash
etherus --datadir ~/.etherus init
```

and then you will have to invoke `tendermint` in another shell with the command:

```bash
tendermint init --home ~/.etherus/tendermint
```

For simplicity, we'll have ethermint start tendermint as a subprocess with the
flag `--with-tendermint`:

```bash
etherus --with-tendermint --datadir ~/.etherus --rpc --rpcaddr=0.0.0.0 --ws --wsaddr=0.0.0.0 --rpcapi eth,net,web3,personal,admin
```

Note: The **password** for the default account is *1234*.

There you have it, Ethereum on Tendermint!

## Contributing

Thank you for considering making contributions to Ethermint!

Check out the [contributing guidelines](.github/CONTRIBUTING.md) for information
on getting starting with contributing.

See the [open issues](https://github.com/tendermint/ethermint/issues) for
things we need help with!

## Support

Check out the [community page](https://tendermint.com/community) for various resources.

## License

[GPLv3](LICENSE)

Before you can run ethermint you need to initialise tendermint and ethermint with their respective genesis states.
Please switch into the folder where you have the initialisation files. If you installed from source you can just follow
these instructions.

```bash
ethermint --datadir ~/.ethermint --with-tendermint init
```

which will also invoke `tendermint init --home ~/.ethermint/tendermint`. You can prevent Tendermint from
being starting by excluding the flag `--with-tendermint` for example:

```bash
ethermint --datadir ~/.ethermint init
```

and then you will have to invoke `tendermint` in another shell with the command:

```bash
tendermint init --home ~/.ethermint/tendermint
```

For simplicity, we'll have ethermint start tendermint as a subprocess with the
flag `--with-tendermint`:

```bash
ethermint --with-tendermint --datadir ~/.ethermint --rpc --rpcaddr=0.0.0.0 --ws --wsaddr=0.0.0.0 --rpcapi eth,net,web3,personal,admin
```

Note: The **password** for the default account is *1234*.

There you have it, Ethereum on Tendermint! For details on what to do next,
check out [the documentation](http://ethermint.readthedocs.io/en/master/)

## Contributing

Thank you for considering making contributions to Ethermint!

Check out the [contributing guidelines](.github/CONTRIBUTING.md) for information
on getting starting with contributing.

See the [open issues](https://github.com/ya-enot/etherus/issues) for
things we need help with!

## Support

Check out the [community page](https://tendermint.com/community) for various resources.

## License

[GPLv3](LICENSE)
