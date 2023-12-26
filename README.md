# Bitcoin BLockchain in Go (Kinda) 
Since this is first try at this implementation, I tried to approach at it with my own perpective but that won't mean it lacks basic functionalities.

### Some of the Key Features include Consensus Algorithm, Wallets, Database, Merkle Tree and many more.

## Functionalities and Design Choices : 
1. Difficulty of signed blocks remains constant unlike the real implementaiton.
2. Instead of deriving the public key using complicated methods like in the bitcoin blockchain, this implementation uses arbitrary keys.
3. Constant Reward with no fancy algo.
4. Since badgerDB only accepts array of bytes or slices of bytes, we have the serialize and deserialize.
5. Though in the original blockchain implementation each block has it's own separate file for efficiency, this project won't use that mainly because it's a small project
6. Rather than Badger's natve iterator, implemented my own.
7. Used Goexit() instead of OSexits so as to give time to badger to proerly garbage collects as the go routines are being shut down. this prevents data corruption.
