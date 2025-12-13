(async () => {
    try {
        web3.eth.getChainId().then(console.log);
        let methodAbi = {
            "inputs": [],
            "name": "c02",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        };
        //var contract = new Contract([methodAbi],'0xe621aBEa69C75dA07C3850eeA3965DE0599d4B3D');
        const accounts = await web3.eth.getAccounts()

        let contract = new web3.eth.Contract([methodAbi], '0xe621aBEa69C75dA07C3850eeA3965DE0599d4B3D');
        // web3.js in constructor calculates the method signature, but we have to reset it mannually
        console.log(contract.options.jsonInterface[0].signature);
        console.log(contract._jsonInterface[0].signature);

        contract.options.jsonInterface[0].signature = '0x42424242';
        contract._jsonInterface[0].signature = '0x42424242';

        console.log(contract.options.jsonInterface[0].signature);
        console.log(contract._jsonInterface[0].signature);

        console.log(accounts)

        const signer = accounts[0];

        // now we use 'someMethod' name to call or send to the 0x2b958b8 method

        let tx = await contract.methods.c02();
        const receipt = await tx.send({
            // used first account from your wallet. 
            from: (await web3.eth.getAccounts())[0],
            gas: await tx.estimateGas(),
        });
    } catch (e) {
        console.log(e.message)
    }
})()