package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// FfmpegArgs is the object to save to bolt
type FfmpegArgs struct {
	Cmd  string   `json:"cmd"`
	Args []string `json:"args"`
}

// Callffmpeg route to create shorten entry
func Callffmpeg(w http.ResponseWriter, req *http.Request) {
	var fargs FfmpegArgs
	err := json.NewDecoder(req.Body).Decode(&fargs)
	switch {
	case err == io.EOF:
		// empty body
		w.WriteHeader(404)
		fmt.Fprint(w, "exit status 1\n\n")
		fmt.Fprint(w, "Err: Missing arguments")

		err = errors.New("Err: Missing arguments")
		log.Printf("%s\n", err)
		return
	case err != nil:
		w.WriteHeader(404)
		fmt.Fprint(w, "exit status 1\n\n")
		fmt.Fprint(w, "Err: JSON format error (should be {args:['arg1', 'arg2']})")

		err = errors.New("Err: JSON format error")
		log.Printf("%v\n", req.Body)
		return
	}

	cmd := "ffmpeg"
	switch {
	case fargs.Cmd == "ffprobe":
		cmd = "ffprobe"
		break
	case fargs.Cmd == "qt-faststart":
		cmd = "qt-faststart"
		break
	case fargs.Cmd == "ffmpeg":
	default:
		cmd = "ffmpeg"
		break
	}
	if len(fargs.Args) > 0 {
		args := strings.Join(fargs.Args[:], " ")
		arga := strings.Split(args, " ")
		cmd := exec.Command(cmd, arga...)
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, err)
			fmt.Fprint(w, "\n\n")
			fmt.Fprint(w, string(stdoutStderr))

			log.Printf("Err:\n")
			log.Printf(string(stdoutStderr))
			return
		}
		w.WriteHeader(200)
		fmt.Fprint(w, string(stdoutStderr))
		return
	}
	w.WriteHeader(404)
	fmt.Fprint(w, "Page Not Found")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/execute", Callffmpeg).Methods("POST")
	log.Fatal(http.ListenAndServe(":32468", router))
}
