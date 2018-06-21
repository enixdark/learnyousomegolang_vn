pragma solidity ^0.4.18;

contract owned {
    function owned() public { owner = msg.sender; }
    address owner;

    modifier onlyOwner {
        require(
            msg.sender == owner,
            "Only owner can call this function."
        );
        _;
    }
}


contract Voting is owned {
    mapping (bytes32 => int) public votes;
    bytes32[] public candidateList;
    
    function Voting(bytes32[] candidateNames) public {
        candidateList = candidateNames;
    }

    function voteForCandidate(bytes32 candidateName) public {
        require(validCandidate(candidateName));
        votes[candidateName] = votes[candidateName] + 1;
    }

    function totalVotesFor(bytes32 candidateName) view public returns(int) {
        require(validCandidate(candidateName));
        return votes[candidateName];
    }
    
    function validCandidate(bytes32 candidateName) view public returns (bool) {
        for(uint i = 0; i < candidateList.length; i++) {
          if (candidateList[i] == candidateName) {
            return true;
          }
        }
        return false;
    }
    
    function getCandidateList() public constant returns (bytes32[]) {
        return candidateList;
    }

    function kill() public onlyOwner {
        selfdestruct(owner);
    }
}