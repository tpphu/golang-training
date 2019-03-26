package helper

import (
	"fmt"

	tm "github.com/buger/goterm"
)

type Watcher struct {
	//load info
	DBLoadUrlReq   int
	DBLoadUrlRes   int
	DBLoadUrlErr   int
	DBLoadUrlTotal int
	//download info
	NumHTTPReq int
	NumHTTPRes int
	NumHTTPErr int
	//update info
	DBInsArticleReq int
	DBInsArticleRes int
	DBInsArticleErr int
	//save info
}

func (self *Watcher) GetDBLoadUrlInfo() string {
	return fmt.Sprintf("%d/%d/%d", self.DBLoadUrlReq, self.DBLoadUrlRes, self.DBLoadUrlErr)
}

func (self *Watcher) GetHTTPDownloadInfo() string {
	return fmt.Sprintf("%d/%d/%d", self.NumHTTPReq, self.NumHTTPRes, self.NumHTTPErr)
}

func (self *Watcher) GetDBInsArticleInfo() string {
	return fmt.Sprintf("%d/%d/%d", self.DBInsArticleReq, self.DBInsArticleRes, self.DBInsArticleErr)
}

func (self *Watcher) Out() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	// Based on http://golang.org/pkg/text/tabwriter
	totals := tm.NewTable(0, 10, 5, ' ', 0)
	fmt.Fprintf(totals, "[DB]LoadUrl\t[DB]LoadUrlTotal\t[HTTP]Download\t[DB]InsArticle\n")
	fmt.Fprintf(totals, "%s\t%d\t%s\t%s\n",
		self.GetDBLoadUrlInfo(),
		self.DBLoadUrlTotal,
		self.GetHTTPDownloadInfo(),
		self.GetDBInsArticleInfo())
	tm.Println(totals)

	tm.Flush()
}
