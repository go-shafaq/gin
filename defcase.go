package gin

import (
	"github.com/gin-gonic/gin/internal/defcase"
	"github.com/gin-gonic/gin/internal/json"
)

type DefCase = defcase.DefCase

func SetDefCase(dcase DefCase) {
	defcase.SetDefCase(dcase)
	json.SetDefCase(dcase)
}
