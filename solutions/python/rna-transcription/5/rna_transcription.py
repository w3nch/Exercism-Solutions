_TABLE = str.maketrans("GCTA", "CGAU")
_translate = str.translate

def to_rna(dna_strand):
    return _translate(dna_strand, _TABLE)
