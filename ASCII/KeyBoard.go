package ascii_art

import "fmt"

// Declare Sentence object as:
type Sentence struct {
	Input  [][]string
	Output string
}

// Instantiate a new Sentence:
func (file *File) NewSentence(input string) {
	sentence := new(Sentence)
	sentence.Input = make([][]string, len(input))
	for i, char := range input {
		if char < 32 || char > 126 {
			fmt.Println("Wrong input!!!")
			sentence.Input = [][]string{}
			break
		}
		sentence.Input[i] = file.GetLetter(uint(char))
	}
	sentence.Merge()
	file.Parsed = sentence
}

// Form a string from slices:
func (sentence *Sentence) Merge() {
	output := ""
	if len(sentence.Input) == 0 {
		sentence.Output = ""
		return
	}
	merged := make([]string, 9)
	for i := 0; i < len(sentence.Input); i++ {
		for j := 0; j < len(sentence.Input[i]); j++ {
			if merged[j] != "" {
				merged[j] = merged[j][:len(merged[j])-1] + sentence.Input[i][j]
			} else {
				merged[j] = sentence.Input[i][j]
			}
		}
	}

	for _, char := range merged {
		output += string(char)
	}
	sentence.Output = output[1:]
}
