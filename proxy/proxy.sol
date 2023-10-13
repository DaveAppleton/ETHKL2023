pragma solidity ^0.8.13;

// start OMIT
contract theproxy {
    address immutable            destination;
    constructor(address _destination) {
        destination = _destination;
    }

    fallback(bytes calldata b) external payable returns (bytes memory)  {
        (bool success, bytes memory returnedData) = destination.delegatecall(b);
        if (!success) {
            assembly {
                revert(add(returnedData,32),mload(returnedData))
            }
        }
        return returnedData; 
    }

}
// end OMIT