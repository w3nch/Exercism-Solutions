def to_rna(dna_strand):
    trans = {"G":"C", "C":"G", "T":"A", "A":"U"}
    return "".join(trans.get(v, v) for v in dna_strand)
