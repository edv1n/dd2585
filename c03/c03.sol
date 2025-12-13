pragma solidity >=0.7.0 <0.9.0;

contract c03 {
    address internal _owner;

    constructor() {
        _owner = msg.sender;
    }

    function getGithubId() public pure returns (string memory) {
        return "edv1n";
    }

    function owner() public view returns (address) {
        return _owner;
    }
}