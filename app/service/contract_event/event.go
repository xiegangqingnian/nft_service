/**
@author: 铁柱
@date:2021/6/25
@note: 合约事件通知
**/
package contract_event

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"nft_service/app/global/variable"
	"nft_service/app/utils/node"
)

func GetAllEventLog() {
	contractAddress := common.HexToAddress(variable.ContractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	//contractAbi,err := abi.JSON(strings.NewReader(_contractAbi))
	//if err != nil {
	//	variable.ZapLog.Error("event"+ err.Error())
	//}

	logs := make(chan types.Log)

	sub, err := node.NodeClient.SubscribeFilterLogs(context.Background(),
		query,
		logs)
	if err != nil {
		variable.ZapLog.Error("event" + err.Error())
	}
	variable.ZapLog.Info("listening ！！！！！")

	for {
		select {
		case err := <-sub.Err():
			variable.ZapLog.Error("event" + err.Error())
		case vLog := <-logs:
			switch vLog.Topics[0].Hex() {
			case AddOrderSigHash:
				fmt.Println("add order")
			case FinishOrderSigHash:
				fmt.Println("finish order")
			case CancelOrderSigHash:
				fmt.Println("cancel order")
			case GameOverSigHash:
				fmt.Println("game over")
			case VoteSigHash:
				fmt.Println("new vote")
			}
		}
	}

}
