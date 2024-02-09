package collections

type Assigner interface {
	AssignPriority(item interface{}) func() float32
}
