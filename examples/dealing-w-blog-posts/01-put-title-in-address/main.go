package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

var runtimeOS string

// Post is capitalized, isn't it exported?
type Post struct {
	Title   string
	Date    string
	Summary string
	Body    string
	File    string
	Address string
}

func main() {
	runtimeOS = runtime.GOOS
	fmt.Println("Starting Test Blog")
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.ListenAndServe(":6060", nil)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("icon")
	fmt.Fprintf(w, "icon")
}

// getPosts creates and returns a slice of Posts from md files under the posts directory.
func handler(w http.ResponseWriter, r *http.Request) {
	// a := []Post{}
	files := getPathsWMarkdownFiles("posts")
	sort.Sort(sort.Reverse(sort.StringSlice(files)))
	// fmt.Println(files)
	for _, fullpath := range files {

		// get the post name
		linkAddress := filenameFromPath(fullpath)
		fmt.Println("DEV filename:", linkAddress)

		// get the title
		title := titleFromFullpath(fullpath)
		fmt.Println("DEV title:", title)

		// get the pretty date
		date := prettyDateFromPath(fullpath)
		fmt.Println("DEV date:", date)

		// get the summary
		summary := summaryFromFile(fullpath)
		fmt.Println("DEV summary:", summary)

		// get the body
		body := bodyFromFile(fullpath)
		fmt.Println("DEV body:", body)

		// a = append(a, Post{title, date, summary, body, fullpath, linkAddress})
	}
	// return a
}

// getPathsWMarkdownFiles takes a path and will return all full paths with .md files in string slice.
func getPathsWMarkdownFiles(path string) []string {
	// recursively get paths from 'path'
	var paths []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		paths = append(paths, path)
		return nil
	})
	if err != nil {
		log.Printf("error walking the path %q: %v\n", path, err)
	}

	// from 'paths', get paths with .md in them, discard all others
	var files []string
	for _, p := range paths {
		_, file := filepath.Split(p)
		if strings.Contains(file, ".md") {
			// now I have the paths with .md
			files = append(files, p)
		}
	}

	// return the full paths with filesname ending in md
	return files
}

// filenameFromPath uses the filename to create the link address used.
func filenameFromPath(fullpath string) string {
	_, filename := filepath.Split(fullpath)
	filename = strings.ToLower(strings.Replace(filename, ".md", "", -1))
	filename = strings.ToLower(strings.Replace(filename, "---", "-", -1))
	filename = strings.ToLower(strings.Replace(filename, "--", "-", -1))
	filename = strings.ToLower(strings.Replace(filename, "(", "", -1))
	filename = strings.ToLower(strings.Replace(filename, ")", "", -1))
	return filename
}

// titleFromFullpath uses the path to create the title used in the blog post.
func titleFromFullpath(fullpath string) string {
	_, filename := filepath.Split(fullpath)
	filename = strings.Replace(filename, ".md", "", -1)
	filename = strings.Replace(filename, "---", "++", -1)
	filename = strings.Replace(filename, "--", "+", -1)
	filename = strings.Replace(filename, "-", " ", -1)
	filename = strings.Replace(filename, "++", " - ", -1)
	filename = strings.Replace(filename, "+", "-", -1)
	return filename
}

// prettyDateFromPath uses the path to the file to create a date to use on the post, ie /2018/01/01/ = January 1, 2018.
func prettyDateFromPath(fullpath string) string {
	dir, _ := filepath.Split(fullpath)
	splitAtString := "/"
	if runtimeOS == "windows" {
		splitAtString = "\\"
	}
	splitDir := strings.Split(dir, splitAtString)

	// drop the first element which is posts
	splitDir = append(splitDir[:0], splitDir[1:]...)

	// set date variables
	year := splitDir[0]
	month := splitDir[1]
	day := splitDir[2]

	day = strings.TrimPrefix(day, "0")

	switch month {
	case "01":
		month = "January"
	case "02":
		month = "February"
	case "03":
		month = "March"
	case "04":
		month = "April"
	case "05":
		month = "May"
	case "06":
		month = "Jume"
	case "07":
		month = "July"
	case "08":
		month = "August"
	case "09":
		month = "September"
	case "10":
		month = "October"
	case "11":
		month = "November"
	case "12":
		month = "December"
	default:
		log.Println("received month \"", month, "\"")
	}

	return month + " " + day + ", " + year
}

// summaryFromFile opens the file, reads in the first line, and returns a string.
func summaryFromFile(f string) string {
	fileread, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println("ERROR: summary ioutil.ReadFile", err)
	}
	lines := strings.Split(string(fileread), "\n")
	return string(lines[0])
}

// bodyFromFile opens the file, reads from the 3rd line down, blackfriday parses it, and returns a string.
func bodyFromFile(f string) string {
	fileread, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println("ERROR: body ioutil.ReadFile", err)
	}
	lines := strings.Split(string(fileread), "\n")
	body := strings.Join(lines[2:len(lines)], "\n")
	return string(blackfriday.Run([]byte(body)))
}
