package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/route/endpoint", getPage)
	e.Logger.Fatal(e.Start(":2711"))
}

type request struct {
	Param string `json:"param" xml:"param" form:"param" query:"param"`
}

type response struct {
	Field string `json:"field" xml:"field" form:"field" query:"field"`
}

// define handler function
func getPage(c echo.Context) error {
	req := new(request)
	if err := c.Bind(req); err != nil {
		return err
	}

	// TODO: add some work
	//	Bad example, but for example calling some shell cmd
	// var stdout, stderr bytes.Buffer
	// cmd := exec.Command("ls", "./")
	// cmd.Stdout = &stdout
	// cmd.Stderr = &stderr
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal()
	// }

	res := new(response)
	res.Field = req.Param

	return c.JSON(http.StatusOK, res)
}
