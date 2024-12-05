package main

import "fmt"

func main() {
        var votes [20]int
        var totalVotes, validVotes int

        fmt.Println("Masukkan data suara (akhiri dengan 0):")

        for {
                var vote int
                fmt.Scan(&vote)

                if vote == 0 {
                        break
                }

                if vote >= 1 && vote <= 20 {
                        validVotes++
                        votes[vote-1]++
                }
                totalVotes++
        }

        fmt.Println("Total suara:", totalVotes)
        fmt.Println("Suara sah:", validVotes)

        fmt.Println("Perolehan suara:")
        for i := 1; i <= 20; i++ {
                if votes[i-1] > 0 {
                        fmt.Printf("Calon %d: %d suara\n", i, votes[i-1])
                }
        }
}