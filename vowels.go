package main

import {
	"fmt"
}

func main() {
   	// Declaring a variable
   var sentence string

   // declaring the variables to store the count of vowels
   var vowelsCount int

   // initializing the variable sentence
   sentence = "Ek baar phir se"
   
   // initializing the variable vowelsCount
   vowelsCount = 0

   fmt.Println("Program to find the count of vowels within the function.")
   
   // iterating over the sentence
   for i := 0; i < len(sentence); i++ {
      // skipping the spaces in the sentence
      if sentence[i] == ' ' {
         continue
      }
   
      // comparing the current character with the vowels
      if sentence[i] == 'a' || sentence[i] == 'e' || sentence[i] == 'i' || sentence[i] == 'o' || sentence[i] == 'u' ||
         sentence[i] == 'A' || sentence[i] == 'E' || sentence[i] == 'I' || sentence[i] == 'O' || sentence[i] == 'U' {

            // if the current character is a vowel it's increase the count variable
            vowelsCount++
      } 
   }

   // printing the count of vowels
   fmt.Println("Sentence:- \n", sentence)

   // printing the count of vowels in the sentence
   fmt.Println("Result:- \nThe total number of vowels in the above sentence are", vowelsCount)
}