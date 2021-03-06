#include <proxygen/httpserver/ResponseBuilder.h>

#include "ServiceWorkerHandler.h"

using namespace proxygen;

ServiceWorkerHandler::ServiceWorkerHandler(
    std::shared_ptr<MasternodeConfig> config,
    std::shared_ptr<ServiceWorker> sw):
    config_(config),
    sw_(sw) {
        CHECK(config_) << "Config object was null";
        CHECK(sw_) << "Service Worker object was null";
    }

void ServiceWorkerHandler::onRequest(
    std::unique_ptr<proxygen::HTTPMessage> headers) noexcept {
    // reply with service worker content
    ResponseBuilder(downstream_)
        .status(200, "OK")
        .header("Content-Type", "application/javascript")
        .body(sw_->getPayload())
        .sendWithEOM();
}

void ServiceWorkerHandler::onBody(
    std::unique_ptr<folly::IOBuf> body) noexcept {}

void ServiceWorkerHandler::onUpgrade(
    proxygen::UpgradeProtocol protocol) noexcept {}

void ServiceWorkerHandler::onEOM() noexcept {}

void ServiceWorkerHandler::requestComplete() noexcept { delete this; }

void ServiceWorkerHandler::onError(proxygen::ProxygenError err) noexcept { delete this; }
