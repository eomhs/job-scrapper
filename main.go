package main

import (
	"job-scrapper/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const FILE_NAME string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.HTML")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(FILE_NAME)
	term := strings.ToLower(scrapper.Cleanstring(c.FormValue("term")))
	scrapper.Scrape(term)
	return c.Attachment(FILE_NAME, FILE_NAME)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
