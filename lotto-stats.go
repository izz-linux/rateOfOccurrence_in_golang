package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// file IO error check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// main function with all the guts...taken from a few functions I previously wrote in python3
func main () {
	// an array with index of 1 - 69 [index + 1] is our 5 numbers
	var normalNums [69]int
	// an array with index of 1 - 26 is our powerball number
	var powerballNums [26]int
	// get path to lotto history file
	fmt.Println("Please enter the full (or relative) path to the powerball history: ")
	var lottoFilePath string
	fmt.Scanln(&lottoFilePath)

	// pull our powerball history into memory and check for errors
	file, err := os.Open(lottoFilePath)
	check(err)

	// read lottory numbers from file into scanner
	scanner := bufio.NewScanner(file)
	// split on all whitespace to get each number
	scanner.Split(bufio.ScanWords)

	// will be used to check for normal lottery numbers and powerball numbers in the draw
	i := 0
	// the number of draws we have history for
	count := 0

	// scan through the text, determine which number set it is and increment the appropriate array
	for scanner.Scan() {
		// get next "word" / number as string
		lottoNumAsString := scanner.Text()
		// convert the string to integer
		lottoNum, _ := strconv.Atoi(lottoNumAsString)
		// up to here we have a single lotto number in the int variable lottoNum

		// increment i priod to following if / else conditional
		i++
		// this is not the powerball for current drawing
		if !((i % 6) == 0) {
			// since the range of i begins at 1, we decrement the index to increment, same goes for the powerball
			normalNums[lottoNum-1]++
		} else {
			powerballNums[lottoNum-1]++
			// we've reached the end of this "draw" and therefore want to reset i
			i = 0
			// increment the counter for total number of drawings
			count++
		}

	}

	// close the file since we're done using it
	file.Close()

	// print out the results
	fmt.Printf("\n\nThe percentage of each number being chosen over the last %v times is\n", count)
	pad := false
	for i := 0; i < 69; i++ {
		// pad digits 1-9 (indices 0 - 8)
		if (i < 9) {
			pad = true
		}

		// round the current value's probability to the nearest 1000th
		s := fmt.Sprintf("%.3f", float64(normalNums[i])*100/float64(count))

		// convert to string for possible padding and go ahead and assign the next number in the array of 0-68
		lottoNum := strconv.Itoa(i+1)

		// pad the string and reset pad to false
		if pad {
			lottoNum = "0" + lottoNum
			pad = false
		}

		fmt.Printf("%v:\t%v%%\n", lottoNum, s)
	}

	fmt.Printf("\n\nThe percentage of each number being chosen as the powerball over the last %v times is\n",
		count)

	for i := 0; i < 26; i++ {
		// pad digits 1-9 (indices 0 - 8)
		if (i < 9) {
			pad = true
		}

		// round the current value's probability to the nearest 1000th
		s := fmt.Sprintf("%.3f", float64(powerballNums[i])*100/float64(count))

		// convert to string for possible padding and go ahead and assign the next number in the array of 0-68
		lottoNum := strconv.Itoa(i+1)

		// pad the string and reset pad to false
		if pad {
			lottoNum = "0" + lottoNum
			pad = false
		}

		fmt.Printf("%v:\t%v%%\n", lottoNum, s)
	}

}
