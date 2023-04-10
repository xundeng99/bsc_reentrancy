## ECE1770 Final project

This repo is the final project of ECE1770 Winter2023. The goal of the final project is to develop a system that can detect reentrancy attacks.

Most of the added code has the comment **// Reentrancy detection**

## Building the source

Many of the below are the same as or similar to go-ethereum. Consult main branch README for more information. 

For prerequisites and detailed build instructions please read the [Installation Instructions](https://geth.ethereum.org/docs/getting-started/installing-geth).

Building `geth` requires both a Go (version 1.18 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make geth
```

or, to build the full suite of utilities:

```shell
make all
```

                                                                                                                    
### Steps to Run
To run the project, you need to start a full node of the bsc chain.

#### 1. Rebase the code onto latest release

The latest release can be found here(https://github.com/bnb-chain/bsc/releases). And build geth client.

```shell
make geth
```

#### 2. Download the config files

Download the config file of the latest release.

```shell
wget $(curl -s https://api.github.com/repos/bnb-chain/bsc/releases/latest |grep browser_ |grep mainnet |cut -d\" -f4)
unzip mainnet.zip
```

#### 3. Download snapshot
Follow the instructions here(https://github.com/48Club/bsc-snapshots) to download the snapshot. 

Move the uncompressed geth folder to your run directory.

Alternative snapshot can be downloaded from here(https://github.com/bnb-chain/bsc-snapshots). 

#### 4. Start a full node
```shell
## Add flags suggested in the snapshot github 
./geth --config ./config.toml --datadir ./node  --cache 8000 --rpc.allow-unprotected-txs --txlookuplimit 0 --tries-verify-mode none
```

#### 5. Monitor node status

Monitor the log from **./run/bsc.log** by default where **./run** is your run directory. When the node has started syncing, should be able to see the following output:
```shell
t=2022-09-08T13:00:27+0000 lvl=info msg="Imported new chain segment"             blocks=1    txs=177   mgas=17.317   elapsed=31.131ms    mgasps=556.259  number=21,153,429 hash=0x42e6b54ba7106387f0650defc62c9ace3160b427702dab7bd1c5abb83a32d8db dirty="0.00 B"
t=2022-09-08T13:00:29+0000 lvl=info msg="Imported new chain segment"             blocks=1    txs=251   mgas=39.638   elapsed=68.827ms    mgasps=575.900  number=21,153,430 hash=0xa3397b273b31b013e43487689782f20c03f47525b4cd4107c1715af45a88796e dirty="0.00 B"
t=2022-09-08T13:00:33+0000 lvl=info msg="Imported new chain segment"             blocks=1    txs=197   mgas=19.364   elapsed=34.663ms    mgasps=558.632  number=21,153,431 hash=0x0c7872b698f28cb5c36a8a3e1e315b1d31bda6109b15467a9735a12380e2ad14 dirty="0.00 B"
```

#### 6. Reentrancy Detection
Detection results will be printed in the terminal. To enable/disable the reentrancy detection, change **core/vm/evm.go**

```shell
func NewEVM(...){
  evm.detect = true
}
```
You should see:
```shell
 reentrancies detected in block 0 26947415 110
Address 0x76e8bfF910e693dd05CcEb9A7C10FE3793643359  loses  2  in 0.0001 USD unit
Address 0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF  gains  2  in 0.0001 USD unit
Address 0x0f338Ec12d3f7C3D77A4B9fcC1f95F3FB6AD0EA6  gains  601199  in 0.0001 USD unit
Address 0xEDD346853333F65D57b97726c87a42be14DEe68a  gains  33  in 0.0001 USD unit
Address 0xFFfFfFffFFfffFFfFFfFFFFFffFFFffffFfFFFfF  loses  601232  in 0.0001 USD unit
Reentrancy detected: ,  0x7720b74120142544babbb6e12b78c0652ed334ac0619e6e1574c7ff416843cdc
 reentracy detected
 reentrancies detected in block 1 26947416 110
```
#### 7. Notes
1. The detection will only start after the chain has started syncing. You may need to wait 1-3min to see outputs in terminal.
2. If cannot pass a certain transaction, disable detection, restart syncing for a while. And then restart detection.
3. Comment/uncomment print statements to display more/less info. 
