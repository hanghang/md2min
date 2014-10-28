package md2min

var templContent = `<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<meta http-equiv="X-UA-Compatible" content="chrome=1">
        <meta name="viewport" content="width=device-width, initial-scale=1.0, user-scalable=yes">
		<title></title>
<link rel="stylesheet" href="/css/style.css" type="text/css"/> 
<script type="text/javascript" src='/js/highlight.pack.js'></script>
<script>hljs.initHighlightingOnLoad();</script>
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
