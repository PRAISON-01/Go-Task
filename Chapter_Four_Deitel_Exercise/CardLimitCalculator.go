package main

import "fmt"

func main() {

    for{

        var accountNumber int
        fmt.Print("\nEnter account number (enter -1 to quit) >> ")
        fmt.Scanf("%d\n,", &accountNumber)


        if accountNumber == -1{
            break
        }

        var beginningBalance int
        fmt.Print("Enter beginning Balance >> ")
        fmt.Scanf("%d\n", beginningBalance)


        var totalCharges int
        fmt.Print("Enter total charges >> ")
        fmt.Scanf("%d\n", &totalCharges)

        var totalCredits int
        fmt.Print("Enter total credits applied >> ")
        fmt.Scanf("%d\n", &totalCredits)

        var creditLimit int
        fmt.Print("Enter credit limit >> ")
        fmt.Scanf("%d\n", &creditLimit)

        newBalance := beginningBalance + totalCharges - totalCredits

        fmt.Printf("New Balance >> %d\n", newBalance

        if newBalance > creditLimit {
            fmt.Println("Credit limit exceeded")
        }

    )
    }
}
