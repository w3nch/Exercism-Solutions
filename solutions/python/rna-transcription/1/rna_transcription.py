def to_rna(dna_strand):
    rna = ""
    for val in dna_strand:
        if val == "G":
            rna  += "C"
        elif  val == "C":
            rna += "G"
        elif  val == "T":
            rna += "A"
        elif  val == "A":
            rna += "U"
        else :
            rna += val
    return rna