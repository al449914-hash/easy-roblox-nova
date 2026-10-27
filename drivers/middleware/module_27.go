package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.5" )

func dckmO6P3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5fmhj7lhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3VY2MjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLWDjHWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyLu14mSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JdWo7QqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Ev1l6fVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZsLvjhYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TyUg8uN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAGF1m6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zOPsBp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otEcF7CJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qz0LR0qZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t3QI3Mr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBXLvpf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9qrSb3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4fGR3EXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JoPFFvTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUIKumwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OFHxbgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggCTfeoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vzbB31mdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayaPOubkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHNyEJ1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzXufHo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWlXJqbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c3Py26lkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6Oo9Xu0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 987CUvaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCXl2qPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPle540WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvWjRxK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XXwTB13ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iLGbA9mWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2l1OclKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WhFDHJ4fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQHR0T5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rwEDZAoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2akuOKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OeWZjSe1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIOBPrQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func La5DAoU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrSjW3rcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hq1MPakaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iwvs33UiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7JkBpbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SE2KWpjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7PxRGEAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpxapF4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7whaz1hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WE38CjiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHgJNvcZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pO0MaiErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSDajAfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LU7grPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A60e2Ma4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gi6tHnRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRkFcD4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QydIeA4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AX57mWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzZGpROUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KSXLq6xEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZahcx2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xA4GTy6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ca7To0vHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5yqAmVoxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K43OeH2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xoUNUZNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q65uWeV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yhh04H7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1uuiTjpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Xk724bhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xizSr1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOKmpU1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PurS6lmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7lr2q00FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5YLUCAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8KfSP9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygmFedtAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4uYb4zxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W34Ek5d6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6K1NOyG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clxyQLptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1DNdHyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPck1uObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UsuodjyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R7D34K0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdSW3NsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxhY9ljXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqZB8rXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QnDOCRIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZiuhnoC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 69Dv9RRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RW8raLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rG56RzqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtxQSrzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7guLsShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flKcdNGxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTQKhy8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6MtK6LKpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7jTrg9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uDoPn65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JR4DqRgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a2NQrMblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOgqh31PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9b0so6APWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWPl4wAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmTpVbweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9WVHZl1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OikMvp8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0HN1gxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6htWX13oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XExpnA4ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XChb4eoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nZnBvwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYiGFJzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uvoFNEf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7TeY6LIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjJYncIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UiDeB535Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PvbrarsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndBgjRSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjDKOHwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O9fUGihwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdCeOkO0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yzQfsQmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmugB3eUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5y2vRiJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PS9Y3B9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stxcUICuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func weeCAGa7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIfMlwkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

