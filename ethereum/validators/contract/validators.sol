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

    mapping(bytes32 => Validator) validators;
    bytes32[] validatorsIndex;
    bytes32[] validatorsCompacted;

    function () payable public {
        revert(); //Should not pay directly
    }

    function hasDeposit(bytes32 vPub) public view returns (bool) {
        return validators[vPub].deposit > 0;
    }

    function isPaused(bytes32 vPub) public view returns (bool) {
        return validators[vPub].pauseCause > 0;
    }

    function getNodeAddr(bytes32 vPub) public view returns (address) {
        return validators[vPub].nodeAddr;
    }

    function pauseValidation(bytes32 vPub, bytes32 vFrom, uint8 cause) public {
        require(!isPaused(vPub) && hasDeposit(vPub));
        require(cause >= PAUSE_CAUSE_VOLUNTARILY);
        require(getNodeAddr(vFrom) == msg.sender);

        if(vPub != vFrom){
            require(hasDeposit(vFrom) && !isPaused(vFrom));
            require(cause != PAUSE_CAUSE_VOLUNTARILY);
        }else{
            require(cause == PAUSE_CAUSE_VOLUNTARILY);
        }

        Validator storage v = validators[vPub];
        v.pauseBlockNumber = uint48(block.number);
        v.pauseCause = cause;

        validatorsIndex[v.idx] = compactValidator(vPub, v);
    }

    function resumeValidation(bytes32 vPub) public {
        require(isPaused(vPub));

        Validator storage v = validators[vPub];
        require(v.nodeAddr == msg.sender);
        require(v.pauseCause == PAUSE_CAUSE_VOLUNTARILY);
        v.pauseBlockNumber = 0;
        v.pauseCause = 0;

        validatorsIndex[v.idx] = compactValidator(vPub, v);
    }

    function withdraw(bytes32 vPub) public {
        require(hasDeposit(vPub));
        require(isPaused(vPub));
        
        Validator storage v = validators[vPub];
        require(v.nodeAddr == msg.sender);
        require(v.pauseCause == PAUSE_CAUSE_VOLUNTARILY);
        require(v.pauseBlockNumber + DEPOSIT_LOCK_BLOCKS <= uint48(block.number));

        uint dep = v.deposit;
        deleteDeposit(vPub);

        msg.sender.transfer(dep);
    }

    function addDeposit(bytes32 vPub, address nodeAddr) public payable {
        Validator storage v = validators[vPub];
        v.deposit += uint96(msg.value);
        require(v.deposit >= MIN_DEPOSIT);
        if(v.nodeAddr == 0)
            v.nodeAddr = nodeAddr;

        bytes32 compactedValidator = compactValidator(vPub, v);
        if(v.idx > 0) {
            validatorsIndex[v.idx] = vPub;
            validatorsCompacted[v.idx] = compactedValidator;
        }else{
            validatorsIndex.push(vPub);
            validatorsCompacted.push(compactedValidator);
            v.idx = uint32(validatorsIndex.length);
        }
    }

    function deleteDeposit(bytes32 key) internal {
        mapping(bytes32 => Validator) map = validators;
        bytes32[] storage arr = validatorsIndex;
        bytes32[] storage ar1 = validatorsCompacted;

        Validator storage item = map[key];
        require(item.deposit > 0); //Should exist!

        assert(arr.length > 0); //If we are here then there must be table in array
        uint32 idx = item.idx - 1;
        if (arr.length > 1 && idx != arr.length-1) {
            arr[idx] = arr[arr.length-1];
            ar1[idx] = ar1[ar1.length-1];
            bytes32 addr = arr[idx];
            map[addr].idx = idx + 1;
        }

        delete arr[arr.length-1];
        arr.length--;
        delete ar1[ar1.length-1];
        ar1.length--;

        delete map[key];
    }

    //Returns validators compacted as bytes32
    function getCompactedValidators() public view returns (bytes32[] ValidatorsCompacted, bytes32[] ValidatorsIndex){
        ValidatorsCompacted = validatorsCompacted;
        ValidatorsIndex = validatorsIndex;
    }

    /**
        Compacts validator into bytes32
        |reserved 24 bits|pauseCause 8 bits|senior 64 bits of deposit|validator address 160 bits|
    */
    function compactValidator(bytes32 vPub, Validator storage v) internal view returns (bytes32) {
        uint256 hash = uint256(sha256(vPub));
         return bytes32(
             (hash >> 96)
             | ((uint256(v.deposit) >> 32) << 160)
             | (uint256(v.pauseCause) << 224));
    }

    function extractAddress(bytes32 compactedValidator) internal pure returns (address) {
        return address(compactedValidator);
    }
}
