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

type DocumentFile struct {
	ID       string
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

	ID := getMD5Checksum(buf.Bytes())
	//Encode JSON
	documentJSON, err := json.Marshal(DocumentFile{
		ID:       ID,
		Filename: handler.Filename,
		Bytes:    buf.Bytes(),
		Size:     handler.Size,
		Type:     "create",
	})
	if err != nil {
		panic(err)
	}

	//log.Print(documentJSON)
	// save Document
	saved := saveDocument(Document{Name: handler.Filename, Size: handler.Size})
	if !saved {
		log.Println("ERROR : create document")
	}
	// send queue  RabbitMQ
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
	id :=req.FormValue("id")
	document := Document{
		ID: id,
	}

	documentJSON, err := json.Marshal(document)

	if err != nil {
		panic(err)
	}
	log.Print(documentJSON)
	// send command Rabbit
	user:=getUserById(id)
	comand, err := json.Marshal(DocumentFile{
		ID: id  ,
		Filename:user.Name,
		Type: "delete",
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

func getMD5Checksum(content []byte) string {
	hasher := md5.New()
	hasher.Write(content)
	return hex.EncodeToString(hasher.Sum(nil))
}
