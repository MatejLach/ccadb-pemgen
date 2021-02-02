## ccadb-pemgen

This is a simple script to generate fresh certificate store from [CCADB](https://www.ccadb.org) provided CSV PEM files. 
Both Root and intermediate certificates are processed and the output is stored in the `ca-certs` directory. 

Generate hashes for the directory with:

`c_rehash ca-certs\` 

(`c_rehash` is provided by `openssl`)

Test with:

`curl --capath ca-certs/ https://uk.store.asus.com`

Useful for your system trust store to match your browser store and avoid any unknown authority SSL errors.
