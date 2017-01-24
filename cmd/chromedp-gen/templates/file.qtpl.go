// This file is automatically generated by qtc from "file.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line templates/file.qtpl:1
package templates

//line templates/file.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line templates/file.qtpl:1
import (
	"sort"

	. "github.com/knq/chromedp/cmd/chromedp-gen/internal"
)

// FileHeader is the file header template.

//line templates/file.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line templates/file.qtpl:8
func StreamFileHeader(qw422016 *qt422016.Writer, pkgName string, d *Domain) {
	//line templates/file.qtpl:8
	qw422016.N().S(`
`)
	//line templates/file.qtpl:9
	if d != nil {
		//line templates/file.qtpl:9
		qw422016.N().S(`// Package `)
		//line templates/file.qtpl:9
		qw422016.N().S(d.PackageName())
		//line templates/file.qtpl:9
		qw422016.N().S(` provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome `)
		//line templates/file.qtpl:10
		qw422016.N().S(d.String())
		//line templates/file.qtpl:10
		qw422016.N().S(` domain.
// `)
		//line templates/file.qtpl:11
		if desc := d.Description; desc != "" {
			//line templates/file.qtpl:11
			qw422016.N().S(`
`)
			//line templates/file.qtpl:12
			qw422016.N().S(formatComment(desc, "", ""))
			//line templates/file.qtpl:12
			qw422016.N().S(`
//`)
			//line templates/file.qtpl:13
		}
		//line templates/file.qtpl:13
		qw422016.N().S(`
// Generated by the chromedp-gen command.`)
		//line templates/file.qtpl:14
	}
	//line templates/file.qtpl:14
	qw422016.N().S(`
package `)
	//line templates/file.qtpl:15
	qw422016.N().S(pkgName)
	//line templates/file.qtpl:15
	qw422016.N().S(`

// AUTOGENERATED. DO NOT EDIT.
`)
//line templates/file.qtpl:18
}

//line templates/file.qtpl:18
func WriteFileHeader(qq422016 qtio422016.Writer, pkgName string, d *Domain) {
	//line templates/file.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/file.qtpl:18
	StreamFileHeader(qw422016, pkgName, d)
	//line templates/file.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line templates/file.qtpl:18
}

//line templates/file.qtpl:18
func FileHeader(pkgName string, d *Domain) string {
	//line templates/file.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/file.qtpl:18
	WriteFileHeader(qb422016, pkgName, d)
	//line templates/file.qtpl:18
	qs422016 := string(qb422016.B)
	//line templates/file.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/file.qtpl:18
	return qs422016
//line templates/file.qtpl:18
}

// FileLocalImportTemplate is a template for local imports.

//line templates/file.qtpl:21
func StreamFileLocalImportTemplate(qw422016 *qt422016.Writer, out string) {
	//line templates/file.qtpl:21
	qw422016.N().S(`
import (
	. "`)
	//line templates/file.qtpl:23
	qw422016.N().S(out)
	//line templates/file.qtpl:23
	qw422016.N().S(`"
)
`)
//line templates/file.qtpl:25
}

//line templates/file.qtpl:25
func WriteFileLocalImportTemplate(qq422016 qtio422016.Writer, out string) {
	//line templates/file.qtpl:25
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/file.qtpl:25
	StreamFileLocalImportTemplate(qw422016, out)
	//line templates/file.qtpl:25
	qt422016.ReleaseWriter(qw422016)
//line templates/file.qtpl:25
}

//line templates/file.qtpl:25
func FileLocalImportTemplate(out string) string {
	//line templates/file.qtpl:25
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/file.qtpl:25
	WriteFileLocalImportTemplate(qb422016, out)
	//line templates/file.qtpl:25
	qs422016 := string(qb422016.B)
	//line templates/file.qtpl:25
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/file.qtpl:25
	return qs422016
//line templates/file.qtpl:25
}

// FileImportTemplate is a general import template.

//line templates/file.qtpl:28
func StreamFileImportTemplate(qw422016 *qt422016.Writer, m map[string]string) {
	//line templates/file.qtpl:29
	var keys []string
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	//line templates/file.qtpl:34
	qw422016.N().S(`
import (`)
	//line templates/file.qtpl:35
	for _, k := range keys {
		//line templates/file.qtpl:36
		v := m[k]

		//line templates/file.qtpl:37
		qw422016.N().S(`
	`)
		//line templates/file.qtpl:38
		if k != v {
			//line templates/file.qtpl:38
			qw422016.N().S(v)
			//line templates/file.qtpl:38
			qw422016.N().S(` `)
			//line templates/file.qtpl:38
		}
		//line templates/file.qtpl:38
		qw422016.N().S(`"`)
		//line templates/file.qtpl:38
		qw422016.N().S(k)
		//line templates/file.qtpl:38
		qw422016.N().S(`"`)
		//line templates/file.qtpl:38
	}
	//line templates/file.qtpl:38
	qw422016.N().S(`
)
`)
//line templates/file.qtpl:40
}

//line templates/file.qtpl:40
func WriteFileImportTemplate(qq422016 qtio422016.Writer, m map[string]string) {
	//line templates/file.qtpl:40
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/file.qtpl:40
	StreamFileImportTemplate(qw422016, m)
	//line templates/file.qtpl:40
	qt422016.ReleaseWriter(qw422016)
//line templates/file.qtpl:40
}

//line templates/file.qtpl:40
func FileImportTemplate(m map[string]string) string {
	//line templates/file.qtpl:40
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/file.qtpl:40
	WriteFileImportTemplate(qb422016, m)
	//line templates/file.qtpl:40
	qs422016 := string(qb422016.B)
	//line templates/file.qtpl:40
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/file.qtpl:40
	return qs422016
//line templates/file.qtpl:40
}

// FileEmptyVarTemplate is a template that imports empty variable names.

//line templates/file.qtpl:43
func StreamFileEmptyVarTemplate(qw422016 *qt422016.Writer, m ...string) {
	//line templates/file.qtpl:44
	sort.Strings(m)

	//line templates/file.qtpl:45
	qw422016.N().S(`
var (`)
	//line templates/file.qtpl:46
	for _, typ := range m {
		//line templates/file.qtpl:46
		qw422016.N().S(`
	_  `)
		//line templates/file.qtpl:47
		qw422016.N().S(typ)
		//line templates/file.qtpl:47
	}
	//line templates/file.qtpl:47
	qw422016.N().S(`)
`)
//line templates/file.qtpl:48
}

//line templates/file.qtpl:48
func WriteFileEmptyVarTemplate(qq422016 qtio422016.Writer, m ...string) {
	//line templates/file.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/file.qtpl:48
	StreamFileEmptyVarTemplate(qw422016, m...)
	//line templates/file.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line templates/file.qtpl:48
}

//line templates/file.qtpl:48
func FileEmptyVarTemplate(m ...string) string {
	//line templates/file.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/file.qtpl:48
	WriteFileEmptyVarTemplate(qb422016, m...)
	//line templates/file.qtpl:48
	qs422016 := string(qb422016.B)
	//line templates/file.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/file.qtpl:48
	return qs422016
//line templates/file.qtpl:48
}

// FileVarTemplate is a template for a single variable declaration.

//line templates/file.qtpl:51
func StreamFileVarTemplate(qw422016 *qt422016.Writer, name, value, desc string) {
	//line templates/file.qtpl:51
	qw422016.N().S(`
`)
	//line templates/file.qtpl:52
	if desc != "" {
		//line templates/file.qtpl:52
		qw422016.N().S(formatComment(desc, "", ""))
		//line templates/file.qtpl:52
	}
	//line templates/file.qtpl:52
	qw422016.N().S(`
var `)
	//line templates/file.qtpl:53
	qw422016.N().S(name)
	//line templates/file.qtpl:53
	qw422016.N().S(` = `)
	//line templates/file.qtpl:53
	qw422016.N().S(value)
	//line templates/file.qtpl:53
	qw422016.N().S(`
`)
//line templates/file.qtpl:54
}

//line templates/file.qtpl:54
func WriteFileVarTemplate(qq422016 qtio422016.Writer, name, value, desc string) {
	//line templates/file.qtpl:54
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line templates/file.qtpl:54
	StreamFileVarTemplate(qw422016, name, value, desc)
	//line templates/file.qtpl:54
	qt422016.ReleaseWriter(qw422016)
//line templates/file.qtpl:54
}

//line templates/file.qtpl:54
func FileVarTemplate(name, value, desc string) string {
	//line templates/file.qtpl:54
	qb422016 := qt422016.AcquireByteBuffer()
	//line templates/file.qtpl:54
	WriteFileVarTemplate(qb422016, name, value, desc)
	//line templates/file.qtpl:54
	qs422016 := string(qb422016.B)
	//line templates/file.qtpl:54
	qt422016.ReleaseByteBuffer(qb422016)
	//line templates/file.qtpl:54
	return qs422016
//line templates/file.qtpl:54
}