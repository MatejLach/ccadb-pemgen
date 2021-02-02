package main

import (
	"log"
	"os"
)

const (
	rootPemCsvUrl         = "https://ccadb-public.secure.force.com/mozilla/IncludedRootsPEMCSV?TrustBitsInclude=Websites"
	intermediatePemCsvUrl = "https://ccadb-public.secure.force.com/mozilla/PublicAllIntermediateCertsWithPEMCSV"
)

func main() {
	if _, err := os.Stat("ca-certs"); os.IsNotExist(err) {
		err = os.Mkdir("ca-certs", 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	intrCacsvReader, err := downloadCertificatesBundleCsv(intermediatePemCsvUrl)
	if err != nil {
		log.Fatal(err)
	}
	parseIntermediateCertificates(intrCacsvReader)

	rootCacsvReader, err := downloadCertificatesBundleCsv(rootPemCsvUrl)
	if err != nil {
		log.Fatal(err)
	}
	parseRootCertificates(rootCacsvReader)
}
