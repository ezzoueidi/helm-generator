// GENERATED CODE - DO NOT EDIT
// This file is the run file for Revel.
// It registers all the controllers and provides details for the Revel server engine to
// properly inject parameters directly into the action endpoints.
package run

import (
	"reflect"
	"github.com/revel/revel"
	controllers "github.com/revel/modules/static/app/controllers"
	_ "github.com/revel/modules/testrunner/app"
	controllers0 "github.com/revel/modules/testrunner/app/controllers"
	_ "helm-generator-ui/app"
	controllers1 "helm-generator-ui/app/controllers"
	tests "helm-generator-ui/tests"
	"github.com/revel/revel/testing"
)

var (
	// So compiler won't complain if the generated code doesn't reference reflect package...
	_ = reflect.Invalid
)

// Register and run the application
func Run(port int) {
	Register()
	revel.Run(port)
}

// Register all the controllers
func Register() {
	revel.AppLog.Info("Running revel server")
	
	revel.RegisterController((*controllers.Static)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Serve",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModule",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "ServeModuleDir",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers0.TestRunner)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					76: []string{ 
						"testSuites",
					},
				},
			},
			&revel.MethodType{
				Name: "Suite",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			&revel.MethodType{
				Name: "Run",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					125: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "List",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
				},
			},
			
		})
	
	revel.RegisterController((*controllers1.App)(nil),
		[]*revel.MethodType{
			&revel.MethodType{
				Name: "Index",
				Args: []*revel.MethodArg{ 
					&revel.MethodArg{Name: "chart_name", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "chartt0", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "chartt1", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "chartt2", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "chartt3", Type: reflect.TypeOf((*string)(nil)) },
					&revel.MethodArg{Name: "chartt4", Type: reflect.TypeOf((*string)(nil)) },
				},
				RenderArgNames: map[int][]string{ 
					38: []string{ 
						"chart_name",
						"chartt0",
						"chartt1",
						"chartt2",
						"chartt3",
						"chartt4",
					},
				},
			},
			&revel.MethodType{
				Name: "Helm",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					49: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "Contact",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					53: []string{ 
					},
				},
			},
			&revel.MethodType{
				Name: "About",
				Args: []*revel.MethodArg{ 
				},
				RenderArgNames: map[int][]string{ 
					56: []string{ 
					},
				},
			},
			
		})
	
	revel.DefaultValidationKeys = map[string]map[int]string{ 
		"helm-generator-ui/app/controllers.App.Index": { 
			13: "chart_name",
			14: "chartt0",
			15: "chartt1",
			18: "chart_name",
			19: "chart_name",
			21: "chartt0",
			22: "chartt0",
			24: "chartt1",
			25: "chartt1",
			30: "chart_name",
			31: "chart_name",
		},
	}
	testing.TestSuites = []interface{}{ 
		(*tests.AppTest)(nil),
	}
}
