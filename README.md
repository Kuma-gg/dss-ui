# dss-ui
Final Project UI component


# dss-ui
Component made with Go, displays UI with user and document register. Sends data to [dss-email component](https://github.com/Kuma-gg/dss-email) and [dss-storage component](https://github.com/Kuma-gg/dss-storage) using Brokers

## Getting Started

Just follow the instruction to copy and run the project.

### Prerequisites

* Install Go from its page at https://golang.org/.
* Install PostgreSQL from its page at https://www.postgresql.org/
* Create a Database in PostgreSQL and run this script:
```
CREATE TABLE documents (
   id integer NOT NULL,
   name character(200),
   size integer
);

ALTER TABLE documents OWNER TO postgres;

CREATE SEQUENCE documents_id_seq
   START WITH 1
   INCREMENT BY 1
   NO MINVALUE
   NO MAXVALUE
   CACHE 1;

ALTER TABLE documents_id_seq OWNER TO postgres;

ALTER SEQUENCE documents_id_seq OWNED BY documents.id;

ALTER TABLE ONLY documents ALTER COLUMN id SET DEFAULT nextval('documents_id_seq'::regclass);

CREATE TABLE users (
   id integer NOT NULL,
   name text,
   age integer,
   email character(50),
   first_name character(50),
   last_name character(50)
);

ALTER TABLE users OWNER TO postgres;

CREATE SEQUENCE users_id_seq
   START WITH 1
   INCREMENT BY 1
   NO MINVALUE
   NO MAXVALUE
   CACHE 1;

ALTER TABLE users_id_seq OWNER TO postgres;

ALTER SEQUENCE users_id_seq OWNED BY users.id;

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);

ALTER TABLE ONLY users
   ADD CONSTRAINT users_pkey PRIMARY KEY (id);
```

### Installing

* Clone the project to your src Golang project folder, on windows its on C:\Users\go\src. 
* In the folder open a terminal an write the following:
```
go get -u github.com/golang/dep/cmd/dep
```
```
dep init -v
```
```
dep ensure -v 
```
* Configure RabbitMQ and Postgresql credentials in main.go file
* Execute the http server:
```
go run main.go connector-sql.go document-controller.go user-controller.go index-controller.go mailsender.go mailreceiver.go receiver.go sender.go repo.go
```
* Go to: [http://localhost:3000](http://localhost:3000)

### Output

* The UI server recieves data from dss-email after inserting or removing documents and prints them in the console
* The UI server recieves data from dss-storage after creating a new document and prints it in the console

Example:

```
2018/12/05 11:16:11 mails sent to alejandro2222 Mail : luis@gmail.com                                 Event : created
2018/12/05 11:16:11 mails sent to josue2222 Mail : josue_147_15@hotmail.com                           Event : created
2018/12/05 11:16:11 mails sent to apagar-MV Mail : luis@gmail.com                                     Event : created
2018/12/05 11:16:11 INFO : Sent successfully
```

## Built With

* [Go](https://golang.org/) - Programming langauge
* [Rabbitmq](https://www.rabbitmq.com/) - Queue Messages

## Authors

* **Luis Daniel Lopez** - [lolpez](https://github.com/lolpez)

### To Do
* Test
* Socket.io implementation to display messages from Brokers in UI
