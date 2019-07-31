{{define "navbar"}}
<nav class="navbar navbar-default navbar-fixed-top" role="navigation">
    <div class="container">
        <a class="navbar-brand" href="#">我的博客</a>
        <ul class="nav navbar-nav">
            <li {{if .isHome}}class="active"{{end}}><a href="/">首页</a></li>
            <li {{if .isCategory}}class="active"{{end}}><a href="/category">分类</a></li>
            <li {{if .isTopic}}class="active"{{end}}><a href="topic">文章</a></li>
        </ul>

        <ul class="nav navbar-nav navbar-right">
            {{if .isLogin}}
            <li><a href="/login?exit=true">退出</a></li>
            {{else}}
            <li><a href="/login">登录</a></li>
            {{end}}
        </ul>
    </div>
</nav>
{{end}}