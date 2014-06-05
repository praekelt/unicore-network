package server

import (
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"os"
	"testing"
)

func TestIdentityLoading(t *testing.T) {
	temp_file, _ := ioutil.TempFile("", "uc")
	defer os.Remove(temp_file.Name())

	ident := Ident{Signature: "signature", Hostname: "localhost", DisplayName: "UCN"}
	data, _ := yaml.Marshal(&ident)
	temp_file.Write(data)
	temp_file.Close()

	i, _ := LoadIdentity(temp_file.Name())
	if i.Signature != "signature" {
		t.Errorf("Unexpected Signature: %s", i.Signature)
	}

	if i.Hostname != "localhost" {
		t.Errorf("Unexpected Hostname: %s", i.Hostname)
	}

	if i.DisplayName != "UCN" {
		t.Errorf("Unexpected DisplayName: %s", i.DisplayName)
	}
}

func TestIdentityCreation(t *testing.T) {
	temp_file, _ := ioutil.TempFile("", "uc")
	defer os.Remove(temp_file.Name())

	ident := CreateIdentity("signature", "hostname", "display name")
	ident, _ = SaveIdentity(ident, temp_file.Name())

	reloaded_ident, _ := LoadIdentity(temp_file.Name())
	if ident != reloaded_ident {
		t.Errorf("Unexpected ident: %s", reloaded_ident)
	}
}

func TestGetOrCreateIdentityCreating(t *testing.T) {
	// create a tempfile and delete it immediately but continue
	// using the name to make sure we've got a unique file name
	temp_file, _ := ioutil.TempFile("", "uc")
	os.Remove(temp_file.Name())

	identity, _ := GetOrCreateIdentity(temp_file.Name())
	if identity.Signature == "" {
		t.Error("Unexpected empty Signature")
	}
	if identity.Hostname == "" {
		t.Error("Unexpected empty Hostname")
	}
	if identity.DisplayName == "" {
		t.Error("Unexpected empty DisplayName")
	}
}

func TestGetOrCreateIdentityGetting(t *testing.T) {
	temp_file, _ := ioutil.TempFile("", "uc")
	defer os.Remove(temp_file.Name())
	identity, _ := SaveIdentityFile(CreateIdentity("foo", "bar", "baz"), temp_file)
	reloaded_identity, _ := GetOrCreateIdentity(temp_file.Name())
	if identity != reloaded_identity {
		t.Errorf("Unexpected identity when reloading. %s vs %s", identity, reloaded_identity)
	}
}
