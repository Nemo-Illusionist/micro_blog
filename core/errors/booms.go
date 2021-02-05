package errors

// Booms can contain multiple boom errors
type Booms struct {
	Errors []Boom `json:"errors"`
}

type Boom struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

func (b *Booms) Add(e Boom) {
	b.Errors = append(b.Errors, e)
}
