# macOSのネットワークインターフェース

```bash
# loopback
lo0: flags=8049<UP,LOOPBACK,RUNNING,MULTICAST> mtu 16384

# Hardware Port: Ethernet
en0: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500

# Hardware Port: Wi-Fi
en1: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500

# Hardware Port: Thunderbolt 1
en2: flags=8963<UP,BROADCAST,SMART,RUNNING,PROMISC,SIMPLEX,MULTICAST> mtu 1500

# Hardware Port: Thunderbolt 2
en3: flags=8963<UP,BROADCAST,SMART,RUNNING,PROMISC,SIMPLEX,MULTICAST> mtu 1500

# Hardware Port: Ethernet Adapter (en4)
en4: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500

# Hardware Port: Ethernet Adapter (en5)
en5: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500

# Hardware Port: Thunderbolt Bridge
bridge0: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500

# WiFiテザリング (Access Point)
ap1: flags=8802<BROADCAST,SIMPLEX,MULTICAST> mtu 1500

# AirDrop (Apple WireLess Direct Link)
awdl0: flags=8843<UP,BROADCAST,RUNNING,SIMPLEX,MULTICAST> mtu 1500

# 低遅延WiFi (llw  Low Latency WLAN)
llw0: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500

# IPv6/IPv4トンネリングを行うときに使うインターフェース (generic tunnel interface)
gif0: flags=8010<POINTOPOINT,MULTICAST> mtu 1280

# IPv6パケットをIPv4ネットワークにルーティングするためのインターフェース (6to4 tunnel interface)
stf0: flags=0<> mtu 1280

# ???
anpi0: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500
anpi1: flags=8863<UP,BROADCAST,SMART,RUNNING,SIMPLEX,MULTICAST> mtu 1500

# ???
utun0: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 1380
utun1: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 2000
utun2: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 1000
utun3: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 1500
utun4: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 1380
utun5: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 1380
utun6: flags=8051<UP,POINTOPOINT,RUNNING,MULTICAST> mtu 16000
```

![[other/macos-net-interface.png]]
