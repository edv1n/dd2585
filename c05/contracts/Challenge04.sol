contract Challenge04 {
    struct Data {
        bytes32 tx_challenge01;
        bytes32 tx_challenge02;
        address contract_challenge03;
        address contract_challenge04;
    }

    mapping(address => Data) public data;

    address[] public registeredAddresses;

    function registerData(
        bytes32 _tx_challenge01,
        bytes32 _tx_challenge02,
        address _contract_challenge03,
        address _contract_challenge04
    ) public {
        require(msg.sender == _contract_challenge04, "Sender must match challenge04 address");

        registeredAddresses.push(tx.origin);

        data[tx.origin] = Data({
            tx_challenge01: _tx_challenge01,
            tx_challenge02: _tx_challenge02,
            contract_challenge03: _contract_challenge03,
            contract_challenge04: _contract_challenge04
        });
    }

    function getRegisteredAddresses() public view returns (address[] memory) {
        return registeredAddresses;
    }

    function getData(address user) public view returns (Data memory) {
        return data[user];
    }
}