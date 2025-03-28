<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# api

```go
import "github.com/nlatham1999/go-agent/pkg/api"
```

## Index

- [type Api](<#Api>)
  - [func NewApi\(models map\[string\]ModelInterface, settings ApiSettings\) \(\*Api, error\)](<#NewApi>)
  - [func \(a \*Api\) HomeHandler\(w http.ResponseWriter, r \*http.Request\)](<#Api.HomeHandler>)
  - [func \(a \*Api\) ModelPageHandler\(w http.ResponseWriter, r \*http.Request\)](<#Api.ModelPageHandler>)
  - [func \(a \*Api\) Serve\(\)](<#Api.Serve>)
- [type ApiSettings](<#ApiSettings>)
- [type Color](<#Color>)
- [type Link](<#Link>)
- [type Model](<#Model>)
- [type ModelInterface](<#ModelInterface>)
- [type Patch](<#Patch>)
- [type Turtle](<#Turtle>)
- [type Widget](<#Widget>)
  - [func NewFloatSliderWidget\(prettyName, targetVariable, minValue, maxValue, defaultValue, stepAmount string, valuePointer \*float64\) Widget](<#NewFloatSliderWidget>)


<a name="Api"></a>
## type [Api](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/api.go#L13-L31>)



```go
type Api struct {
    // contains filtered or unexported fields
}
```

<a name="NewApi"></a>
### func [NewApi](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/api.go#L40>)

```go
func NewApi(models map[string]ModelInterface, settings ApiSettings) (*Api, error)
```



<a name="Api.HomeHandler"></a>
### func \(\*Api\) [HomeHandler](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/handlers.go#L141>)

```go
func (a *Api) HomeHandler(w http.ResponseWriter, r *http.Request)
```



<a name="Api.ModelPageHandler"></a>
### func \(\*Api\) [ModelPageHandler](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/handlers.go#L162>)

```go
func (a *Api) ModelPageHandler(w http.ResponseWriter, r *http.Request)
```



<a name="Api.Serve"></a>
### func \(\*Api\) [Serve](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/api.go#L59>)

```go
func (a *Api) Serve()
```



<a name="ApiSettings"></a>
## type [ApiSettings](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/api.go#L33-L38>)



```go
type ApiSettings struct {
    ButtonTitles       map[string]string
    ButtonDescriptions map[string]string
    StoreSteps         bool // Whether to store steps
    MaxSteps           int  // Maximum number of steps to store. Default is 1000
}
```

<a name="Color"></a>
## type [Color](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/structs.go#L34-L39>)



```go
type Color struct {
    Red   int `json:"r"`
    Green int `json:"g"`
    Blue  int `json:"b"`
    Alpha int `json:"a"`
}
```

<a name="Link"></a>
## type [Link](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/structs.go#L41-L56>)



```go
type Link struct {
    End1       int         `json:"end1"`
    End2       int         `json:"end2"`
    End1X      float64     `json:"end1X"`
    End1Y      float64     `json:"end1Y"`
    End2X      float64     `json:"end2X"`
    End2Y      float64     `json:"end2Y"`
    End1Size   float64     `json:"end1Size"`
    End2Size   float64     `json:"end2Size"`
    Directed   bool        `json:"directed"`
    Color      Color       `json:"color"`
    Label      interface{} `json:"label"`
    LabelColor Color       `json:"labelColor"`
    Size       int         `json:"size"`
    Hidden     bool        `json:"hidden"`
}
```

<a name="Model"></a>
## type [Model](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/structs.go#L3-L14>)



```go
type Model struct {
    Patches     []Patch  `json:"patches"`
    Turtles     []Turtle `json:"turtles"`
    Links       []Link   `json:"links"`
    Ticks       int      `json:"ticks"`
    WorldWidth  int      `json:"width"`
    WorldHeight int      `json:"height"`
    MinPxCor    int      `json:"minPxCor"`
    MaxPxCor    int      `json:"maxPxCor"`
    MinPyCor    int      `json:"minPyCor"`
    MaxPyCor    int      `json:"maxPyCor"`
}
```

<a name="ModelInterface"></a>
## type [ModelInterface](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/modelInterface.go#L5-L14>)



```go
type ModelInterface interface {
    Init()        // runs at the very beginning
    SetUp() error // sets up the model
    Go()          // runs the model

    Model() *model.Model           // returns the model
    Stats() map[string]interface{} //returns the stats of the model
    Stop() bool                    // on whether to stop the model
    Widgets() []Widget             // returns the widgets of the model
}
```

<a name="Patch"></a>
## type [Patch](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/structs.go#L16-L20>)



```go
type Patch struct {
    X     int   `json:"x"`
    Y     int   `json:"y"`
    Color Color `json:"color"`
}
```

<a name="Turtle"></a>
## type [Turtle](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/structs.go#L22-L32>)



```go
type Turtle struct {
    X          float64     `json:"x"`
    Y          float64     `json:"y"`
    Color      Color       `json:"color"`
    Size       float64     `json:"size"`
    Who        int         `json:"who"`
    Shape      string      `json:"shape"`
    Heading    float64     `json:"heading"`
    Label      interface{} `json:"label"`
    LabelColor Color       `json:"labelColor"`
}
```

<a name="Widget"></a>
## type [Widget](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/widgets.go#L5-L18>)



```go
type Widget struct {
    PrettyName         string   `json:"prettyName"`
    TargetVariable     string   `json:"targetVariable"`
    WidgetType         string   `json:"widgetType"`
    WidgetValueType    string   `json:"widgetValueType"`
    MinValue           string   `json:"minValue"`
    MaxValue           string   `json:"maxValue"`
    DefaultValue       string   `json:"defaultValue"`
    StepAmount         string   `json:"stepAmount"`
    Target             func()   `json:"target"` // this is a function that will be called when the widget is interacted with if the type is a button
    ValuePointerInt    *int     `json:"valuePointerInt"`
    ValuePointerFloat  *float64 `json:"valuePointerFloat"`
    ValuePointerString *string  `json:"valuePointerString"`
}
```

<a name="NewFloatSliderWidget"></a>
### func [NewFloatSliderWidget](<https://github.com/nlatham1999/go-agent/blob/main/pkg/api/widgets.go#L20>)

```go
func NewFloatSliderWidget(prettyName, targetVariable, minValue, maxValue, defaultValue, stepAmount string, valuePointer *float64) Widget
```



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
