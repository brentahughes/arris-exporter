package arris

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/brentahughes/arris-exporter/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	infoEndpoint   = "/RgSwInfo.asp"
	statusEndpoint = "/RgConnect.asp"
)

type Scrapper struct {
	host   string
	client *http.Client
}

func NewScrapper(host string) *Scrapper {
	return &Scrapper{
		host: host,
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (s *Scrapper) Parse(ctx context.Context) error {
	if err := s.getModemInfo(ctx); err != nil {
		return err
	}
	if err := s.getStatus(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Scrapper) getModemInfo(ctx context.Context) error {
	doc, err := s.doRequest(ctx, fmt.Sprintf("http://%s%s", s.host, infoEndpoint))
	if err != nil {
		return err
	}

	metrics.Info.With(prometheus.Labels{
		"model":            doc.Find("#thisModelNumberIs").First().Text(),
		"software_version": doc.Find("#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(4) > td:nth-child(2)").First().Text(),
		"hardware_version": doc.Find("#bg3 > div.container > div.content > table:nth-child(2) > tbody > tr:nth-child(3) > td:nth-child(2)").First().Text(),
	}).Set(1)

	uptimeText := doc.Find("#bg3 > div.container > div.content > table:nth-child(5) > tbody > tr:nth-child(2) > td:nth-child(2)").First().Text()

	parts := strings.SplitAfterN(uptimeText, "days ", 2)
	daysStr := strings.TrimSuffix(parts[0], " days ")
	days, err := strconv.Atoi(daysStr)
	if err != nil {
		return err
	}

	durationStr := strings.ReplaceAll(parts[1], ":", "")
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return err
	}

	totalDuration := duration + (time.Duration(days) * time.Hour * 24)
	metrics.Uptime.Set(totalDuration.Seconds())
	return nil
}

func (s *Scrapper) getStatus(ctx context.Context) error {
	doc, err := s.doRequest(ctx, fmt.Sprintf("http://%s%s", s.host, statusEndpoint))
	if err != nil {
		return err
	}

	var caughtErr error

	channelTable := doc.Find("#bg3 > div.container > div.content > form > center:nth-child(4) > table > tbody > tr")
	channelTable.Each(func(i int, s *goquery.Selection) {
		if i < 2 {
			return
		}

		channel := s.Find("td:nth-child(1)").Text()

		channelLabels := prometheus.Labels{
			"channel":    channel,
			"status":     s.Find("td:nth-child(2)").Text(),
			"modulation": s.Find("td:nth-child(3)").Text(),
			"channel_id": s.Find("td:nth-child(4)").Text(),
			"frequency":  s.Find("td:nth-child(5)").Text(),
		}

		powerStr := strings.TrimSpace(strings.TrimSuffix(s.Find("td:nth-child(6)").Text(), " dBmV"))
		power, err := strconv.ParseFloat(powerStr, 64)
		if err != nil {
			caughtErr = fmt.Errorf("error getting power on channel %s: %v", channel, err)
			return
		}

		snrStr := strings.TrimSuffix(s.Find("td:nth-child(7)").Text(), " dB")
		snr, err := strconv.ParseFloat(snrStr, 64)
		if err != nil {
			caughtErr = fmt.Errorf("error getting snr on channel %s: %v", channel, err)
			return
		}

		correctedStr := s.Find("td:nth-child(8)").Text()
		corrected, err := strconv.ParseFloat(correctedStr, 64)
		if err != nil {
			caughtErr = fmt.Errorf("error getting corrected on channel %s: %v", channel, err)
			return
		}

		uncorrectedStr := s.Find("td:nth-child(9)").Text()
		uncorrected, err := strconv.ParseFloat(uncorrectedStr, 64)
		if err != nil {
			caughtErr = fmt.Errorf("error getting uncorrected on channel %s: %v", channel, err)
			return
		}

		metrics.ChannelPower.With(channelLabels).Set(power)
		metrics.ChannelSNR.With(channelLabels).Set(snr)
		metrics.ChannelCorrected.With(channelLabels).Set(corrected)
		metrics.ChannelUncorrectable.With(channelLabels).Set(uncorrected)
	})

	return caughtErr
}

func (s *Scrapper) doRequest(ctx context.Context, url string) (*goquery.Document, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	return goquery.NewDocumentFromReader(resp.Body)
}
