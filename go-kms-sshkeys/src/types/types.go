package types

type JsonFile struct {
	Aws     Aws     `json:"aws"`
	Project Project `json:"project"`
}

type Project struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Author       string `json:"author"`
	Version      string `json:"version"`
	Pem_file     string `json:"pem_file"`
	Encrypt_file string `json:"encrypt_file"`
	Decrypt_file string `json:"decrypt_file"`
}

type Aws struct {
	Kms_arn     string `json:"kms_arn"`
	Secret_name string `json:"secret_name"`
	Region      string `json:"region"`
}

type Args struct {
	Import string `arg:"-i" help:"Import your selected sshkeypair to aws secretsmanager, values: true or false"`
	Export string `arg:"-e" help:"Export from aws secretsmanager your stored sshkeypair, values: true or false"`
}
