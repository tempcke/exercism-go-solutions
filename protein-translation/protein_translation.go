package protein

import "errors"

//ErrInvalidBase invalid sequence in string
var ErrInvalidBase = errors.New("invalid")

//ErrStop encountered stop sequence
var ErrStop = errors.New("stop")

var codonProteinMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

//FromCodon converts a 3 char codon string to a protein name
func FromCodon(codon string) (string, error) {
	p, valid := codonProteinMap[codon]
	if !valid {
		return "", ErrInvalidBase
	}
	if p == "STOP" {
		return "", ErrStop
	}
	return p, nil
}

//FromRNA converts an rna string into a slice of proteins
func FromRNA(rna string) ([]string, error) {
	var proteins []string
	for i := 0; i < len(rna); i = i + 3 {
		c := rna[i : i+3]
		p, err := FromCodon(c)
		if err == ErrStop {
			return proteins, nil
		}
		if err != nil {
			return proteins, err
		}
		proteins = append(proteins, p)
	}
	return proteins, nil
}
