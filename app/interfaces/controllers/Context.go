package controllers

// Context is a type
type Context interface {
	Param(key string) string
	JSON(code int, obj interface{})
	PostForm(key string) string
	PostFormArray(key string) []string
	PostFormMap(key string) map[string]string
	BindJSON(obj interface{}) error
	Bind(obj interface{}) error
}
