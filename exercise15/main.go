package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

//main function initiate the code execution. Different handles are created
//Each handle has its execution. The ListenAndServe takes the address and handler.
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/debug", sourceCodeHandler)
	mux.HandleFunc("/panic", panicHandler)
	log.Fatal(http.ListenAndServe(":3000", devMw(mux)))
}

func sourceCodeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.FormValue("path")
	lineStr := r.FormValue("line")
	line, _ := strconv.Atoi(lineStr)
	b, err := ioutil.ReadFile(path)
	fmt.Println("Path", path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return

	}
	str := string(b)

	var lines [][2]int
	if line > 0 {
		lines = append(lines, [2]int{line, line})
	}

	lexer := lexers.Get("go")
	iterator, _ := lexer.Tokenise(nil, str)

	style := styles.Get("github")

	formatter := html.New(html.TabWidth(2), html.WithLineNumbers(), html.LineNumbersInTable(), html.HighlightLines(lines))
	w.Header().Set("Content-Type", "text/html")
	formatter.Format(w, style, iterator)

	// err = quick.Highlight(w, str, "go", "html", "monokai")
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

//the function checks for panic if any. And if found defer is called. The function returns handler
func devMw(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				//log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, createLinks(string(stack)))
				//fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, makeLinks(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

//the default handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")

}

// panicHandler handle the request at /panic url
// and call the panic method
func panicHandler(w http.ResponseWriter, r *http.Request) {
	funcThatPanic()
}

// funcThatPanic call panic
func funcThatPanic() {
	panic("ho no!")
}

func createLinks(Trace string) string {

	lines := strings.Split(Trace, "\n")
	for li, line := range lines {
		if !strings.HasPrefix(line, "\t") {
			continue
		}
		val := strings.SplitAfter(lines[li], "\t")
		output := strings.SplitAfter(val[1], ":")[0]
		//lineNo := strings.SplitAfter(val[1], ":")
		lineNo := strings.SplitAfter(strings.SplitAfter(val[1], ":")[1], " ")[0]
		output = strings.TrimRight(output, ":")
		v := url.Values{}
		v.Set("path", output)
		lineNo = strings.Trim(lineNo, " ")
		v.Set("line", lineNo)
		//	lines[li] = "<a href=\"/debug?line=" + lineNo + "&path=" + output + "\">" + output + "</a>" + " lineNo: " + lineNo
		lines[li] = "<a href=\"/debug?" + v.Encode() + "\">" + output + "</a>" + " lineNo: " + lineNo

		var lineNumber strings.Builder
		lineNumber.WriteString(lineNo)
	}
	result := strings.Join(lines, "\n")
	//fmt.Println("output:", result)
	return result

}
