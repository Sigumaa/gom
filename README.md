A tool to make `go mod init` a little easier  

Please use `go install` to install  

## Usage
```bash
gom init
```

This tool serves to simplify the process of 'go mod init'.

Have you ever found yourself thinking, "It's such a hassle to type 'go mod init github.com...'!"?

Well, this tool is here to address that concern. If your repository has already been initialized using 'git init', it will automatically retrieve the URL using 'git remote -v' and perform the 'go mod init' process for you!
