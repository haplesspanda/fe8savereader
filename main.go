package main

import (
	"bufio"
	"log"
	"os"

	"github.com/haplesspanda/fe8savereader/format"
)

func main() {
	log.Println("Starting up save reader")
	if len(os.Args) < 2 {
		log.Println("Missing args, exiting")
		PrintUsage()
		return
	}
	if os.Args[1] == "read" {
		if len(os.Args) < 4 {
			log.Println("Missing read args, exiting")
			PrintUsage()
			return
		}

		filename := os.Args[2]
		outputFilename := os.Args[3]
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		log.Printf("Writing to %s", outputFilename)
		output, err := os.Create(outputFilename)
		if err != nil {
			panic(err)
		}
		defer output.Close()

		format.Read(f, output)
	} else if os.Args[1] == "compare" {
		if len(os.Args) < 5 {
			log.Println("Missing compare args, exiting")
			PrintUsage()
			return
		}

		oldFilename := os.Args[2]
		newFilename := os.Args[3]
		outputFilename := os.Args[4]
		log.Printf("Reading old file %s", oldFilename)
		oldF, err := os.Open(oldFilename)
		if err != nil {
			panic(err)
		}
		defer oldF.Close()

		log.Printf("Reading new file %s", newFilename)
		newF, err := os.Open(newFilename)
		if err != nil {
			panic(err)
		}
		defer newF.Close()

		output, err := os.Create(outputFilename)
		if err != nil {
			panic(err)
		}
		defer output.Close()

		writer := bufio.NewWriter(output)

		format.Diff(oldF, newF, writer)

		err = writer.Flush()
		if err != nil {
			panic(err)
		}

	} else {
		PrintUsage()
		return
	}
	log.Println("Done")
}

func PrintUsage() {
	log.Printf("Usage: %s read savefile.sav output.txt", os.Args[0])
	log.Println("OR")
	log.Printf("Usage: %s compare oldsave.sav newsave.sav output.txt", os.Args[0])
}
