def to_rna(dna):
    dna_nuc = ['G', 'C', 'T', 'A']
    rna_nuc = ['C', 'G', 'A', 'U']
    rna = []
    for char in dna:
        try:
            rna.append(rna_nuc[dna_nuc.index(char)])
        except ValueError:
            return ""
    return ''.join(rna)
