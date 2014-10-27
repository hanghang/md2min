package md2min

var templContent = `<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="chrome=1">
		<title></title>
		<style text="text/css">
body {
    overflow-x: hidden;
}
.container {
		margin: 0 auto;
		width: 1020px;
}
.nav-wrap {
    display: block;
    height: 650px;
    float: left;
    position: fixed;
		padding: 3px;
		margin: 0;
		border: 0;
		font: 13px Helvetica, arial, freesans, clean, sans-serif;
		line-height: 1.4;
    {{.MenuWrapStyle}}
}
.nav {
		display: block;
		float: left;
		position: fixed;
		padding: 3px;
		border-radius: 2px;
		margin: 40px 0 0 0;
		border: 0;
		font: 13px Helvetica, arial, freesans, clean, sans-serif;
		line-height: 1.4;
    {{.MenuStyle}}
}
.nav ul {
		background: #fafafb;
		border-radius: 2px;
		border: 1px solid #d8d8d8;
		margin: 0;
		padding: 0;
		list-style: none;
		display: block;
}
.nav ul li {
		display: list-item;
		border-top: 1px solid #fff;
		border-bottom: 1px solid #eee;
}
.nav ul li a {
		display: block;
		padding: 8px 10px 8px 8px;
		text-shadow: 0 1px 0 #fff;
		border-left: 2px solid #fafafb;
		color: #4183c4;
		text-decoration: none;
}
.nav ul li a:-webkit-any-link {
		cursor: auto;
}
.markdown-body {
		font: 13px Helvetica, arial, freesans, clean, sans-serif;
		display: block;
		margin: 20px;
		background-color: #F9F9F5;
    		color: #2C3E50;
    {{.ContentStyle}}
    font-size: 14px;
    line-height: 1.6;
}
.markdown-body > *:first-child {
    margin-top: 0 !important;
}
.markdown-body > *:last-child {
    margin-bottom: 0 !important;
}
.markdown-body a {
    text-decoration: none;
}
.markdown-body a:hover {
    text-decoration: underline;
}
.markdown-body a.absent {
    color: #CC0000;
}
.markdown-body a.anchor {
    bottom: 0;
    cursor: pointer;
    display: block;
    left: 0;
    margin-left: -30px;
    padding-left: 30px;
    position: absolute;
    top: 0;
}
.markdown-body h1, .markdown-body h2, .markdown-body h3, .markdown-body h4, .markdown-body h5, .markdown-body h6 {
    cursor: text;
    font-weight: bold;
    margin: 20px 0 10px;
    padding: 0;
    position: relative;
}
.markdown-body h1 .mini-icon-link, .markdown-body h2 .mini-icon-link, .markdown-body h3 .mini-icon-link, .markdown-body h4 .mini-icon-link, .markdown-body h5 .mini-icon-link, .markdown-body h6 .mini-icon-link {
    color: #000000;
    display: none;
}
.markdown-body h1:hover a.anchor, .markdown-body h2:hover a.anchor, .markdown-body h3:hover a.anchor, .markdown-body h4:hover a.anchor, .markdown-body h5:hover a.anchor, .markdown-body h6:hover a.anchor {
    line-height: 1;
    margin-left: -22px;
    padding-left: 0;
    text-decoration: none;
    top: 15%;
}
.markdown-body h1:hover a.anchor .mini-icon-link, .markdown-body h2:hover a.anchor .mini-icon-link, .markdown-body h3:hover a.anchor .mini-icon-link, .markdown-body h4:hover a.anchor .mini-icon-link, .markdown-body h5:hover a.anchor .mini-icon-link, .markdown-body h6:hover a.anchor .mini-icon-link {
    display: inline-block;
}
.markdown-body h1 tt, .markdown-body h1 code, .markdown-body h2 tt, .markdown-body h2 code, .markdown-body h3 tt, .markdown-body h3 code, .markdown-body h4 tt, .markdown-body h4 code, .markdown-body h5 tt, .markdown-body h5 code, .markdown-body h6 tt, .markdown-body h6 code {
    font-size: inherit;
}
.markdown-body h1 {
    color: #000000;
    font-size: 28px;
}
.markdown-body h2 {
    border-bottom: 1px solid #CCCCCC;
    color: #000000;
    font-size: 24px;
}
.markdown-body h3 {
    font-size: 18px;
}
.markdown-body h4 {
    font-size: 16px;
}
.markdown-body h5 {
    font-size: 14px;
}
.markdown-body h6 {
    color: #777777;
    font-size: 14px;
}
.markdown-body p, .markdown-body blockquote, .markdown-body ul, .markdown-body ol, .markdown-body dl, .markdown-body table, .markdown-body pre {
    margin: 15px 0;
}
.markdown-body li ul, .markdown-body li ol {
    margin-top: 0;
}
.markdown-body hr {
    border: 0 none;
    color: #CCCCCC;
    height: 4px;
    padding: 0;
}
.markdown-body > h2:first-child, .markdown-body > h1:first-child, .markdown-body > h1:first-child + h2, .markdown-body > h3:first-child, .markdown-body > h4:first-child, .markdown-body > h5:first-child, .markdown-body > h6:first-child {
    margin-top: 0;
    padding-top: 0;
}
.markdown-body a:first-child h1, .markdown-body a:first-child h2, .markdown-body a:first-child h3, .markdown-body a:first-child h4, .markdown-body a:first-child h5, .markdown-body a:first-child h6 {
    margin-top: 0;
    padding-top: 0;
}
.markdown-body h1 + p, .markdown-body h2 + p, .markdown-body h3 + p, .markdown-body h4 + p, .markdown-body h5 + p, .markdown-body h6 + p {
    margin-top: 0;
}
.markdown-body li p.first {
    display: inline-block;
}
.markdown-body ul, .markdown-body ol {
    padding-left: 30px;
}
.markdown-body ul.no-list, .markdown-body ol.no-list {
    list-style-type: none;
    padding: 0;
}
.markdown-body ul li > *:first-child, .markdown-body ol li > *:first-child {
    margin-top: 0;
}
.markdown-body ul ul, .markdown-body ul ol, .markdown-body ol ol, .markdown-body ol ul {
    margin-bottom: 0;
}
.markdown-body dl {
    padding: 0;
}
.markdown-body dl dt {
    font-size: 14px;
    font-style: italic;
    font-weight: bold;
    margin: 15px 0 5px;
    padding: 0;
}
.markdown-body dl dt:first-child {
    padding: 0;
}
.markdown-body dl dt > *:first-child {
    margin-top: 0;
}
.markdown-body dl dt > *:last-child {
    margin-bottom: 0;
}
.markdown-body dl dd {
    margin: 0 0 15px;
    padding: 0 15px;
}
.markdown-body dl dd > *:first-child {
    margin-top: 0;
}
.markdown-body dl dd > *:last-child {
    margin-bottom: 0;
}
.markdown-body blockquote {
    border-left: 4px solid #DDDDDD;
    color: #777777;
    padding: 0 15px;
}
.markdown-body blockquote > *:first-child {
    margin-top: 0;
}
.markdown-body blockquote > *:last-child {
    margin-bottom: 0;
}
.markdown-body table th {
    font-weight: bold;
}
.markdown-body table th, .markdown-body table td {
    border: 1px solid #CCCCCC;
    padding: 6px 13px;
}
.markdown-body table tr {
    background-color: #FFFFFF;
    border-top: 1px solid #CCCCCC;
}
.markdown-body table tr:nth-child(2n) {
    background-color: #F8F8F8;
}
.markdown-body img {
    max-width: 100%;
}
.markdown-body span.frame {
    display: block;
    overflow: hidden;
}
.markdown-body span.frame > span {
    border: 1px solid #DDDDDD;
    display: block;
    float: left;
    margin: 13px 0 0;
    overflow: hidden;
    padding: 7px;
    width: auto;
}
.markdown-body span.frame span img {
    display: block;
    float: left;
}
.markdown-body span.frame span span {
    clear: both;
    color: #333333;
    display: block;
    padding: 5px 0 0;
}
.markdown-body span.align-center {
    clear: both;
    display: block;
    overflow: hidden;
}
.markdown-body span.align-center > span {
    display: block;
    margin: 13px auto 0;
    overflow: hidden;
    text-align: center;
}
.markdown-body span.align-center span img {
    margin: 0 auto;
    text-align: center;
}
.markdown-body span.align-right {
    clear: both;
    display: block;
    overflow: hidden;
}
.markdown-body span.align-right > span {
    display: block;
    margin: 13px 0 0;
    overflow: hidden;
    text-align: right;
}
.markdown-body span.align-right span img {
    margin: 0;
    text-align: right;
}
.markdown-body span.float-left {
    display: block;
    float: left;
    margin-right: 13px;
    overflow: hidden;
}
.markdown-body span.float-left span {
    margin: 13px 0 0;
}
.markdown-body span.float-right {
    display: block;
    float: right;
    margin-left: 13px;
    overflow: hidden;
}
.markdown-body span.float-right > span {
    display: block;
    margin: 13px auto 0;
    overflow: hidden;
    text-align: right;
}
.markdown-body code, .markdown-body tt {
		font-family: Consolas, "Liberation Mono", Courier, monospace;
    background-color: #F8F8F8;
    border: 1px solid #EAEAEA;
    border-radius: 3px 3px 3px 3px;
    margin: 0 2px;
    padding: 0 5px;
    white-space: nowrap;
}
.markdown-body pre > code {
		font-size: 14px;
    background: none repeat scroll 0 0 transparent;
    border: medium none;
    margin: 0;
    padding: 0;
    white-space: pre;
}
.markdown-body .highlight pre, .markdown-body pre {
    background-color: #F8F8F8;
    border: 1px solid #CCCCCC;
    border-radius: 3px 3px 3px 3px;
    font-size: 13px;
    line-height: 19px;
    overflow: auto;
    padding: 6px 10px;
}
.markdown-body pre code, .markdown-body pre tt {
    background-color: transparent;
    border: medium none;
}
.container .logo {
    display: block;
    text-align: right;
    font-size: 10px;
}
{{.ScrollBar}}
		</style>
	</head>
	<body>
  <div class="container">
    <div class="nav-wrap">
      <div class="nav">{{.ListMenu}}</div>
      {{.MenuLogo}}
    </div>
  	<div class="markdown-body">{{.Content}}{{.ContentLogo}}</div>
  <div>
	</body>
</html>
`
