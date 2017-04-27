/*
 * ----------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * <yazgazan@gmail.com> wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return.
 * ----------------------------------------------------------------------------
 */

package start

import (
	"encoding/json"
	"net/http"
	"os"
	"sort"
)

type indexContext struct {
	Path           string
	FullPath       string
	DownloadPrefix string
	Files          []indexItem
	JSONContext    string
	Sort           string
	UserDefined    map[string]string
}

const (
	sortAlpha = "alpha"
	sortDate  = "date"
	sortSize  = "size"
)

func (context *indexContext) InitSort(req *http.Request) {
	cookie, err := req.Cookie("sort")
	if err != nil {
		context.Sort = sortAlpha
		return
	}
	switch cookie.Value {
	case sortDate:
		context.Sort = sortDate
	case sortSize:
		context.Sort = sortSize
	case sortAlpha:
		context.Sort = sortAlpha
	}
	if context.Sort == "" {
		context.Sort = sortAlpha
	}
}

func (context *indexContext) InitContext(method *methodConfig, config Config) error {
	fullPath := method.FullPath

	context.Path = method.Path
	context.FullPath = method.FullPath
	context.DownloadPrefix = config.DownloadPrefix
	context.UserDefined = config.UserDefined

	if err := context.ReadDirInfos(config); err != nil {
		return err
	}

	switch context.Sort {
	case sortDate:
		sort.Sort(dateSortFiles(context.Files))
	case sortAlpha:
		sort.Sort(alphaSortFiles(context.Files))
	case sortSize:
		sort.Sort(sizeSortFiles(context.Files))
	}

	if !config.ShowFullPath {
		context.FullPath = context.Path
		for id := range context.Files {
			context.Files[id].FullPath = context.Files[id].Path
		}
	}
	context.CreateJSON()
	if !config.ShowFullPath {
		context.FullPath = fullPath
		for id := range context.Files {
			context.Files[id].FullPath = context.Files[id].fullPath
		}
	}

	return nil
}

func (context *indexContext) CreateJSON() {
	bytes, err := json.Marshal(context)

	if err != nil {
		panic(err) // shouldn't happen
	}
	context.JSONContext = string(bytes)
}

func isExcluded(file os.FileInfo) bool {
	name := file.Name()

	return (name != "" && name[0] == '.') || name == "indexof.toml"
}

func countFiles(files []os.FileInfo) int {
	count := 0

	for _, info := range files {
		if isExcluded(info) {
			continue
		}
		count++
	}

	return count
}

func (context *indexContext) ReadDirInfos(config Config) error {
	file, err := os.Open(context.FullPath)

	if err != nil {
		return err
	}

	fileInfos, err := file.Readdir(0)
	if err != nil {
		return err
	}
	filesCount := countFiles(fileInfos)
	context.Files = make([]indexItem, filesCount)

	i := 0
	for _, info := range fileInfos {
		if isExcluded(info) {
			continue
		}
		context.Files[i].Populate(info, context, config)
		i++
	}
	return nil
}
