package transport

import (
	"github.com/gin-gonic/gin"
)

/* utiliza el framework de manejo de solicitudes web Gin para crear un servidor que maneje solicitudes HTTP
de una manera estructurada y reutilizable */

func GinServer(
	endpoint Endpoint,
	decode func(c *gin.Context) (interface{}, error),
	encode func(c *gin.Context, resp interface{}),
	encodeError func(c *gin.Context, err error)) func(c *gin.Context) {

	return func(c *gin.Context) {
		/* Llama a la función de decodificación para extraer los datos de la solicitud. Si hay un error,
		se maneja con encodeError y se detiene la ejecución.*/
		data, err := decode(c)
		if err != nil {
			encodeError(c, err)
			return
		}

		/* Si la decodificación fue exitosa, se llama al endpoint con los datos decodificados.
		Si el endpoint devuelve un error, se llama nuevamente a encodeError y se detiene la ejecución. */
		res, err := endpoint(c.Request.Context(), data)
		if err != nil {
			encodeError(c, err)
			return
		}

		/* Si todo va bien, se llama a la función encode para devolver la respuesta al cliente. */
		encode(c, res)

	}
}
