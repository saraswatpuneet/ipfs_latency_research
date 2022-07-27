package main

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/bradhe/stopwatch"
)

func main() {
	log.Println("Hello World")
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
