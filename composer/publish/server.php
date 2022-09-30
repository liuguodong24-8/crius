<?php

declare(strict_types=1);

/**
 * 第三方配置
 */
return [
    'kuaidi100' => [ // 快递100查询
        'key' => env('KUAIDI100_KEY' ,'hjaJeuzv4770'),
        'customer' => env('KUAIDI100_CUSTOMER', 'AFEF09FEFBE8BAC4CAFC876D91153B6E'),
        'callback_poll' => env('KUAIDI100_CALLBACK_POLL', '') //回调地址
    ],
];