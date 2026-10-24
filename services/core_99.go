package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.7" )

func htQ5vG6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5LgMp0MZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4C9WxSbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6X0KBSNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35Taqym8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9PkkIxAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yAUt2nyIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRIUqsy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzXRVDlaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZ9CJHDUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXTZQ8A9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func itgysGTiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzjVAp7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNqOaxwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fb71gujrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wSnplvH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6b6pmz2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zJbU27LbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCVlgKnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1aPsr0ebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fGUihPcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INt7VWEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lIlfEblYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvV8PmUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YkZPaFvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIGbDurYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgp96COJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bAWYdYhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLBWieQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Htj1TRiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70VHVyBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDXUeZozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1bQP8sqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZNGwPoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32JCbzZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0uK2hkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiRAPVuRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsivIN4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJEWDVekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZxyGOgafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jiBmo5qBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVYOIn2gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8hJKM5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e5QGQe0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sg29zomDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WkkVXHPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5nUR8olWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nanEsuHAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZhyAefeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ko4FdCEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ldDuOW39Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUStkqOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIiodq7MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3OkD2YGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rPfPn6AxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5drR8lKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isYJI7NGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dH0xrPUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRNsGCRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TF2V3XiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ci9kkgRKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YiZqJ75NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPQBGNH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZN9SBm6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bH6bDsFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Af4XePmbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9fNLMVbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHFqBs4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcIu1nJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dkNrasTbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43Eky6DAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7S2G8MwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VW0iepBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEkGBkRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2CIlPr9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ADnEmxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFXfT0u2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z3SAQzdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXReVrJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CydHzZKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSVp2YcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iHNtHN0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1KmH8eQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3UYLrr4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x5OhkombWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1wL2olbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7VMKGc85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPXaJFYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vwgo7Yk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wlQgPkLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGoNQ3FRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bNHgX0jgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6EIHd9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGNZM2gdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7AgF3U9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8F2g2yrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WVB5eA8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VNWER7g2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KWFEIOn7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wU2t4sInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqrLs5gFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SFdtJneWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzA4aK1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KPJgiY9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sVAaWxCwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZroLKIlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jy7HS3AWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dIPdQVS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func smAJpS3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tq1uuGHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r9qAHEnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teXiTQ9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TiuNqp3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5afkL5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zX32BJMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 83jwLVjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SomcPvBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUzLIfMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fqMufi8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ivad6bFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YGKoryWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DPXMMiIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drujiARhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L731QgSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FgH1yLPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQA592XQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WqW65mkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuI2BNm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClwrSB2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCmQ0yGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WtEd9QycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lmohmt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDvaBQVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IvumOMCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0p4ydkdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DBJUxJYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXOlDZxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1mSaCV7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mntrhBOTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DE1faJzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijB7VP4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8eU1xuaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UDaxpjDyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSshuAo1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bVMqs3hWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 11AYa2xxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6LfKiY3iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wr6LaOkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ioxsw8z4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efLDopKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 71tWdRkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7tADfXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNHEk18TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUKEra9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ioyySEnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BJ730pEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wFNIiflEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03ZSWJvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0dg9fjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLN5OzobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5FhWHDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDItYEFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRgp8uoZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGpCOMpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnsjCpWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPuRoydLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0KI3OPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVFENEsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9EpvUqvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxEpFBH2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0OnZsHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QqiQ32qSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aiPj16JZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZqSgWddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6BAlyicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgTgIY01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xK6MAlgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CTmOP2RyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPG4upk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nE2EPOqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgRqjaatWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSrRLgzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HXKdpuwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eMlFgdXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KVI4x2THWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAEGIrfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LkdIYtXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sraIbcibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RInyVwyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8oM2DzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XXBK4vPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUX3JVfQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zvPKwjziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EohbtNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JO2gYNBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mBEiyya6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cBE2HHO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSOKcLjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b9XRk9rCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SMhrcu5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S66VhwSFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsIMC4iSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IrrY8bv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dRrt8WRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKifRdpgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHpLTZM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgPgQm5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0MHr6aTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpKH2rAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsOVKX6VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7V03Kt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lgtyx61CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xE7fGtUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJ2kSCn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3vfsfmRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ibnZsD5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOr44QKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OaGFoiMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 340gaf8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0U0M1azDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XoewEVhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZbXwUbwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdCtrhbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7iVMSvkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7T4euSEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6OSJFa1pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6aSB0wwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkOwbGRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSozcCF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUbLq8dmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBsAfQFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VcqG2Z4jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9yJeQzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NL949FETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vunQu1dDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdXQupjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cW4fnTmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEEZm0WFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BfgZI09UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USdOPiOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1I7LblVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mk2uDWKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoUvM58FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xiUComoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KGjpeHCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LXiP5jrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DLea85WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4I8WY3RjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QQzMf5wkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZJGFm9QHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMxgbojCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjq12YfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YQa7pTFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8iww2QnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66257YJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZS7KZ7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5JJBCauAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FH11bdBQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12HA54Z2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r09NKw8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgRKE6SNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPq3HYujWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eDugOwb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LrwyuQJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7p9tqIPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUldpWyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TiEzuciqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wDhJkzfdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evuEGqS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 628WZL1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jL5VbgvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8In3HQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func otX1nYeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5KfBvrJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TPacHDvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LJwGcMpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IWGKHYoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hURbqTeRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tfwu8ksGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmF7L7QRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func viG0PyYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F6aiBDs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ylty08DMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mtRS9wNfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QTduPc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0URSZ9DlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRTkti3tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7jTsxbAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 81Q6UoKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tes2soZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjcJL8GjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Gz6u6sSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOeTrHWkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdLqPBFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3UtoRgtcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YqBTw2DVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUZciMUZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eidDbK3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oaomyoF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ge2DYEYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u03MKIzDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpvGXXlDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nXFk1XuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2f2xqIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sfsfFA0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHTMHSVnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rC1SnczuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLenag5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b0G2N80xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nCayt1sNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3SteuhDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UO5rMp4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2yqlNNvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyJ8VgSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmSACMf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I5RzROSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PX6cECttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ye2RBkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eFl8jra4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cz7BPS0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4WglGj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u1w4DoxEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 55f7155nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtwkinTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nivF75YbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jk3DzB9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XqyitlOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a3yRYuHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func evUaValUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vopQN5LjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVud89PfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1FT1oXE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlkFb2EvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyAfSRLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func akNdJmQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func roFRtR3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OOkRRHLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wgLvyI0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uy6iJ1ZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pll4s7C5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxfOQPutWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HU4tfzwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RlCBGg5fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8ey9powWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBCIGCriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thRHwLTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MY9RlsLuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tuvp7kGEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ragq0mhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lh1DJYxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XSuZYpXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A70px7wYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6Lr93mOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y12xJlzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JUPRlAbKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7gJ8HTglWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKaI7iLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ASFWgS6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Gxt8w0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GBkDKtFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nvIYiR0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QlnuFPHoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jTUaVKyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PLd0jeTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Kn30Ch9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tesQjxWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cnd6Ck0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWaqJ6ffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KOH8fJIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgudTZabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DG72DY1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TD1cxi1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AmwgoRCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sm7kHlEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZzyb9JSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQPgDLl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HFOHZEzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4AKVfNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4g4OdpsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dygxmCbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WJcKPSDgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8B2z9WfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9n106NRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MO1ojFgEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MpDAwjQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CfzZBaTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdz1CxgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6i1JTGMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZ1pbtbfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwWGRpFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7X0gfyAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SLMYfX1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N97O6RYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsX6Rs60Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lFWQz2wBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dnT3ZOcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 87SSz7urWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtzYMm0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QDOckU0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MV6goirbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUbhQZtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GrCEZhLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4TZTtCgCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMkLNdPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQUBK4pXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PWhIzU2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DH0A7hDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4TqHp3laWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ojtnn81zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FejiseoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTe26gd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3yU77XaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHbTwZy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OFO4W8m0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBJjx98YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdECWVUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOYX6DZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjgSDH8lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFwEmXA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y5tR6mWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e83ssj8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhil49VQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7JqQJ92OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 10K10A1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PV7cDLdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ftIuXabTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3IHboAqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlCqSc6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9L4alWYdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O63wlDowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8476OgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9CxEyodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpcMjOzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4Toe8CLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oD8TOl5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSUrvIsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b5VECdsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VO76vVfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIlhjfgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySpdU7o5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z26UGZYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npcxMO4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWgcxKtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zexMiJXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxjDXJ47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GupW12CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CT2FO6PTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PXglAREEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wNaadBFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrN1DsmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4DHqXKg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o0eZ8TucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jUkCaUGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cNI8NS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFSiOz2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LB16Ch9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awbVU9BQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func foSELa6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7xfB3DdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TftUUNeYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZWsf8JuVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PvC6ayTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBeUSOvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkwh6fv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2OEZzUWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOmoevGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTLwczAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8cCUHegAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAvT9rifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6GQlv5LfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5nikkWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c54k8LyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HPxKOl8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RQgBfdYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTz1ITFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTKRQuAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2sodJcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JkJLCpOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KVsMRjQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wIuhQFlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NfAm4eKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqOgyZA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wrWdAW08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrBpXdCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbkJr8brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SomsTxu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njGHzlYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqTHGKGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDdhPPazWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0E0m45ssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNrYr6CcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSRSudYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXPNbv5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ZKxny4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnRqZYS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhnRGuYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKM2pNXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQ0ii204Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1zR7mzJpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func So6YpEMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24ZZPIXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7a7NBcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IpbxiAJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjOfbobcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoDX9vFzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r2YbLqCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hbMvSlYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajJQZZcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lm5BfvyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZjLLUpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ul6t1qSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6UOrRA5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8mslV3nwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzSOFl5QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PrATRBquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHuwThjYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v9V3u2uRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8sT1leUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBbcLY2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dxlOSKcYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y1m95115Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Uk6H80IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCJ9aFquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwkgAW5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWB7YxMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JneZY7w2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljSwmdJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LMVGAFkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func scOqp4eCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AOFLJjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nx98qwPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cgbo0MuyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DIQQTZcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQC8qZAhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3SdZLr3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apFOv8pGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXTVLmdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1EBUT1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjFnbuTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEwp96RXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3G2ku9nNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aI2hW7WuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ChxT9foxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWKcuMamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8uc2vTKYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBEAYyZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RSXKAvPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtYa7fg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHc3DYIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YFOLjiD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sO5sWOK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67q0UwFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PnjkZVjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPabROgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZQfr1iCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdI3nn6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wUrneYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w5g6hr2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ecmAOxaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hLizmfDfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7zo0VKDwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P08FuiLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func llxVd3wrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Coul9qRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykcjsvcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M38vitBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yi1bbXsvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUNPt7O1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YaXtQ8v7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQ2loW8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZreFo5OsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OD9C82diWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q85LwmvnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HiuywoPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKzbDIBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sEncMWNjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTJTHXqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IneQjRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0nqS7t5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c9a4kjrzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaC4RzWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h4hJFy6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vvqJYNCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yHhpPHrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MePbIpicWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zeiBZM3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGbKNYMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40iR82XuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wAQkK2tWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dC1l85YvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hRdwJFkIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P28LaO14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7BxmeO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvjJiLk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func URL4ph3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6pTVI6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sZYiDOiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KRCnP2URWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2xSBKm8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0u6y6x2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8sioC6g3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRHqnqjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBw792E8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rmu9YmksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJiq2VjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJRVklymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqWT9lRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3lJUZyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jtF3VDkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KEYqxmSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RxdFq17QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fLQtmLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhCfk1MdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UM7FGru2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbUWmtY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MGuq1hDkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZzcQEft5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7ex5XfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJcVWTK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KE4Cb9D8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1VWWNrTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGhtX0QDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPLETvCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SobqBByWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ey7nxxb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2ivSzGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7Btg9ZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjC1PgGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kc7GraEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsgqbedlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JftULwlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2E4nPHnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3B795BoRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FcgKnMSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57AUB5IUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWsAaSjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lNgJXlOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNqYGVL2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qyd74xyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ryCDSigBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSYRVmlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RoIhFV6rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecqOtdnqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjziXuIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AueWKdyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rYnwhcOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFQS0kxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UYGwjMRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFjtwpedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQJNtisGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPcvmsQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VONcf7cWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EPLWw5AKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dlDOngcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfSSce0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYJQcMSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xW9Wp1b0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZpgJwFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljmPyWRmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zDe33mmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F6fsZAKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xTJdLfNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GY7GsM1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ksdVshzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3JMB9TQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dVjzA0FMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jzxKfqntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fgvXJYYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSxmtKFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s6JNHwm0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCIV6jy0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t4FC1e0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9dWYf6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5JfLaWgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Utq4eHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNw2QUAKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyhXEGrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89Mhgso7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DLAwzzClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xaVjJTZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSbSj2GHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPIWqVHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VP3MYXGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxPFAe2EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZqRWDD2yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 794DyojgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KA7ZHK6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mttmswK7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQfmMb5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func maY8WAodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JT1fsQqpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lASQKndNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjoQyKgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fn2oeEYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mnjyitFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAH8wp7XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N7q07IZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGiP3AwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PbW4Zj2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RuTslxf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bODbWUbnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Md829xCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJPKS6IjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ElOJB5UQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIQM21cIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgUKpCxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNsV7UhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CDceeJGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j4Rv88GxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFfKoVf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0d1iqkdcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMlGAulNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dNBvbQfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDB5dFE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xsyMH74CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0JLxYBBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihZuGpZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lc1DnJD5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uhCYO0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rf3tY2HwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iJ62Wl5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDZNQRGfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKGmjHszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8WMtvGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQ8J1HmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w65WihXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bdTXebRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0waDKTCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z44cK837Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8qN6AClNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7VzfX0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoAinvtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBCcwTkrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQ81923ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0keS5jbdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXbM134qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p1TfOG5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WHzlacxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Gd1bAjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 53AYAbZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HaPcZseAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B2SU1d30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tfwjuMRcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AnHKYNXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 479XE0IpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lipv85qKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhYFVs3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsZa3Q18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PyC8Fsv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g7MMRklIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hawXt5PMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaMlHHFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V5ueRDiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyvsEoPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JxbHGlS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y1GVuc0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Lf9KP3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eO7ZqMOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8KNhwygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APvHOEkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CUScpxQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YHJJugaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Ee9TAbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2K8uwxtVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8vrsepNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTzKLIRuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMwNWFWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nU5NnJr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func piIB2OlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GDkR7jEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C3tXXLzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljAhe4nuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7WftFdXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kNJB8MmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZHdrEIMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GwmEP2SfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7okbxy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vYZ97OdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lp8co7KZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9Tr2AwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZNroEBpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Yy5MbFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uR5jmHJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DU6NWL5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kPsvXpktWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDKR7HI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qBGgb51VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q18o9ISXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7TCMwdUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8i739QZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FJkc9zjnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jryomtbMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TXFmoLpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrGn2kBjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gON7B3Q0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SPMXMy3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13fdGnfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1AG23REzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Ifg6krMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17rsV3MyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pu9VMsnWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GkU1IYwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u0SrqoUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zMhxoE87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fkcra2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9DWAp30DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXC0CT8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qN2AvSVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5A7bzviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fiJiELMHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gvh88u8iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXL04uJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kj4xYDnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WhJmtZlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TufGjlQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iMRfJGPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HA01pbyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PXcmbEqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FmE8w5LuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jkh8Eo1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rNHE3DudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tvUW7svDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRv2NmEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fY6wA2MCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3UUsinoyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9w0iwFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iGBlm0qTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZF6AhTLAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3X7lfKIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzFcql7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2xXLEPeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IdjGTOOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 10W7MO0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lY0acTQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4eMl0vl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXxBKkL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RatQGOspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lp5gAAfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XPNZ9oQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YMt0PQ4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzHKEr9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cU2ehYk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJP7GstaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rDuMePPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9sO6FWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZcGRSExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fv4iZvfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDTTkJjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3adllIL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HJq35ZsMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCjTeq29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wYKHtZtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZERqTrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZzZQTgvcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JHMMqUDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3reJsY4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QnPVXt6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asfTQwomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YP0OAk7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EviEM1wVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlRvEBa2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dbs2jIfJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VgXEN6uNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U9RpS3T4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsrxePiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jBcU8kq3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiIB2YlIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OL4MveQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rav1Rsr7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWAXbQHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcPxV2n9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wA0itCV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func af4YyogwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeudHIx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xW61dL8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzQDloOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3khikaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xjNYuMTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8vnxR7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qePLB6waWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ckdVJww0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8qYvtDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdZbXBdbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q1D0oRyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

