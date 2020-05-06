package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const promNamespace = "arris_modem"

var (
	Uptime = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: promNamespace,
			Name:      "uptime",
			Help:      "Time in seconds the modem has been booted",
		},
	)

	Info = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: promNamespace,
			Name:      "info",
			Help:      "Hardware and software information about the modem",
		},
		[]string{"model", "hardware_version", "software_version"},
	)

	BootStatus = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: promNamespace,
			Name:      "boot_status",
			Help:      "Info about the boot sequence",
		},
		[]string{"type"},
	)

	DownstreamChannelPower = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: promNamespace,
			Subsystem: "downstream",
			Name:      "channel_power",
			Help:      "Power info in dBmV for a channel",
		},
		[]string{"channel", "status", "modulation", "channel_id", "frequency"},
	)

	DownstreamChannelSNR = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: promNamespace,
			Subsystem: "downstream",
			Name:      "channel_snr",
			Help:      "SNR info in dB for a channel",
		},
		[]string{"channel", "status", "modulation", "channel_id", "frequency"},
	)

	DownstreamChannelCorrected = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: promNamespace,
			Subsystem: "downstream",
			Name:      "channel_corrected",
			Help:      "Packets corrected for a channel",
		},
		[]string{"channel", "status", "modulation", "channel_id", "frequency"},
	)

	DownstreamChannelUncorrectable = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: promNamespace,
			Subsystem: "downstream",
			Name:      "channel_uncorrectable",
			Help:      "Packets uncorrectable for a channel",
		},
		[]string{"channel", "status", "modulation", "channel_id", "frequency"},
	)

	UpstreamChannelPower = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: promNamespace,
			Subsystem: "upstream",
			Name:      "channel_power",
			Help:      "Power info in dBmV for a channel",
		},
		[]string{"channel", "status", "type", "channel_id", "symbol_rate", "frequency"},
	)
)
