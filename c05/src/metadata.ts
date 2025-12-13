export default { 
    "abi": [
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "",
                    "type": "address"
                }
            ],
            "name": "data",
            "outputs": [
                {
                    "internalType": "bytes32",
                    "name": "tx_challenge01",
                    "type": "bytes32"
                },
                {
                    "internalType": "bytes32",
                    "name": "tx_challenge02",
                    "type": "bytes32"
                },
                {
                    "internalType": "address",
                    "name": "contract_challenge03",
                    "type": "address"
                },
                {
                    "internalType": "address",
                    "name": "contract_challenge04",
                    "type": "address"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "address",
                    "name": "user",
                    "type": "address"
                }
            ],
            "name": "getData",
            "outputs": [
                {
                    "components": [
                        {
                            "internalType": "bytes32",
                            "name": "tx_challenge01",
                            "type": "bytes32"
                        },
                        {
                            "internalType": "bytes32",
                            "name": "tx_challenge02",
                            "type": "bytes32"
                        },
                        {
                            "internalType": "address",
                            "name": "contract_challenge03",
                            "type": "address"
                        },
                        {
                            "internalType": "address",
                            "name": "contract_challenge04",
                            "type": "address"
                        }
                    ],
                    "internalType": "struct Challenge04.Data",
                    "name": "",
                    "type": "tuple"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [],
            "name": "getRegisteredAddresses",
            "outputs": [
                {
                    "internalType": "address[]",
                    "name": "",
                    "type": "address[]"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "_tx_challenge01",
                    "type": "bytes32"
                },
                {
                    "internalType": "bytes32",
                    "name": "_tx_challenge02",
                    "type": "bytes32"
                },
                {
                    "internalType": "address",
                    "name": "_contract_challenge03",
                    "type": "address"
                },
                {
                    "internalType": "address",
                    "name": "_contract_challenge04",
                    "type": "address"
                }
            ],
            "name": "registerData",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "uint256",
                    "name": "",
                    "type": "uint256"
                }
            ],
            "name": "registeredAddresses",
            "outputs": [
                {
                    "internalType": "address",
                    "name": "",
                    "type": "address"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        }
    ]
}