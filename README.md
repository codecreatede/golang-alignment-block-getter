go-alignment-block-getter

- etxracting a particular block of alignment from the genome alignment
- extract a region specific block and write the block alignment.

```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang-alignment-block-getter% \
go run main.go -h
This estimates the site proportion in your whole genome or gene specific alignment

Usage:
  flags [flags]

Flags:
  -a, --alignmentfile string   a alignment file (default "align")
  -e, --endcoordinate int      end of the alignment block (default 40)
  -h, --help                   help for flags
  -s, --startcoordinate int    start of the alignment block (default 1)
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang-alignment-block-getter% \
go run main.go -a ./samplefile/samplealignment.fasta -s 1 -e 10
>ENA|OX291461|OX291461.1         AACTATC--
>ENA|OX291509|OX291509.1         CTC----TC

```


Gaurav Sablok
