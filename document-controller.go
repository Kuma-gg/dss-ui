package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func documentSave(writer http.ResponseWriter, req *http.Request) {
	type DocumentFile struct {
		ID       string
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

	ID := getMD5Checksum(buf.Bytes())
	//Encode JSON
	documentJSON, err := json.Marshal(DocumentFile{
		ID:       ID,
		Filename: handler.Filename,
		Bytes:    buf.Bytes(),
		Size:     handler.Size,
	})
	if err != nil {
		panic(err)
	}

	log.Print(documentJSON)
	// save Document
	saved := saveDocument(Document{Name:handler.Filename,Size:handler.Size})
	if !saved {
		log.Println("ERROR : create document")
	}
	// send queue  RabbitMQ
	sendFileMessage(documentJSON)
	//Decode JSON
	var documentNormal DocumentFile
	errDecoding := json.Unmarshal(documentJSON, &documentNormal)
	if errDecoding != nil {
		panic(errDecoding)
	}
	log.Print(documentNormal.Bytes)
	log.Print(documentNormal.ID)
	log.Print(documentNormal.Filename)
	log.Print(documentNormal.Size)

	http.Redirect(writer, req, "/#documents", http.StatusMovedPermanently)
}

func documentDelete(writer http.ResponseWriter, req *http.Request) {
	//Encode JSON
	document := Document{
		ID: req.FormValue("id"),
	}

	documentJSON, err := json.Marshal(document)
	if err != nil {
		panic(err)
	}
	log.Print(documentJSON)
	// Delete document
	deleteDocument(document)
	//Decode JSON
	var documentNormal Document
	errDecoding := json.Unmarshal(documentJSON, &documentNormal)
	if errDecoding != nil {
		panic(errDecoding)
	}
	log.Print(documentNormal.ID)
	http.Redirect(writer, req, "/", http.StatusMovedPermanently)
}

func getMD5Checksum(content []byte) string {
	hasher := md5.New()
	hasher.Write(content)
	return hex.EncodeToString(hasher.Sum(nil))
}


