## Hi there, I'm Fernando! 👋

![visitors](https://visitor-badge.laobi.icu/badge?page_id=ferch5003.ferch5003)

### 💬About me

- 💻 Backend developer (Go, Gin Gonic, Ruby, Ruby on Rails)
- 🎨Frontend developer (JavaScript, TypeScript, Vue)
- 🌱 I’m currently learning Design Patterns and gRPC
- ⚡ Love video games, anime, music and movies!

### 👾Frameworks and languages commonly used:

<a target="_blank" href="https://icons8.com/icon/20906/git">
    <img src="https://img.icons8.com/color/48/null/git.png" alt="git" width="32" height="32" />
</a>
<a target="_blank" href="https://icons8.com/icon/61466/intellij-idea">
    <img src="https://img.icons8.com/color/48/null/intellij-idea.png" alt="intellijieda" width="32" height="32" />
</a>
<a target="_blank" href="https://icons8.com/icon/9OGIyU8hrxW5/visual-studio-code-2019">
    <img src="https://img.icons8.com/color/48/null/visual-studio-code-2019.png" alt="visualstudiocode" width="32" height="32" />
</a>
<a target="_blank" href="https://icons8.com/icon/44442/golang">
    <img src="https://img.icons8.com/color/48/null/golang.png" alt="go" width="32" height="32" />
</a>
<a target="_blank" href="https://icons8.com/icon/108784/javascript">
    <img src="https://img.icons8.com/color/48/null/javascript--v1.png" alt="javascript" width="32" height="32" />
</a>
<a target="_blank" href="https://icons8.com/icon/uJM6fQYqDaZK/typescript">
    <img src="https://img.icons8.com/color/48/null/typescript.png" alt="typescript" width="32" height="32" />
</a>
<a target="_blank" href="https://icons8.com/icon/rY6agKizO9eb/vue-js">
    <img src="https://img.icons8.com/color/48/null/vue-js.png" alt="vue" width="32" height="32" />
</a>
<a target="_blank" href="https://icons8.com/icon/e2hIFBAN6UIe/ruby-programming-language">
    <img src="https://img.icons8.com/fluency/48/null/ruby-programming-language.png" alt="ruby" width="32" height="32" />
</a>
<a target="_blank" href="https://icons8.com/icon/ZMFmFsekpKfY/ruby-on-rails">
    <img src="https://img.icons8.com/windows/32/null/ruby-on-rails.png" alt="ruby on rails" width="32" height="32" />
</a>

### 🏆 My GitHub Stats:

<!--
![GitHub stats](https://github-readme-stats.vercel.app/api?username=ferch5003&show_icons=true&theme=tokyonight)
![Top Langs](https://github-readme-stats.vercel.app/api/top-langs/?username=ferch5003&theme=tokyonight)
-->
<div style="display: block;">
    <a href="https://github-readme-stats.vercel.app/api?username=ferch5003&show_icons=true&theme=tokyonight">
      <img style="padding: 10px;" src="https://github-readme-stats.vercel.app/api?username=ferch5003&show_icons=true&theme=tokyonight" />
    </a>
    <a href="https://github-readme-stats.vercel.app/api/top-langs/?username=ferch5003&theme=tokyonight">
      <img style="padding: 10px;" src="https://github-readme-stats.vercel.app/api/top-langs/?username=ferch5003&theme=tokyonight" />
    </a>
</div>

### Music Status🎵
![Music Staus](https://mut.mfd-web.online/spotify/track-to-show)

### Do you know...

<div>
{{- if .IsVideoFormat}}
    <video src="{{.Nasa.APOD.Url}}" width="400" height="500"></video>
{{- else}}
    <a href="{{.Nasa.APOD.Url}}">
        <img align="left" src="{{.Nasa.APOD.Url}}" width="400" height="500" alt="{{.Nasa.APOD.Copyright}}">
    </a>
{{- end}}
    <div>
        <h4>{{.Nasa.APOD.Title}}</h4>
        <time>{{.Nasa.APOD.Date}}</time>
        <p>{{.Nasa.APOD.Explanation}}</p>
        <strong><em>{{.Nasa.APOD.Copyright}}</em></strong>
    </div>
</div>
