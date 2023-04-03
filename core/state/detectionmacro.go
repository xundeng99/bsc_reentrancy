package state 


import (
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

type TokenInfo struct {
	price    common.Hash
	decimals common.Hash
}

var WHITE_LIST = map[common.Address]bool{
	common.HexToAddress("0x10ED43C718714eb63d5aA57B78B54704E256024E"):true,
	common.HexToAddress("0xb45a2dda996c32e93b8c47098e90ed0e7ab18e39"):true,

}

var ERC_TOKEN_INFORMATION_MAP = map[common.Address]TokenInfo{
	// Nerve 18
	common.HexToAddress("0x42F6f551ae042cBe50C739158b4f0CAC0Edb9096"): {common.BigToHash(big.NewInt(61)), common.BigToHash(big.NewInt(1000000000000000000))},
	// Gen Token 18
	common.HexToAddress("0xB0F2939A1c0e43683E5954c9Fe142F7df9F8D967"): {common.BigToHash(big.NewInt(1)), common.BigToHash(big.NewInt(1000000000000000000))},
	// Venus BUSD 18
	common.HexToAddress("0x95c78222B3D6e262426483D42CfA53685A67Ab9D"): {common.BigToHash(big.NewInt(216)), common.BigToHash(big.NewInt(100000000))},
	// Ellipsis 18
	common.HexToAddress("0xA7f552078dcC247C2684336020c03648500C6d9F"): {common.BigToHash(big.NewInt(911)), common.BigToHash(big.NewInt(1000000000000000000))},
	// JULb 18
	common.HexToAddress("0x32dFFc3fE8E3EF3571bF8a72c0d0015C5373f41D"): {common.BigToHash(big.NewInt(1154700)), common.BigToHash(big.NewInt(1000000000000000000))},
	// Bunny Token 18
	common.HexToAddress("0xC9849E6fdB743d08fAeE3E34dd2D1bc69EA11a51"): {common.BigToHash(big.NewInt(3088)), common.BigToHash(big.NewInt(1000000000000000000))},
	// PancakeSwap Token 18
	common.HexToAddress("0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"): {common.BigToHash(big.NewInt(94746)), common.BigToHash(big.NewInt(1000000000000000000))},
	// Binance-Peg BSC-USD 18
	common.HexToAddress("0x55d398326f99059fF775485246999027B3197955"): {common.BigToHash(big.NewInt(10000)), common.BigToHash(big.NewInt(1000000000000000000))},
	// Binance-Peg BUSD Token 18
	common.HexToAddress("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"): {common.BigToHash(big.NewInt(10000)), common.BigToHash(big.NewInt(1000000000000000000))},
	// AlpacaToken Token 18
	common.HexToAddress("0x8F0528cE5eF7B51152A59745bEfDD91D97091d2F"): {common.BigToHash(big.NewInt(6700)), common.BigToHash(big.NewInt(1000000000000000000))},
	//"Token A" from erc.sol for testing TODO: remove after done
	common.HexToAddress("0xb4c79daB8f259C7Aee6E5b2Aa729821864227e84"): {common.BigToHash(big.NewInt(2000)), common.BigToHash(big.NewInt(1))},
	//"Token B" from erc.sol for testing TODO: remove after done
	common.HexToAddress("0xee35211C4D9126D520bBfeaf3cFee5FE7B86F221"): {common.BigToHash(big.NewInt(1000)), common.BigToHash(big.NewInt(1))},
	// Wrapped BNB 18
	common.HexToAddress("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"): {common.BigToHash(big.NewInt(2962400)), common.BigToHash(big.NewInt(1000000000000000000))},
	//BNB 18
	common.HexToAddress("0x0000000000000000000000000000000000000001"): {common.BigToHash(big.NewInt(2962400)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Wrapped BNB (WBNB)
	common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c"): {common.BigToHash(big.NewInt(2186300)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Tether USD (USDT)
	common.HexToAddress("0x55d398326f99059ff775485246999027b3197955"): {common.BigToHash(big.NewInt(10000)), common.BigToHash(big.NewInt(1000000000000000000))},
	//BUSD Token (BUSD)
	common.HexToAddress("0xe9e7cea3dedca5984780bafc599bd69add087d56"): {common.BigToHash(big.NewInt(10000)), common.BigToHash(big.NewInt(1000000000000000000))},
	//USD Coin (USDC)
	common.HexToAddress("0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d"): {common.BigToHash(big.NewInt(10000)), common.BigToHash(big.NewInt(1000000000000000000))},
	//FGDTOKEN (FGD)
	common.HexToAddress("0x0566b9a8ffb8908682796751eed00722da967be0"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//LUNA (Wormhole) (LUNA)
	common.HexToAddress("0x156ab3346823b651294766e23e6cf87254d68962"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000))},
	//LITE (LITE)
	common.HexToAddress("0x4a846d300f793752ee8bd579192c477130c4b369"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//BTCB Token (BTCB)
	common.HexToAddress("0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c"): {common.BigToHash(big.NewInt(200509199)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Ethereum Token (ETH)
	common.HexToAddress("0x2170ed0880ac9a755fd29b2688956bd959f933f8"): {common.BigToHash(big.NewInt(700)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Wrapped UST Token (UST)
	common.HexToAddress("0x23396cf899ca06c4472205fc903bdb4de249d6fc"): {common.BigToHash(big.NewInt(10914000)), common.BigToHash(big.NewInt(1000000000000000000))},
	//HASH (HASH)
	common.HexToAddress("0x75aa08642963f160bf88b58a7d654bb6fc527a49"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//ORYX (ORYX)
	common.HexToAddress("0x10bb58010cb58e7249099ef2efdffe342928b639"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//PancakeSwap Token (Cake)
	common.HexToAddress("0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82"): {common.BigToHash(big.NewInt(30400)), common.BigToHash(big.NewInt(1000000000000000000))},
	//TiFi Token (TIFI)
	common.HexToAddress("0x17e65e6b9b166fb8e7c59432f0db126711246bc0"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//BitTorrent (BTT)
	common.HexToAddress("0x352cb5e19b12fc216548a2677bd0fce83bae434b"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Day Of Rights (DOR)
	common.HexToAddress("0x23766cb8a96ff2f46f664bc7d088a6306de73618"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//FGCA (FGCA)
	common.HexToAddress("0x58ca2873c9091ad2eab4f00cc415b846a286c080"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000))},
	//AscentN Token (ASN)
	common.HexToAddress("0xec077b020da5935b293701cf6072bcbb0eb1ed18"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(100000000))},
	//Freedom Protocol Token (FREE)
	common.HexToAddress("0x880bce9321c79cac1d290de6d31dde722c606165"): {common.BigToHash(big.NewInt(8500)), common.BigToHash(big.NewInt(100000000))},
	//Green Metaverse Token (GMT)
	common.HexToAddress("0x3019bf2a2ef8040c242c9a4c5c4bd4c81678b2a1"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000))},
	//DIAOS (DIAOS)
	common.HexToAddress("0x2adc3468fc4149932dd16d0244b7228aa8eead5c"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//BabyGODZ (BabyGODZ)
	common.HexToAddress("0xea2339b4ce4a6c384d6bb418cad59175d9b18c57"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(100000000))},
	//Super Soccer (SPS)
	common.HexToAddress("0x4539a864cfe6c489f219dbfb7624e2a43bf71aff"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//GreenSatoshiToken (GST)
	common.HexToAddress("0x4a2c860cec6471b9f5f5a336eb4f38bb21683c98"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(100000000))},
	//Brayzin Heist (BRZH)
	common.HexToAddress("0xbefd18d6dd7e5b98fbd57fcb61a7cb7a2fc82c68"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(100000000))},
	//BinaryX (BNX)
	common.HexToAddress("0x8c851d1a123ff703bd1f9dabe631b69902df5f97"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000))},
	//BlockAura BEP (TBAC)
	common.HexToAddress("0x2326c7395d02a8c89a9d7a0b0c1cf159d49ce51c"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//BabyGST (BabyGST)
	common.HexToAddress("0xcd3df672d4bd806593f4adf4442fe8e1f33a57f1"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//DEBT (DEBT)
	common.HexToAddress("0xc632f90affec7121120275610bf17df9963f181c"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//METAGALAXY LAND (MEGALAND)
	common.HexToAddress("0x7cd8c22d3f4b66230f73d7ffcb48576233c3fe33"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//FistToken (FIST)
	common.HexToAddress("0xc9882def23bc42d53895b8361d0b1edc7570bc6a"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//HOUND (HOUND)
	common.HexToAddress("0x3b9661287c23e1ea44eeb1c2a606b51432d70863"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//LUCA (LUCA)
	common.HexToAddress("0x51e6ac1533032e72e92094867fd5921e3ea1bfa0"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000))},
	//Radio Caca V2 (RACA)
	common.HexToAddress("0x12bb890508c125661e03b09ec06e404bc9289040"): {common.BigToHash(big.NewInt(5600)), common.BigToHash(big.NewInt(1000000000000000000))},
	//CryptoGodz (GODZ)
	common.HexToAddress("0xf0a8ecbce8caadb7a07d1fcd0f87ae1bd688df43"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000))},
	//TeddyDoge (TEDDY)
	common.HexToAddress("0x10f6f2b97f3ab29583d9d38babf2994df7220c21"): {common.BigToHash(big.NewInt(152400)), common.BigToHash(big.NewInt(100000))},
	//SpacePi Token (SpacePi)
	common.HexToAddress("0x69b14e8d3cebfdd8196bfe530954a0c226e5008e"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//HASH (HASH)
	common.HexToAddress("0x555d8355a31d62c68e13074bae90dc548ab6faf9"): {common.BigToHash(big.NewInt(1000)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Mobox (MBOX)
	common.HexToAddress("0x3203c9e46ca618c8c1ce5dc67e7e9d75f5da2377"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000))},
	//FileSystemVideo (FSV)
	common.HexToAddress("0xe9c7a827a4ba133b338b844c19241c864e95d75f"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Safuu (SAFUU)
	common.HexToAddress("0xe5ba47fd94cb645ba4119222e34fb33f59c7cd90"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//10Set Token (10SET)
	common.HexToAddress("0x1ae369a6ab222aff166325b7b87eb9af06c86e57"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//CALO (CALO)
	common.HexToAddress("0xb6b91269413b6b99242b1c0bc611031529999999"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Fidlecoin (FIDLE)
	common.HexToAddress("0xc861534efeb8a1c5daf31f1c13c440a7f86984f1"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000))},
	//Everdome (DOME)
	common.HexToAddress("0x475bfaa1848591ae0e6ab69600f48d828f61a80e"): {common.BigToHash(big.NewInt(61700)), common.BigToHash(big.NewInt(1000000000000000000))},
	//MetFi (MFI)
	common.HexToAddress("0xeb5bb9d14d27f75c787cf7475e7ed00d21dc7279"): {common.BigToHash(big.NewInt(200)), common.BigToHash(big.NewInt(1000000000000000000))},
	//TRON (TRX)
	common.HexToAddress("0x85eac5ac2f758618dfa09bdbe0cf174e7d574d5b"): {common.BigToHash(big.NewInt(400)), common.BigToHash(big.NewInt(1000000000000000000))},
	//MetaSwap (MSC)
	common.HexToAddress("0xeacad6c99965cde0f31513dd72de79fa24610767"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Amazing doge (Adoge)
	common.HexToAddress("0x0ebc30459551858e81306d583025d12c7d795fa2"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(100000000))},
	//Open Protocol Token (OPP)
	common.HexToAddress("0x5992bcd0145f8a7bcd5a1cc12138964bba08cb1e"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//ChainLink Token (LINK)
	common.HexToAddress("0xf8a0bf9cf54bb92f17374d9e9a321e6a111a51bd"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//MOVEZ.me (MOVEZ)
	common.HexToAddress("0x039cd22cb49084142d55fcd4b6096a4f51ffb3b4"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Splintershards (SPS)
	common.HexToAddress("0x1633b7157e7638c4d6593436111bf125ee74703f"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000))},
	//ELVES (ELV )
	common.HexToAddress("0xb78b7e82b074c267dd487db293a8faf831ae2d71"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(100000))},
	//Binance-Peg Pandora Chain DAO (PCD)
	common.HexToAddress("0xe4f1ae07760b985d1a94c6e5fb1589afaf44918c"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000))},
	//Doxy Finance (DOXY)
	common.HexToAddress("0xd6c7fbd02752e41d9b6000193668c16470fefd7d"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(100000000))},
	//BBK (BBK)
	common.HexToAddress("0x855290057797044f0d8e68544f4f736b444ef43c"): {common.BigToHash(big.NewInt(400)), common.BigToHash(big.NewInt(1000000000000000000))},
	//DarkShield (DKS)
	common.HexToAddress("0x121235cff4c59eec80b14c1d38b44e7de3a18287"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Square Token (SQUA)
	common.HexToAddress("0xb82beb6ee0063abd5fc8e544c852237aa62cbb14"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Mas (MAS)
	common.HexToAddress("0xffa905bb70b6d43ce7a8b84f2ba8e0a334b9765d"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000))},
	// *Artificial Intelligence Technology Network (AITN)
	common.HexToAddress("0xda3d20e21caeb1cf6dd84370aa0325087326f07a"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Game Coin [via ChainPort.io] (GAME)
	common.HexToAddress("0x66109633715d2110dda791e64a7b2afadb517abb"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//ZenFi AI (ZENFI)
	common.HexToAddress("0xa84d7a90bdbbe6de3fffe9b7f549366320ef90d3"): {common.BigToHash(big.NewInt(3500)), common.BigToHash(big.NewInt(1000000000000000000))},
	//ANTA (ANTA)
	common.HexToAddress("0x9eaf5369c9a9809bad8716591f9b2f68124ccd63"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//SingMon Token (SM)
	common.HexToAddress("0xbb1c91ad6cdf2e9604a664cbbe2f64f1c131e2f1"): {common.BigToHash(big.NewInt(2100)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Meta X DAO (XDAO)
	common.HexToAddress("0xfbe5c35ae60bf96bb91de653f32741840a582321"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//SquidGrow (SquidGrow)
	common.HexToAddress("0x88479186bac914e4313389a64881f5ed0153c765"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//MOVEZ.me (BURNZ)
	common.HexToAddress("0xa1b4c67ab08dceab25e7b22b4a11897ea042a1a8"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//FIT Token (FIT)
	common.HexToAddress("0x77922a521182a719a48ba650ac2a040269888888"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Fight Of The Ages (FOTA)
	common.HexToAddress("0x0a4e1bdfa75292a98c15870aef24bd94bffe0bd4"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Tidex Token (TDX)
	common.HexToAddress("0x317eb4ad9cfac6232f0046831322e895507bcbeb"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//CEEK (CEEK)
	common.HexToAddress("0xe0f94ac5462997d2bc57287ac3a3ae4c31345d66"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//CyberDragon Gold (Gold)
	common.HexToAddress("0xb3a6381070b1a15169dea646166ec0699fdaea79"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Coins Maker Coin (CMX)
	common.HexToAddress("0xe3d2d7552295e3d1d3fa151a44e10ec304eb0689"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Baby Doge Coin (BabyDoge)
	common.HexToAddress("0xc748673057861a797275cd8a068abb95a902e8de"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//pTokens $ANRX ($ANRX)
	common.HexToAddress("0xe2e7329499e8ddb1f2b04ee4b35a8d7f6881e4ea"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//CloudChat Token (CC)
	//common.HexToAddress("0x0c2bfa54d6d4231b6213803df616a504767020ea"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(10000000000000000000))},
	//Project Galaxy (GAL)
	common.HexToAddress("0xe4cc45bb5dbda06db6183e8bf016569f40497aa5"): {common.BigToHash(big.NewInt(2963300)), common.BigToHash(big.NewInt(1000000000000000000))},
	//DarkLight Gold (Gold)
	common.HexToAddress("0x1970b4e13ea8d07ed46571782dd3988d8569a857"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(100000000))},
	//TIME (TIME)
	common.HexToAddress("0x26619fa1d4c957c58096bbbeca6588dcfb12e109"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000))},
	//MarsDAO (MDAO)
	common.HexToAddress("0x60322971a672b81bcce5947706d22c19daecf6fb"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//ALAD Token (ALAD)
	common.HexToAddress("0xb1c3a83c74f52b245dc3a91462b1341b85d3fec7"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Heavenly Book (HBK)
	common.HexToAddress("0x613813fafa8b2ab299c04ff59fdbf8b8d65a1ea1"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//PIGS Token (AFP)
	common.HexToAddress("0x9a3321e1acd3b9f6debee5e042dd2411a1742002"): {common.BigToHash(big.NewInt(69100)), common.BigToHash(big.NewInt(1000000000000000000))},
	//BTCPi Token (BTCPi)
	common.HexToAddress("0x4aaad68be1a2ac9886b72b9dae474f3edd2132d9"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000))},
	//ESHARE (ESHARE)
	common.HexToAddress("0xdb20f6a8665432ce895d724b417f77ecac956550"): {common.BigToHash(big.NewInt(4600)), common.BigToHash(big.NewInt(1000000000000000000))},
	//ULTI Coin (ULTI)
	common.HexToAddress("0x42bfe4a3e023f2c90aebffbd9b667599fa38514f"): {common.BigToHash(big.NewInt(400)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Meta Channel Network (MCN)
	common.HexToAddress("0x1ecd7007beb60992fb16ee1d43a0d05edb139718"): {common.BigToHash(big.NewInt(2000)), common.BigToHash(big.NewInt(1000000000000000000))},
	//lizardtoken.finance (LIZ)
	common.HexToAddress("0xfcb520b47f5601031e0eb316f553a3641ff4b13c"): {common.BigToHash(big.NewInt(700)), common.BigToHash(big.NewInt(100000000))},
	//WingStep Token (WST)
	common.HexToAddress("0xd68f9f6769f68cb30505aa3f175f9e81e58503c8"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//WeWay Token (WWY)
	common.HexToAddress("0x9ab70e92319f0b9127df78868fd3655fb9f1e322"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//DD (DD)
	common.HexToAddress("0x7f7a036aba49122dbbdb3da9bd67b45f10fcd765"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Dark Token (DARK)
	common.HexToAddress("0x12fc07081fab7de60987cad8e8dc407b606fb2f8"): {common.BigToHash(big.NewInt(6899)), common.BigToHash(big.NewInt(1000000000000000000))},
	//MIC (MIC)
	common.HexToAddress("0x71fc2c893e41eabdf9c4afda3b2cdb46b93cd8aa"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Metis Token (Metis)
	common.HexToAddress("0xe552fb52a4f19e44ef5a967632dbc320b0820639"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000000000000000))},
	//LAN (LAN)
	common.HexToAddress("0xffecf080207264cd6a32883c74bfbed0e10a1096"): {common.BigToHash(big.NewInt(50199)), common.BigToHash(big.NewInt(1000000000000000000))},
	//Polkadot Token (DOT)
	common.HexToAddress("0x7083609fce4d1d8dc0c979aab8c869ea2c873402"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(100000000))},
	//Trust Wallet (TWT)
	common.HexToAddress("0x4b0f1812e5df2a09796481ff14017e6005508003"): {common.BigToHash(big.NewInt(3300)), common.BigToHash(big.NewInt(1000000000000000000))},
	//DEXShare (DEXShare)
	common.HexToAddress("0xf4914e6d97a75f014acfcf4072f11be5cffc4ca6"): {common.BigToHash(big.NewInt(12100)), common.BigToHash(big.NewInt(1000000000))},
	//Cardano Token (ADA)
	common.HexToAddress("0x3ee2200efb3400fabb9aacf31297cbdd1d435d47"): {common.BigToHash(big.NewInt(0)), common.BigToHash(big.NewInt(1000000))},
}