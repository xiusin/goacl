package goacl

type Role struct {
    _name string
    _description string
}

func NewRole(name,description string) Role {
    return Role{_name:name,_description:description}
}

func (r *Role) GetName() string {
    return r._name
}

func (r *Role) GetDescription() string  {
    return r._description
}

func(r *Role) ToString() string{
    return r._name
}
