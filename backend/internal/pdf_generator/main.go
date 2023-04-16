package pdf_generator

import (
	"context"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

var chromeCtx context.Context
var allocCtxCancel context.CancelFunc
var chromeCtxCancel context.CancelFunc

func InitChrome(opts []func(*chromedp.ExecAllocator)) error {
	var allocCtx context.Context

	if opts == nil {
		opts = append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
			chromedp.Flag("disable-gpu", true),
			chromedp.Flag("no-sandbox", true),
			chromedp.Flag("disable-extensions", true),
			chromedp.Flag("disable-setuid-sandbox", true),
			chromedp.Flag("disable-dev-shm-usage", true),
		)
	}

	allocCtx, allocCtxCancel = chromedp.NewExecAllocator(context.Background(), opts...)
	chromeCtx, chromeCtxCancel = chromedp.NewContext(allocCtx)

	return nil
}

func CloseChrome() {
	chromeCtxCancel()
	allocCtxCancel()
}

func GeneratePdf(url string) ([]byte, error) {
	var pdfBuffer []byte

	if err := chromedp.Run(
		chromeCtx,
		chromedp.Tasks{
			chromedp.Navigate(url),
			chromedp.ActionFunc(func(ctx context.Context) error {
				buf, _, err := page.PrintToPDF().WithPrintBackground(true).Do(ctx)
				if err != nil {
					return err
				}
				pdfBuffer = buf
				return nil
			}),
		},
	); err != nil {
		return nil, err
	}

	return pdfBuffer, nil
}
