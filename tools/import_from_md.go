package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	//"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func walk(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// if path contrains space, replace it with '_' and rename the file
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func parse(file_name string) ([]byte, []byte) {
	mdBytes, err := ioutil.ReadFile(file_name)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	// remove front empty line and write back to file
	for mdBytes[0] == '\n' {
		mdBytes = mdBytes[1:]
	}
	ioutil.WriteFile(file_name, mdBytes, 0644)

	frontMatter := []byte{}
	content := []byte{}
	// get front matter from markdown file if exist
	// if start with --- and end with ---, then it is front matter
	if mdBytes[0] == '-' && mdBytes[1] == '-' && mdBytes[2] == '-' {
		// get the end of front matter
		var end int

		find := false
		for i := 3; i < len(mdBytes)-3; i++ {
			if mdBytes[i] == '-' && mdBytes[i+1] == '-' && mdBytes[i+2] == '-' {
				end = i + 3
				find = true
				break
			}
		}
		// get front matter
		if find {
			frontMatter = mdBytes[3 : end-3]
			content = mdBytes[end:]
			//fmt.Println(string(frontMatter))
		} else {
			//fmt.Println("no front matter")
			log.Println("no front matter found")
			content = mdBytes
		}
	}
	return frontMatter, content

}

type BlogFrontMatter struct {
	Title      string
	Date       time.Time
	Author     string
	Tags       []string
	Categories []string
}

// func main() {
//     db.Init()
//     root := "posts"
//     files, _ := walk(root)
//     for _, f := range files {
//         if strings.HasSuffix(f, ".md") {
//             frontMatter, content := parse(f)
//             // parse front front matter use yaml
//             var blogFrontData BlogFrontMatter
//             yaml.Unmarshal(frontMatter, &blogFrontData)
//             //log.Println(blogFrontData)
//             _ = content
//             _ = blogFrontData
//             log.Println(string(content))
//             var dbBlogData db.BlogData
//             dbBlogData.Title = blogFrontData.Title
//             dbBlogData.Content = string(content)
//             dbBlogData.Datetime = blogFrontData.Date
//             dbBlogData.Author = blogFrontData.Author
//             dbBlogData.Tags = blogFrontData.Tags
//             dbBlogData.Categories = blogFrontData.Categories
//             db.InsertPost(dbBlogData)
//         }
//     }
// }

const dbTypev1 = "sqlite3"
const dbPathv1 = "./blogv1.db"
