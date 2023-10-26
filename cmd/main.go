package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func main() {
	// 初始化以太坊客户端
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2ed8f259187948c299d2d63352385b42")
	if err != nil {
		panic(err)
	}

	account := common.HexToAddress("0x1f889dac7f0686467e6826df51dcd3966b21a537")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Println(ethValue) // 25.729324269165216041

	//// 导入账户 keystore 文件
	//keystoreFile, err := os.Open("/Users/chao/GolandProjects/keystore.json")
	//if err != nil {
	//	panic(err)
	//}
	//buff := new(bytes.Buffer)
	//_, _ = io.Copy(buff, keystoreFile)
	//
	//key, err := keystore.DecryptKey(buff.Bytes(), "Helloi23")
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 获取账户地址
	//address := common.HexToAddress(key.Address.Hex())
	//
	//// 查询账户余额
	//balance, err := client.BalanceAt(context.Background(), address, nil)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("账户余额：", balance.String())
	//
	//// 转账
	//toAddress := "0x1234567890abcdef01234567890abcdef01234567"
	//amount := big.NewInt(1000000000000000000)
	//gasPrice := big.NewInt(20000000000)
	//gasLimit := uint64(21000)
	//
	//transferTx, err := types.NewTransaction(
	//	context.Background(),
	//	key.PrivateKey,
	//	toAddress,
	//	amount,
	//	gasPrice,
	//	gasLimit,
	//	nil,
	//)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = client.SendTransaction(context.Background(), transferTx)
	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println("转账成功！")
}
