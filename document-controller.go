package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func documentSave(writer http.ResponseWriter, req *http.Request) {
	type DocumentFile struct {
		Filename string
		Bytes    []byte
		Size     int64
	}

	req.ParseMultipartForm(32 << 20)
	file, handler, err := req.FormFile("file")
	if err != nil {
		log.Print(handler, err)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		panic(err)
	}
	//Encode JSON
	documentJSON, err := json.Marshal(DocumentFile{
		Filename: handler.Filename,
		Bytes:    buf.Bytes(),
		Size:     handler.Size,
	})
	if err != nil {
		panic(err)
	}
	log.Print(documentJSON)

	//Decode JSON
	var documentNormal DocumentFile
	errDecoding := json.Unmarshal(documentJSON, &documentNormal)
	if errDecoding != nil {
		panic(errDecoding)
	}
	log.Print(documentNormal.Bytes)
	log.Print(documentNormal.Filename)
	log.Print(documentNormal.Size)

	http.Redirect(writer, req, "/#documents", http.StatusMovedPermanently)
}

func documentDelete(writer http.ResponseWriter, req *http.Request) {
	//Encode JSON
	documentJSON, err := json.Marshal(Document{
		ID: req.FormValue("id"),
	})
	if err != nil {
		panic(err)
	}
	log.Print(documentJSON)

	//Decode JSON
	var documentNormal Document
	errDecoding := json.Unmarshal(documentJSON, &documentNormal)
	if errDecoding != nil {
		panic(errDecoding)
	}
	log.Print(documentNormal.ID)
	http.Redirect(writer, req, "/", http.StatusMovedPermanently)
}
