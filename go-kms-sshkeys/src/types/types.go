package types

type JsonFile struct {
	Aws     Aws     `json:"aws"`
	Project Project `json:"project"`
}

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Version     string `json:"version"`
}

type Aws struct {
	Kmsid       string `json:"kmsid"`
	Secret_name string `json:"secret_name"`
	Region      string `json:"region"`
}
