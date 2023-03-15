// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	responses := make(chan string)
	go answerQuestions(questions, responses)
	go predictions(responses) 
	go printResponse(responses)
	
	return questions
}

func answerQuestions(questions chan string, responses chan string) {
	for question := range questions {
		go prophecy(question, responses)
	}
}

func predictions(predictions chan string) {
	for {
		RandomSleep(10)
		go prophecy("mosh mosh", predictions)
	}
}

func printResponse(responses <- chan string) {
	for response := range responses {
		fmt.Println(star)
		for _, char := range response {
			fmt.Print(string(char))
			time.Sleep(100 * time.Millisecond)
		}
		fmt.Println()
	}
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	RandomSleep(5)
	
	// Find the longest word.
	//TODO what is the point of this? 
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
		"Face is the place.",
		"Mace to the face.",
		"Face is a flat circle.",
		"Flat circle is life.",
		"Time is a flat circle.",
		"Life is a fart.",
		"Life is a face.",
		"Life is a flat circle.",
		"Life is a flat fart.",
		"Life is a flat face.",
		"Life is a flat moon.",
		"Better to be a fart than a face.",
		"Better to be a face than a fart.",
	}

	// Legit wise responses to questions.
	responses := []string{
		"Yes.",
		"No.",
		"Maybe.",
		"Probably.",
		"Probably not.",
		"Perhaps.",
		"Perhaps not.",
		"Perhaps you should ask again later.",
		"I don't know.",
		"I don't care.",
		"I don't understand.",
		"I don't think so.",
		"I don't think you should ask that.",
		"I should not answer that.",
		"I should think so",
		"Perhaps you should ask someone else.",
		"Perhaps you should ask yourself.",
		"Perhaps you should ask your mother.",
		"Perhaps you should ask your therapist.",
		"Perhaps you should ask your doctor.",
		"Perhaps you should ask your lawyer.",
		"Perhaps you should ask your priest.",
		"Perhaps you should ask your witch doctor.",
		"Perhaps you should ask a shaman.",
		"Perhaps you should ask a wizard.",
	}

	if strings.Fields(question) == nil {
		answer <- "You must ask a question."
	} else if strings.Contains(question, "mosh mosh") {
		RandomSleep(10)
		answer <- nonsense[rand.Intn(len(nonsense))]
	} else if strings.Contains(question, "What is the answer to life, the universe and everything?") {
		answer <- "This one's obvious. 42, dummy."
	} else { 
		answer <- responses[rand.Intn(len(responses))]
	}
}

// RandomSleep waits for x seconds, where x is a random number, 0 < x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Second)
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
