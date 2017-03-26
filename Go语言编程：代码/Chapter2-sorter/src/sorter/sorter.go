package main

import (
	"algorithm/bubblesort"
	"algorithm/qsort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

var (
	infile    *string = flag.String("i", "infile", "File contains values for sorting")
	outfile   *string = flag.String("o", "outfile", "File to receive sorted values")
	algorithm *string = flag.String("a", "qsort", "Sort algorithm")
)

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file ", infile)
		return
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
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Filed to create the output file", outfile)
		return err
	}

	defer file.Close()
	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil
}

var ch chan int = make(chan int, 10)

func sort(sort string, values []int) {
	t1 := time.Now()
	switch sort {
	case "qsort":
		qsort.QuickSort(values)
	case "bubblesort":
		bubblesort.BubbleSort(values)
	default:
		fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
	}
	t2 := time.Now()
	fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
	writeValues(values, *outfile)
	ch <- 1
}

func main() {
	//命令行参数处理
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
	}
	//读取文件
	values, err := readValues(*infile)
	if err == nil {
		go sort("qsort", values)
		go sort("bubblesort", values)
		<-ch
		<-ch
	} else {
		fmt.Println(err)
	}

}
