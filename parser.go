package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func parseIntermediateCertificates(reader io.ReadCloser) {
	defer reader.Close()

	csvReader := csv.NewReader(reader)

	// ignore headers
	_, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		entry, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			continue
		}

		filename := fmt.Sprintf("%s.pem", entry[8])
		err = ioutil.WriteFile(fmt.Sprintf("ca-certs/%s", filename), []byte(strings.Trim(entry[23], "'")), 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func parseRootCertificates(reader io.ReadCloser) {
	defer reader.Close()

	csvReader := csv.NewReader(reader)

	// ignore headers
	_, err := csvReader.Read()
	if err != nil {
		log.Fatal(err)
	}

	for {
		entry, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			continue
		}

		rand.Seed(time.Now().UnixNano())
		filename := fmt.Sprintf("%d.pem", rand.Intn(2000000000-1000000002+1)+1000000002)
		err = ioutil.WriteFile(fmt.Sprintf("ca-certs/%s", filename), []byte(strings.Trim(entry[0], "'")), 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func downloadCertificatesBundleCsv(uri string) (io.ReadCloser, error) {
	if _, err := url.Parse(uri); err != nil {
		return nil, err
	}

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
