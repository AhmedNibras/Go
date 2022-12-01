package main

import (
    "fmt"
    "io/ioutil"
)

func main() {


	// Read the file
	word, err := ioutil.ReadFile("word.txt")
    if err != nil {
        panic(err)
    }
	// Convert the file to a string
    wordText := string(word)
   	

   // declaring the variables to store the count of vowels
   var vowelsCount int


   // initializing the variable vowelsCount
   vowelsCount = 0

   fmt.Println("Program to find the count of vowels within the function.")
   
   // iterating over the sentence
   for i := 0; i < len(word); i++ {
      // skipping the spaces in the sentence
      if word[i] == ' ' {
         continue
      }
   
      // comparing the current character with the vowels
      if word[i] == 'a' || word[i] == 'e' || word[i] == 'i' || word[i] == 'o' || word[i] == 'u' ||
         word[i] == 'A' || word[i] == 'E' || word[i] == 'I' || word[i] == 'O' || word[i] == 'U' {

      // if the current character is a vowel it's increase the count variable
            vowelsCount++
      } 
   }

   // printing the count of vowels
   fmt.Println("Sentence:- \n", wordText)

   // printing the count of vowels in the sentence
   fmt.Println("Result:- \nThe total number of vowels in the above sentence are", vowelsCount)
}