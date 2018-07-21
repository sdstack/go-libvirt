// +build ignore

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"text/template"

	introspect "github.com/godbus/dbus/introspect"
)

func normalizeMethodName(name string) string {
	words := strings.Split(name, "_")
	normal := ""
	for _, w := range words {
		normal += strings.ToTitle(w)
	}
	return normal
}

func ifc2obj(ifc string) string {
	return "/" + strings.Replace(ifc, ".", "/", -1)
}

func GuessType(val string, arg string, obj string) (string, string) {
	var rtype string
	var robj string
	switch arg[0] {
	case 'a':
		if arg[1] == '{' {
			rtype += "map["
			dtype, _ := GuessType(val, arg[2:], obj)
			rtype += dtype
			rtype += "]"
			dtype, dobj := GuessType(val, arg[3:], obj)
			rtype += dtype
			robj = dobj
		} else if arg[1] == 'a' {
			rtype += "[]"
			dtype, dobj := GuessType(val, arg[1:], obj)
			rtype += dtype
			robj = dobj
		} else {
			rtype += "[]"
			dtype, dobj := GuessType(val, arg[1:], obj)
			//if arg[1] == 'o' {
			//				rtype += dtype[:len(dtype)-1]
			//		} else {
			rtype += dtype
			//	}
			robj = dobj
		}
	case 'y':
		rtype = "byte"
		robj = obj
	case 'h':
		rtype = "uint32"
		robj = obj
	case '(':
		rtype = "interface{}"
		robj = obj
	case 'n':
		rtype = "int16"
		robj = obj
	case 'q':
		rtype = "uint16"
		robj = obj
	case 'i':
		rtype = "int32"
		robj = obj
	case 'x':
		rtype = "int64"
		robj = obj
	case 't':
		rtype = "uint64"
		robj = obj
	case 'u':
		rtype = "uint32"
		robj = obj
	case 's':
		rtype = "string"
		robj = obj
	case 'v':
		rtype = "interface{}"
		robj = obj
	case 'b':
		rtype = "bool"
		robj = obj
	case 'o':
		rtype = "dbus.ObjectPath"
		/*
			switch val {
			case "dev":
				rtype = "*NodeDevice"
			case "devs":
				rtype = "*NodeDevices"
			case "nwfilter":
				rtype = "*NWFilter"
			case "nwfilters":
				rtype = "*NWFilters"
			default:
				rtype = "*" + strings.Title(val)
			}
		*/
		robj = obj
	}
	return rtype, robj
}
func main() {
	var err error

	tplbuf, err := ioutil.ReadFile("template.tpl")
	if err != nil {
		panic(err)
	}
	httpPref := "https://raw.githubusercontent.com/libvirt/libvirt-dbus/master/data/"
	ifaces := []string{"Connect", "Domain", "NWFilter", "Network", "NodeDevice", "Secret", "StoragePool", "StorageVol"}
	for _, iface := range ifaces {
		res, err := http.Get(httpPref + "org.libvirt." + iface + ".xml")
		if err != nil {
			panic(err)
		}
		dec := xml.NewDecoder(res.Body)
		var node introspect.Node
		err = dec.Decode(&node)
		res.Body.Close()
		if err != nil {
			panic(err)
		}

		funcs := template.FuncMap{
			"Lower": strings.ToLower,
			"Upper": strings.ToUpper,
			"PkgName": func() string {
				return "libvirt"
				//				parts := strings.Split(node.Name, "/")
				//			return parts[len(parts)-1]
			},
			"OBJ_NAME":       func() string { return "obj" },
			"ExportName":     func() string { return strings.Split(node.Interfaces[0].Name, ".")[2] },
			"DbusDest":       func() string { return "org.libvirt" },
			"DbusObjectPath": func() string { return "/org/libvirt/QEMU" },
			"DbusInterface":  func() string { return node.Interfaces[0].Name },
			"Normalize":      normalizeMethodName,
			"Ifc2Obj":        ifc2obj,
			"AnnotationComment": func(s string) string {
				var ret []string
				parts := strings.Split(s, "\n")
				for _, part := range parts {
					ret = append(ret, strings.TrimSpace(part))
				}
				return strings.Join(ret, " ")
			},
			"PropWritable": func(prop introspect.Property) bool { return prop.Access == "readwrite" },
			"GuessType": func(val string, arg string, obj string) string {
				dtype, _ := GuessType(val, arg, obj)
				return dtype
			},
			"GetIns": func(args []introspect.Arg) []introspect.Arg {
				ret := make([]introspect.Arg, 0)
				for _, a := range args {
					if a.Direction != "out" {
						ret = append(ret, a)
					}
				}
				return ret
			},
			"GetOuts": func(args []introspect.Arg) []introspect.Arg {
				ret := make([]introspect.Arg, 0)
				for _, a := range args {
					if a.Direction == "out" {
						ret = append(ret, a)
					}
				}
				return ret
			},
			"CalcArgNum": func(args []introspect.Arg, direction string) (r int) {
				for _, arg := range args {
					if arg.Direction == direction {
						r++
					}
				}
				return
			},
			"Repeat": func(str string, sep string, times int) (r string) {
				for i := 0; i < times; i++ {
					if i != 0 {
						r += sep
					}
					r += str
				}
				return
			},
			"GetParamterNames": func(args []introspect.Arg) (ret string) {
				for _, arg := range args {
					if arg.Direction == "in" {
						ret += ", "
						if getKeyword(arg.Name) {
							ret += "i" + arg.Name
						} else {
							ret += arg.Name
						}
					}
				}
				return
			},
			"GetParamterOuts": func(args []introspect.Arg) (ret string) {
				var notFirst = false
				for _, arg := range args {
					if arg.Direction == "out" {
						if notFirst {
							ret += " ,"
						}
						notFirst = true
						if getKeyword(arg.Name) {
							ret += "&o" + arg.Name
						} else {
							ret += "&" + arg.Name
						}
					}
				}
				return
			},
			"GetParamterOutsProto": func(args []introspect.Arg) (ret string) {
				var notFirst = false
				for _, arg := range args {
					if arg.Direction == "out" || arg.Direction == "" {
						if notFirst {
							ret += ", "
						}
						notFirst = true
						dtype, _ := GuessType(arg.Name, arg.Type, "")
						if getKeyword(arg.Name) {
							ret += "o" + arg.Name + " " + dtype
						} else {
							ret += arg.Name + " " + dtype
						}
					}
				}
				return
			},
			"GetParamterInsProto": func(args []introspect.Arg) (ret string) {
				var notFirst = false
				for _, arg := range args {
					if arg.Direction == "in" || arg.Direction == "" {
						if notFirst {
							ret += ", "
						}
						notFirst = true
						if strings.Contains(arg.Type, "(") {
							if getKeyword(arg.Name) {
								ret += "i" + arg.Name + " interface{}"
							} else {
								ret += arg.Name + " interface{}"
							}
						} else {
							dtype, _ := GuessType(arg.Name, arg.Type, "")
							if getKeyword(arg.Name) {
								ret += "i" + arg.Name + " " + dtype
							} else {
								ret += arg.Name + " " + dtype
							}
						}
					}
				}
				return
			},
			/*
				"TryConvertObjectPath": func(prop introspect.Property) string {
					if v := getObjectPathConvert("Property", prop.Annotations); v != "" {
						return tryConvertObjectPathGo(node, prop.Type, v)
					}
					return ""
				},
				"GetObjectPathType": func(prop introspect.Property) (ret string) {
					if v := getObjectPathConvert("Property", prop.Annotations); v != "" {
						ret, _ = guessTypeGo(node, prop.Type, v)
						return
					}
					return dbus.TypeFor(prop.Type)
				},
			*/
		}

		parts := strings.Split(node.Name, "/")
		fname := parts[len(parts)-1]
		tpl, err := template.New(fname).Funcs(funcs).Parse(string(tplbuf))
		if err != nil {
			panic(err)
		}
		fmt.Printf("writing %s\n", fname)
		fp, err := os.OpenFile(fname+".go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.FileMode(0600))
		if err != nil {
			panic(err)
		}
		for _, ifc := range node.Interfaces {
			if err = tpl.Execute(fp, ifc); err != nil {
				log.Println("executing template:", err)
			}
		}
		fp.Close()
		cmd := exec.Command("goimports", "-w", fname+".go")
		if err = cmd.Run(); err != nil {
			panic(err)
		}
	}
}

func getKeyword(arg string) bool {
	var r bool
	switch arg {
	case "break":
		r = true
	case "default":
		r = true
	case "func":
		r = true
	case "interface":
		r = true
	case "select":
		r = true
	case "case":
		r = true
	case "defer":
		r = true
	case "go":
		r = true
	case "map":
		r = true
	case "struct":
		r = true
	case "chan":
		r = true
	case "else":
		r = true
	case "goto":
		r = true
	case "package":
		r = true
	case "switch":
		r = true
	case "const":
		r = true
	case "fallthrough":
		r = true
	case "if":
		r = true
	case "range":
		r = true
	case "type":
		r = true
	case "continue":
		r = true
	case "for":
		r = true
	case "import":
		r = true
	case "return":
		r = true
	case "var":
		r = true
	case "uint8":
		r = true
	case "uint16":
		r = true
	case "uint32":
		r = true
	case "uint64":
		r = true
	case "int8":
		r = true
	case "int16":
		r = true
	case "int32":
		r = true
	case "int64":
		r = true
	case "float32":
		r = true
	case "float64":
		r = true
	case "complex64":
		r = true
	case "complex128":
		r = true
	case "byte":
		r = true
	case "rune":
		r = true
	case "uint":
		r = true
	case "int":
		r = true
	case "uintptr":
		r = true
	case "string":
		r = true
	case "class":
		r = true
	case "bool":
		r = true
	}
	return r
}
