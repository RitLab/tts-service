// Package pdf provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package pdf

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
	Slug    string `json:"slug"`
}

// JoinPdfFiles defines model for JoinPdfFiles.
type JoinPdfFiles struct {
	Files []openapi_types.File `json:"files"`
}

// SignPdfFile defines model for SignPdfFile.
type SignPdfFile struct {
	File openapi_types.File `json:"file"`
	Key  string             `json:"key"`
}

// Success defines model for Success.
type Success struct {
	Message string `json:"message"`
}

// TtsResponse defines model for TtsResponse.
type TtsResponse struct {
	Data *struct {
		Url *string `json:"url,omitempty"`
	} `json:"data,omitempty"`
}

// VerifyPdfFile defines model for VerifyPdfFile.
type VerifyPdfFile struct {
	File openapi_types.File `json:"file"`
	Key  string             `json:"key"`
}

// JoinPdfFilesMultipartRequestBody defines body for JoinPdfFiles for multipart/form-data ContentType.
type JoinPdfFilesMultipartRequestBody = JoinPdfFiles

// SignPdfFileMultipartRequestBody defines body for SignPdfFile for multipart/form-data ContentType.
type SignPdfFileMultipartRequestBody = SignPdfFile

// VerifyPdfFileMultipartRequestBody defines body for VerifyPdfFile for multipart/form-data ContentType.
type VerifyPdfFileMultipartRequestBody = VerifyPdfFile
