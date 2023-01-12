package main

import "fmt"

func main() {
	// if-else
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	// switch case
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Very Good"
		score 10	=> "Execellent"
		otherwise 	=> "Invalid Score"
	*/
	/*
		score := 6
		switch score {
	*/
	/*
		switch score := 6; score {
		case 0:
			fmt.Printf("Score : %d, Terrible\n", score)
		case 1:
			fmt.Printf("Score : %d, Terrible\n", score)
		case 2:
			fmt.Printf("Score : %d, Terrible\n", score)
		case 3:
			fmt.Printf("Score : %d, Terrible\n", score)
		case 4:
			fmt.Printf("Score : %d, Not Bad\n", score)
		case 5:
			fmt.Printf("Score : %d, Not Bad\n", score)
		case 6:
			fmt.Printf("Score : %d, Not Bad\n", score)
		case 7:
			fmt.Printf("Score : %d, Not Bad\n", score)
		case 8:
			fmt.Printf("Score : %d, Very Good\n", score)
		case 9:
			fmt.Printf("Score : %d, Very Good\n", score)
		case 10:
			fmt.Printf("Score : %d, Excellent\n", score)
		default:
			fmt.Printf("Score : %d, Invalid Score\n", score)
		}
	*/

	/*
		switch score := 6; score {
		case 0, 1, 2, 3:
			fmt.Printf("Score : %d, Terrible\n", score)
		case 4, 5, 6, 7:
			fmt.Printf("Score : %d, Not Bad\n", score)
		case 8, 9:
			fmt.Printf("Score : %d, Very Good\n", score)
		case 10:
			fmt.Printf("Score : %d, Excellent\n", score)
		default:
			fmt.Printf("Score : %d, Invalid Score\n", score)
		}
	*/

	// simulating nested if else
	switch score := 6; {
	case score >= 0 && score <= 3:
		fmt.Printf("Score : %d, Terrible\n", score)
	case score <= 7:
		fmt.Printf("Score : %d, Not Bad\n", score)
	case score <= 9:
		fmt.Printf("Score : %d, Very Good\n", score)
	case score == 10:
		fmt.Printf("Score : %d, Excellent\n", score)
	default:
		fmt.Printf("Score : %d, Invalid Score\n", score)
	}

	// using fallthrough
	fmt.Println("using fallthrough")
	switch no := 5; no {
	case 0:
		fmt.Println("no == 0")
		fallthrough
	case 1:
		fmt.Println("no <= 1")
		fallthrough
	case 2:
		fmt.Println("no <= 2")
		fallthrough
	case 3:
		fmt.Println("no <= 3")
		fallthrough
	case 4:
		fmt.Println("no <= 4")
		fallthrough
	case 5:
		fmt.Println("no <= 5")
		fallthrough
	case 6:
		fmt.Println("no <= 6")
		fallthrough
	case 7:
		fmt.Println("no <= 7")
		fallthrough
	case 8:
		fmt.Println("no <= 8")
		fallthrough
	case 9:
		fmt.Println("no <= 9")
	}

	fmt.Println("fallthrough - applied")
	switch plan := "Pro"; plan {
	case "Super":
		fmt.Println("[Super] License for a family of 3")
		fallthrough
	case "Premium":
		fmt.Println("[Premium] Downloads for offline usage")
		fallthrough
	case "Pro":
		fmt.Println("[Pro] No ads")
		fmt.Println("[Pro] Custom play list")
		fallthrough
	case "Free":
		fmt.Println("[Free] Listen to music")
	}

	/* loop - for */
	fmt.Println("for - [v1.0]")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("for (while) - [v2.0]")
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println("numSum =", numSum)

	fmt.Println("for (infinte) - [v3.0]")
	x := 1
	for {
		x += x
		if x >= 100 {
			break
		}
	}
	fmt.Println("x =", x)

	fmt.Println("for - using labels")

OUTER_LOOP:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				fmt.Println("==================")
				continue OUTER_LOOP // break
			}
		}
	}

}
