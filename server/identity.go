package server

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/json"
	"gopkg.in/yaml.v1"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Ident struct {
	Signature   string `json:"signature"`
	Hostname    string `json:"hostname"`
	DisplayName string `json:"display_name"`
}

func (i *Ident) ToString() (string, error) {
	b, err := i.ToJson()
	return string(b), err
}

func (i *Ident) ToJson() ([]byte, error) {
	return json.Marshal(i)
}

func NewIdentFromString(s string) (Ident, error) {
	ident := Ident{}
	err := json.Unmarshal([]byte(s), &ident)
	return ident, err
}

func NewIdentFromReader(reader io.Reader) (Ident, error) {
	ident := Ident{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&ident)
	return ident, err
}

func CreateIdentity(signature string, hostname string, display_name string) Ident {
	return Ident{Signature: signature, Hostname: hostname, DisplayName: display_name}
}

func GetOrCreateIdentity(filename string) (Ident, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Printf("Creating identity file: %s", filename)
		identity := CreateIdentity(uuid.New(), "localhost", "UC Network Node")
		return SaveIdentity(identity, filename)
	}
	log.Printf("Loading identity file: %s", filename)
	return LoadIdentity(filename)
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
