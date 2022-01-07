package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const Letters = "abcdefghijklmnopqrstuvwxyz"

func count(url string, freq *[26]int32) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		i := strings.Index(Letters, c)
		if i >= 0 {
			freq[i] += 1
		}
	}
}

func main() {
	var freq [26]int32
	start := time.Now()
	for i := 1000; i <= 1200; i++ {
		count(fmt.Sprintf("https://www.rfc-editor.org/rfc/rfc%d.txt", i), &freq)
	}
	elapsed := time.Since(start)
	fmt.Println("Done")
	fmt.Printf("Processing took: %s\n", elapsed)
	for i, c := range freq {
		fmt.Printf("%s -> %d\n", string(Letters[i]), c)
	}
}

/**
Processing took: 1m3.570868098s
a -> 533257
b -> 112398
c -> 302791
d -> 272584
e -> 913683
f -> 165911
g -> 130183
h -> 249034
i -> 539690
j -> 16454
k -> 45356
l -> 257378
m -> 216497
n -> 533842
o -> 518718
p -> 206754
q -> 14651
r -> 531012
s -> 514864
t -> 695963
u -> 187718
v -> 68991
w -> 83089
x -> 27718
y -> 95829
z -> 7647
*/
