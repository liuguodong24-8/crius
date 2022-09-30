<?php

declare(strict_types=1);

namespace  Omy\MicroServers\ThirdParty;

use GuzzleHttp\Client;
use GuzzleHttp\Exception\GuzzleException;
use RuntimeException;

class Kuaidi100Server
{
    private int $timeout = 10;

    private string $contentType = 'application/x-www-form-urlencoded';

    // 实时查询请求地址
    private string $queryUrl = 'http://poll.kuaidi100.com/poll/query.do';

    // 查询地图轨迹
    private string $mapTrack = 'https://poll.kuaidi100.com/poll/maptrack.do';

    //智能识别
    private string $autoUrl = 'http://www.kuaidi100.com/autonumber/auto';

    //订阅请求
    private string $poll = 'https://poll.kuaidi100.com/poll';

    private array $company = [
        'yuantong' => '圆通速递',
        'yunda' => '韵达快递',
        'zhongtong' => '中通快递',
        'shentong' => '申通快递',
        'huitongkuaidi' => '百世快递',
        'shunfeng' => '顺丰速运',
        'youzhengguonei' => '邮政快递包裹',
        'ems' => 'EMS',
        'jd' => '京东物流',
        'jtexpress' => '极兔速递',
        'youzhengbk' => '邮政标准快递',
        'debangwuliu' => '德邦',
        'debangkuaidi' => '德邦快递',
        'yuantongkuaiyun' => '圆通快运',
        'youshuwuliu' => '优速快递',
        'zhaijisong' => '宅急送',
        'yundakuaiyun' => '韵达快运',
        'baishiwuliu' => '百世快运',
        'zhongtongkuaiyun' => '中通快运',
        'zhongtongguoji' => '中通国际',
    ];
    /**
     * 快递查询
     * https://api.kuaidi100.com/document/5f0ffb5ebc8da837cbd8aefc.html
     * @param string $com 快递公司编码
     * @param string $num 快递单号
     * @return mixed
     */
    public function query(string $com, string $num)
    {
        $param = array(
            'com' => $com,
            'num' => $num,
        );
        $params = [
            'customer' => config('server.kuaidi100.customer'),
            'param' => json_encode($param),
            'sign' => strtoupper(md5(json_encode($param) . config('server.kuaidi100.key') . config('server.kuaidi100.customer'))),
        ];

        $response = $this->request($this->queryUrl, $params);

        return $response;
    }

    /**
     * 识别单号
     * @param string $num
     * @return mixed
     */
    public function getAutoComCode(string $num)
    {
        $params = [
            'num' => $num,
            'key' => config('server.kuaidi100.key')
        ];
        $response = $this->request($this->autoUrl, $params, 'GET');

        return $response;
    }

    /**
     * 订阅
     * @param array $attributes
     *  string $attributes.company // 物流公司编码
     *  string $attributes.number // 订单号
     *  string $attributes.from // 发货地
     *  string $attributes.to // 收货地货地
     * callback_poll 回调地址
     * @return mixed
     */
    public function poll(array $attributes)
    {
        $keys = ['company', 'number'];
        $param = [];
        foreach ($keys as $key) {
            $param[$key] = $attributes[$key] ?? '';
        }
        //回调地址
        $param['key'] = config('server.kuaidi100.key');
        $param['parameters']["callbackurl"] =  $attributes['callback_poll'];
        $params = [
            'schema' => 'json',
            'param' => json_encode($param),
        ];
        $response = $this->request($this->poll, $params, 'POST');

        return $response;
    }

    /**
     * 查询地图轨迹
     * @param array $attributes
     * @return mixed
     */
    public function getMapTrack(array $attributes)
    {
        $keys = ['com', 'num', 'phone', 'from', 'to', 'show', 'order','orderTime'];
        $param = [];
        foreach ($keys as $key) {
            $param[$key] = $attributes[$key] ?? '';
        }
        $params = [
            'customer' => config('server.kuaidi100.customer'),
            'param' => $param,
            'sign' => strtoupper(md5(json_encode($param) . config('server.kuaidi100.key') . config('server.kuaidi100.customer'))),
        ];
        $response = $this->request($this->mapTrack, $params);
        return $response;
    }


    /**
     * 获取物流公司
     * @return mixed
     */
    public function getlogisticsCompany()
    {
        return $this->company;
    }
    /**
     * @param $url
     * @param $params
     * @param string $method
     * @return mixed
     */
    private function request($url, $params, $method = 'POST')
    {
        $client = new Client();
        $body = [];

        if ($method === 'POST') {
            $body = [
                'header' => ['Content-Type' => $this->contentType],
                'form_params' => $params,
                'timeout' => $this->timeout
            ];
        }

        if ($method === 'GET') {
            $body = [
                'header' => ['Content-Type' => $this->contentType],
                'query' => $params,
                'timeout' => $this->timeout
            ];
        }
        try {
            $result = $client->request($method, $url, $body);

            return json_decode($result->getBody()->getContents(), true);
        } catch (GuzzleException $e) {
            throw new RuntimeException(sprintf('请求快递100错误:%s', $e->getMessage()));
        }
    }
}
