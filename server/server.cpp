#include <functional>
#include <string>

#ifdef _MSC_VER
#pragma comment(lib, "ws2_32.lib")
#endif

#include <ladder/Buffer.h>
#include <ladder/Connection.h>
#include <ladder/Logging.h>

#include "server.h"

using namespace ladder;
using namespace std::placeholders;

void OnMessage(const ConnectionPtr& conn, Buffer* buffer) {
  std::string message = buffer->ReadAll();
  LOG_INFO("Received: " + message);
  conn->Send("Echoback~ " + message);
}

void OnConnection(const ConnectionPtr& conn) {
  conn->Send("Hello -- This is a ladder server.");
}

Server::Server(unsigned short port) {
  SocketAddr addr("0.0.0.0", static_cast<uint16_t>(port), false);
  server_ = new TcpServer(addr, false);
  server_->SetConnectionCallback(std::bind(OnConnection, _1));
  server_->SetReadCallback(std::bind(OnMessage, _1, _2));
}

Server::~Server() {
  delete server_;
}

void Server::Start() {
  server_->Start();
}
