package meta

import (
	"fmt"
	"net/http"

	"github.com/gregbuehler/tilr/modules/setting"

	"github.com/julienschmidt/httprouter"
)

func MetaHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Tilr v%s", setting.AppVer)
}
