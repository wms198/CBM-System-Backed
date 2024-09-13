package main

import (
	"fmt"
	"main/controllers"
	"main/initD"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func init() {
	initD.LoadEnvVariables()
	initD.ConnectDB()
	initD.SyncDatabase()
}

// GET retrieves data.
// POST creates data.
// PUT updates data entirely.
// PATCH allows partially updating data.
// DELETE removes data.
func main() {
	r := gin.Default()
	r.Use(RequestLogger())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("user", controllers.CreateUser)

	r.GET("/users", controllers.ReadUsers)
	r.GET("/user/:id", controllers.ReadUser)
	r.PATCH("/user/:id", controllers.UpdateUser)
	r.DELETE("/user/:id", controllers.DeleteUser)
	r.POST("/appointment", controllers.CreateAppointment)
	r.GET("/appointments", controllers.ReadAppointments)
	r.PATCH("/appointment/:id", controllers.UpdateAppointment)
	r.DELETE("/appointment/:id", controllers.DeleteAppointment)
	r.GET("/appointment/:id", controllers.Readappointment)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", Authorized(), controllers.Validate)
	r.Run()
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Request.Host, c.Request.RemoteAddr, c.Request.RequestURI)

		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(c.Request, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(requestDump))

		c.Next()
	}
}

func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			return
		}
		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Unauthorized"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["foo"], claims["nbf"])
		} else {
			fmt.Println(err)
		}
		c.Next()
	}
}
