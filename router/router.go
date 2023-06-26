package router

import (
	graphctl "LoreGit/controller/graph"
	graphmodels "LoreGit/models/graph"
	//graphmodels "LoreGit/models/graph"
	"github.com/gin-gonic/gin"
)

type szx interface {
}

func Route(router *gin.Engine) {
	api := router.Group("/loregit")
	{
		api.GET("/getGraph", graphctl.GetGraph)
		//api.GET("/NAME/:name", func(c *gin.Context) {
		//	name := c.Param("name")
		//	//typeOfA := reflect.TypeOf(name)
		//	//.Println(typeOfA.Name(), typeOfA.Kind())
		//	s := graphmodels.Find_Node(name)
		//	c.JSON(200, s)
		//	//c.String(http.StatusOK, "the phone of %s", name)
		//})
		api.GET("/Insert", func(c *gin.Context) {

			title := "C"
			name := "CC"
			s := graphmodels.InsertPernson(title, name)
			c.JSON(200, s)

			//c.String(http.StatusOK, "the phone of %s %s", title, name)
		})
		api.GET("/Delete", func(c *gin.Context) {

			title := "B"
			id := 3
			s := graphmodels.DeletePernson(title, id)

			c.JSON(200, s)
			//c.String(http.StatusOK, "the phone of %s %s", title, name)
		})
		api.GET("/Modify", func(c *gin.Context) {

			attributes := map[string]string{
				"age": "12",
				"sex": "girl", // 注意这里逗号不可缺少，否则会报语法错误
			}
			title := "C"
			id := 24
			s := graphmodels.ModifyPernson(title, id, attributes)

			c.JSON(200, s)
		})
		api.GET("/Search", func(c *gin.Context) {

			title := "C"
			id := 24
			s := graphmodels.SearchPernson(title, id)

			c.JSON(200, s)
			//c.String(http.StatusOK, "the phone of %s %s", title, name)
		})
	}
}
