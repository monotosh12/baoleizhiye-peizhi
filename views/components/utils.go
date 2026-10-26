package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.2" )

func UcOnR2dCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLNVOkrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbHEp838Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCcWqTuEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ExVNgptKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bm4SMfvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PeAN0iZ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NhcBdZ0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PT9lCXE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0SWSSQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GMs7vMUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IeQgOFQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 154DiXdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qRYuFgXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 10GhlLDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYyf45dYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wPzwh99RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tAHH9yJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3YyvZH1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3knFW9ahWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func unNKWm0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9fRXkJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYu4e4beWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tjLgpypsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WhVQHM3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63odUJeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKn2N5F2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A4uaNbIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func afM6tiTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1zWj90eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p68yxZIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BSBdZ0RUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zbJ6qkNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZfTHsBCIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdHgY0C2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0b5y2aKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2tJvaGpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFpIZ53pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDIuN1Z2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cC8YjtOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fn0xZhwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DzXr0E96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IsbXjvOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkwX3Kf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T5fBT8FmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6e5IGgU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWK0bJqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mf9jT4jiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jG0mJ8s6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbAEJmunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0559VNzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func okf4ARkvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qOOOcZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1eL09OmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k7K2MrXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYJC3HguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w52EANruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qp4RdGXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7SNe6OrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q7QyeTQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqH3equGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MHLiIXgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRRfgflxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcqHejmMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAzvdy9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9pDMqktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toVcKuynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1IRyeax2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpz9h1GTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ghPCmxMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S097eT5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OHQUg66QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gtJad2hIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rALMTFq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AWkB4f3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqs4WoEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1X7ZWxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqBztUxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCgzCO7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xkbPEvHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sKTVKa7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JhwZqYd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gGp836NoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrqGtfyAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q1Zxkpi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BDUEsYUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNxOvL22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGMMWjivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jj9m0kjjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qoHYv5dNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tzp0JBj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqpAUrdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LTBhhZRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLvHwVxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jd9FlQfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xquJgpq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tAmCnMC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jn5fM5VjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwEPiR0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bn6wiRDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CltkxOruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0a2tcAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ui6EsiLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zIypSK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uae0rryXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qtdhuorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsQwy1LxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QY2Nzs0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ah82APGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cws4OvZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWwpB8bDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6F75pxviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXHg7x7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pu8xvoXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2hsQOnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdUnTsiEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fY1QAZG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wTWAbWwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2gLnPENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6CfAg3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdL54HiwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NS18Xc49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkxiCj2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hcnnmx5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67j9vsnhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xEtmoqAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xyjq461JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTeYybgiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1yeYnvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwuBHeTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TRcwwdD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUh1k1uEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtoBuJ89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hkcUXlRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pJ8qoCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEyJcL2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpfjmJ3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5O2LCdmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKSv0zOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func me7SsrWLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xP4PIqbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func olkAZbRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jlnqp7phWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSjqwajQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KaTfKl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gp144kWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q40hknQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eoLE3NpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4RoqaVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I54RZ8LjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgfUV58gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w5I8sAbkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lWWmHajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Dl7rSpQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTfyKC91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFSA9P1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mMkdG2VrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTLUZGT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWdn7DxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1unN50oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AUXAj7YKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrWUQ0RkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXOpX92CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbHgCzgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSkAsz9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3Hbr42xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V32xkLOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktA6tBBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IOiaMHraWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFDNRx55Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AAfxUJcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AjmkMfifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BlOfvL7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y5arcqlfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbOYJSqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rwDgyrZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rkp95RlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rax8EFbiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSNoG8b5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbv2xRS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdNSg2VuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2c7W7eoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0sXG1McWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Zl7jCVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBTpAusMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgBKzI6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EhLJDPFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TJmI0XjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9VSNwmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fmPA0QUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GOrj6u62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jHMmcaaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXHWXLqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jIvP1QINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V3WJsSPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmGC2JwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9egeT9tXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtOFSWDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1mW1WKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Or5DhHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtOWWv8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwYscS9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hd2OidC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DJSygntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EjRNE6G0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vm3OJBcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mj9Lpr1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oK5Z1B9eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFmcCsPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VkEfYokgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cbFcFhdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qWmp8D62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zHFyS75sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcrkFucgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FrrVbTlXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvYLmex0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdCVnHnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIlggLhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ADy0g20tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeRH81mFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B7osyw2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFSvIdeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

