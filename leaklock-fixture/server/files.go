// Fixture for Leak Lock: intentionally insecure code. Never deploy or copy.
package server

import (
	"crypto/md5"
	"crypto/tls"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func download(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("file")
	data, _ := os.ReadFile(filepath.Join("/srv/files", name))
	w.Write(data)
}

func convert(r *http.Request) {
	exec.Command("sh", "-c", "convert "+r.URL.Query().Get("img")).Run()
}

func lookup(db *sql.DB, id string) {
	db.Query(fmt.Sprintf("SELECT * FROM orders WHERE id = '%s'", id))
}

func client() *http.Client {
	return &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
}

func etag(b []byte) [16]byte { return md5.Sum(b) }
