# wt-video-audio-transmission-example

# Installation for compilation
```bash
sudo dnf install libadwaita-devel gobject-introspection-devel gtk4-devel x264-devel v4l-utils
```

# Initialization

```bash
bash create_certs.sh
```

```bash
sudo sysctl -w net.core.rmem_max=2500000
sudo sysctl -w net.core.wmem_max=2500000
```

# Development

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
sudo dnf install protobuf-compiler

protoc -I=app/connections --go_out=app/connections app/connections/messagepb/message.proto
```