<?php

declare(strict_types=1);

namespace Omy\MicroServers\Client;

use Omy\Crius\Service;
use Hyperf\Utils\Context;
use Omy\Crius\DiscoverRequest;
use Hyperf\GrpcClient\BaseClient;
use Omy\Crius\DiscoverServersRequest;
use Omy\Crius\DiscoverServersResponse;
use Hyperf\GrpcServer\Exception\GrpcException;
use Psr\Http\Message\ServerRequestInterface;

class GrpcClient extends BaseClient
{
    protected array $metadata = [];

    public function __construct(string $hostname, array $options = [])
    {
        parent::__construct($hostname, $options);
        $request = Context::get(ServerRequestInterface::class);
        if ($request) {
            $metadata = $request->getAttribute('metadata', []);
        } else {
            $metadata = Context::get('metadata', []);
        }
        foreach ($metadata as $k => $v) {
            $this->metadata[$k] = $v;
        }
        $this->metadata['merchant_id'] = config('merchant_id');
        $this->metadata['sleuth_code'] = Context::get('sleuth_code');
    }

    public function discover(DiscoverRequest $request, $metadata = [])
    {
        return $this->_simpleRequest(
            '/proto.CriusService/Discover',
            $request,
            [Service::class, 'decode'],
            array_merge($this->metadata, $metadata)
        );
    }

    public function discoverServers(DiscoverServersRequest $request, $metadata = [])
    {
        return $this->_simpleRequest(
            '/proto.CriusService/DiscoverServers',
            $request,
            [DiscoverServersResponse::class, 'decode'],
            array_merge($this->metadata, $metadata)
        );
    }

    public function request(string $method, $request, $metadata = [])
    {
        if (! isset($this->actionMap[$method])) {
            throw new GrpcException(sprintf('grpc method %s error', $method));
        }
        list($requestObj, $response) = $this->actionMap[$method];
        if (! $request instanceof $requestObj) {
            throw new GrpcException(sprintf('request method should be instanceof %s', $requestObj));
        }

        return $this->_simpleRequest(
            $this->getBasePath() . $method, //组装method,
            $request,
            [$response, 'decode'],
            array_merge($this->metadata, $metadata),
            $options = []
        );
    }
}
