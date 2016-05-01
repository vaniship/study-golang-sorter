package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"algorithms/bubblesort"
  "algorithms/qsort"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "infile", "input file")
var outfile *string = flag.String("o", "outfile", "output file")
var algorithm *string = flag.String("a", "qsort", "sore algorithm")

func main() {
	flag.Parse()

	if infile != nil {
		fmt.Println("infile = ", *infile, ", outfile = ", *outfile, ", algorithm = ", *algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values: ", values)

		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("unkown algorithm ", *algorithm)
		}

		t2 := time.Now()

    fmt.Println("Sort values: ", values)
    fmt.Println("time ", t2.Sub(t1))

		err = writeValues(values, *outfile)
		if err == nil {
			fmt.Println("Write file ok.")
		} else {
			fmt.Println(err)
		}

	} else {
		fmt.Println(err)
		return
	}
}

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to read the input file ", infile)
	}

	defer file.Close()

	br := bufio.NewReader(file)

	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, semms unexpected.")
			return
		}

		str := string(line)

		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}
