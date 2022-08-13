package api

// Import packages
import (
	"encoding/json"
	"fmt"

	"github.com/realTristan/The_University_of_Waterloo/global"
	scraper "github.com/realTristan/The_University_of_Waterloo/scraper"
	"github.com/valyala/fasthttp"
)

// Fasthttp request client
var RequestClient *fasthttp.Client = &fasthttp.Client{}

// The CourseDataHandler() function handles the incoming requests
// with the /courses?course={course_code} path.
// The function is used to scrape the data of a subject using the
// ScrapeCourseData() function, then return it as a json string
//
// The function takes the ctx *fasthttp.RequestCtx parameter
func CourseDataHandler(ctx *fasthttp.RequestCtx) {
	// Scrape the course data
	var (
		course      []byte = ctx.QueryArgs().Peek("course")
		result, err        = scraper.ScrapeCourseData(RequestClient, string(course))
	)
	// Handle the error
	if err != nil {
		ctx.SetStatusCode(500)
		fmt.Fprintf(ctx, "{\"error\": \"%v\"}", err)
	} else {
		// Marshal the data result
		_json, _ := json.Marshal(result)

		// Set the response body
		fmt.Fprint(ctx, string(_json))
	}
}

// The SubjectCodesHandler() function handles the incoming
// requests with the /subjects path
//
// The function returns the global SubjectCodes array
func SubjectCodesHandler(ctx *fasthttp.RequestCtx) {
	// Marshal the subject codes map
	_json, _ := json.Marshal(map[string][]string{
		"subjects": global.SubjectCodes,
	})

	// Set the response body
	fmt.Fprint(ctx, string(_json))
}

// The SubjectCodesWithNamesHandler() function handles the incoming
// requests with the /subjects/names path
//
// The function returns the global SubjectNames
func SubjectCodesWithNamesHandler(ctx *fasthttp.RequestCtx) {
	// Marshal the codes and names map
	_json, _ := json.Marshal(global.SubjectNames)

	// Set the response body
	fmt.Fprint(ctx, string(_json))
}