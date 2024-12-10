# Latihan modul 11
Penjelasan singkat pada tiap latihan soal modul 11

## Setiap soal
Pada setiap soal saya menggunakan beberapa komponen :
- `Package main` agar program dapat dieksekusi
- `import fmt` agar dapat menggunakan beberapa operasi dasar bahasa program `go`
- `main()` sebagai titik awal program dan dieksekusi ketika program dijalankan

## Soal 1
```go
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
```
Operasi diatas merupakan operasi menginput suara. Jika inputan bernilai 0 maka proses inputan suara akan berakhir. Jika inputan bernilai 1 sampai 20 maka akan dimasukkan ke suara yang valid

```go

        fmt.Println("Perolehan suara:")
        for i := 1; i <= 20; i++ {
                if votes[i-1] > 0 {
                        fmt.Printf("Calon %d: %d suara\n", i, votes[i-1])
                }
        }
```
Operasi diatas merupakan operasi memperlihatkan Perolehan suara untuk setiap calon.

## Soal 2
```go
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
```
Operasi diatas merupakan operasi menginput suara. Jika inputan bernilai 0 maka proses inputan suara akan berakhir. Jika inputan bernilai 1 sampai 20 maka akan dimasukkan ke suara yang valid

```go
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
```
Operasi diatas merupakan operasi Menentukan Ketua dan Wakil Ketua. Jika jumlah suara calon lebih besar dari maxVotes, maka calon tersebut menjadi ketua baru, dan nilai maxVotes dan secondMaxVotes diperbarui. Jika jumlah suara calon lebih besar dari secondMaxVotes tetapi tidak lebih besar dari maxVotes, maka calon tersebut menjadi wakil ketua baru, dan nilai secondMaxVotes diperbarui.

## Soal 3
```go
func isiArray(n int) {
        for i := 0; i < n; i++ {
                fmt.Scan(&data[i])
        }
}
```
Operasi diatas merupakan operasi menginputkan nilai pada tiap elemen array sampai elemen array ke-n
```go
func posisi(n, k int) int {
        left, right := 0, n-1
        for left <= right {
                mid := (left + right) / 2
                if data[mid] == k {
                        return mid
                } else if data[mid] < k {
                        left = mid + 1
                } else {
                        right = mid - 1
                }
        }
        return -1
}
```
Operasi diatas merupakan operasi Pencarian indeks array yang memiliki nilai dicari oleh user menggunakan binary search.
