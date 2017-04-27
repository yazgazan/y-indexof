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
	"os"
	"path"

	"github.com/docker/go-units"
)

type indexItem struct {
	Name         string
	Path         string
	FullPath     string
	DownloadPath string
	Size         int64
	HumanSize    string
	Mode         string
	ModTime      string
	fileInfo     os.FileInfo
	Type         fileType
	fullPath     string
	IsDir        bool
}

func (file *indexItem) Populate(
	info os.FileInfo,
	context *indexContext,
	config Config) {
	file.Name = info.Name()
	file.Path = path.Join(context.Path, info.Name())
	file.FullPath = path.Join(context.FullPath, info.Name())
	file.fullPath = file.FullPath
	file.Size = info.Size()
	file.HumanSize = units.HumanSize(float64(info.Size()))
	file.Mode = info.Mode().String()
	file.ModTime = info.ModTime().Format("02-01-2006 03:04")
	file.IsDir = info.IsDir()
	file.fileInfo = info
	file.ResolveType(config)
	file.DownloadPath = generateDownloadPath(file, info, context, config)
}

func generateDownloadPath(file *indexItem, info os.FileInfo, context *indexContext, config Config) string {
	if info.IsDir() {
		return file.Path + "/"
	}
	if len(context.DownloadPrefix) != 0 {
		return path.Join(context.DownloadPrefix, file.Path)
	}
	return file.Path
}

func (file *indexItem) ResolveType(config Config) {
	allType, ok := config.Types["All"]
	if ok {
		file.Type = allType
	}
	if file.IsDir {
		_type, ok := config.Types["Folder"]
		if ok {
			file.Type.Merge(_type)
		}
	} else {
		defaultType, ok := config.Types["Default"]
		if ok {
			file.Type.Merge(defaultType)
		}
		ext := path.Ext(file.Name)
		if ext == "" {
			return
		}
		for _, curType := range config.Types {
			for _, typeExt := range curType.Exts {
				if typeExt == ext {
					file.Type.Merge(curType)
					return
				}
			}
		}
	}
}
