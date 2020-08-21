package ydapp

type YdDepartment struct {
	DeptId   int    `json:"deptId"`   //部门id
	Position string `json:"position"` //职务名称
	Weight   int    `json:"weight"`   //职务权重
	SortId   int    `json:"sortId"`   //用户在部门中的排序，值越大排序越靠前
}

type YdUserInfo struct {
	Gender       int             `json:"gender"`
	UserId       string          `json:"userId"`
	Name         string          `json:"name"`
	Mobile       string          `json:"mobile"`
	Phone        string          `json:"phone"`
	Email        string          `json:"email"`
	DepartmentID []int           `json:"dept"`
	DepDetail    []*YdDepartment `json:"deptDetail"`
}
