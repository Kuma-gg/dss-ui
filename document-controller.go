package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

//DocumentFile structure
type DocumentFile struct {
	ID       int
	Filename string
	Bytes    []byte
	Size     int64
	Type     string
}

func documentSave(writer http.ResponseWriter, req *http.Request) {

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

	if err != nil {
		panic(err)
	}

	//log.Print(documentJSON)
	// save Document
	idSaved := saveDocument(Document{Name: handler.Filename, Size: handler.Size})
	if idSaved == 0 {
		log.Println("ERROR : create document")
	}
	// send queue  RabbitMQ
	documentJSON, err := json.Marshal(DocumentFile{
		ID:       idSaved,
		Filename: handler.Filename,
		Bytes:    buf.Bytes(),
		Size:     handler.Size,
		Type:     "create",
	})
	sendFileStorageMessage(documentJSON)
	//Decode JSON
	var documentNormal DocumentFile
	errDecoding := json.Unmarshal(documentJSON, &documentNormal)
	if errDecoding != nil {
		panic(errDecoding)
	}
	http.Redirect(writer, req, "/#documents", http.StatusMovedPermanently)
}

func documentDelete(writer http.ResponseWriter, req *http.Request) {
	//Encode JSON
	id := req.FormValue("id")
	document := Document{
		ID: id,
	}

	documentJSON, err := json.Marshal(document)

	if err != nil {
		panic(err)
	}
	log.Print(documentJSON)
	// send command Rabbit
	i, err := strconv.Atoi(id)
	user := getUserByID(id)
	comand, err := json.Marshal(DocumentFile{
		ID:       i,
		Filename: user.Name,
		Type:     "delete",
	})

	if err != nil {
		panic(err)
	}
	sendFileStorageMessage(comand)
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
