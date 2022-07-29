![thumbnail](./resources/thumb.png)

# Golang Blochain 🔓
I wanted to get into Go development, but I wanted to start with something interesting so I could get up to speed while doing something useful. That's why I decided to try to implement blockchains in Go. The plan is to then use this knowledge to build something more useful.

## How does it work?

A blockchain is a structure that stores sequencial groups of data called blocks. These blocks have a header field, which has a timestamp, data (which, although not specified, can be of any kind as long as it can be stored as a byte slice), the hash of the previous block, and a value called "nonce", which can be changed to try to compute a valid hash for the block; and a Hash field, which stores the hash generated by the join of the fields of the header.
<br></br>
To create a new block, you generate the header and pass it through a hash function, starting with a value of 0 for the nonce. If the generated hash is smaller than the target, the hash is valid, otherwise it's not valid and u must regenerate the header with a different value of nonce and try again until you suceed. Therefore the target value determines how hard it is to generate new blocks.

![sample](resources/sample.png)

## Features

- Create Blocks
- Verify if Blocks are valid
- Generate valid blocks to add upon the blockchain
