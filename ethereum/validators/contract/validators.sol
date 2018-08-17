pragma solidity ^0.4.20;


contract Validators {
    struct Validator {
        address nodeAddr;
        uint96 deposit;
        uint48 pauseBlockNumber;
        uint8  pauseCause;
        uint32 idx;
    }

    uint public constant MIN_DEPOSIT = 1 ether;
    uint48 public constant DEPOSIT_LOCK_BLOCKS = 10;

    uint8 public constant PAUSE_CAUSE_VOLUNTARILY = 1;
    uint8 public constant PAUSE_CAUSE_PUNISHMENT = 2;

    mapping(address => Validator) validators;
    bytes32[] validatorsIndex;

    function () payable public {
        revert(); //Should not pay directly
    }

    function hasDeposit(address vAddr) public view returns (bool) {
        return validators[vAddr].deposit > 0;
    }

    function isPaused(address vAddr) public view returns (bool) {
        return validators[vAddr].pauseCause > 0;
    }

    function getNodeAddr(address vAddr) public view returns (address) {
        return validators[vAddr].nodeAddr;
    }

    function pauseValidation(address vAddr, address vFrom, uint8 cause) public {
        require(!isPaused(vAddr) && hasDeposit(vAddr));
        require(cause >= PAUSE_CAUSE_VOLUNTARILY);
        require(getNodeAddr(vFrom) == msg.sender);

        if(vAddr != vFrom){
            require(hasDeposit(vFrom) && !isPaused(vFrom));
            require(cause != PAUSE_CAUSE_VOLUNTARILY);
        }else{
            require(cause == PAUSE_CAUSE_VOLUNTARILY);
        }

        Validator storage v = validators[vAddr];
        v.pauseBlockNumber = uint48(block.number);
        v.pauseCause = cause;

        validatorsIndex[v.idx] = compactValidator(vAddr, v);
    }

    function resumeValidation(address vAddr) public {
        require(isPaused(vAddr));

        Validator storage v = validators[vAddr];
        require(v.nodeAddr == msg.sender);
        require(v.pauseCause == PAUSE_CAUSE_VOLUNTARILY);
        v.pauseBlockNumber = 0;
        v.pauseCause = 0;

        validatorsIndex[v.idx] = compactValidator(msg.sender, v);
    }

    function withdraw(address vAddr) public {
        require(hasDeposit(vAddr));
        require(isPaused(vAddr));
        
        Validator storage v = validators[vAddr];
        require(v.nodeAddr == msg.sender);
        require(v.pauseCause == PAUSE_CAUSE_VOLUNTARILY);
        require(v.pauseBlockNumber + DEPOSIT_LOCK_BLOCKS <= uint48(block.number));

        uint dep = v.deposit;
        deleteDeposit(vAddr);

        msg.sender.transfer(dep);
    }

    function addDeposit(address vAddr, address nodeAddr) public payable {
        Validator storage v = validators[vAddr];
        v.deposit += uint96(msg.value);
        require(v.deposit >= MIN_DEPOSIT);
        if(v.nodeAddr == 0)
            v.nodeAddr = nodeAddr;

        bytes32 compactedValidator = compactValidator(vAddr, v);
        if(v.idx > 0) {
            validatorsIndex[v.idx] = compactedValidator;
        }else{
            validatorsIndex.push(compactedValidator);
            v.idx = uint32(validatorsIndex.length);
        }
    }

    function deleteDeposit(address key) internal {
        mapping(address => Validator) map = validators;
        bytes32[] storage arr = validatorsIndex;

        Validator storage item = map[key];
        require(item.deposit > 0); //Should exist!

        assert(arr.length > 0); //If we are here then there must be table in array
        uint32 idx = item.idx - 1;
        if (arr.length > 1 && idx != arr.length-1) {
            arr[idx] = arr[arr.length-1];
            address addr = extractAddress(arr[idx]);
            map[addr].idx = idx + 1;
        }

        delete arr[arr.length-1];
        arr.length--;

        delete map[key];
    }

    //Returns validators compacted as bytes32
    function getCompactedValidators() public view returns (bytes32[]){
        return validatorsIndex;
    }

    /**
        Compacts validator into bytes32
        |reserved 24 bits|pauseCause 8 bits|senior 64 bits of deposit|validator address 160 bits|
    */
    function compactValidator(address vAddr, Validator storage v) internal view returns (bytes32) {
        return bytes32(uint256(vAddr)
            & ((uint256(v.deposit) >> 32) << 160)
            & (uint256(v.pauseCause) << 224));
    }

    function extractAddress(bytes32 compactedValidator) internal pure returns (address) {
        return address(compactedValidator);
    }
}
