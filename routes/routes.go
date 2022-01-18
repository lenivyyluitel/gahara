package routes

import (
	"github.com/labstack/echo/v4"
	"gitlab.com/lenivyyluitel/gahara/pkg/api"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Site struct {
	Title string
	Link  string
}

type HomePage struct {
	PageTitle string
	Image     string
	Sites     []Site
	About     string
}

func Home(c echo.Context) error {
	waifuRandom := api.RandomTags()
	data := HomePage{
		PageTitle: "Your daily dose of waifu",
		//Sites: []Site{
		//	{Title: "Waifupics", Link: "https://waifu.pics"},
		//	{Title: "Zerochan", Link: "https://zerochan.net"},
		//},
		Image: api.ImagesWaifu("sfw", waifuRandom),
		About: waifuRandom,
	}
	return c.Render(http.StatusOK, "home.tmpl", data)
}

func RandomFiles(c echo.Context) error {
	files, err := ioutil.ReadDir("./assets")
	if err != nil {
		panic(err)
	}
	// randomize
	rand.Seed(time.Now().Unix())
	num := rand.Int()
	z := "assets/" + files[num%len(files)].Name()

	return c.File(z)
}
