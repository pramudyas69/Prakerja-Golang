package luas

func CalculateTrapezoidArea(base1, base2, height float64) float64 {
	if height < 0 {
		return 0 // Mengembalikan 0 jika tinggi negatif
	}
	return 0.5 * (base1 + base2) * height
}
