//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-11-30
//  Update on 2013-12-01
//  Email  slowfei@foxmail.com
//  Home   http://www.slowfei.com

//	go html design static server
package main

import (
	"flag"
	"fmt"
	"github.com/slowfei/gosfcore/utils/filemanager"
	"github.com/slowfei/leafveingo/template"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	port     = flag.Int("p", 8080, "http port default 8080")
	suffix   = flag.String("s", "tpl", "request url suffix default tpl")
	charset  = flag.String("c", "utf-8", "content-type charset default utf-8")
	template = LVTemplate.SharedTemplate()
)

func HtmlOut(rw http.ResponseWriter, req *http.Request) {

	reqPath := req.URL.Path

	if '/' == reqPath[len(reqPath)-1] {
		reqPath += "index." + *suffix
	}

	filePath := filepath.Join(template.BaseDir(), reqPath)

	isExists, isDir, _ := SFFileManager.Exists(filePath)
	if !isExists || isDir {
		http.NotFound(rw, req)
		return
	}

	fmt.Println("requet page:", filePath)

	if strings.HasSuffix(reqPath, *suffix) {
		e := req.ParseForm()
		if nil != e {
			fmt.Println("parse form error:", e)
			return
		}
		rw.Header().Set("Content-Type", "text/html; charset="+*charset)
		err := template.Execute(rw, LVTemplate.NewTemplateValue(reqPath, req.Form))
		if nil != err {
			fmt.Println("template error: ", err)
		}
	} else {
		http.ServeFile(rw, req, filePath)
	}

}

func main() {
	flag.Parse()
	http.HandleFunc("/", HtmlOut)
	template.SetBaseDir(SFFileManager.GetExecDir())
	template.SetCache(false)
	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
