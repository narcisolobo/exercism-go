package strand

func ToRNA(dna string) string {
	if len(dna) == 0 {
		return ""
	}

	rnaMap := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	rna := ""

	for _, char := range dna {
		rna += string(rnaMap[char])
	}

	return rna
}
