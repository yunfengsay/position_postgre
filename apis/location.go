package apis

type AddLocationApiForm struct {
	Imgs    []string  `binding:"required"`
	Point   []float64 `binding:"required"`
	Content string
	L_type  []int `binding:"required" json:"l_type"`
}
