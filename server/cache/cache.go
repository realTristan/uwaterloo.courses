package cache

// Import modules
import (
	"encoding/json"
	"fmt"
	"strings"
)

/*

	Some of you might be wondering why I decided to use a string cache
	instead of a map cache. For starters iterating over a map cache takes
	too long. To solve that you can start a bunch of goroutines which returns
	the same speed, but destroys your memory usage, especially if a lot of
	users are calling the api. Secondly, in most cases the string cache is
	faster.

*/

// Hold Course data in memory cache map
var Cache string

// The Set() function sets the data for the
// given key in the cache
func Set(value map[string]string) error {
	tmp, _ := json.Marshal(value)
	Cache += string(tmp)
	return nil
}

// The GetSimilarCourses() function iterates through the
// cache and gets any courses that contain the query args
func GetSimilarCourses(query string, subject string) []map[string]string {
	// Define variables
	var (
		courseMapStart    int = -1
		closeBracketCount int = 0
		subjectResult     []map[string]string
		similarResult     []map[string]string
		TempCache         string = strings.ToLower(Cache)
	)

	// Iterate over the lowercase cache string
	for i := 0; i < len(TempCache); i++ {

		// Check if current index is the start of
		// the course data map
		if TempCache[i] == '{' {
			if courseMapStart == -1 {
				courseMapStart = i
			}
			closeBracketCount++
		} else

		// Check if the current index is the end of
		// the course data map
		if TempCache[i] == '}' {
			if closeBracketCount == 1 {
				// Check if the map contains the subject code
				if strings.Contains(
					Cache[courseMapStart:i+1], fmt.Sprintf(`,"title":"%s `, subject)) {

					// Convert the string to a map
					var data map[string]string
					json.Unmarshal([]byte(Cache[courseMapStart:i+1]), &data)

					// Append the map to the result array
					subjectResult = append(subjectResult, data)
				} else

				// Check if the map contains the query string
				if strings.Contains(TempCache[courseMapStart:i+1], query) {
					// Convert the string to a map
					var data map[string]string
					json.Unmarshal([]byte(Cache[courseMapStart:i+1]), &data)

					// Append the map to the result array
					similarResult = append(similarResult, data)
				}
				// Reset indexing variables
				closeBracketCount = 0
				courseMapStart = -1
			} else {
				closeBracketCount--
			}
		}
	}
	// Return the combined arrays
	return append(subjectResult, similarResult...)
}
