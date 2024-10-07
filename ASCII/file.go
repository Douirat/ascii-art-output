package ascii_art

import (
	"fmt"
	"os"
)

// Declare a file object:
type File struct {
	Name    string
	Data    []byte
	Symbols []string
	Parsed  *Sentence
	Result  string
}

// Declare an other object of output:
type Output struct {
	Option, String, Banner string
}

// Instantiate a new output object:
func TheOutput() *Output {
	return new(Output)
}

// Instantiate a new file object:
func NewFile(name string) *File {
	new_file := new(File)
	new_file.Name = name
	data, err := os.ReadFile(new_file.Name)
	if err != nil {
		data = []byte{}
	}
	new_file.Data = data
	new_file.Symbols = new_file.Extract()
	new_file.Parsed = nil
	new_file.Result = ""
	return new_file
}

// file length validation {}
func (file *File) IsValidLength() bool {
	return len(file.Data) == 7463 || len(file.Data) == 6625 || len(file.Data) == 5558
}

// Create a method to extract the leters from my file:
func (file *File) Extract() []string {
	fragment := ""
	symbols := []string{}
	if !file.IsValidLength() {
		fmt.Println(len(file.Data))
		return []string{}
	}
	for _, bit := range file.Data {
		if bit == 13 {
			continue
		}
		fragment += string(bit)
		if bit == 10 {
			if fragment != "" {
				symbols = append(symbols, fragment)
				fragment = ""
			}
		}
	}
	return symbols
}

func (file *File) GetLetter(key uint) []string {
	hash := 855 - ((127 - key) * 9)
	return file.Symbols[hash : hash+9]
}

// split the user input based on the occurrence of the <\n>
func (file *File) SplitAndExec(in string) string {
	str := ""
	result := ""
	happend := false
	for i := 0; i < len(in); i++ {
		if i < len(in)-1 && rune(in[i]) == 92 && in[i+1] == 110 {
			happend = true
			continue
		}
		if happend {
			file.NewSentence(str)
			if str == "" {
				result += "\n"
			}
			result += file.Parsed.Output
			str = ""
			happend = false
			continue
		}
		str += string(rune(in[i]))
	}
	if !happend {
		file.NewSentence(str)
		if str != "" {
			if file.Parsed.Output != "" {
				result += file.Parsed.Output
			}
			str = ""
		}
	}
	file.Result = result
	return file.Result
}

// Validate the arguments:
// => Check the validation of the option:
func (output *Output) IsValidOption() bool {
	return output.Option[:len(output.Option)-len("--output=")-1] == "--output=" && output.Option[len(output.Option)-len(".txt"):] == ".txt"
}

// => Check the validation of the banner:
func (output *Output) IsValidBanner() bool {
	return output.Banner == "standard" || output.Banner == "shadow" || output.Banner == "thinkertoy"
}

// Handle the output:
func GenerateOutput(option, sentence, banner string) {
	fmt.Printf("the option is: %v the string is: %v and the banner is: %v\n", option, sentence, banner)
	output := TheOutput()
	output.Option = option
	output.String = sentence
	output.Banner = banner
	if output.IsValidOption() && output.IsValidBanner() {
		fmt.Println("The input is valid")
		file := NewFile("Symbols/" + output.Banner + ".txt")
		if len(file.Symbols) == 0 {
			fmt.Println("The file or data is missing!!!")
			return
		}
		file.SplitAndExec(output.String)
		fs, err := os.Create(output.Option[len("--output="):])
		fmt.Print(file.Result)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer fs.Close()
		fs.WriteString(file.Result)

	} else {
		fmt.Println("You should respect the formula: --output=<fileName.txt and the banner should be a valid file name: standard, shadow or thinkertoy")
	}
}
