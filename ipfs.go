package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bradhe/stopwatch"
	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	uploadToIPFSFromDirectory()
}

func uploadToIPFSFromDirectory() {
	ipfsShell1 := shell.NewShell("ip")
	ipfsShell2 := shell.NewShell("ip")
	// start a timer
	sw := stopwatch.Start()
	// loop of fils in files10mb directory and upload to ipfs shell
	// get list of files in directory
	files, err := ioutil.ReadDir("file10mb")
	if err != nil {
		log.Fatal(err)
	}
	hashes := make([]string, 0)
	// loop through files
	for _, file := range files {
		// get file name
		fileName := file.Name()
		// open file
		file, err := os.Open("file10mb/" + fileName)
		if err != nil {
			log.Fatal(err)
		}
		// upload file to ipfs
		file2, err := os.Open("file10mb/" + fileName)
		if err != nil {
			log.Fatal(err)
		}
		hash, err := ipfsShell1.Add(file)
		if err != nil {
			log.Fatal(err)
		}
		_, err = ipfsShell2.Add(file2)
		if err != nil {
			log.Fatal(err)
		}
		// close file
		file.Close()
		// add hash to list of hashes
		hashes = append(hashes, hash)
	}
	// stop timer
	sw.Stop()
	// print time
	log.Println("Time:", sw.Milliseconds())
	// save hashes to file
	err = ioutil.WriteFile("hashes.txt", []byte(strings.Join(hashes, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	}
	// save time to file
	file, err := os.Create("times_upload.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(sw.Milliseconds().String())
	// restart timer
	sw_total_file := stopwatch.Start()
	// loop through hashes and get file from ipfs
	file_times_each, _ := os.Create("times_cat_each.txt")
	file_times_total, _ := os.Create("times_cat_total.txt")
	file_times_read, _ := os.Create("times_read.txt")
	// ipfs cloud flare
	cloudFlareUrl := "https://cloudflare-ipfs.com/ipfs/"
	for _, hash := range hashes {
		err := fmt.Errorf("starting error")
		sw_each_file := stopwatch.Start()
		// get file from ipfs
		// while error not nil, try again
		var resp *http.Response
		for err != nil {
			// get file from ipfs
			resp, err = http.Get(cloudFlareUrl + hash)
			if err != nil {
				log.Printf("%v", err.Error())
			}
			if resp.StatusCode == http.StatusPermanentRedirect || resp.StatusCode == http.StatusFound {
				// get new hash
				hash = resp.Header.Get("Location")
				// close response
				resp.Body.Close()
			}
		}
		sw_each_file.Stop()
		// add to file
		file_times_each.WriteString(sw_each_file.Milliseconds().String() + "\n")

		sw_read_file := stopwatch.Start()
		for {
			buffer := make([]byte, 1024)
			_, err := resp.Body.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Println("Error:", err)
			}
		}
		sw_read_file.Stop()
		file_times_read.WriteString(sw_read_file.Milliseconds().String() + "\n")
	}
	sw_total_file.Stop()
	file_times_total.WriteString(sw_total_file.Milliseconds().String())
	file_times_each.Close()
	file_times_total.Close()

}

type IPFSResponse struct {
	URL                       string
	Started                   time.Time
	Finished                  time.Time
	TotalTime                 time.Duration
	TotalThroughput           time.Duration
	Redirects                 int
	DownloadedLength          int
	ContentLength             int64
	DownloadSpeed             float64
	DownloadedSpeedThroughput float64
	Err                       string
}

func downloadInformation(url string) (IPFSResponse, error) {
	log.Println("Downloading information from", url)
	started := time.Now()
	sw := stopwatch.Start()
	sw.Start()
	httpClient := &http.Client{Timeout: time.Second * 10}
	resp, err := httpClient.Get(url)
	if err != nil {
		log.Println("Error:", err)
		return IPFSResponse{}, err
	}
	redirects := 0
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusPermanentRedirect || resp.StatusCode == http.StatusTemporaryRedirect {
		redirects++
		location := resp.Header.Get("Location")
		log.Println("Redirecting to", location)
		resp_redirect, err := httpClient.Get(location)
		if err != nil {
			log.Println("Error:", err)
			return IPFSResponse{}, err
		}
		defer resp_redirect.Body.Close()
		if resp_redirect.StatusCode == http.StatusPermanentRedirect || resp_redirect.StatusCode == http.StatusTemporaryRedirect {
			redirects++
			return IPFSResponse{
				URL:                       url,
				Started:                   started,
				Finished:                  time.Now(),
				TotalTime:                 sw.Milliseconds(),
				TotalThroughput:           -1,
				Redirects:                 redirects,
				DownloadedLength:          -1,
				ContentLength:             -1,
				DownloadSpeed:             -1,
				DownloadedSpeedThroughput: -1,
				Err:                       "",
			}, nil
		}
	}
	contentLength := resp.ContentLength
	throughPutMbps := sw.Milliseconds()
	oldthroughPutMbps := throughPutMbps
	accumulatedBytes := 0
	progress := 0
	for {
		buffer := make([]byte, 1024)
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Error:", err)
			return IPFSResponse{}, err
		}
		accumulatedBytes += n
		p := int(float64(accumulatedBytes) / float64(contentLength) * 100)
		el := sw.Milliseconds()
		if p != progress || (el-oldthroughPutMbps) > 1000 {
			progress = p
			throughPutMbps = sw.Milliseconds()
			oldthroughPutMbps = throughPutMbps
			log.Println("Progress:", p, "%")
		}
	}
	finalResult := IPFSResponse{
		URL:                       url,
		Started:                   started,
		Finished:                  time.Now(),
		TotalTime:                 sw.Milliseconds(),
		TotalThroughput:           throughPutMbps,
		Redirects:                 redirects,
		ContentLength:             contentLength,
		DownloadedLength:          accumulatedBytes,
		DownloadSpeed:             float64(accumulatedBytes) / float64(sw.Milliseconds()),
		DownloadedSpeedThroughput: float64(accumulatedBytes) / float64(throughPutMbps),
		Err:                       "",
	}
	log.Println("Finished downloading information from", url)
	return finalResult, nil
}
