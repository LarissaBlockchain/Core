## Larissa Network
The Larissa Network is a decentralized blockchain platform designed for secure, scalable, and fast transactions. It supports smart contracts, decentralized applications (dApps), and peer-to-peer transactions on a permissionless network. Built for interoperability, Larissa allows seamless communication with other blockchain networks. It is optimized for high throughput and low latency, making it suitable for enterprise and public blockchain use cases. Larissa enables developers to build decentralized financial systems and other applications with a focus on security and decentralization.
### Building the source

Building `geth` requires both a Go (version 1.14 or later) and a C compiler. You can install them using your favorite package manager. Once the dependencies are installed, run:
```shell
make geth
```
or, to build the full suite of utilities:
```shell
make all
```
### How to deploy node and export RPC interface (Linux)
#### 1. Create a non-privileged user
```shell
useradd -m larissa
```
#### 2. Create directories to store blockchain and node software
```shell
mkdir /home/larissa/data/geth
mkdir /home/larissa/opt/geth
```
#### 3. Download node software for Ubuntu
```shell
curl https://github.com/LarissaBlockchain/Core/releases/download/v1.12.0/geth-ubuntu-x86_64 -o /home/larissa/opt/geth/geth-ubuntu-x86_64
chmod +x /home/larissa/opt/geth/geth-ubuntu-x86_64
chown -R larissa:larissa /home/larissa
```
#### 4. Create a systemd unit file
```shell
cat > /etc/systemd/system/larissa-node.service << EOF
[Unit]
Description=Larissa Node
After=network.target
[Service]
User=larissa
Group=larissa
Restart=on-failure
RestartSec=10
ExecStart=/home/larissa/opt/geth/geth-ubuntu-x86_64 --datadir /home/larissa/data/geth \
    --port 38000 \
    --http \
    --http.addr 127.0.0.1 \
    --http.port 8545 \
    --http.vhosts * \
    --http.api "eth,net,web3" \
    --ws \
    --ws.addr 127.0.0.1 \
    --ws.port 8546 \
    --ws.origins * \
    --ws.api "eth,net,web3" \
    --cache 512 \
    --maxpeers 128
[Install]
WantedBy=default.target
EOF
```
#### 5. Enable and start the service
```shell
systemctl daemon-reload
systemctl enable larissa-node
systemctl start larissa-node
```
#### 6. Check the status and logs
```shell
systemctl status larissa-node
```
```shell
journalctl -u larissa-node
```
---
### How to deploy node and export RPC interface (Windows)
#### 1. Download the node software for Windows
Download the node software from the link below and save it to a desired folder:
[Download geth-windows-x64.exe](https://github.com/LarissaBlockchain/Core/releases/download/v1.12.0/geth-windows-x64.exe)
#### 2. Create directories to store blockchain data
In the command prompt, create the necessary directories for blockchain data storage:
```cmd
mkdir C:\larissa\data\geth
mkdir C:\larissa\node
```
#### 3. Move the downloaded `geth-windows-x64.exe` file
Move the `geth-windows-x64.exe` file to `C:\larissa\node`.
#### 4. Run the Larissa node
Open a command prompt with Administrator privileges, navigate to the `C:\larissa\node` folder, and start the node using the following command:
```cmd
geth-windows-x64.exe --datadir C:\larissa\data\geth ^
    --port 38000 ^
    --http ^
    --http.addr 127.0.0.1 ^
    --http.port 8545 ^
    --http.vhosts * ^
    --http.api "eth,net,web3" ^
    --ws ^
    --ws.addr 127.0.0.1 ^
    --ws.port 8546 ^
    --ws.origins * ^
    --ws.api "eth,net,web3" ^
    --cache 512 ^
    --maxpeers 128
```
#### 5. Check status and logs
Check the node's status and logs by reviewing the output in the command prompt window.
---
### Common RPC Commands
Once the node is running, you can access the JSON-RPC interface and WebSocket interface on the following ports:
- **HTTP RPC**: `http://127.0.0.1:8545`
- **WebSocket RPC**: `ws://127.0.0.1:8546`
You can interact with the node using the following APIs: `eth`, `net`, `web3`.
