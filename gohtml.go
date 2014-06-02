//	Copyright 2013 slowfei And The Contributors All rights reserved.
//
//	Software Source Code License Agreement (BSD License)
//
//  Create on 2013-11-30
//  Update on 2014-06-01
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
	isCmdDir = flag.Bool("cmddir", true, "is current cmd directory find files? default true. false is execute file directory find.")
	path     = flag.String("path", "", "specify the run path")
	compact  = flag.Bool("compact", false, "is compact html code? default false")
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

	if 0 == len(*path) {
		if *isCmdDir {
			template.SetBaseDir(SFFileManager.GetCmdDir())
		} else {
			template.SetBaseDir(SFFileManager.GetExecDir())
		}
	} else {
		template.SetBaseDir(*path)
	}

	template.SetCache(false)
	template.SetCompactHTML(*compact)

	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
