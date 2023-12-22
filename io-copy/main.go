package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"io"
	"log"
	"time"
)

func copyBytes(dst io.Writer, src io.Reader) (written int64, err error) {
	for {
		b := make([]byte, 128*1024) // 32 KB
		n := 0

		n, err = src.Read(b)
		if err != nil {
			break
		}

		n, err = dst.Write(b[:n])
		if err != nil {
			break
		}

		written += int64(n)
	}

	if err == io.EOF {
		err = nil
	}

	return written, err
}

var mode = "io.Copy"

func init() {
	flag.StringVar(&mode, "mode", mode, "copy mode")
	flag.Parse()
}

func main() {
	const dataSize = 1 << 30 // 1 GB
	data := make([]byte, dataSize)
	if _, err := rand.Read(data); err != nil {
		return
	}

	switch mode {
	case "copyBytes":
		src := bytes.NewReader(data)
		dst := bytes.NewBuffer([]byte{})

		startTime := time.Now()

		if _, err := copyBytes(dst, src); err != nil {
			log.Panicf("copy error: %s", err)
		}

		duration := time.Since(startTime)
		log.Printf("copyBytes: Copied %d bytes in %s", dataSize, duration)
	case "io.Copy":
		src := bytes.NewReader(data)
		dst := bytes.NewBuffer([]byte{})

		startTime := time.Now()

		if _, err := io.Copy(dst, src); err != nil {
			log.Panicf("io.Copy error: %s", err)
		}

		duration := time.Since(startTime)
		log.Printf("io.Copy: Copied %d bytes in %s", dataSize, duration)
	case "io.CopyBuffer":
		src := bytes.NewReader(data)
		dst := bytes.NewBuffer([]byte{})
		buf := make([]byte, 32*1024) // 32 KB

		startTime := time.Now()

		if _, err := io.CopyBuffer(dst, src, buf); err != nil {
			log.Panicf("io.CopyBuffer error: %s", err)
		}

		duration := time.Since(startTime)
		log.Printf("io.CopyBuffer: Copied %d bytes in %s", dataSize, duration)
	}
}
