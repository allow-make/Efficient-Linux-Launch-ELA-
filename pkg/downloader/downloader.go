package downloader

import (
"crypto/sha256"
"encoding/hex"
"fmt"
"io"
"net/http"
"os"
"path/filepath"
"time"
)

type Progress struct {
Total      int64
Downloaded int64
Percent    float64
Speed      float64
}

type Distro struct {
Name     string
Version  string
Arch     string
URL      string
Size     int64
Checksum string
}

type Downloader struct {
client *http.Client
}

func NewDownloader() *Downloader {
return &Downloader{
client: &http.Client{Timeout: time.Hour},
}
}

func (d *Downloader) Download(distro Distro, destDir string, progress chan Progress) error {
resp, err := d.client.Get(distro.URL)
if err != nil {
return err
}
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK {
return fmt.Errorf("HTTP error: %s", resp.Status)
}
total := resp.ContentLength
if total == -1 {
total = distro.Size
}
destFile := filepath.Join(destDir, fmt.Sprintf("%s-%s-%s.iso", distro.Name, distro.Version, distro.Arch))
out, err := os.Create(destFile)
if err != nil {
return err
}
defer out.Close()
hasher := sha256.New()
writer := io.MultiWriter(out, hasher)
buf := make([]byte, 1024*1024)
var downloaded int64
var startTime time.Time
for {
n, err := resp.Body.Read(buf)
if n > 0 {
if _, werr := writer.Write(buf[:n]); werr != nil {
return werr
}
downloaded += int64(n)
if progress != nil {
elapsed := time.Since(startTime)
var speed float64
if elapsed.Seconds() > 0 {
speed = float64(downloaded) / elapsed.Seconds()
}
progress <- Progress{
Total:      total,
Downloaded: downloaded,
Percent:    float64(downloaded) / float64(total) * 100,
Speed:      speed,
}
}
}
if err == io.EOF {
break
}
if err != nil {
return err
}
}
if distro.Checksum != "" {
hash := hex.EncodeToString(hasher.Sum(nil))
if hash != distro.Checksum {
return fmt.Errorf("checksum mismatch: expected %s, got %s", distro.Checksum, hash)
}
}
return nil
}

func (d *Downloader) ListDistros() []Distro {
return []Distro{
{Name: "Ubuntu", Version: "22.04", Arch: "amd64", URL: "https://releases.ubuntu.com/22.04/ubuntu-22.04-desktop-amd64.iso", Size: 4800000000},
{Name: "Debian", Version: "12", Arch: "amd64", URL: "https://cdimage.debian.org/debian-cd/current/amd64/iso-cd/debian-12.0.0-amd64-netinst.iso", Size: 600000000},
{Name: "Arch", Version: "latest", Arch: "x86_64", URL: "https://archlinux.org/releng/releases/latest/torrent/archlinux-x86_64.iso", Size: 800000000},
{Name: "Kali", Version: "2024.1", Arch: "amd64", URL: "https://cdimage.kali.org/kali-2024.1/kali-linux-2024.1-installer-amd64.iso", Size: 4000000000},
}
}
