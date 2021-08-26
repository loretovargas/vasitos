package domain

type Numero struct {
	Value int
}

func (n Numero) EsMenorOIgualACero() bool {
	return n.Value <= 0
}

func (n Numero) EsMayorASeis() bool {
	return n.Value > 6
}
