package main

import "fmt"

func main() {
	for count := 1; count <= 12; count++ {
		switch count {
		case 1:
			fmt.Println("On the first Day of christmas my true love sent to me....\n\n")
			nextVerses(count)

		case 2:
			fmt.Println("\n\nOn the second day of christmas my tru love send to me...\n\n")
			nextVerses(count)

		case 3:
			fmt.Println("\n\nOn the third Day of christmas my true love sent to me...\n\n")
			nextVerses(count)

		case 4:
			fmt.Println("\n\nOn the fourth Day of christmas my true love sent to me...\n\n")
			nextVerses(count)

		case 5:
			fmt.Println("\n\nOn the fifth Day of christmas my true love sent to me...\n\n")
			nextVerses(count)

		case 6:
			fmt.Println("\n\nOn the sixth Day of christmas my true love sent to me...\n\n")
			nextVerses(count)

		case 7:
			fmt.Println("\n\nOn the seventh Day of christmas my true love sent to me...\n\n")
			nextVerses(count)

		case 8:
			fmt.Println("\n\nOn the eighth Day of christmas my true love sent to me...\n\n")
			nextVerses(count)

		case 9:
			fmt.Println("\n\nOn the ninth Day of christmas my true love sent to me...\n\n")
			nextVerses(count)

		case 10:
			fmt.Println("\n\nOn the tenth Day of christmas my true love sent to me\n\n")
			nextVerses(count)

		case 11:
			fmt.Println("\n\nOn the eleventh Day of christmas my true love sent to me\n\n")
			nextVerses(count)

		case 12:
			fmt.Println("\n\nOn the twelfth Day of christmas my true love sent to me\n\n")
			nextVerses(count)

		}
	}
}

func nextVerses(count int) {
	switch count {
	case 12:
		print("twelve drummers drumming")
		fallthrough
	case 11:
		print("eleven pipers piping")
		fallthrough

	case 10:
		print("ten lords a-leaping")
		fallthrough

	case 9:
		print("nine ladies dancing")
		fallthrough

	case 8:
		print("eight maids a-milking")
		fallthrough

	case 7:
		print("seven swans a-swimming")
		fallthrough

	case 6:
		print("six geese a-laying")
		fallthrough

	case 5:
		print("five gold rings")
		fallthrough

	case 4:
		print("Four Colly birds")
		fallthrough

	case 3:
		print("Three French hens")
		fallthrough

	case 2:
		print("Two Turtle doves")
		fallthrough

	case 1:
		print("And a patridge in a pear tree!")

	}
}

func print(message string) {
	fmt.Println(message)

}
