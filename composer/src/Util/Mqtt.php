<?php


namespace Omy\MicroServers\Util;


use Simps\MQTT\Client as MQTTClient;
use Simps\MQTT\Config\ClientConfig;

class Mqtt
{
    private static $client;

    const SWOOLE_MQTT_CONFIG = [
        'open_mqtt_protocol' => true,
        'package_max_length' => 2 * 1024 * 1024,
        'connect_timeout' => 5.0,
        'write_timeout' => 5.0,
        'read_timeout' => 5.0,
    ];

    public static function getClient()
    {
        if (!self::$client) {
            $client = new MQTTClient(config('mqtt.host'), config('mqtt.port'), self::getConfig());
            $will = [
                'topic' => 'omy',
                'qos' => 1,
                'retain' => 0,
                'message' => 'byebye',
            ];
            $client->connect(true, $will);
            self::$client = $client;
        }

        return self::$client;
    }



    private static function getConfig()
    {
        $config = new ClientConfig();

        return $config->setUserName('')
            ->setPassword('')
            ->setClientId(MQTTClient::genClientID())
            ->setKeepAlive(10)
            ->setDelay(3000) // 3s
            ->setMaxAttempts(5)
            ->setSwooleConfig(self::SWOOLE_MQTT_CONFIG);
    }
}