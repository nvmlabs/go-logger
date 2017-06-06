package logging

import (
	"net/http"

	"github.com/Sirupsen/logrus"
)

// LogHTTPRequest logs an info message with standard fields and extract HTTP specific data from the request
func LogHTTPRequest(req *http.Request, msg string) {
	fields := logrus.Fields{
		"request_id":     req.Header.Get("x-request-id"),
		"external_id":    req.Header.Get("x-external-id"),
		"forwarded_for":  req.Header.Get("x-forwarded-for"),
		"protocol":       req.Proto,
		"remote_address": req.RemoteAddr,
		"url":            req.RequestURI,
		"method":         req.Method,
		"content_length": req.ContentLength,
	}

	logger.WithFields(fields).Info(msg)
}

// LogHTTPResponse logs an info message with standard fields and extract HTTP specific data from the response
func LogHTTPResponse(res *http.Response, msg string) {
	fields := logrus.Fields{
		"request_id":     res.Request.Header.Get("x-request-id"),
		"external_id":    res.Request.Header.Get("x-external-id"),
		"forwarded_for":  res.Request.Header.Get("x-forwarded-for"),
		"protocol":       res.Proto,
		"remote_address": res.Request.RemoteAddr,
		"url":            res.Request.RequestURI,
		"method":         res.Request.Method,
		"content_length": res.ContentLength,
		"status":         res.StatusCode,
	}

	logger.WithFields(fields).Info(msg)
}
