# arris-exporter
Scrap the arris interface and export prometheus metrics

### Acceptable Signal Levels
from https://arris.secure.force.com/consumers/articles/General_FAQs/SB6183-Cable-Signal-Levels

| **Recommended Max Downstream <br>Power Level (DPL)** | **Recommended Min Downstream <br>Power Level (DPL)** |
| -------------------------------------------- | ---------------------------------------------------- |
| +15 dBmV	                                   | -15 dBmV                                             |


| **Downstream Signal to Noise Ratio (SNR)** | | |
| -------------------------------------- | - | - |
| **Modulation** | **Downstream Power Level** |	**Acceptable DS SNR** |
| 64 QAM	| n/a |	| 23.5 dB or greater |
| 256 QAM | -6 dBmV to +15 dBmV <br> -6 dBmV to -15 dBmV | 30 dB or greater<br>33 dB or greater |

| **Upstream Transmit Power Level** | | | | |
| - | - | - | - | - |
| **Channel** | **US Channel Type** | **Symbol Rate** |	**Recommended Max<br> US Power Level** | **Recommended Min <br>US Power Level** |
| Single |	TDMA | 1280 Ksym/sec | +61 dBmV | 45 dBmV |
|        | ATDMA | 2560 Ksym/sec | +58 dBmV | 45 dBmV |
|        |       | 5120 Ksym/sec | +57 dBmV | 45 dBmV |
| Two    |	TDMA | 1280 Ksym/sec | +58 dBmV | 45 dBmV |
|        | ATDMA | 2560 Ksym/sec | +55 dBmV | 45 dBmV |
|        |       | 5120 Ksym/sec | +54 dBmV | 45 dBmV |
| Three  |	TDMA | 1280 Ksym/sec | +55 dBmV | 45 dBmV |
|        | ATDMA | 2560 Ksym/sec | +52 dBmV | 45 dBmV |
|        |       | 5120 Ksym/sec | +51 dBmV | 45 dBmV |
