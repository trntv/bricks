package blueprint

type DefaultBlueprint struct {

}

func (d DefaultBlueprint) Collect() {
	panic("implement me")
}

type Blueprint interface {
	Collect()
}

type Generator interface {
	Generate(string, Blueprint)
}
