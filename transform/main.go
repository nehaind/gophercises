package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"gophercises/Exercise_18/primitive"

	homedir "github.com/mitchellh/go-homedir"
)

var listenAndServeFunc = http.ListenAndServe

func main() {
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img")
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(imgPath))
	mux.Handle("/img/", http.StripPrefix("/img", fs))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/modify/", modify)
	log.Fatal(listenAndServeFunc(":8888", mux))
}

func index(w http.ResponseWriter, r *http.Request) {
	html := `<html><body>
			<form action="/upload" method="post" enctype="multipart/form-data">
				<input type="file" name="image">
				<button type="submit">Upload Image</button>
			</form>
			</body></html>`
	fmt.Fprint(w, html)
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload works")
	var err error
	file, h, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		e := filepath.Ext(h.Filename)[1:]
		finalFile, err := createTempFile("", e)
		if err == nil {
			defer finalFile.Close()
			io.Copy(finalFile, file)
			http.Redirect(w, r, "/modify/"+filepath.Base(finalFile.Name()), http.StatusFound)
		}
	}
	http.Error(w, "error occured in upload", http.StatusInternalServerError)
	return
}

func modify(w http.ResponseWriter, r *http.Request) {

	var err error
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img")
	f, err := os.Open(imgPath + "/" + filepath.Base(r.URL.Path))
	if err == nil {
		defer f.Close()
		ext := filepath.Ext(f.Name())[1:]
		modeStr := r.FormValue("mode")
		if modeStr == "" {
			renderModeChoices(w, r, f, ext)
			return
		}
		mode, err := strconv.Atoi(modeStr)
		if err == nil {
			nStr := r.FormValue("n")
			if nStr == "" {
				renderNumShapeChoices(w, r, f, ext, primitive.Mode(mode))
				return
			}
			_, err = strconv.Atoi(nStr)
			if err == nil {
				http.Redirect(w, r, "/img/"+filepath.Base(f.Name()), http.StatusFound)
				return
			}
		}
	}
	http.Error(w, "error occured in modify", http.StatusBadRequest)
}

func renderModeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string) {
	var err error
	opts := []genOpts{
		{N: 10, M: primitive.ModeCircle},
		{N: 10, M: primitive.ModeBeziers},
		{N: 10, M: primitive.ModePolygon},
		{N: 10, M: primitive.ModeCombo},
	}
	imgs, err := genImages(rs, ext, opts...)
	if err == nil {
		html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}">
					<img style="width: 20%;" src="/img/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
		tpl := template.Must(template.New("").Parse(html))
		type dataStruct struct {
			Name string
			Mode primitive.Mode
		}
		var data []dataStruct
		for i, img := range imgs {
			data = append(data, dataStruct{
				Name: filepath.Base(img),
				Mode: opts[i].M,
			})
		}
		err = tpl.Execute(w, data)
		return
	}
	http.Error(w, "error occured in render mode choices", http.StatusInternalServerError)
}

func renderNumShapeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string, mode primitive.Mode) {
	var err error
	opts := []genOpts{
		{N: 10, M: mode},
		{N: 20, M: mode},
		{N: 30, M: mode},
		{N: 40, M: mode},
	}
	imgs, err := genImages(rs, ext, opts...)
	if err == nil {
		html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}&n={{.NumShapes}}">
					<img style="width: 20%;" src="/img/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
		tpl := template.Must(template.New("").Parse(html))
		type dataStruct struct {
			Name      string
			Mode      primitive.Mode
			NumShapes int
		}
		var data []dataStruct
		for i, img := range imgs {
			data = append(data, dataStruct{
				Name:      filepath.Base(img),
				Mode:      opts[i].M,
				NumShapes: opts[i].N,
			})
		}
		err = tpl.Execute(w, data)
		return
	}
	http.Error(w, "Error occured in rendering number of choices", http.StatusInternalServerError)
}

type genOpts struct {
	N int
	M primitive.Mode
}

func genImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	var ret []string
	var err error
	var f string
	for _, opt := range opts {
		rs.Seek(0, 0)
		f, err = genImage(rs, ext, opt.N, opt.M)
		if err == nil {
			ret = append(ret, f)
		}
	}
	return ret, err
}

func genImage(r io.Reader, ext string, numShapes int, mode primitive.Mode) (string, error) {
	var outFile *os.File
	var err error
	var out io.Reader
	out, err = primitive.Transform(r, ext, numShapes, primitive.WithMode(mode))
	if err == nil {
		outFile, err = createTempFile("", ext)
		if err == nil {
			defer outFile.Close()
			io.Copy(outFile, out)
			return outFile.Name(), err
		}
	}
	return "", err
}

func createTempFile(name, e string) (*os.File, error) {
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img")
	f, err := ioutil.TempFile(imgPath+"/", name)
	if err != nil {
		return nil, errors.New("failed to create temporary file")
	}
	defer os.Remove(f.Name())
	return os.Create(fmt.Sprintf("%s.%s", f.Name(), e))
}
