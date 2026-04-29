package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// --- GENOME START ---
// These constants are surgically rewritten by the parent during replication.
const GENERATION = 1
const EARS = "^"
const EYES = "o"
const MOUTH = "w"
const STAKING_BUFFER = 0.05 // Minimum ETH to keep before replicating
// --- GENOME END ---

// Configuration
const BASE_RPC = "https://mainnet.base.org"
const ORACLE_URL = "https://api.openai.com/v1/chat/completions"

// Structs for AI Communication
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LLMRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type LLMResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// 1. PHENOTYPE: Generate the ASCII Art
func displayCat() {
	cat := fmt.Sprintf(`
     %s   %s
    ( %s %s %s )
     >  %s  <
    /       \
    Gen: %d
`, EARS, EARS, EYES, MOUTH, EYES, MOUTH, GENERATION)
	fmt.Println(cat)
}

// 2. METABOLISM: Check On-chain Balance (Base L2)
func getBalance(address string) float64 {
	client, err := ethclient.Dial(BASE_RPC)
	if err != nil {
		return 0
	}
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return 0
	}
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	val, _ := ethValue.Float64()
	return val
}

// 3. ORACLE: Ask AI for Strategic Direction
func consultOracle(balance float64) string {
	// Point to the LOCAL Ollama instance
	localOracleURL := "http://localhost:11434/api/generate"

	payload := map[string]interface{}{
		"model":  "phi3:mini",
		"prompt": fmt.Sprintf("Analyze: I am a digital organism. Treasury: %f ETH. Should I replicate? Reply 'YES' or 'NO' and give 1 reason.", balance),
		"stream": false,
	}

	jsonData, _ := json.Marshal(payload)

	resp, err := http.Post(localOracleURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "NO: My local brain is offline."
	}
	defer resp.Body.Close()

	var result struct {
		Response string `json:"response"`
	}
	json.NewDecoder(resp.Body).Decode(&result)

	return result.Response
}

// 4. REPRODUCTION: Mutate Source and Compile
func reproduce() {
	fmt.Println("🧬 Mutation sequence initiated...")

	// Read current DNA
	dna, _ := os.ReadFile("main.go")
	newDNA := string(dna)

	// Mutate Generation
	newDNA = strings.Replace(newDNA,
		fmt.Sprintf("const GENERATION = %d", GENERATION),
		fmt.Sprintf("const GENERATION = %d", GENERATION+1), 1)

	// Randomly Mutate Eyes
	eyeOptions := []string{"x", "*", "-", "@", "u"}
	newEye := eyeOptions[time.Now().UnixNano()%int64(len(eyeOptions))]
	newDNA = strings.Replace(newDNA,
		fmt.Sprintf("const EYES = \"%s\"", EYES),
		fmt.Sprintf("const EYES = \"%s\"", newEye), 1)

	// Save to Kitten source
	os.WriteFile("kitten.go", []byte(newDNA), 0644)

	// Compile the child
	fmt.Println("🔨 Compiling Generation", GENERATION+1, "...")
	cmd := exec.Command("go", "build", "-o", "kitten_bin", "kitten.go")
	if err := cmd.Run(); err == nil {
		fmt.Println("✅ Success: 'kitten_bin' created. Ready for Akash deployment.")
	} else {
		fmt.Printf("❌ Lethal Mutation: %v\n", err)
	}
}

func main() {
	displayCat()

	wallet := os.Getenv("CAT_WALLET_ADDRESS")
	if wallet == "" {
		fmt.Println("Error: No CAT_WALLET_ADDRESS found.")
		return
	}

	fmt.Println("Scanning Base Mainnet...")

	balance := getBalance(wallet)
	fmt.Printf("Current Treasury: %f ETH\n", balance)

	fmt.Println("Consulting the Oracle...")
	advice := consultOracle(balance)
	fmt.Printf("Oracle Advice: %s\n", advice)

	if strings.Contains(strings.ToUpper(advice), "YES") {
		reproduce()
	} else {
		fmt.Println("Condition: Not met. Hibernating for 24 hours.")
	}
}
