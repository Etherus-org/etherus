pragma solidity ^0.5.0;


contract Validators {

    event ValidatorUpdated(
        bytes32 vPub,
        uint96 deposit
    );

    event ValidatorRemoved(
        bytes32 vPub,
        uint8 cause
    );

    struct Validator {
        address payable nodeAddr;
        uint96 deposit;
        uint48 pauseBlockNumber;
        uint8  pauseCause;
        uint32 idx;
        address payable receiver;
    }

    uint private constant MIN_DEPOSIT_INCREMENT = 10 ether;
    uint private constant MIN_DEPOSIT = 2500 ether;
    uint48 private constant DEPOSIT_LOCK_BLOCKS = 120;
    //Число блоков, на которые блокируется нода, получившее предупреждение
    uint48 private constant NODE_LOCK_BLOCKS = 1200;

    uint8 private constant PAUSE_NOT_PAUSED = 0;      //Активное состояние
    uint8 private constant PAUSE_CAUSE_VOLUNTARILY = 1;      //Добровольная пауза
    uint8 private constant PAUSE_CAUSE_SEMI_VOLUNTARILY = 2; //Предупреждение
    uint8 private constant PAUSE_CAUSE_PUNISHMENT = 3;       //Наказание

    mapping(bytes32 => Validator) validators;
    //Array of validators public keys
    bytes32[] validatorsPubKeys;
    //Array of compacted validators info
    bytes32[] validatorsCompacted;

    function () external payable {
        revert(); //Should not pay directly
    }

    function hasDeposit(bytes32 vPub) public view returns (bool) {
        return validators[vPub].deposit > 0;
    }

    function isPaused(bytes32 vPub) public view returns (bool) {
        return validators[vPub].pauseCause > 0;
    }

    function getNodeAddr(bytes32 vPub) public view returns (address payable) {
        return validators[vPub].nodeAddr;
    }

    function getNodeReceiver(bytes32 vPub) public view returns (address payable) {
        address payable recv = validators[vPub].receiver;
        if(recv == address(0))
            recv = getNodeAddr(vPub);
        assert(recv != address(0)); //receiver always should be not 0 by initialization
        return recv;
    }

    function pauseValidation(bytes32 vPub, bytes32 vFrom, uint8 cause) public {
        require(!isPaused(vPub) && hasDeposit(vPub), "Node should not be paused and should have deposit");
        require(cause >= PAUSE_CAUSE_VOLUNTARILY, "You should specify cause");
        require(getNodeAddr(vFrom) == msg.sender, "Node should correctly pass its validator public key");

        if(vPub != vFrom){
            require(hasDeposit(vFrom) && !isPaused(vFrom), "You should not be paused to pause others");
            require(cause != PAUSE_CAUSE_VOLUNTARILY, "This pausing is not voluntary!");
        }else{
            require(cause == PAUSE_CAUSE_VOLUNTARILY, "You are pausing valuntarily");
        }

        Validator storage v = validators[vPub];
        v.pauseBlockNumber = uint48(block.number);
        v.pauseCause = cause;

        compactAndSave(v);

        emit ValidatorRemoved(vPub, cause);
    }

    function resumeValidation(bytes32 vPub) public {
        require(isPaused(vPub), "This node is not paused!");

        Validator storage v = validators[vPub];
        require(v.nodeAddr == msg.sender, "You can only unpause yourself");
        require(v.pauseCause == PAUSE_CAUSE_VOLUNTARILY ||
        v.pauseCause == PAUSE_CAUSE_SEMI_VOLUNTARILY, "You can only unpause voluntary paused node");
        require(v.pauseCause != PAUSE_CAUSE_SEMI_VOLUNTARILY ||
        v.pauseBlockNumber + NODE_LOCK_BLOCKS <= uint48(block.number), "You should wait some blocks before unpause");

        require(v.deposit >= MIN_DEPOSIT, "You can not unpause before deposit exceeds min value");

        v.pauseBlockNumber = 0;
        v.pauseCause = 0;

        compactAndSave(v);

        emit ValidatorUpdated(vPub, v.deposit);
    }

    function withdraw(bytes32 vPub) public {
        require(hasDeposit(vPub), "You should have deposit");
        require(isPaused(vPub), "You should be paused");

        Validator storage v = validators[vPub];

        //Checking that this is called from the node address
        require(v.nodeAddr == msg.sender, "Only node itself can withdraw!");
        require(v.pauseCause == PAUSE_CAUSE_VOLUNTARILY, "You should have voluntarily paused before withdraw");
        require(v.pauseBlockNumber + DEPOSIT_LOCK_BLOCKS <= uint48(block.number), "You should wait some blocks before withdraw");

        uint dep = v.deposit;
        address payable rcv = getNodeReceiver(vPub);

        deleteDeposit(vPub);
        rcv.transfer(dep);
    }

    function addDeposit(bytes32 vPub) public payable {
        addInitialDeposit(vPub, address(0), address(0));
    }

    function addInitialDeposit(bytes32 vPub, address payable nodeAddr, address payable receiver) public payable {
        require(msg.value >= MIN_DEPOSIT_INCREMENT, "Too small value to add to the deposit");
        
        Validator storage v = validators[vPub];
        v.deposit += uint96(msg.value);
        if(v.nodeAddr == address(0)){
            require(nodeAddr != address(0), "Please specify node address");
            v.nodeAddr = nodeAddr;
            v.receiver = receiver;
            //Always create new validator in paused state
            v.pauseCause = PAUSE_CAUSE_VOLUNTARILY;
            v.pauseBlockNumber = uint48(block.number);
        }

        if(v.idx > 0) {
            compactAndSave(v);
        }else{
            bytes32 compactedValidator = compactValidator(v);
            validatorsPubKeys.push(vPub);
            validatorsCompacted.push(compactedValidator);
            v.idx = uint32(validatorsPubKeys.length);
        }

        if(msg.value > 0)
            emit ValidatorUpdated(vPub, v.deposit);
    }

    function deleteDeposit(bytes32 key) internal {
        mapping(bytes32 => Validator) storage map = validators;
        bytes32[] storage arr = validatorsPubKeys;
        bytes32[] storage ar1 = validatorsCompacted;

        Validator storage item = map[key];
        require(item.deposit > 0, "Deposit should exist!"); //Should exist!

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

    /**
        Compacts validator into bytes32
        |reserved 24 bits|pauseCause 8 bits|senior 64 bits of deposit|node address 160 bits|
    */
    function compactValidator(Validator storage v) internal view returns (bytes32) {
        uint256 hash = uint256(v.nodeAddr); //Too much gas on private testnet
        return bytes32(
            hash
            | ((uint256(v.deposit) >> 32) << 160)
            | (uint256(v.pauseCause) << 224));
    }

    function compactAndSave(Validator storage v) internal returns (bytes32) {
        validatorsCompacted[v.idx - 1] = compactValidator(v);
    }

    //Returns validators compacted as bytes32
    function getCompactedValidators() public view returns (bytes32[] memory ValidatorsCompacted, bytes32[] memory ValidatorsPubKeys){
        ValidatorsCompacted = validatorsCompacted;
        ValidatorsPubKeys = validatorsPubKeys;
    }

    //Returns active validators compacted as bytes32
    function getActiveCompactedValidators() public view returns (bytes32[] memory ValidatorsCompacted, bytes32[] memory ValidatorsPubKeys){
        uint count = getActiveCount();
        ValidatorsCompacted = new bytes32[](count);
        ValidatorsPubKeys = new bytes32[](count);
        uint len = validatorsCompacted.length;
        uint idx = 0;
        for(uint i=0; i<len; ++i){
            bytes32 cv = validatorsCompacted[i];
            if((uint256(cv) & (0xff << 224)) == 0){ //PAUSE_NOT_PAUSED
                ValidatorsCompacted[idx] = cv;
                ValidatorsPubKeys[idx] = validatorsPubKeys[i];
                ++idx;
            }
        }
    }
    
    function getActiveCount() public view returns (uint count) {
        uint len = validatorsCompacted.length;
        for(uint i=0; i<len; ++i){
            bytes32 cv = validatorsCompacted[i];
            if((uint256(cv) & (0xff << 224)) == 0){ //PAUSE_NOT_PAUSED
                ++count;
            }
        }
    }

    function getValidator(bytes32 vPub) public view returns (address nodeAddr, uint96 deposit, uint48 pauseBlockNumber,
        uint8  pauseCause, address receiver) {

        Validator storage v = validators[vPub];

        nodeAddr = v.nodeAddr;
        deposit = v.deposit;
        pauseBlockNumber = v.pauseBlockNumber;
        pauseCause = v.pauseCause;
        receiver = v.receiver;
    }

}
