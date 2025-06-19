package routes
import(
 controllers "go-rms/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users",controllers.GetUsers())
	incomingRoutes.GET("users/:user_id",controllers.GetUser())
	incomingRoutes.POST("/user/signup",controllers.Signup())
	incomingRoutes.POST("/user/login",controllers.Login())


}