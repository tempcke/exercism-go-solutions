package protein

import "errors"

var ErrInvalidBase = errors.New("invalid")
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

func FromRNA(rna string) (interface{}, error) {
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
