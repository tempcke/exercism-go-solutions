// Package strand deals with dna strands
package strand

var dnaToRna = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

// ToRNA converts a dna sequence into an rna sequence
func ToRNA(dna string) string {
	var rna string
	for _, d := range dna {
		rna += dnaToRna[d]
	}
	return rna
}
