package ts

import (
	_ "embed"
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	apiutil "github.com/zeromicro/go-zero/tools/goctl/api/util"
	"github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

const (
	pathTagKey = "path"
	pathPrefix = "prefix"
)

//go:embed tpl/api.tpl
var apiTemplate string

func genAPI(dir string, api *spec.ApiSpec) error {
	filename := strings.Replace(api.Service.Name, "-api", "", 1) + ".ts"
	if err := pathx.RemoveIfExist(path.Join(dir, filename)); err != nil {
		return err
	}
	fp, created, err := apiutil.MaybeCreateFile(dir, "", filename)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()

	imports := `import axios, { AxiosRequestConfig } from "axios"` + pathx.NL +
		`import qs from "qs"`

	if len(api.Types) != 0 {
		if len(imports) > 0 {
			imports += pathx.NL
		}
		outputFile := apiutil.ComponentName(api)
		imports += fmt.Sprintf(`import * as components from "%s"`, "./"+outputFile)
		imports += fmt.Sprintf(`%sexport * from "%s"`, pathx.NL, "./"+outputFile)
	}

	apis, err := genAPIs(api, "http")
	if err != nil {
		return err
	}

	t := template.Must(template.New("apiTemplate").Parse(apiTemplate))
	return t.Execute(fp, map[string]string{
		"imports": imports,
		"apis":    strings.TrimSpace(apis),
	})
}

func genAPIs(api *spec.ApiSpec, caller string) (string, error) {
	var builder strings.Builder
	for _, group := range api.Service.Groups {
		prefix := strings.Replace(group.Annotation.Properties["prefix"], "/", "_", -1)
		for _, route := range group.Routes {
			handler := prefix + "_" + route.Handler
			handler = strings.TrimPrefix(handler, "_")
			if len(handler) == 0 {
				return "", fmt.Errorf("missing handler annotation for route %q", route.Path)
			}

			handler = util.Untitle(handler)
			handler = strings.Replace(handler, "Handler", "", 1)
			comment := commentForRoute(route)
			if len(comment) > 0 {
				fmt.Fprintf(&builder, "%s\n", comment)
			}
			fmt.Fprintf(&builder, "export function %s(%s) {\n", handler, paramsForRoute(route))
			writeIndent(&builder, 1)
			responseGeneric := "<null>"
			if len(route.ResponseTypeName()) > 0 {
				val, err := goTypeToTs(route.ResponseType, true)
				if err != nil {
					return "", err
				}

				responseGeneric = fmt.Sprintf("<%s>", val)
			}
			fmt.Fprintf(&builder, `return %s.%s%s(%s)`, caller, strings.ToLower(route.Method),
				util.Title(responseGeneric), callParamsForRoute(route, group))
			builder.WriteString("\n}\n\n")
		}
	}

	apis := builder.String()
	return apis, nil
}

func paramsForRoute(route spec.Route) string {
	if route.RequestType == nil {
		return "config?: AxiosRequestConfig"
	}
	rt, err := goTypeToTs(route.RequestType, true)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	var params []string
	if hasRequestForm(route) {
		params = append(params, fmt.Sprintf("query: %sParams", rt))
	}
	if hasRequestBody(route) {
		params = append(params, fmt.Sprintf("req: %s", rt))
	}
	if hasRequestPath(route) {
		ds, ok := route.RequestType.(spec.DefineStruct)
		if !ok {
			fmt.Printf("invalid route.RequestType: {%v}\n", route.RequestType)
		}
		members := ds.GetTagMembers(pathTagKey)
		for _, member := range members {
			tags := member.Tags()

			if len(tags) > 0 && tags[0].Key == pathTagKey {
				valueType, err := goTypeToTs(member.Type, false)
				if err != nil {
					fmt.Println(err.Error())
					return ""
				}
				params = append(params, fmt.Sprintf("%s: %s", tags[0].Name, valueType))
			}
		}
	}

	params = append(params, "config?: AxiosRequestConfig")

	return strings.Join(params, ", ")
}

func commentForRoute(route spec.Route) string {
	var builder strings.Builder
	comment := route.JoinedDoc()
	builder.WriteString("/**")
	builder.WriteString("\n * @description " + comment)

	if hasRequestForm(route) {
		builder.WriteString("\n * @param query")
	}
	if hasRequestBody(route) {
		builder.WriteString("\n * @param req")
	}

	builder.WriteString("\n */")
	return builder.String()
}

func callParamsForRoute(route spec.Route, group spec.Group) string {
	var params = []string{pathForRoute(route, group)}

	if hasRequestForm(route) {
		params[0] += " + `?` + qs.stringify(query)"
	}
	if hasRequestBody(route) {
		params = append(params, "req")
	}

	params = append(params, "config")

	return strings.Join(params, ", ")
}

func pathForRoute(route spec.Route, group spec.Group) string {
	prefix := group.GetAnnotation(pathPrefix)

	routePath := route.Path
	if strings.Contains(routePath, ":") {
		pathSlice := strings.Split(routePath, "/")
		for i, part := range pathSlice {
			if strings.Contains(part, ":") {
				pathSlice[i] = fmt.Sprintf("${%s}", part[1:])
			}
		}
		routePath = strings.Join(pathSlice, "/")
	}
	if len(prefix) == 0 {
		return "`" + routePath + "`"
	}

	prefix = strings.TrimPrefix(prefix, `"`)
	prefix = strings.TrimSuffix(prefix, `"`)
	return fmt.Sprintf("`%s/%s`", prefix, strings.TrimPrefix(routePath, "/"))
}

func hasRequestBody(route spec.Route) bool {
	if len(strings.Split(",post,put,patch,", route.Method)) == 1 {
		return false
	}

	ds, ok := route.RequestType.(spec.DefineStruct)
	if !ok {
		return false
	}

	return len(route.RequestTypeName()) > 0 && len(ds.GetBodyMembers()) > 0
}

func hasRequestPath(route spec.Route) bool {
	ds, ok := route.RequestType.(spec.DefineStruct)
	if !ok {
		return false
	}

	return len(route.RequestTypeName()) > 0 && len(ds.GetTagMembers(pathTagKey)) > 0
}

func hasRequestForm(route spec.Route) bool {
	ds, ok := route.RequestType.(spec.DefineStruct)
	if !ok {
		return false
	}

	return len(route.RequestTypeName()) > 0 && len(ds.GetFormMembers()) > 0
}
