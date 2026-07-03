package middleware

import (
	"ToDoList/internal/service/auth"
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		uid, err := auth.ParseToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("uid", uid)
		ctx.Next()
	}
}

type gzipWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

func (gz *gzipWriter) Write(data []byte) (int, error) {
	contentType := gz.Header().Get("Content-Type")
	// Сжимаем только если это JSON или HTML
	if strings.Contains(contentType, "application/json") || strings.Contains(contentType, "text/html") {
		return gz.Writer.Write(data)
	}
	return gz.ResponseWriter.Write(data)
}

func GzipCompressMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Чтение сжатых запросов
		if ctx.GetHeader("Content-Encoding") == "gzip" {
			gzReader, err := gzip.NewReader(ctx.Request.Body)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid gzip body"})
				return
			}

			defer gzReader.Close()

			ctx.Request.Body = http.MaxBytesReader(ctx.Writer, gzReader, ctx.Request.ContentLength)
		}

		//Отправка сжатых запросов
		str := ctx.GetHeader("Accept-Encoding")
		if !strings.Contains(str, "gzip") {
			ctx.Next()
			return
		}
		gz, err := gzip.NewWriterLevel(ctx.Writer, gzip.DefaultCompression)
		if err != nil {
			ctx.Next()
			return
		}
		defer gz.Close()

		ctx.Header("Content-Encoding", "gzip")
		ctx.Writer = &gzipWriter{ResponseWriter: ctx.Writer, Writer: gz}

		ctx.Next()
	}
}
