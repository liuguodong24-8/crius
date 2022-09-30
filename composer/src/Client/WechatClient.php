<?php

declare(strict_types=1);

namespace Omy\MicroServers\Client;

use Omy\Crius\WechatServer\MiniProgramCodeToSessionRequest;
use Omy\Crius\WechatServer\MiniProgramCodeToSessionResponse;
use Omy\Crius\WechatServer\MiniProgramDecryptRequest;
use Omy\Crius\WechatServer\MiniProgramDecryptResponse;
use Omy\Crius\WechatServer\OfficialAuthURLRequest;
use Omy\Crius\WechatServer\OfficialAuthURLResponse;
use Omy\Crius\WechatServer\OfficialCodeGetUserRequest;
use Omy\Crius\WechatServer\OfficialCodeGetUserResponse;
use Omy\Crius\WechatServer\OfficialJsSdkRequest;
use Omy\Crius\WechatServer\OfficialJsSdkResponse;
use Omy\Crius\WechatServer\OfficialOpenidGetUserRequest;
use Omy\Crius\WechatServer\OfficialOpenidGetUserResponse;
use Omy\Crius\WechatServer\SendTemplateMessageRequest;
use Omy\Crius\WechatServer\SendTemplateMessageResponse;

class WechatClient extends GrpcClient
{
    private const BASE_PATH = '/wechat.WechatService/';

    public const OFFICIAL_GET_AUTH_URL = 'OfficialGetAuthURL';

    public const OFFICIAL_CODE_GET_USER = 'OfficialCodeGetUser';

    public const SEND_TEMPLATE_MESSAGE = 'SendTemplateMessage';

    public const OFFICIAL_JS_SDK = 'OfficialJsSdk';

    public const OFFICIAL_OPENID_GET_USER = 'OfficialOpenidGetUser';

    public const MINI_PROGRAM_CODE_TO_SESSION = 'MiniProgramCodeToSession';

    public const MINI_PROGRAM_DECRYPT = 'MiniProgramDecrypt';

    protected array $actionMap = [
        self::OFFICIAL_GET_AUTH_URL => [
            OfficialAuthURLRequest::class,
            OfficialAuthURLResponse::class,
        ],
        self::OFFICIAL_CODE_GET_USER => [
            OfficialCodeGetUserRequest::class,
            OfficialCodeGetUserResponse::class
        ],
        self::SEND_TEMPLATE_MESSAGE => [
            SendTemplateMessageRequest::class,
            SendTemplateMessageResponse::class
        ],
        self::OFFICIAL_JS_SDK => [
            OfficialJsSdkRequest::class,
            OfficialJsSdkResponse::class
        ],
        self::OFFICIAL_OPENID_GET_USER => [
            OfficialOpenidGetUserRequest::class,
            OfficialOpenidGetUserResponse::class
        ],
        self::MINI_PROGRAM_CODE_TO_SESSION => [
            MiniProgramCodeToSessionRequest::class,
            MiniProgramCodeToSessionResponse::class
        ],

        self::MINI_PROGRAM_DECRYPT => [
            MiniProgramDecryptRequest::class,
            MiniProgramDecryptResponse::class
        ]
    ];

    protected function getBasePath()
    {
        return self::BASE_PATH;
    }
}
