package commune

type Commune struct {
	Nom             string
	Code            string
	CodeDepartement string
	CodeRegion      string
	CodesPostaux    []string
	Population      int
}
