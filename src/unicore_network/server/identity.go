package server

import (
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"os"
)

type Ident struct {
	Signature   string
	Hostname    string
	DisplayName string
}

func CreateIdentity(signature string, hostname string, display_name string) Ident {
	return Ident{Signature: signature, Hostname: hostname, DisplayName: display_name}
}

func SaveIdentity(identity Ident, filename string) (Ident, error) {
	fp, err := os.Create(filename)
	if err != nil {
		return identity, err
	}
	defer fp.Close()
	return SaveIdentityFile(identity, fp)
}

func SaveIdentityFile(identity Ident, file *os.File) (Ident, error) {
	data, err := yaml.Marshal(&identity)
	file.Write(data)
	return identity, err
}

func LoadIdentity(filename string) (Ident, error) {
	ident := Ident{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return ident, err
	}
	err = yaml.Unmarshal(data, &ident)
	return ident, err
}
