package info

// Version 软件版本
const Version = "0.0.1"

// Author 作者
const Author = "Arror"

// Email 作者电子邮件
const Email = "763911422@qq.com"

// AppHelpTemplate 帮助模板
const AppHelpTemplate = `
名称:
    {{.Name}} - {{.Usage}}
作者:
    {{range .Authors}}{{ . }}{{end}}{{if .Commands}}
命令:
{{range .Commands}}{{if not .HideHelp}}    {{join .Names ", "}}{{ "\t" }}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
全局选项:
    {{range .VisibleFlags}}{{.}}{{end}}{{end}}{{if .Copyright }}
版权:
    {{.Copyright}}{{end}}{{if .Version}}
版本:
    {{.Version}}{{end}}
`
