package main

import (
	"fmt"
	"github.com/chrisport/go-lang-detector/langdet"
	"io/ioutil"
	"os"
)

/*
analyzed.json includes:
Arabic, English, French, German, Hebrew, Russian, Turkish
*/
func main() {

	detector := langdet.NewDefaultLanguages()
	testString := "Ich kümmere mich nicht wirklich über Sie"
	result := detector.GetClosestLanguage(testString)
	fmt.Println(result)
	probaly := detector.GetLanguages(testString)
	fmt.Println(probaly)
	language := langdet.Analyze(GetTextFromFile("data/vietnamese.txt"), "vietnamese")
	detector.AddLanguage(language)
	testString = "tôi tội lỗi"
	result = detector.GetClosestLanguage(testString)
	fmt.Println(result)
	probaly = detector.GetLanguages(testString)
	fmt.Println(probaly)
}

// GetTextFromFile returns the content of file (identified by given fileName) as text
func GetTextFromFile(fileName string) string {
	text, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(text)
}

// WriteToFile writes a content into a file with specified name
func WriteToFile(content []byte, fileName string) {
	err := ioutil.WriteFile(fileName, content, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
