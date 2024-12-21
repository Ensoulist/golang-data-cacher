package datacacher

type ICalculator interface {
	Create(ICacheContainer, *Param) (any, error)
}
