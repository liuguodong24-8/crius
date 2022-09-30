<?php


namespace Omy\MicroServers;

use Grpc\ChannelCredentials;
use Omy\Crius\DiscoverRequest;
use Omy\Crius\DiscoverServersRequest;
use Omy\MicroServers\Client\AppointmentClient;
use Omy\MicroServers\Client\AssistClient;
use Omy\MicroServers\Client\BasicClient;
use Omy\MicroServers\Client\GrpcClient;
use Omy\MicroServers\Client\MemberAccountClient;
use Omy\MicroServers\Client\MemberExtensionClient;
use Omy\MicroServers\Client\MemberPrivateClient;
use Omy\MicroServers\Client\MerchantBasicClient;
use Omy\MicroServers\Client\MessageClient;
use Omy\MicroServers\Client\PaymentClient;
use Omy\MicroServers\Client\WechatClient;
use Omy\MicroServers\Exception\GrpcNoServerException;
use Omy\MicroServers\Exception\GrpcParseException;

class Service
{
    public const SERVER_BASIC = 'basic';

    public const SERVER_MERCHANT_BASIC = 'merchant-basic';

    public const SERVER_MESSAGE = 'message';

    public const SERVER_APPOINTMENT = 'appointment';

    public const SERVER_PAYMENT = 'payment';

    public const SERVER_WECHAT = 'wechat';

    public const SERVER_MEMBER_ACCOUNT = 'member-account';

    public const SERVER_MEMBER_PRIVATE = 'member-private';

    public const SERVER_EXTENSION = 'member-extension';

    public const SERVER_ASSIST = 'assist';


    private const CLIENTS = [
        self::SERVER_BASIC => BasicClient::class,
        self::SERVER_MERCHANT_BASIC => MerchantBasicClient::class,
        self::SERVER_MESSAGE => MessageClient::class,
        self::SERVER_APPOINTMENT => AppointmentClient::class,
        self::SERVER_PAYMENT => PaymentClient::class,
        self::SERVER_WECHAT => WechatClient::class,
        self::SERVER_MEMBER_ACCOUNT => MemberAccountClient::class,
        self::SERVER_MEMBER_PRIVATE => MemberPrivateClient::class,
        self::SERVER_EXTENSION => MemberExtensionClient::class,
        self::SERVER_ASSIST => AssistClient::class,
    ];

    /**
     * 发现服务
     * @param string $name
     * @return BasicClient
     * @throws GrpcNoServerException
     */
    public static function discoverServer(string $name)
    {
        $option = ['credentials' => ChannelCredentials::createInsecure(), 'timeout' => 5000];

        if (!array_key_exists($name, self::CLIENTS)) {
            throw new GrpcNoServerException(sprintf('server [%s] client not exists', $name));
        }

        $request = new DiscoverRequest();
        $request->setName($name);
        $rpc = sprintf('%s:%d', config('rpc.host'), config('rpc.port'));
        $server = new GrpcClient($rpc, $option);
        [$response, $status] = $server->discover($request);
        if (0 !== $status) {
            throw new GrpcNoServerException(sprintf('no server [%s]', $name));
        }

        $address = sprintf('%s:%d', $response->getIP(), $response->getPort());
        $client = self::CLIENTS[$name];

        return new $client($address, $option);
    }

    /**
     * 发现多个服务
     * @param array $names
     * @return array
     * @throws GrpcNoServerException
     */
    public static function discoverServers(array $names)
    {
        $rpcRequest = new DiscoverServersRequest();
        $rpcRequest->setNames($names);
        $rpc = sprintf('%s:%d', config('rpc.host'), config('rpc.port'));

        $client = new GrpcClient($rpc, [
            'credentials' => ChannelCredentials::createInsecure(),
            'timeout' => 5000
        ]);
        [$res, $status] = $client->discoverServers($rpcRequest);

        if (0 !== $status) {
            throw new GrpcNoServerException('批量发现服务错误[' . implode('/', $names) . ']');
        }

        $services = $res->getServices();
        $clients = [];
        for ($i = 0; $i < $services->count(); ++$i) {
            $name = $names[$i];
            if (!array_key_exists($name, self::CLIENTS)) {
                throw new GrpcNoServerException(sprintf('server [%s] client not exists', $name));
            }

            $service = $services->offsetGet($i);
            $address = sprintf('%s:%d', $service->getIP(), $service->getPort());
            $clientClass = self::CLIENTS[$name];

            $client = new $clientClass($address, [
                'credentials' => ChannelCredentials::createInsecure(),
            ]);

            $clients[$name] = $client;
        }
        return $clients;
    }

    /**
     * rpc 消息解析
     * @param $rpcData
     * @return array
     * @throws GrpcParseException
     */
    public static function rpcHandler($rpcData)
    {
        [$result, $status, $response] = $rpcData;
        if (0 !== $status) {
            throw new GrpcParseException($response->headers['grpc-message'] ?? '');
        }

        $grpcData = json_decode($result->serializeToJsonString(), true);

        if (json_last_error() !== JSON_ERROR_NONE) {
            throw new GrpcParseException('rpc数据解析失败');
        }

        $grpcData = self::snakeCase($grpcData);

        if (!isset($grpcData['error_code']) || $grpcData['error_code'] == 0) {
            return [null, $grpcData['data'] ?? []];
        }

        return [
            [
                'error_code' => $grpcData['error_code'] ?? '',
                'error_message' => $grpcData['error_message'] ?? '',
            ],
            $grpcData['data'] ?? []
        ];
    }

    /**
     * @param $data
     * @return array
     */
    public static function snakeCase($data)
    {
        if (is_array($data)) {
            $snakeData = [];
            foreach ($data as $k => $v) {

                $snakeData[snake_case($k)] = self::snakeCase($v);
            }
            return $snakeData;
        }
        return $data;
    }
}
