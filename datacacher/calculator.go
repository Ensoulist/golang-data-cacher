package datacacher

type ICalculator[ContainerType ICacheContainer[KeyType, IdType], KeyType comparable, IdType comparable] interface {
	Create(ContainerType, *Param) (any, error)
}
