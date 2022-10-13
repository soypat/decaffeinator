package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	vignettes, err := ParseDirExercises(".")
	if err != nil {
		log.Fatal(err)
	}
	withZig, err := os.Create("tagalong_w_zig.md")
	if err != nil {
		log.Fatal(err)
	}
	tagalong, _ := os.Create("tagalong.md")
	tagalongCodeOnly, _ := os.Create("tagalongCode.md")
	defer withZig.Close()
	const codeLevel = "###"
	for _, vig := range vignettes {
		fmt.Fprintf(withZig, "%s\n", vig.MD)
		fmt.Fprintf(tagalong, "%s\n", vig.MD)
		if vig.Go == "" || vig.Python == "" {
			continue
		}
		fmt.Fprintf(withZig, codeLevel+" Python (%s)\n```python\n%s\n```\n", vig.Name, vig.Python)
		fmt.Fprintf(withZig, codeLevel+" Go (%s)\n```go\n%s\n```\n", vig.Name, vig.Go)

		fmt.Fprintf(tagalong, codeLevel+" Python (%s)\n```python\n%s\n```\n", vig.Name, vig.Python)
		fmt.Fprintf(tagalong, codeLevel+" Go (%s)\n```go\n%s\n```\n", vig.Name, vig.Go)

		fmt.Fprintf(tagalongCodeOnly, "\n# %s\n", vig.Name)
		fmt.Fprintf(tagalongCodeOnly, codeLevel+" Python (%s)\n```python\n%s\n```\n", vig.Name, vig.Python)
		fmt.Fprintf(tagalongCodeOnly, codeLevel+" Go (%s)\n```go\n%s\n```\n", vig.Name, vig.Go)
		if vig.Zig != "" {
			fmt.Fprintf(withZig, codeLevel+" Zig (%s)\n```zig\n%s\n```\n", vig.Name, vig.Zig)
		}
	}
}

type Header struct {
	Num  int
	Name string
}

type Vignette struct {
	Header
	MD     string
	Go     string
	Python string
	Zig    string
}

func ParseDirExercises(dir string) ([]Vignette, error) {
	found, err := ParseDir(dir)
	if err != nil {
		return nil, err
	}
	vignettes := make([]Vignette, len(found))
	for i := range vignettes {
		vignettes[i].Header = found[i]
		subdir := filepath.Join(dir, vignettes[i].Code())
		entries, err := os.ReadDir(subdir)
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			filename := entry.Name()
			if entry.IsDir() {
				continue
			}
			path := filepath.Join(subdir, filename)
			fp, err := os.Open(path)
			if err != nil {
				return nil, err
			}
			b, err := io.ReadAll(fp)
			if err != nil {
				return nil, err
			}
			switch {
			case filename == "README.md":
				vignettes[i].MD = string(b)
			case filename == vignettes[i].Name+".zig":
				vignettes[i].Zig = string(b)
			case filename == vignettes[i].Name+".go":
				vignettes[i].Go = string(b)
			case filename == vignettes[i].Name+".py":
				vignettes[i].Python = string(b)
			}
		}
	}
	sort.Sort(ByFilenameNumber(vignettes))
	return vignettes, nil
}

// ParseDir parses filenames and sorts according to number.
func ParseDir(dir string) ([]Header, error) {
	if dir == "" {
		dir = "."
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var found []Header
	for _, entry := range entries {
		filename := entry.Name()
		if !entry.IsDir() {
			continue
		}
		path := filepath.Join(dir, filename)
		exercise, err := ParseHeader(path)
		if err != nil {
			continue
			return nil, errors.New("parsing directory \"" + dir + "\": " + err.Error())
		}
		found = append(found, exercise)
	}
	sort.Sort(ByHeaderNumber(found))

	return found, nil
}

type ByFilenameNumber []Vignette

func (a ByFilenameNumber) Len() int           { return len(a) }
func (a ByFilenameNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFilenameNumber) Less(i, j int) bool { return a[i].Num < a[j].Num }

func (f Header) Code() string {
	return fmt.Sprintf("%03d-%s", f.Num, f.Name)
}

type ByHeaderNumber []Header

func (a ByHeaderNumber) Len() int           { return len(a) }
func (a ByHeaderNumber) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHeaderNumber) Less(i, j int) bool { return a[i].Num < a[j].Num }

func ParseHeader(code string) (Header, error) {
	if len(code) < 8 {
		return Header{}, errors.New(code + " exercise filename name must be more than 8 characters long")
	}
	n, err := strconv.Atoi(code[0:3])
	if err != nil {
		return Header{}, err
	}
	en := Header{Num: n, Name: code[4:]}
	if en.Code() != code {
		return Header{}, fmt.Errorf("generated path %q does not match argument path %q", en.Code(), code)
	}
	return en, nil
}
