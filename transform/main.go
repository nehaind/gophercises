package main

import (
	"errors"
	"fmt"
	"gophercises/transform/transform"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./img"))
	mux.Handle("/img/", http.StripPrefix("/img", fs))
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/upload", upLoadHandler)

	log.Fatal(http.ListenAndServe(":3000", mux))

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `<html><body>
	<form action="/upload" method="post" enctype="multipart/form-data">
		<input type="file" name="image">
		<button type="submit">Upload Image</button>
	</form>
	</body></html>`
	fmt.Fprint(w, html)
}

func upLoadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		ext := filepath.Ext(header.Filename)[1:]
		a, err := generateImage(file, ext, 33, transform.ModeCircle)
		if err != nil {
			panic(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		b, err := generateImage(file, ext, 33, transform.ModeEllipse)
		if err != nil {
			panic(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		c, err := generateImage(file, ext, 33, transform.ModePolygon)
		if err != nil {
			panic(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		file.Seek(0, 0)
		d, err := generateImage(file, ext, 33, transform.ModeCombo)
		if err != nil {
			panic(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		html := `<html><body>
			{{range .}}
				<img src="/{{.}}">
			{{end}}
			</body></html>`
		tpl := template.Must(template.New("").Parse(html))
		images := []string{a, b, c, d}
		tpl.Execute(w, images)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func generateImage(file io.Reader, ext string, number int, mode transform.Mode) (string, error) {
	out, err := transform.Transform(file, ext, number, mode)

	//	out, err := primitive.Transform(r, ext, numShapes, primitive.WithMode(mode))
	if err != nil {
		return "", err
	}
	outFile, err := tempfile("", ext)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	io.Copy(outFile, out)
	return outFile.Name(), nil

	//some changes required
	// Outfile, err := os.Open(out) // For read access.
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// temp, err := tempfile("", ext)
	// defer Outfile.Close()
	// io.Copy(temp, Outfile)
	// return temp.Name(), nil
}

func tempfile(prefix, ext string) (*os.File, error) {
	in, err := ioutil.TempFile("./img/", prefix)
	if err != nil {
		return nil, errors.New("main: failed to create temporary file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}

// w.Header().Set("Content-Type", "image/png")
// switch ext {
// case "jpg":
// 	fallthrough
// case "jpeg":
// 	w.Header().Set("Content-Type", "image/jpeg")
// case "png":
// 	w.Header().Set("Content-Type", "image/png")
// default:
// 	http.Error(w, fmt.Sprintf("invalid image type %s", ext), http.StatusBadRequest)
// 	return
// }

//defer os.Remove(in.Name())
// outputBuffer, err := ioutil.ReadFile(out)
// ioutil.WriteFile("out1.png", outputBuffer, 0644)
// w.Write(outputBuffer)
// fmt.Println(w)

// io.Copy(w, out)

// inFile, err := os.Open("/home/neha/dev/src/gophercises/out.png")
// if err != nil {
// 	panic(err)
// }
// defer inFile.Close()
// _, err = transform.Transform(inFile, 33)
// if err != nil {
// 	fmt.Println("error in transform", err)
// }
