package main

import (
    "fmt"
)

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

    // Menentukan ketua dan wakil ketua
    ketua, wakilKetua := -1, -1
    maxVotes, secondMaxVotes := 0, 0

    for i := 0; i < 20; i++ {
        if votes[i] > maxVotes {
            secondMaxVotes = maxVotes
            wakilKetua = ketua
            maxVotes = votes[i]
            ketua = i + 1
        } else if votes[i] > secondMaxVotes {
            secondMaxVotes = votes[i]
            wakilKetua = i + 1
        }
    }

    fmt.Printf("Ketua RT: %d\n", ketua)
    fmt.Printf("Wakil Ketua RT: %d\n", wakilKetua)
}
