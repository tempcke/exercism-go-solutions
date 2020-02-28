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
	p := codonProteinMap[codon]
	switch p {
	case "STOP":
		return "", ErrStop
	case "":
		return "", ErrInvalidBase
	default:
		return p, nil
	}
}

//FromRNA converts an rna string into a slice of codons
func FromRNA(rna string) ([]string, error) {
	var codon string
	var output []string
	for _, char := range rna {
		codon += string(char)
		if len(codon) == 3 {
			p, err := FromCodon(codon)
			if err == ErrInvalidBase {
				return output, err
			}
			if err == ErrStop {
				return output, nil
			}
			output = append(output, p)
			codon = ""
		}
	}
	return output, nil
}
