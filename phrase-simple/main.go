package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/MaxHalford/eaopt"
)

type Phrase eaopt.IntSlice

func (p Phrase) Evaluate() (misses float64, err error) {
	for n := 0; n < len(p); n++ {
		if p[n] != targetPhrase[n] {
			misses++
		}
	}
	return
	//val := float64(matched) / float64(len(targetPhrase))
	//if val > 0.0 {
	//	fmt.Println(p, "fitness", val)
	//}
	//return -val, nil
	//return float64(matched) / float64(len(p)), nil
}

func (p Phrase) Mutate(rng *rand.Rand) {
	//eaopt.MutSpliceInt(p, rng)
	//fmt.Println("Before mutation", p)
	//eaopt.MutPermuteInt(p, 1, rng)
	//fmt.Println("After:         ", p)

	for i := 0; i < 2; i++ {
		var (
			element = rng.Intn(95) + 32
			pos     = rng.Intn(len(p))
		)
		p[pos] = element
	}
}

func (p Phrase) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	//eaopt.CrossCXInt(p, Y.(Phrase))
	//eaopt.CrossOXInt(p, Y.(Phrase), rng)
	//fmt.Println("Before X", p, Y.(Phrase))
	eaopt.CrossGNXInt(p, Y.(Phrase), 2, rng)
	//fmt.Println("After  X", p, Y.(Phrase))
}

func (p Phrase) Clone() eaopt.Genome {
	phrase := make(Phrase, len(p))
	copy(phrase, p)
	return phrase
}

func (p Phrase) String() string {
	var str strings.Builder
	for n := 0; n < len(p); n++ {
		str.WriteByte(byte(p[n]))
	}
	return str.String()
}

// ••••••

func PhraseFactory(rng *rand.Rand) eaopt.Genome {
	phrase := make(Phrase, len(targetPhrase))
	for n := 0; n < len(phrase); n++ {
		// generate a character between 32 (space) and 127
		phrase[n] = rng.Intn(95) + 32
	}
	return phrase
}

// ••••••

var targetPhrase Phrase

func main() {
	var phraseArg string
	flag.StringVar(&phraseArg, "phrase", "To be or not to be", "Phrase to solve")
	flag.Parse()

	targetPhrase = make(Phrase, len(phraseArg))
	for n := 0; n < len(phraseArg); n++ {
		targetPhrase[n] = int(phraseArg[n])
	}
	fmt.Printf("Solving for '%s'\n", targetPhrase)
	ga, err := eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		panic(err)
	}
	ga.NGenerations = 10000000
	now := time.Now()
	ga.Callback = func(ga *eaopt.GA) {
		if ga.Generations%100 == 0 {
			fmt.Printf("Best fitness at generation %d: %f (%s)\n", ga.Generations, ga.HallOfFame[0].Fitness, ga.HallOfFame[0].Genome.(Phrase))
		}
		if ga.HallOfFame[0].Fitness == 0.0 {
			fmt.Println("Found solution, generation count", ga.Generations)
			fmt.Println("Elapsed:", time.Since(now))
			os.Exit(0)
		}
		//}
	}
	//ga.ParallelEval = true
	err = ga.Minimize(PhraseFactory)
	if err != nil {
		panic(err)
	}
}
