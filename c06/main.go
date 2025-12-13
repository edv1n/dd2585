package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

/* Example response:
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": [
    {
      "action": {
        "from": "0xe73f6a4e2e09b22862eb4cd21baadfc099dfc857",
        "callType": "call",
        "gas": "0x48fb8",
        "input": "0x42035482",
        "to": "0x8a694d2362c9189fc53f4d9b20c53e931dda7860",
        "value": "0x0"
      },
      "blockHash": "0xa96885866f505bdb2b68d4d0516103dccc648193f2b32c865e006d79f2798a73",
      "blockNumber": 9827567,
      "result": {
        "gasUsed": "0x1d5dc",
        "output": "0x"
      },
      "subtraces": 1,
      "traceAddress": [],
      "transactionHash": "0xed5b0db032c095c04b00f27d3751d29c2ca543fe2d9e70cea205872c441859a3",
      "transactionPosition": 14,
      "type": "call"
    },
    {
      "action": {
        "from": "0x8a694d2362c9189fc53f4d9b20c53e931dda7860",
        "callType": "call",
        "gas": "0x4719b",
        "input": "0x21f3f819a9abc4dfddee92253832274bee2ab6c759cd75e929f6e5e992dda6e54391fd9e6010212f0a836df5ccad51452f8c0674d73411141aac5627ba688e16b99e06400000000000000000000000006366f9686ecab28ef2c82989c2b92b5633b3a7860000000000000000000000008a694d2362c9189fc53f4d9b20c53e931dda7860",
        "to": "0x3819c7071f2bc39c83187bf5b5aea79fa3e37c42",
        "value": "0x0"
      },
      "blockHash": "0xa96885866f505bdb2b68d4d0516103dccc648193f2b32c865e006d79f2798a73",
      "blockNumber": 9827567,
      "result": {
        "gasUsed": "0x1c96e",
        "output": "0x"
      },
      "subtraces": 0,
      "traceAddress": [
        0
      ],
      "transactionHash": "0xed5b0db032c095c04b00f27d3751d29c2ca543fe2d9e70cea205872c441859a3",
      "transactionPosition": 14,
      "type": "call"
    }
  ]
}
*/

type TransactionTraceResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Result  []TraceResult `json:"result"`
}

type TraceResult struct {
	Action      TraceAction     `json:"action"`
	BlockHash   string          `json:"blockHash"`
	BlockNumber int             `json:"blockNumber"`
	Result      TraceResultData `json:"result"`
	Subtraces   int             `json:"subtraces"`
}

type TraceAction struct {
	From     string `json:"from"`
	CallType string `json:"callType"`
	Gas      string `json:"gas"`
	Input    string `json:"input"`
	To       string `json:"to"`
	Value    string `json:"value"`
}

type TraceResultData struct {
	GasUsed string `json:"gasUsed"`
	Output  string `json:"output"`
}

func main() { // Application entry point
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <ethereum-debug-enabled-rpc-url> <tx-hash>")
	}

	url := os.Args[1]
	txHash := os.Args[2]

	debug := false
	if os.Getenv("DEBUG") != "" {
		debug = true
	}

	contract, err := abi.JSON(strings.NewReader(abiContent))
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Contract methods: %+v", contract.Methods)

	resp, err := getTransactionTrace(url, txHash)
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.Result) == 0 {
		log.Fatal("No trace results found")
	}

	for i, trace := range resp.Result {
		if debug {
			fmt.Printf("Trace %d:\n", i)
			fmt.Printf("  From: %s\n", trace.Action.From)
			fmt.Printf("  To: %s\n", trace.Action.To)
			fmt.Printf("  Value: %s\n", trace.Action.Value)
			fmt.Printf("  Input: %s\n", trace.Action.Input)
			fmt.Printf("  Gas: %s\n", trace.Action.Gas)
			fmt.Printf("  Gas Used: %s\n", trace.Result.GasUsed)
			fmt.Printf("  Output: %s\n", trace.Result.Output)
			fmt.Println()
		}

		if trace.Action.To != "0x3819c7071f2bc39c83187bf5b5aea79fa3e37c42" {
			continue
		}

		if debug {
			fmt.Printf("Found relevant trace to address %s:\n", trace.Action.To)
		}

		txInput := trace.Action.Input

		decodedSig, err := hex.DecodeString(txInput[2:10])
		if err != nil {
			log.Print(err)

			continue
		}

		method, err := contract.MethodById(decodedSig)
		if err != nil {
			if debug {
				fmt.Printf("Error getting method by id %x: %s\n", decodedSig, err.Error())
			}

			continue
		}

		if method == nil {
			if debug {
				fmt.Printf("No method found for signature %x\n", decodedSig)
			}

			continue
		}

		if method.Name != "registerData" {
			if debug {
				fmt.Printf("Method found is %s, not registerData\n", method.Name)
			}

			continue
		}

		fmt.Printf("The transaction has an internal call to 0x3819c7071f2bc39c83187bf5b5aea79fa3e37c42 and this call is to %s(bytes32,bytes32,address,address)\n", method.Name)

		os.Exit(0)
	}

	fmt.Println("The transaction has no internal call to 0x3819c7071f2bc39c83187bf5b5aea79fa3e37c42 or has no call to registerData(bytes32,bytes32,address,address)")
}

func getTransactionTrace(url, txHash string) (*TransactionTraceResponse, error) {
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"method":"trace_transaction","params":["%s"],"id":1,"jsonrpc":"2.0"}`, txHash))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	resp := &TransactionTraceResponse{}
	if err := json.NewDecoder(res.Body).Decode(resp); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}

func getTransactionTraceDummy(_, _ string) (*TransactionTraceResponse, error) {
	result := `{"jsonrpc":"2.0","id":1,"result":[{"action":{"from":"0xe73f6a4e2e09b22862eb4cd21baadfc099dfc857","callType":"call","gas":"0x48fb8","input":"0x42035482","to":"0x8a694d2362c9189fc53f4d9b20c53e931dda7860","value":"0x0"},"blockHash":"0xa96885866f505bdb2b68d4d0516103dccc648193f2b32c865e006d79f2798a73","blockNumber":9827567,"result":{"gasUsed":"0x1d5dc","output":"0x"},"subtraces":1,"traceAddress":[],"transactionHash":"0xed5b0db032c095c04b00f27d3751d29c2ca543fe2d9e70cea205872c441859a3","transactionPosition":14,"type":"call"},{"action":{"from":"0x8a694d2362c9189fc53f4d9b20c53e931dda7860","callType":"call","gas":"0x4719b","input":"0x21f3f819a9abc4dfddee92253832274bee2ab6c759cd75e929f6e5e992dda6e54391fd9e6010212f0a836df5ccad51452f8c0674d73411141aac5627ba688e16b99e06400000000000000000000000006366f9686ecab28ef2c82989c2b92b5633b3a7860000000000000000000000008a694d2362c9189fc53f4d9b20c53e931dda7860","to":"0x3819c7071f2bc39c83187bf5b5aea79fa3e37c42","value":"0x0"},"blockHash":"0xa96885866f505bdb2b68d4d0516103dccc648193f2b32c865e006d79f2798a73","blockNumber":9827567,"result":{"gasUsed":"0x1c96e","output":"0x"},"subtraces":0,"traceAddress":[0],"transactionHash":"0xed5b0db032c095c04b00f27d3751d29c2ca543fe2d9e70cea205872c441859a3","transactionPosition":14,"type":"call"}]}`
	resp := &TransactionTraceResponse{}
	if err := json.Unmarshal([]byte(result), resp); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resp, nil
}
