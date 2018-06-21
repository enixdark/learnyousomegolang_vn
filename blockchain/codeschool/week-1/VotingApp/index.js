(function() {
    var web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
    var interface = JSON.parse(`[{"constant":true,"inputs":[{"name":"","type":"bytes32"}],"name":"votes","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"candidateName","type":"bytes32"}],"name":"totalVotesFor","outputs":[{"name":"","type":"int256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"candidateName","type":"bytes32"}],"name":"validCandidate","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[],"name":"kill","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"","type":"uint256"}],"name":"candidateList","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"candidateName","type":"bytes32"}],"name":"voteForCandidate","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"getCandidateList","outputs":[{"name":"","type":"bytes32[]"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"candidateNames","type":"bytes32[]"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"}]`);
    var VotingContract = web3.eth.contract(interface);
    // var contractInstance = VotingContract.at(
    //     "0x0dcd2f752394c41875e259e00bb44fd505297caf"
    // );
    var byteCode = '608060405234801561001057600080fd5b5060405161064238038061064283398101806040528101908080518201929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060029080519060200190610089929190610090565b5050610108565b8280548282559060005260206000209081019282156100d2579160200282015b828111156100d15782518290600019169055916020019190600101906100b0565b5b5090506100df91906100e3565b5090565b61010591905b808211156101015760008160009055506001016100e9565b5090565b90565b61052b806101176000396000f300608060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632b38cd96146100885780632f265cf7146100cd578063392e66781461011257806341c0e1b51461015b578063b13c744b14610172578063cc9ab267146101bb578063fdbc4006146101ec575b600080fd5b34801561009457600080fd5b506100b76004803603810190808035600019169060200190929190505050610258565b6040518082815260200191505060405180910390f35b3480156100d957600080fd5b506100fc6004803603810190808035600019169060200190929190505050610270565b6040518082815260200191505060405180910390f35b34801561011e57600080fd5b5061014160048036038101908080356000191690602001909291905050506102a9565b604051808215151515815260200191505060405180910390f35b34801561016757600080fd5b50610170610308565b005b34801561017e57600080fd5b5061019d6004803603810190808035906020019092919050505061042c565b60405180826000191660001916815260200191505060405180910390f35b3480156101c757600080fd5b506101ea600480360381019080803560001916906020019092919050505061044f565b005b3480156101f857600080fd5b506102016104a3565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b83811015610244578082015181840152602081019050610229565b505050509050019250505060405180910390f35b60016020528060005260406000206000915090505481565b600061027b826102a9565b151561028657600080fd5b600160008360001916600019168152602001908152602001600020549050919050565b600080600090505b6002805490508110156102fd5782600019166002828154811015156102d257fe5b90600052602060002001546000191614156102f05760019150610302565b80806001019150506102b1565b600091505b50919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001807f4f6e6c79206f776e65722063616e2063616c6c20746869732066756e6374696f81526020017f6e2e00000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b60028181548110151561043b57fe5b906000526020600020016000915090505481565b610458816102a9565b151561046357600080fd5b6001806000836000191660001916815260200190815260200160002054016001600083600019166000191681526020019081526020016000208190555050565b606060028054806020026020016040519081016040528092919081815260200182805480156104f557602002820191906000526020600020905b815460001916815260200190600101908083116104dd575b50505050509050905600a165627a7a72305820bb355bf5cd28f684b3af10a8a77f1d346aa2038e5c6fdf03032508e1c63b9d770029'
    var deployedContract = VotingContract.new(['Rama','Nick','Jose'],{data: byteCode, from: web3.eth.accounts[0], gas: 4700000});
    var contractInstance = VotingContract.at(deployedContract.address);
    
    // Default account is used if you don't specify from in function call.
    web3.eth.defaultAccount = web3.eth.accounts[0];
    // web3.personal.unlockAccount(web3.eth.accounts[0], "123456", 0);
    
    var tableElem = document.getElementById("table-body");
    var candidateOptions = document.getElementById("candidate-options");
    var voteForm = document.getElementById("vote-form");
    
    function handleVoteForCandidate(evt) {
      var candidate = new FormData(evt.target).get("candidate");
      contractInstance.voteForCandidate(candidate, {from: web3.eth.accounts[0]}, function() {
        var votes = contractInstance.totalVotesFor.call(candidate);
    
        // Updates the vote element.
        document.getElementById("vote-" + candidate).innerText = votes;
      });
    }
    
    voteForm.addEventListener("submit", handleVoteForCandidate, false);
    
    function populateCandidates() {
      var candidateList = contractInstance.getCandidateList.call(); // call() is used for sync read only calls.
      candidateList.forEach((candidate) => {
        var candidateName = web3.toUtf8(candidate);
        var votes = contractInstance.totalVotesFor.call(candidate);
        
        // Creates a row element.
        var rowElem = document.createElement("tr");
    
        // Creates a cell element for the name.
        var nameCell = document.createElement("td");
        nameCell.innerText = candidateName;
        rowElem.appendChild(nameCell);
    
        // Creates a cell element for the votes.
        var voteCell = document.createElement("td");
        voteCell.id = "vote-" + candidate; 
        voteCell.innerText = votes;
        rowElem.appendChild(voteCell);
    
        // Adds the new row to the voting table.
        tableElem.appendChild(rowElem);
    
        // Creates an option for each candidate
        var candidateOption = document.createElement("option");
        candidateOption.value = candidate;
        candidateOption.innerText = candidateName;
        candidateOptions.appendChild(candidateOption);
      });
    }
    
    populateCandidates();
})();