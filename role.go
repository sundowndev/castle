package castle

type RoleInterface interface {
	HasPermission(*Permission) bool
	String() string
	InheritFrom(RoleInterface) RoleInterface
	GetRoles() []*Permission
}

// Role ...
type Role struct {
	Name        string
	Permissions []*Permission
}

func (r *Role) GetRoles() []*Permission {
	return r.Permissions
}

func (r *Role) String() string {
	return r.Name
}

func (r *Role) HasPermission(role *Permission) bool {
	for _, v := range r.Permissions {
		if v.UUID == role.UUID {
			return true
		}
	}

	return false
}

func (r *Role) InheritFrom(role RoleInterface) RoleInterface {
	r.Permissions = append(r.Permissions, role.GetRoles()...)

	return r
}
