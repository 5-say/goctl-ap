package doc

import (
	_ "embed"
)

var (
	//go:embed swagger.json
	SwaggerJson []byte
)

/**

internal/logic/api/swaggerDocLogic.go

func (l *SwaggerDocLogic) SwaggerDoc() error {
	// todo: add your logic here and delete this line
	l.w.Write(doc.SwaggerJson)

	return nil
}

*/
