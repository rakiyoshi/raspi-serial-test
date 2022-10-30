# raspi-serial-test
This is a test repository for serial communication


## Prerequisite

### Let UART0 available

Reference: [Raspberry Pi Documentation \- Configuration](https://www.raspberrypi.com/documentation/computers/configuration.html)

#### Disable Serial Console
1. execure the command below
```bash
sudo raspi-config
```

1. Select `3 Interface Options    Configure connections to peripherals`

1. Select `I6 Serial Port   Enable/disable shell messages on the serial connection`

1. Select `<No>` for `Would you like a login shell to be accessible over serial?`

1. Select `<Yes>` for `Would you like the serial port hardware to be enabled?`

#### Disable bluetooth
```bash
echo dtoverlay=disable-bt | tee -a /boot/config.txt
```

#### Reboot
```bash
sudo reboot
```


## Build

### Resolve dependencies

```bash
bazelisk run //:gazelle -- update-repos -from_file=go.mod -to_macro=deps.bzl%go_dependencies
```

### Build

```bash
bazelisk build //cmd/transmitter
bazelisk build //cmd/receiver
```

### Run

```bash
bazelisk run //cmd/transmitter
bazelisk run //cmd/receiver
```

## TODO
https://github.com/aspect-build/aspect-cli/blob/main/release/release.bzl
