package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"
)

func trainModel(wordMap map[string][]string, words []string) {
	start := time.Now()
	for i := 0; i < len(words)-2; i++ {
		key := words[i] + " " + words[i+1]
		wordMap[key] = append(wordMap[key], string(words[i+2]))
	}
	jsonData, err := json.Marshal(wordMap)
	if err != nil {
		fmt.Println("Model saving failed.")
	}
	err = os.WriteFile("model.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Model saving failed.")
	}
	fmt.Println("Training took:", time.Since(start))
}
func runModel(wordMap map[string][]string, outputLength int, words []string) {
	index := rand.IntN(len(words) - 2)
	word1 := words[index]
	for string(word1[len(word1)-1]) != "." {
		index = rand.IntN(len(words) - 2)
		word1 = words[index]
	}
	index += 1
	word1 = words[index]
	word2 := words[index+1]
	fmt.Print(word1 + " " + word2 + " ")
	for i := 0; i < outputLength; i++ {
		followingWords := wordMap[word1+" "+word2]
		if len(followingWords) == 0 {
			break
		}
		nextWord := followingWords[rand.IntN(len(followingWords))]
		fmt.Print(nextWord + " ")
		word1 = word2
		word2 = nextWord
		time.Sleep(50 * time.Millisecond)
	}
}
func main() {
	trainingData, err := os.ReadFile("training.txt")
	outputLength := 20
	if err != nil {
		fmt.Println("Failed to read training file.", err)
	}
	words := strings.Fields(string(trainingData))
	//Make the map!
	wordMap := make(map[string][]string)
	mapFil, err := os.ReadFile("model.json")
	if err != nil {
		fmt.Println("No trained data found, training now!")
		trainModel(wordMap, words)
	} else {
		json.Unmarshal(mapFil, &wordMap)
	}
	runModel(wordMap, outputLength, words)
}
