package filesystem

import (
	"sort"
	"strings"
)

type Dir struct {
	files map[string]string
	dirs  map[string]*Dir
}

func initDir() *Dir {
	d := new(Dir)
	d.files = make(map[string]string)
	d.dirs = make(map[string]*Dir)
	return d
}

type FileSystem struct {
	root *Dir
}

/*
 题解：
	可以这样理解，文件系统本质上分为两种对象，一个是文件，一个是目录（*nix系统下面）
	那么对于一个目录的ls， 就是返回目录的名称，对于一个文件的ls，就是返回文件名
	mkdir 在这里的实现是，更像linux 下面的mkdir -p 的命令,下面的实现中有多一个函数是not p 的实现

*/

func Constructor() FileSystem {
	root := initDir()
	return FileSystem{
		root: root,
	}
}

func (this *FileSystem) Ls(path string) []string {
	r := this.root
	files := make([]string, 0)
	if path != "/" {
		d := strings.Split(path, "/")
		for i := 1; i < len(d); i++ {
			r = r.dirs[d[i]]
		}
		if val, ok := r.files[d[len(d)-1]]; ok {
			files = append(files, val)
		} else {
			r = r.dirs[d[len(d)-1]]
		}
	}
	for item := range r.files {
		files = append(files, item)
	}
	for item := range r.dirs {
		files = append(files, item)
	}
	sort.Strings(files)
	return files
}

func (this *FileSystem) Mkdir(path string) {
	r := this.root
	d := strings.Split(path, "/")
	for i := 1; i < len(d); i++ {
		if _, ok := r.dirs[d[i]]; !ok {
			r.dirs[d[i]] = initDir()
		}
		r = r.dirs[d[i]]
	}
}

func (this *FileSystem) MkdirNotParent(path string) bool {
	r := this.root
	d := strings.Split(path, "/")
	for i := 1; i < len(d); i++ {
		if _, ok := r.dirs[d[i]]; !ok {
			if i != len(d)-1 {
				return false
			}
			r.dirs[d[i]] = initDir()
		}
		r = r.dirs[d[i]]
	}
	return true
}

func (this *FileSystem) AddContentToFile(filePath string, content string) {
	r := this.root
	d := strings.Split(filePath, "/")
	for i := 1; i < len(d)-1; i++ {
		if _, ok := r.files[d[i]]; !ok {
			r = r.dirs[d[i]]
		}
	}
	files := r.files[d[len(d)-1]]
	files += content
}

func (this *FileSystem) ReadContentFromFile(filePath string) string {
	r := this.root
	d := strings.Split(filePath, "/")
	for i := 1; i < len(d)-1; i++ {
		if _, ok := r.files[d[i]]; !ok {
			r = r.dirs[d[i]]
		}
	}
	return r.files[d[len(d)-1]]
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
