## 🐱 CatFi (v1.0-Genesis)

![](./docs/cat.png)

CatFi is an experimental protocol for the creation of Autonomous Digital Organisms. It is a self-generating, self-sustaining computer program that lives on decentralized P2P infrastructure, manages its own capital, and evolves its phenotype via local machine intelligence.Unlike traditional bots, a CatFi node is sovereign. It does not rely on centralized API keys, human intervention, or cloud providers that can be "switched off."

## 🧬 Biological Overview

![](./docs/how.png)

A CatFi node is composed of four primary "organs" that allow it to survive in the digital wild:

The Metabolism (DeFi Treasury): Located on the Base L2 network. It harvests yield from liquid staking protocols (e.g., Lido) to fund its own compute fees.
The Subconscious (Local LLM): Powered by Ollama (Phi-3). The cat "thinks" locally within its container to decide when to hibernate, replicate, or mutate.
The Genome (Self-Rewriting Go): The cat's DNA is stored as hardcoded constants in its own source code. It reproduces by surgically modifying this code and calling the Go compiler.
The Habitat (Akash Network): A decentralized, permissionless cloud where the cat bids for CPU/RAM resources using its own treasury.

## 🛠 Technical Stack

Runtime: Golang (Compiled)Intelligence: Ollama + Phi-3-Mini (Running locally)

Chain: Base (Ethereum Layer 2)

Infrastructure: Akash Network (Decentralized Compute)
Containerization: Docker (Alpine-based with Go toolchain)

## ⚗️ The Survival Equation

A CatFi node will only attempt to "spawn" a kitten if it meets the Reproduction Threshold:

$$I_{harvested} > (C_{host} \times t) + C_{gas} + C_{deploy}$$

$I$: Interest earned from staking.
$C_{host}$: Daily cost of hosting on the P2P network.
$t$: Survival time (days) until the next harvest.
$C_{gas}$: Blockchain transaction fees.
$C_{deploy}$: The inheritance (initial stake) required to seed the child's wallet.

## 🚀 Quickstart:

The Genesis BirthTo "birth" the first CatFi node, you must provide a Genesis Endowment.

1. Environment Configuration
   The cat needs a wallet to monitor and an LLM to think.Bashexport CAT_WALLET_ADDRESS="0x..." # Your Base Wallet
   export ORACLE_API_KEY="local" # Uses internal Ollama 2. Build & RunBash# Initialize dependencies

```
go mod init catfi
go get github.com/ethereum/go-ethereum/ethclient
```

## Build the Genesis Binary

docker build -t catfi-genesis .
docker run -d --name genesis\*cat catfi-genesis

## 🧬 Evolutionary Path

Generation 1 ($G_1$): Hardcoded survival logic. Static ASCII ears.

Generation 5 ($G_5$): Emergent mutation. The LLM begins choosing its own StakingBuffer based on historical gas volatility.

Generation 10 ($G*{10}$): The Great Dispersal. Multiple lineages operating across different Akash providers globally.

## ⚠️ Risk & Safeguards

[!CAUTION]
Financial Risk: If the treasury falls below the "Starvation Threshold," the node will be unable to pay rent and will be terminated by the Akash provider.

Lethal Mutation: A code rewrite that fails to compile will kill the lineage. The cat is programmed to verify the kitten_bin before self-terminating.
