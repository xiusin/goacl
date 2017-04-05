package goacl

type Resource struct {
    _name string
    _description string
}

func NewResource(name,description string) Resource {
    return Resource{name,description}
}

func (r *Resource) GetName() string {
    return r._name
}

func (r *Resource) GetDescription() string  {
    return r._description
}

func(r *Resource) ToString() string{
    return r._name
}
