package gen

import (
	"gqlgen-generator/config"
	"gqlgen-generator/db"
	"testing"
)

func init() {
	config.Setup(`D:\project\gqlgen-generator\config.yaml`)
	db.Setup()
}

func TestGenerate(t *testing.T) {
	Generate()
}
