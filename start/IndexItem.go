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

	"github.com/yazgazan/y-indexof/utils"
)

type IndexItem struct {
	Name         string
	Path         string
	FullPath     string
	DownloadPath string
	Size         int64
	HumanSize    string
	Mode         string
	ModTime      string
	IsDir        bool
	fileInfo     os.FileInfo
	Type         Type
	fullPath     string
}

func (file *IndexItem) Populate(
	info os.FileInfo,
	context *IndexContext,
	config Config) {
	file.Name = info.Name()
	file.Path = path.Join(context.Path, info.Name())
	file.FullPath = path.Join(context.FullPath, info.Name())
	file.fullPath = file.FullPath
	file.Size = info.Size()
	file.HumanSize = utils.GetHumanReadableSize(info.Size())
	file.Mode = info.Mode().String()
	file.ModTime = info.ModTime().Format("02-01-2006 03:04")
	file.IsDir = info.IsDir()
	file.fileInfo = info
	file.ResolveType(config)
	file.DownloadPath = GenerateDownloadPath(file, info, context, config)
}

func GenerateDownloadPath(
	file *IndexItem,
	info os.FileInfo,
	context *IndexContext,
	config Config) string {
	if info.IsDir() {
		return file.Path + "/"
	}
	if len(context.DownloadPrefix) != 0 {
		return path.Join(context.DownloadPrefix, file.Path)
	}
	return file.Path
}

func (file *IndexItem) ResolveType(config Config) {
	allType, ok := config.Types["All"]
	if ok == true {
		file.Type = allType
	}
	if file.IsDir == true {
		_type, ok := config.Types["Folder"]
		if ok == true {
			file.Type.Merge(_type)
		}
	} else {
		defaultType, ok := config.Types["Default"]
		if ok == true {
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
