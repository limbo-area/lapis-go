package router

import (
	"lapis-go/constants"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

//go:linkname decoderBuilder github.com/gofiber/fiber/v2.decoderBuilder
func decoderBuilder(parserConfig fiber.ParserConfig) interface{}

//go:linkname equalFieldType github.com/gofiber/fiber/v2.equalFieldType
func equalFieldType(out interface{}, kind reflect.Kind, key string) bool

func ParseToStruct(aliasTag string, out interface{}, data map[string][]string) error {
	decoder := decoderBuilder(fiber.ParserConfig{
		SetAliasTag:       aliasTag,
		IgnoreUnknownKeys: true,
	})
	result := reflect.ValueOf(decoder).
		MethodByName("Decode").
		Call([]reflect.Value{reflect.ValueOf(out), reflect.ValueOf(data)})
	switch result[0].Interface().(type) {
	case error:
		return result[0].Interface().(error)
	case nil:
		return nil
	}
	return nil
}

func HeaderParser(c *fiber.Ctx, model interface{}) error {
	headerData := make(map[string][]string)
	c.Request().Header.VisitAll(func(key, val []byte) {
		k := utils.UnsafeString(key)
		v := utils.UnsafeString(val)
		if strings.Contains(v, ",") && equalFieldType(model, reflect.Slice, k) {
			values := strings.Split(v, ",")
			for i := 0; i < len(values); i++ {
				headerData[k] = append(headerData[k], values[i])
			}
		} else {
			headerData[k] = append(headerData[k], v)
		}
	})
	return ParseToStruct(constants.HEADER, model, headerData)
}

func ParamsParser(c *fiber.Ctx, model interface{}) error {
	params := make(map[string][]string)
	for _, param := range c.Route().Params {
		params[param] = append(params[param], c.Params(param))
	}
	return ParseToStruct(constants.URI, model, params)
}

func CookiesParser(c *fiber.Ctx, model interface{}) error {
	params := make(map[string][]string)
	c.Request().Header.VisitAllCookie(func(key, value []byte) {
		params[string(key)] = append(params[string(key)], string(value))
	})
	return ParseToStruct(constants.COOKIE, model, params)
}
