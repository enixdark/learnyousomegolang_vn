const Web3 = require('web3');
const solc = require('solc');
const fs = require('fs');

let code = fs.readFileSync('contracts/Voting.sol').toString();
let compiledCode = solc.compile(code);
let web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));

let byteCode = compiledCode.contracts[':Voting'].bytecode;
let VotingContract = web3.eth.contract(JSON.parse(compiledCode.contracts[':Voting'].interface));

console.log(compiledCode.contracts[':Voting'].interface);


console.log(compiledCode.contracts[':Voting'].bytecode);
// web3.personal.unlockAccount(web3.eth.accounts[0], "123456", 0);

// let deployedContract = VotingContract.new(
//     ['0x60aaaC3eF8ea20d41662FEa90875845DAC28E408', 
//      '0x3A0e072B0EE200ccaa7e154d84857718017283ad', 
//      '0xc2730544Ce488376DeBF5fd61e4537B5a2aC88fc'],
//     {
//         data: byteCode, 
//         from: web3.eth.accounts[0],
//         gas: 1000000
//     }
// );
// let contractInstance = VotingContract.at(deployedContract.address);
// // debugger
// contractInstance.voteForCandidate.sendTransaction("0x692a70d2e424a56d2c6c27aa97d1a86395877b3a")
