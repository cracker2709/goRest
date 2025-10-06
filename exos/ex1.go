package exos

import (
	"log"
)

func moyenne(numbers ... float32) {
	var somme float32 = 0
	var nb_numbers float32 = float32(len(numbers))
	log.Printf("Array size is %v", nb_numbers)
	for _, n := range numbers {
        somme += n
    }
    average := somme / nb_numbers
    log.Printf("Calculated average is %v", average)

}

func main() {
	moyenne(10,20,15.5,2.36,4.56,19.52)}