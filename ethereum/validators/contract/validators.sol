pragma solidity ^0.5.0;


/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
    address public owner;


    event OwnershipRenounced(address indexed previousOwner);
    event OwnershipTransferred(
        address indexed previousOwner,
        address indexed newOwner
    );


    /**
     * @dev The Ownable constructor sets the original `owner` of the contract to the sender
     * account.
     */
    constructor() public {
        owner = msg.sender;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    /**
     * @dev Allows the current owner to relinquish control of the contract.
     * @notice Renouncing to ownership will leave the contract without an owner.
     * It will not be possible to call the functions with the `onlyOwner`
     * modifier anymore.
     */
    function renounceOwnership() public onlyOwner {
        emit OwnershipRenounced(owner);
        owner = address(0);
    }

    /**
     * @dev Allows the current owner to transfer control of the contract to a newOwner.
     * @param _newOwner The address to transfer ownership to.
     */
    function transferOwnership(address _newOwner) public onlyOwner {
        _transferOwnership(_newOwner);
    }

    /**
     * @dev Transfers control of the contract to a newOwner.
     * @param _newOwner The address to transfer ownership to.
     */
    function _transferOwnership(address _newOwner) internal {
        require(_newOwner != address(0));
        emit OwnershipTransferred(owner, _newOwner);
        owner = _newOwner;
    }
}



contract Validators is Ownable {

    event ValidatorUpdated(
        bytes32 vPub,
        uint96 deposit
    );

    event ValidatorRemoved(
        bytes32 vPub,
        uint8 cause
    );

    struct Validator {
        address payable nodeAddr; //Address of the node that can sign transactions
        uint96 deposit;           //Node deposit
        uint48 pauseBlockNumber;  //Block number at which the node was paused
        uint8  pauseCause;        //The cause of pausement
        uint32 idx;               //The index of validator public key
        address payable receiver; //The address of the block reward receiver
        uint96 punishValue;       //The fine that should be paid to unpause this node
    }

    uint public c_MIN_DEPOSIT_INCREMENT = 10 ether;
    uint public c_MIN_DEPOSIT = 2500 ether;

    //Число блоков, на которые блокируется нода перед выводом депозита
    uint48 private constant DEPOSIT_LOCK_BLOCKS = 6000;
    //Число блоков, на которые блокируется нода, получившее предупреждение
    uint48 private constant NODE_LOCK_BLOCKS_MAX = 100e12; //Меньше 2^47

    uint8 private constant PAUSE_NOT_PAUSED = 0;      //Активное состояние
    uint8 private constant PAUSE_CAUSE_VOLUNTARILY = 1;      //Добровольная пауза
    uint8 private constant PAUSE_CAUSE_UNTIL_BLOCK = 2; //Пауза до блока
    uint8 private constant PAUSE_CAUSE_UNTIL_FINE = 3; //Пауза до выплаты штрафа
    uint8 private constant PAUSE_CAUSE_PUNISHMENT = 4;       //Наказание

    mapping(bytes32 => Validator) validators;
    //Array of validators public keys
    bytes32[] validatorsPubKeys;
    //Array of compacted validators info
    bytes32[] validatorsCompacted;
    //Flag who can punish - any node or just owner of the contract
    bool anyoneCanPunish = false;

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

    function pauseValidation(bytes32 vPub, bytes32 vFrom, uint8 cause, uint96 punishValue) public {
        require(!isPaused(vPub) && hasDeposit(vPub), "Node should not be paused and should have deposit");
        require(cause >= PAUSE_CAUSE_VOLUNTARILY, "You should specify cause");
        require(getNodeAddr(vFrom) == msg.sender, "Node should correctly pass its validator public key");
        require(cause == PAUSE_CAUSE_VOLUNTARILY || msg.sender == owner || anyoneCanPunish, "Wrong punisher");
        require(cause != PAUSE_CAUSE_UNTIL_BLOCK || punishValue < NODE_LOCK_BLOCKS_MAX, "Wrong blocks number");

        if(vPub != vFrom){
            require(hasDeposit(vFrom) && !isPaused(vFrom), "You should not be paused to pause others");
            require(cause != PAUSE_CAUSE_VOLUNTARILY, "This pausing is not voluntary!");
        }else{
            require(cause == PAUSE_CAUSE_VOLUNTARILY, "You are pausing voluntarily");
        }

        Validator storage v = validators[vPub];

        v.pauseBlockNumber = uint48(block.number);
        v.pauseCause = cause;
        v.punishValue = punishValue;

        compactAndSave(v);

        emit ValidatorRemoved(vPub, cause);
    }

    function resumeValidation(bytes32 vPub) public {
        require(isPaused(vPub), "This node is not paused!");

        Validator storage v = validators[vPub];
        require(v.nodeAddr == msg.sender, "You can only unpause yourself");
        uint8 pauseCause = v.pauseCause;
        require(pauseCause == PAUSE_CAUSE_VOLUNTARILY ||
                pauseCause == PAUSE_CAUSE_UNTIL_BLOCK ||
                pauseCause == PAUSE_CAUSE_UNTIL_FINE, "You can not unpause from this state");

        if(pauseCause == PAUSE_CAUSE_UNTIL_BLOCK)
            require(uint256(v.pauseBlockNumber) + uint256(v.punishValue) <= block.number, "You should wait some blocks before unpause");

        if(pauseCause == PAUSE_CAUSE_UNTIL_FINE){
            require(v.deposit >= v.punishValue, "Deposit is not enough to cover the fine");
            v.deposit -= v.punishValue;
        }

        require(v.deposit >= c_MIN_DEPOSIT, "You can not unpause before deposit exceeds min value");

        v.pauseBlockNumber = 0;
        v.pauseCause = 0;
        v.punishValue = 0;

        compactAndSave(v);

        emit ValidatorUpdated(vPub, v.deposit);
    }

    function enablePunishers(bool all) onlyOwner public {
        anyoneCanPunish = all;
    }

    function setDepositBounds(uint min_dep, uint min_dep_inc) onlyOwner public{
        c_MIN_DEPOSIT = min_dep;
        c_MIN_DEPOSIT_INCREMENT = min_dep_inc;
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
        require(msg.value >= c_MIN_DEPOSIT_INCREMENT, "Too small value to add to the deposit");

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
        uint8 pauseCause, uint96 punishValue, address receiver) {

        Validator storage v = validators[vPub];

        nodeAddr = v.nodeAddr;
        deposit = v.deposit;
        pauseBlockNumber = v.pauseBlockNumber;
        pauseCause = v.pauseCause;
        receiver = v.receiver;
        punishValue = v.punishValue;
    }

}
