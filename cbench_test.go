package main

// go test -bench=. -memprofile profile-mem.out -cpuprofile profile-cpu.out -blockprofile profile-block.out

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/lzw"
	"compress/zlib"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

var pi []byte
var comp []byte
var results map[string][]int // [type]{}
const xtimes = 100           // how many times to compress and decompress

func TestMain(m *testing.M) {
	data, err := ioutil.ReadFile("pi.txt")
	if err != nil {
		panic(err)
	}

	// prepare compressed data for decompress test
	pi = data
	fmt.Println("pi size", len(pi))

	results = map[string][]int{
		"zlib":  []int{},
		"flate": []int{},
		"gzip":  []int{},
		"lzw":   []int{},
	}
	// exec test
	retCode := m.Run()

	fmt.Printf("tested %d times each\n", len(results["zlib"]))
	// results
	for name, nums := range results {
		total := 0
		for _, num := range nums {
			total += num
		}
		avg := total / len(nums)
		pc := 100 * avg / len(pi)
		// fmt.Printf("%5s avg size %d %dpc      %d %v\n", name, avg, pc, len(nums), nums)
		fmt.Printf("%5s avg size %d %dpc\n", name, avg, pc)
	}
	// exit
	os.Exit(retCode)
}

// ******************
// zlib
// ******************
func BenchmarkZlibCompress(b *testing.B) {
	name := "zlib"
	for i := 0; i < xtimes; i++ {
		var buf bytes.Buffer
		w := zlib.NewWriter(&buf)
		w.Write(pi)
		w.Close()

		comp = buf.Bytes()
		results[name] = append(results[name], len(comp))
	}
}

func BenchmarkZlibDecompress(b *testing.B) {
	for i := 0; i < xtimes; i++ {
		buf := bytes.NewBuffer(comp)
		r, err := zlib.NewReader(buf)
		if err != nil {
			panic(err)
		}
		r.Close()
	}
}

// ******************
// flate
// ******************
func BenchmarkFlateCompress(b *testing.B) {
	name := "flate"
	for i := 0; i < xtimes; i++ {
		var buf bytes.Buffer
		w, err := flate.NewWriter(&buf, flate.DefaultCompression)
		if err != nil {
			panic(err)
		}
		w.Write(pi)
		w.Close()

		comp = buf.Bytes()
		results[name] = append(results[name], len(comp))
	}
}

func BenchmarkFlateDecompress(b *testing.B) {
	for i := 0; i < xtimes; i++ {
		buf := bytes.NewBuffer(comp)
		r := flate.NewReader(buf)
		r.Close()
	}
}

// ******************
// gzip
// ******************
func BenchmarkGzipCompress(b *testing.B) {
	name := "gzip"
	for i := 0; i < xtimes; i++ {
		var buf bytes.Buffer
		w := gzip.NewWriter(&buf)
		w.Write(pi)
		w.Close()

		comp = buf.Bytes()
		results[name] = append(results[name], len(comp))
	}
}

func BenchmarkGzipDecompress(b *testing.B) {
	for i := 0; i < xtimes; i++ {
		buf := bytes.NewBuffer(comp)
		r, err := gzip.NewReader(buf)
		if err != nil {
			panic(err)
		}
		r.Close()
	}
}

// ******************
// lzw
// ******************
func BenchmarkLZWCompress(b *testing.B) {
	name := "lzw"
	for i := 0; i < xtimes; i++ {
		var buf bytes.Buffer
		w := lzw.NewWriter(&buf, lzw.MSB, 8)
		w.Write(pi)
		w.Close()

		comp = buf.Bytes()
		results[name] = append(results[name], len(comp))
	}
}

func BenchmarkLZWDecompress(b *testing.B) {
	for i := 0; i < xtimes; i++ {
		buf := bytes.NewBuffer(comp)
		r := lzw.NewReader(buf, lzw.MSB, 8)
		r.Close()
	}
}
