pragma solidity ^0.4.20;


contract Validators {
    struct Validator {
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
        addDeposit(msg.sender, msg.value);
    }

    function hasDeposit(address addr) public view returns (bool) {
        return validators[addr].deposit > 0;
    }

    function isPaused(address addr) public view returns (bool) {
        return validators[addr].pauseCause > 0;
    }

    function pauseValidation(address addr, uint8 cause) public {
        require(!isPaused(addr) && hasDeposit(addr));
        require(cause >= PAUSE_CAUSE_VOLUNTARILY);

        if(addr != msg.sender){
            require(hasDeposit(msg.sender) && !isPaused(msg.sender));
            require(cause != PAUSE_CAUSE_VOLUNTARILY);
        }else{
            require(cause == PAUSE_CAUSE_VOLUNTARILY);
        }

        Validator storage v = validators[addr];
        v.pauseBlockNumber = uint48(block.number);
        v.pauseCause = cause;

        validatorsIndex[v.idx] = compactValidator(addr, v);
    }

    function resumeValidation() public {
        require(isPaused(msg.sender));

        Validator storage v = validators[msg.sender];
        require(v.pauseCause == PAUSE_CAUSE_VOLUNTARILY);
        v.pauseBlockNumber = 0;
        v.pauseCause = 0;

        validatorsIndex[v.idx] = compactValidator(msg.sender, v);
    }

    function withdraw() public {
        require(hasDeposit(msg.sender));
        require(isPaused(msg.sender));
        
        Validator storage v = validators[msg.sender];
        require(v.pauseCause == PAUSE_CAUSE_VOLUNTARILY);
        require(v.pauseBlockNumber + DEPOSIT_LOCK_BLOCKS <= uint48(block.number));

        uint dep = v.deposit;
        deleteDeposit(msg.sender); 

        msg.sender.transfer(dep);
    }

    function addDeposit(address from, uint value) internal {
        Validator storage val = validators[from];
        val.deposit += uint96(value);
        require(val.deposit >= MIN_DEPOSIT);

        bytes32 compactedValidator = compactValidator(from, val);
        if(val.idx > 0) {
            validatorsIndex[val.idx] = compactedValidator;
        }else{
            validatorsIndex.push(compactedValidator);
            val.idx = uint32(validatorsIndex.length);
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
    function compactValidator(address addr, Validator storage v) internal view returns (bytes32) {
        return bytes32(uint256(addr) 
            & ((uint256(v.deposit) >> 32) << 160)
            & (uint256(v.pauseCause) << 224));
    }

    function extractAddress(bytes32 compactedValidator) internal pure returns (address) {
        return address(compactedValidator);
    }
}
