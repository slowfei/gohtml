go html design static server
======

	htmldir 		
      ├─images 				// 一些网站所使用的公用图片目录
      ├─js 					// javascript 文件目录
      ├─themes 				// css主题存放目录，这样做的目的主要是可以方便切换皮肤
      │  ├─core 			// css核心文件目录
      │  │  └─core.css 	  	// css核心文件
      │  └─default 			// 默认css皮肤主题目录
      │     ├─css 			// 默认皮肤主题的css存放目录
      │     │  ├─images  	// 默认皮肤主题所使用的图片目录，这样style.css访问图片路径时更好操作。
      │     │  └─style.css 	// 默认皮肤主题的css文件
      │     └─js 			// 默认皮肤主题所使用的javascript文件目录
      │        └─init.js 	// 默认皮肤主题所需要初始化的javascript函数或布局所使用的文件
      │
      ├─index.tpl 			// 静态页面模版
      └─gohtml 		 		// gohtml执行文件

Run

	$ ./gohtml

Browser open

      http://localhost:8080/

### Install

	go get github.com/slowfei/gosfcore

	go get github.com/slowfei/leafveingo/template
	
    go get github.com/slowfei/gohtml

### Use

	$go install github.com/slowfei/gohtml

	$gohtml -h
		-c="utf-8": content-type charset default utf-8
 		-p=8080: http port default 8080
 		-s="tpl": request url suffix default tpl


#### [LICENSE](https://github.com/slowfei/gosfcore/blob/master/LICENSE)

Copyright 2013 slowfei And The Contributors All rights reserved.

Software Source Code License Agreement (BSD License)

###
