package main

/*

Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-20

A alignment block getter which gives you the alignment block from the alignment either the whole genome
alignment or the read alignment or the alignment for the phylogenomics. It takes an alignment file and
the alignment block start and the end coordinates.


*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var (
	alignment string
	start     int
	end       int
)

var rootCmd = &cobra.Command{
	Use:  "flags",
	Long: "This estimates the site proportion in your whole genome or gene specific alignment",
	Run:  flagsFunc,
}

func init() {
	rootCmd.Flags().StringVarP(&alignment, "alignmentfile", "a", "align", "a alignment file")
	rootCmd.Flags().IntVarP(&start, "startcoordinate", "s", 1, "start of the alignment block")
	rootCmd.Flags().IntVarP(&end, "endcoordinate", "e", 40, "end of the alignment block")
}

func flagsFunc(cmd *cobra.Command, args []string) {
	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	type alignBlock struct {
		id  string
		seq string
	}

	fOpen, err := os.Open(alignment)
	if err != nil {
		log.Fatal(err)
	}

	alignIDcapture := []alignmentID{}
	alignSeqcapture := []alignmentSeq{}
	sequenceCap := []string{}
	sequenceID := []string{}
	alignmentBlock := []alignBlock{}
	fRead := bufio.NewScanner(fOpen)
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignIDcapture = append(alignIDcapture, alignmentID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignSeqcapture = append(alignSeqcapture, alignmentSeq{
				seq: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			sequenceCap = append(sequenceCap, string(line))
		}
		if strings.HasPrefix(string(line), ">") {
			sequenceID = append(sequenceID, string(line))
		}
	}

	for i := 0; i < len(sequenceID); i++ {
		alignmentBlock = append(alignmentBlock, alignBlock{
			id:  string((sequenceID[i])),
			seq: string((sequenceCap[i][start:end])),
		})
	}

	for i := range alignmentBlock {
		fmt.Println(alignmentBlock[i].id, "\t", alignmentBlock[i].seq)
	}
}
