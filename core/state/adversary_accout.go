package state

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

)

type TransferDirOnly int
type PotentialIdentity int

const (
	From TransferDirOnly = iota //The "Address" is the receiver
	To                          //The "Address" is the sender
)

const (
	NotDecided PotentialIdentity = iota
	Victim
	Neither
	Beneficiary
)

type TransferDir struct {
	tdo  TransferDirOnly
	addr common.Address
}
type TransferAmount struct {
	td  TransferDir
	amt common.Hash
}
type IndividualAdversaryAccountHelper struct {
	id  PotentialIdentity
	amt common.Hash
}
type TransferInfo struct {
	addr1 common.Address
	addr2 common.Address
	amt   common.Hash
	token common.Address
}
type AccountTransferInfoPerToken struct {
	token     common.Address
	earn_flag bool
	amt       common.Hash
}
type AccountTransferInfo struct {
	acct      common.Address
	tokenflow []AccountTransferInfoPerToken
}
type AdversaryAccount struct {
	// account balance trace
	balance_traces map[common.Address]map[common.Address][]TransferAmount
	// store all transfers in order
	transfer_in_order []TransferInfo
	//Track all address flash loan potential attack
	account_information map[common.Address]IndividualAdversaryAccountHelper
	// potential flash loan transaction
	old_tx *types.Message
	//old_tx contract address. It is set when init if old_tx is Call. Set after executing if old_tx is Create
	old_tx_contract_address *common.Address

	//temp contract addresses created in this transcation
	temp_contract_addresses []common.Address
	//target beneficiary addresses to replace in data instead of sender and old_tx_contract_address
	target_beneficiary_addresses []common.Address
}

var TRANSFER_EVENT_HASH = common.HexToHash("ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
var WITHDRAW_EVENT_HASH = common.HexToHash("7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65")
var DEPOSIT_EVENT_HASH = common.HexToHash("e1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c")
var EMPTY_ADDRESS = common.HexToAddress("ffffffffffffffffffffffffffffffffffffffff")

func NewAdversaryAccount(n uint64, t *types.Message) *AdversaryAccount {
	tmp_old_tx_contract_address := t.To()
	aa := &AdversaryAccount{
		balance_traces:               make(map[common.Address]map[common.Address][]TransferAmount),
		transfer_in_order:            []TransferInfo{},
		account_information:       make(map[common.Address]IndividualAdversaryAccountHelper),
		old_tx:                       t,
		old_tx_contract_address:      tmp_old_tx_contract_address,
		temp_contract_addresses:      []common.Address{},
		target_beneficiary_addresses: []common.Address{},
	}
	return aa
}
func (aa *AdversaryAccount) set_balance(addr common.Address, related_addr common.Address, bal common.Hash, token_addr common.Address, sender_receiver bool) {
	if balinfo := aa.balance_traces[addr]; balinfo != nil {
		if val := balinfo[token_addr]; val != nil {
			if sender_receiver {
				val = append(val, TransferAmount{TransferDir{From, related_addr}, bal})
			} else {
				val = append(val, TransferAmount{TransferDir{To, related_addr}, bal})
			}
			aa.balance_traces[addr][token_addr] = val
		} else {
			var tmp []TransferAmount
			if sender_receiver {
				tmp = append(tmp, TransferAmount{TransferDir{From, related_addr}, bal})
			} else {
				tmp = append(tmp, TransferAmount{TransferDir{To, related_addr}, bal})
			}
			aa.balance_traces[addr][token_addr] = tmp
		}
	} else {
		var tmp []TransferAmount
		if sender_receiver {
			tmp = append(tmp, TransferAmount{TransferDir{From, related_addr}, bal})
		} else {
			tmp = append(tmp, TransferAmount{TransferDir{To, related_addr}, bal})
		}
		aa.balance_traces[addr] = make(map[common.Address][]TransferAmount)
		aa.balance_traces[addr][token_addr] = tmp
	}
}

func (aa *AdversaryAccount) set_token_flow(addrfrom common.Address, addrto common.Address, amt common.Hash, token_addr common.Address) {
	if _, ok := ERC_TOKEN_INFORMATION_MAP[token_addr]; ok {
		aa.transfer_in_order = append(aa.transfer_in_order, TransferInfo{addrfrom, addrto, amt, token_addr})
		aa.set_balance(addrfrom, addrto, amt, token_addr, true)
		aa.set_balance(addrto, addrfrom, amt, token_addr, false)
	}
}

func (aa *AdversaryAccount) identify_helper() []AccountTransferInfo {
	var ret []AccountTransferInfo
	for a, b := range aa.balance_traces {
		// assert.Equal(len(b) > 0, true, "TransferAmount per token should not be empty!")
		var inner_ret []AccountTransferInfoPerToken
		for c, d := range b {
			// assert.Equal(len(d) > 0, true, "TransferAmount should not be empty!")
			earn_flag := true
			benefit := common.BigToHash(big.NewInt(0))
			for _, e := range d {
				if e.td.tdo == From {
					if earn_flag {
						if e.amt.Big().Cmp(benefit.Big()) == -1 {
							earn_flag = true
							benefit = common.BigToHash(big.NewInt(0).Sub(benefit.Big(), e.amt.Big()))
						} else {
							earn_flag = false
							benefit = common.BigToHash(big.NewInt(0).Sub(e.amt.Big(), benefit.Big()))
						}
					} else {
						earn_flag = false
						benefit = common.BigToHash(big.NewInt(0).Add(e.amt.Big(), benefit.Big()))
					}
				} else {
					if earn_flag {
						earn_flag = true
						benefit = common.BigToHash(big.NewInt(0).Add(e.amt.Big(), benefit.Big()))
					} else {
						if e.amt.Big().Cmp(benefit.Big()) == -1 {
							earn_flag = false
							benefit = common.BigToHash(big.NewInt(0).Sub(benefit.Big(), e.amt.Big()))
						} else {
							earn_flag = true
							benefit = common.BigToHash(big.NewInt(0).Sub(e.amt.Big(), benefit.Big()))
						}
					}
				}
			}
			inner_ret = append(inner_ret, AccountTransferInfoPerToken{c, earn_flag, benefit})
		}
		ret = append(ret, AccountTransferInfo{a, inner_ret})
	}
	return ret
}

func (aa *AdversaryAccount) anaylsis_net_profit_in_ten_thousandth_usd() {
	ret_vec := aa.identify_helper()
	if len(ret_vec) > 0 {
		for _, pack := range ret_vec {
			addr := pack.acct
			values := pack.tokenflow
			earn_flag := true
			benefit := common.BigToHash(big.NewInt(0))
			for _, flow := range values {
				a := flow.token
				b := flow.earn_flag
				c := flow.amt
				if ti, ok := ERC_TOKEN_INFORMATION_MAP[a]; ok {
					price := ti.price
					decimal := ti.decimals
					net_value := common.BigToHash(big.NewInt(0).Div(big.NewInt(0).Mul(c.Big(), price.Big()), decimal.Big()))
					if earn_flag {
						if b {
							earn_flag = true
							benefit = common.BigToHash(big.NewInt(0).Add(net_value.Big(), benefit.Big()))
						} else {
							if net_value.Big().Cmp(benefit.Big()) == 1 {
								earn_flag = false
								benefit = common.BigToHash(big.NewInt(0).Sub(net_value.Big(), benefit.Big()))
							} else {
								earn_flag = true
								benefit = common.BigToHash(big.NewInt(0).Sub(benefit.Big(), net_value.Big()))
							}
						}
					} else {
						if !b {
							earn_flag = false
							benefit = common.BigToHash(big.NewInt(0).Add(net_value.Big(), benefit.Big()))
						} else {
							if net_value.Big().Cmp(benefit.Big()) == 1 {
								earn_flag = true
								benefit = common.BigToHash(big.NewInt(0).Sub(net_value.Big(), benefit.Big()))
							} else {
								earn_flag = false
								benefit = common.BigToHash(big.NewInt(0).Sub(benefit.Big(), net_value.Big()))
							}
						}
					}
				} else {
					fmt.Println("ERROR: The token address is not recognizable!")
				}
			}
			var val IndividualAdversaryAccountHelper
			if benefit.Big().Cmp(big.NewInt(0)) == 0 {
				val.id = Neither
			} else if earn_flag {
				val.id = Beneficiary
				val.amt = benefit
			} else {
				val.id = Victim
				val.amt = benefit
			}
			aa.account_information[addr] = val
		}
	}
}
func (aa *AdversaryAccount) token_transfer_check() bool {
	
	aa.anaylsis_net_profit_in_ten_thousandth_usd()
	var beneficiary []common.Address
	var victim []common.Address

	for addr, result := range aa.account_information {
		if result.id == Beneficiary {
			beneficiary = append(beneficiary, addr)
		} else if result.id == Victim {
			victim = append(victim, addr)
		}
	}
	if len(beneficiary) == 0 {
		//fmt.Println("No beneficiary found")
		return false
	}else{
		/*
		for addr, result := range aa.account_information {
			if result.id == Beneficiary {
				//fmt.Println("Address", addr, " gains ", result.amt.Big(), " in 0.0001 USD unit")
			} else if result.id == Victim {
				//fmt.Println("Address", addr, " loses ", result.amt.Big(), " in 0.0001 USD unit")
			}
		}
		*/

	}
	// check beneficiary
	for _, i := range beneficiary {
		if i == aa.old_tx.From() || i == *(aa.old_tx_contract_address) {
			return true
		}
	}

	return false
}