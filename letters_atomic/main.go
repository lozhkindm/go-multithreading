package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const Letters = "abcdefghijklmnopqrstuvwxyz"

func count(url string, freq *[26]int32, wg *sync.WaitGroup) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	for n := 0; n <= 20; n++ {
		for _, b := range body {
			c := strings.ToLower(string(b))
			i := strings.Index(Letters, c)
			if i >= 0 {
				atomic.AddInt32(&freq[i], 1)
			}
		}
	}
	wg.Done()
}

func main() {
	var freq [26]int32
	wg := sync.WaitGroup{}
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		wg.Add(1)
		go count(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &freq, &wg)
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took: %s\n", elapsed)
	for i, c := range freq {
		fmt.Printf("%s -> %d\n", string(Letters[i]), c)
	}
}

/**
Processing took: 4.493259401s
a -> 11198397
b -> 2360358
c -> 6358611
d -> 5724264
e -> 19187343
f -> 3484131
g -> 2733843
h -> 5229714
i -> 11333490
j -> 345534
k -> 952476
l -> 5404938
m -> 4546437
n -> 11210682
o -> 10893078
p -> 4341834
q -> 307671
r -> 11151252
s -> 10812144
t -> 14615223
u -> 3942078
v -> 1448811
w -> 1744869
x -> 582078
y -> 2012409
z -> 160587
*/
