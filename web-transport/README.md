# Web Transport Go Example

## UDP Buffer Sizes
https://github.com/quic-go/quic-go/wiki/UDP-Buffer-Sizes

It is recommended to increase the maximum buffer size on Ubuntu by running:
```
sudo sysctl -w net.core.rmem_max=2500000
sudo sysctl -w net.core.wmem_max=2500000
```
This command would increase the maximum send and the receive buffer size to roughly 2.5 MB.

## Certificates
Place your domains SSL certificates in the `certs` folder.
