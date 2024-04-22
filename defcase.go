package gin

import (
	"github.com/go-shafaq/gin/internal/defcase"
	"github.com/go-shafaq/gin/internal/json"
)

type DefCase = defcase.DefCase

func SetDefCase(dcase DefCase) {
	defcase.SetDefCase(dcase)
	json.SetDefCase(dcase)
}
