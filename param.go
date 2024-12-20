package datacacher

type Param struct {
	id    any
	extra map[any]any
}

func NewParam() *Param {
	return &Param{}
}

func (p *Param) SetId(id any) *Param {
	p.id = id
	return p
}

func (p *Param) ReplaceExtra(extra map[any]any) *Param {
	p.extra = extra
	return p
}

func (p *Param) SetExtra(k, v any) *Param {
	if p.extra == nil {
		p.extra = make(map[any]any)
	}
	p.extra[k] = v
	return p
}

func (p *Param) Id() any {
	if p.id == nil {
		return 0
	}
	return p.id
}

func (p *Param) GetExtra(k any) any {
	if p.extra == nil {
		return nil
	}
	return p.extra[k]
}
