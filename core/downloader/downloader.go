// Package downloader is the main module of GO_SPIDER for download page.
package downloader

import (
    "github.com/hu17889/core/common/page"
    "github.com/hu17889/core/common/request"
)

type Downloader interface {
    Download(req *request.Request) *page.Page
}
