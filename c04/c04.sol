// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract c04 {
    constructor() {}

    function callRegistry() external {
        address target = 0x3819C7071f2bc39C83187Bf5B5aeA79Fa3e37C42;
        bytes32 p1 = 0xa9abc4dfddee92253832274bee2ab6c759cd75e929f6e5e992dda6e54391fd9e;
        bytes32 p2 = 0x6010212f0a836df5ccad51452f8c0674d73411141aac5627ba688e16b99e0640;
        address p3 = 0x6366f9686eCAb28Ef2c82989C2B92b5633b3a786;
        address p4 = address(this);

        (bool success, bytes memory returndata) = target.call(
            abi.encodeWithSignature("registerData(bytes32,bytes32,address,address)", p1, p2, p3, p4)
        );

        require(success);
    }
}
