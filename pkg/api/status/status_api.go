package status

import (
	"net/http"
	"runtime"
	"soarca/pkg/models/api"
	"soarca/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
)

var status = api.Status{
	Uptime:  api.Uptime{Since: time.Now(), Milliseconds: 0},
	Mode:    utils.GetEnv("LOG_MODE", "production"),
	Runtime: runtime.GOOS,
}

func SetVersion(version string) {
	status.Version = version
}

// /Status/ping GET handler for handling status api calls
// Returns the status model object for SOARCA
//
//	@Summary	ping to see if SOARCA is up returns pong
//	@Schemes
//	@Description	return SOARCA status
//	@Tags			ping pong
//	@Produce		plain
//	@success		200	string	pong
//	@Router			/status/ping [GET]
func GetPong(g *gin.Context) {
	g.Data(http.StatusOK, "text/plain", []byte("pong"))
}

// /Status GET handler for handling status api calls
// Returns the status model object for SOARCA
//
//	@Summary	gets the SOARCA status
//	@Schemes
//	@Description	return SOARCA status
//	@Tags			status
//	@Produce		json
//	@success		200	{object}	api.Status
//	@failure		400	{object}	api.Error
//	@Router			/status [GET]
func GetApi(g *gin.Context) {
	status.Uptime.Milliseconds = uint64(time.Since(status.Uptime.Since).Milliseconds())
	status.Time = time.Now()

	g.JSON(http.StatusOK, status)
}
