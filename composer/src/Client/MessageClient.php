<?php

declare(strict_types=1);

namespace Omy\MicroServers\Client;

use Omy\Crius\BasicServer\EmptyRequest;
use Omy\Crius\MessageServer\ChangeMessageSettingStatusRequest;
use Omy\Crius\MessageServer\ChangeMessageSettingStatusResponse;
use Omy\Crius\MessageServer\ChangeSmsTemplateStatusRequest;
use Omy\Crius\MessageServer\ChangeSmsTemplateStatusResponse;
use Omy\Crius\MessageServer\CreateMessageSettingRequest;
use Omy\Crius\MessageServer\CreateMessageSettingResponse;
use Omy\Crius\MessageServer\CreateSmsTemplateRequest;
use Omy\Crius\MessageServer\CreateSmsTemplateResponse;
use Omy\Crius\MessageServer\CreateWechatTemplateRequest;
use Omy\Crius\MessageServer\CreateWechatTemplateResponse;
use Omy\Crius\MessageServer\GetBranchTemplateRequest;
use Omy\Crius\MessageServer\GetBranchTemplateResponse;
use Omy\Crius\MessageServer\GetShortURLRequest;
use Omy\Crius\MessageServer\GetShortURLResponse;
use Omy\Crius\MessageServer\ListMessageSettingRequest;
use Omy\Crius\MessageServer\ListMessageSettingResponse;
use Omy\Crius\MessageServer\ListOfficialLinkResponse;
use Omy\Crius\MessageServer\ListSmsHistoryRequest;
use Omy\Crius\MessageServer\ListSmsHistoryResponse;
use Omy\Crius\MessageServer\ListSmsTemplateRequest;
use Omy\Crius\MessageServer\ListSmsTemplateResponse;
use Omy\Crius\MessageServer\ListWechatTemplateRequest;
use Omy\Crius\MessageServer\ListWechatTemplateResponse;
use Omy\Crius\MessageServer\MessageVariableResponse;
use Omy\Crius\MessageServer\PBEmpty;
use Omy\Crius\MessageServer\SendSmsRequest;
use Omy\Crius\MessageServer\SendSmsResponse;
use Omy\Crius\MessageServer\SendWechatTemplateRequest;
use Omy\Crius\MessageServer\SendWechatTemplateResponse;
use Omy\Crius\MessageServer\ShowMessageSettingRequest;
use Omy\Crius\MessageServer\ShowMessageSettingResponse;
use Omy\Crius\MessageServer\ShowSmsTemplateRequest;
use Omy\Crius\MessageServer\ShowSmsTemplateResponse;
use Omy\Crius\MessageServer\ShowWechatTemplateRequest;
use Omy\Crius\MessageServer\ShowWechatTemplateResponse;
use Omy\Crius\MessageServer\StatSmsHistoryRequest;
use Omy\Crius\MessageServer\StatSmsHistoryResponse;
use Omy\Crius\MessageServer\UpdateMessageSettingRequest;
use Omy\Crius\MessageServer\UpdateMessageSettingResponse;
use Omy\Crius\MessageServer\UpdateSmsTemplateRequest;
use Omy\Crius\MessageServer\UpdateSmsTemplateResponse;
use Omy\Crius\MessageServer\UpdateWechatTemplateRequest;
use Omy\Crius\MessageServer\UpdateWechatTemplateResponse;
use Omy\Crius\MessageServer\UpdateWechatTemplateStatusRequest;
use Omy\Crius\MessageServer\UpdateWechatTemplateStatusResponse;
use Omy\Crius\MessageServer\WechatStatRequest;
use Omy\Crius\MessageServer\WechatStatResponse;

class MessageClient extends GrpcClient
{
    private const BASE_PATH = '/message.MessageService/';

    public const GET_SHORT_URL = 'GetShortURL';

    public const SEND_SMS = 'SendSms';

    public const GET_BRANCH_TEMPLATE = 'GetBranchTemplate';

    public const SEND_WECHAT_TEMPLATE = 'SendWechatTemplate';

    const GET_MESSAGE_VARIABLE = 'GetMessageVariable';
    const CREATE_SMS_TEMPLATE = 'CreateSmsTemplate';
    const LIST_SMS_TEMPLATE = 'ListSmsTemplate';
    const UPDATE_SMS_TEMPLATE = 'UpdateSmsTemplate';
    const SHOW_SMS_TEMPLATE = 'ShowSmsTemplate';
    const CHANGE_SMS_TEMPLATE_STATUS = 'ChangeSmsTemplateStatus';
    const CREATE_MESSAGE_SETTING = 'CreateMessageSetting';
    const LIST_MESSAGE_SETTING = 'ListMessageSetting';
    const UPDATE_MESSAGE_SETTING = 'UpdateMessageSetting';
    const SHOW_MESSAGE_SETTING = 'ShowMessageSetting';
    const CHANGE_MESSAGE_SETTING_STATUS = 'ChangeMessageSettingStatus';
    const STAT_SMS_HISTORY = 'StatSmsHistory';
    const LIST_SMS_HISTORY = 'ListSmsHistory';

    //微信模板管理
    const CREATE_WECHAT_TEMPLATE = 'CreateWechatTemplate';
    const UPDATE_WECHAT_TEMPLATE = 'UpdateWechatTemplate';
    const LIST_WECHAT_TEMPLATE = 'ListWechatTemplate';
    const UPDATE_WECHAT_TEMPLATE_STATUS = 'UpdateWechatTemplateStatus';
    const SHOW_WECHAT_TEMPLATE = 'ShowWechatTemplate';
    const LIST_OFFICIAL_LINK = 'ListOfficialLink';
    const WECHAT_STAT = 'WechatStat';

    protected array $actionMap = [
        self::WECHAT_STAT => [
            WechatStatRequest::class,
            WechatStatResponse::class,
        ],
        self::LIST_OFFICIAL_LINK => [
            EmptyRequest::class,
            ListOfficialLinkResponse::class,
        ],
        self::SHOW_WECHAT_TEMPLATE => [
            ShowWechatTemplateRequest::class,
            ShowWechatTemplateResponse::class,
        ],
        self::UPDATE_WECHAT_TEMPLATE_STATUS => [
            UpdateWechatTemplateStatusRequest::class,
            UpdateWechatTemplateStatusResponse::class,
        ],
        self::LIST_WECHAT_TEMPLATE => [
            ListWechatTemplateRequest::class,
            ListWechatTemplateResponse::class,
        ],
        self::UPDATE_WECHAT_TEMPLATE => [
            UpdateWechatTemplateRequest::class,
            UpdateWechatTemplateResponse::class,
        ],
        self::CREATE_WECHAT_TEMPLATE => [
            CreateWechatTemplateRequest::class,
            CreateWechatTemplateResponse::class,
        ],
        self::LIST_SMS_HISTORY => [
            ListSmsHistoryRequest::class,
            ListSmsHistoryResponse::class,
        ],
        self::STAT_SMS_HISTORY => [
            StatSmsHistoryRequest::class,
            StatSmsHistoryResponse::class,
        ],
        self::CHANGE_MESSAGE_SETTING_STATUS => [
            ChangeMessageSettingStatusRequest::class,
            ChangeMessageSettingStatusResponse::class,
        ],
        self::SHOW_MESSAGE_SETTING => [
            ShowMessageSettingRequest::class,
            ShowMessageSettingResponse::class,
        ],
        self::UPDATE_MESSAGE_SETTING => [
            UpdateMessageSettingRequest::class,
            UpdateMessageSettingResponse::class,
        ],
        self::LIST_MESSAGE_SETTING => [
            ListMessageSettingRequest::class,
            ListMessageSettingResponse::class,
        ],
        self::CREATE_MESSAGE_SETTING => [
            CreateMessageSettingRequest::class,
            CreateMessageSettingResponse::class,
        ],
        self::GET_MESSAGE_VARIABLE => [
            PBEmpty::class,
            MessageVariableResponse::class,
        ],
        self::CREATE_SMS_TEMPLATE => [
            CreateSmsTemplateRequest::class,
            CreateSmsTemplateResponse::class,
        ],
        self::LIST_SMS_TEMPLATE => [
            ListSmsTemplateRequest::class,
            ListSmsTemplateResponse::class,
        ],
        self::UPDATE_SMS_TEMPLATE => [
            UpdateSmsTemplateRequest::class,
            UpdateSmsTemplateResponse::class,
        ],
        self::SHOW_SMS_TEMPLATE => [
            ShowSmsTemplateRequest::class,
            ShowSmsTemplateResponse::class,
        ],
        self::CHANGE_SMS_TEMPLATE_STATUS => [
            ChangeSmsTemplateStatusRequest::class,
            ChangeSmsTemplateStatusResponse::class,
        ],
        self::GET_SHORT_URL => [
            GetShortURLRequest::class,
            GetShortURLResponse::class,
        ],
        self::SEND_SMS => [
            SendSmsRequest::class,
            SendSmsResponse::class
        ],
        self::GET_BRANCH_TEMPLATE => [
            GetBranchTemplateRequest::class,
            GetBranchTemplateResponse::class
        ],
        self::SEND_WECHAT_TEMPLATE => [
            SendWechatTemplateRequest::class,
            SendWechatTemplateResponse::class
        ]
    ];

    protected function getBasePath()
    {
        return self::BASE_PATH;
    }
}
