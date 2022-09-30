<?php

declare(strict_types=1);

namespace Omy\MicroServers\Client;

use Omy\Crius\AppointmentServer\AppointmentCancelOpenRoomRequest;
use Omy\Crius\AppointmentServer\AppointmentConfig;
use Omy\Crius\AppointmentServer\AppointmentOpenRoomRequest;
use Omy\Crius\AppointmentServer\ArriveAppointmentRequest;
use Omy\Crius\AppointmentServer\CancelAppointmentRequest;
use Omy\Crius\AppointmentServer\CancelAppointmentResponse;
use Omy\Crius\AppointmentServer\CreateAppointmentRequest;
use Omy\Crius\AppointmentServer\CreateAppointmentResponse;
use Omy\Crius\AppointmentServer\CreateAppointmentThemeCategoryRequest;
use Omy\Crius\AppointmentServer\CreateAppointmentThemeFeatureRequest;
use Omy\Crius\AppointmentServer\CreateAppointmentThemeRequest;
use Omy\Crius\AppointmentServer\CreateCallerRecordRequest;
use Omy\Crius\AppointmentServer\CreateCallerRecordResponse;
use Omy\Crius\AppointmentServer\DeleteAppointmentLockRequest;
use Omy\Crius\AppointmentServer\DeleteAppointmentRequest;
use Omy\Crius\AppointmentServer\GetAppointmentBreachLimitRequest;
use Omy\Crius\AppointmentServer\GetAppointmentBreachLimitResponse;
use Omy\Crius\AppointmentServer\GetAppointmentConfigResponse;
use Omy\Crius\AppointmentServer\GetAppointmentLimitRequest;
use Omy\Crius\AppointmentServer\GetAppointmentLimitResponse;
use Omy\Crius\AppointmentServer\GetAppointmentRecordRequest;
use Omy\Crius\AppointmentServer\GetAppointmentRecordResponse;
use Omy\Crius\AppointmentServer\GetAppointmentRemainingFastRequest;
use Omy\Crius\AppointmentServer\GetAppointmentRemainingFastResponse;
use Omy\Crius\AppointmentServer\GetAppointmentRemainingRequest;
use Omy\Crius\AppointmentServer\GetAppointmentRemainingResponse;
use Omy\Crius\AppointmentServer\GetAppointmentsRequest;
use Omy\Crius\AppointmentServer\GetAppointmentsResponse;
use Omy\Crius\AppointmentServer\GetAppointmentTemplateConfigsRequest;
use Omy\Crius\AppointmentServer\GetAppointmentTemplateConfigsResponse;
use Omy\Crius\AppointmentServer\GetAppointmentTemplatesRequest;
use Omy\Crius\AppointmentServer\GetAppointmentTemplatesResponse;
use Omy\Crius\AppointmentServer\GetAppointmentThemeCategoriesRequest;
use Omy\Crius\AppointmentServer\GetAppointmentThemeCategoriesResponse;
use Omy\Crius\AppointmentServer\GetAppointmentThemeFeaturesRequest;
use Omy\Crius\AppointmentServer\GetAppointmentThemeFeaturesResponse;
use Omy\Crius\AppointmentServer\GetAppointmentThemesByRoomTypeRequest;
use Omy\Crius\AppointmentServer\GetAppointmentThemesByRoomTypeResponse;
use Omy\Crius\AppointmentServer\GetAppointmentThemesRequest;
use Omy\Crius\AppointmentServer\GetAppointmentThemesResponse;
use Omy\Crius\AppointmentServer\GetCallerRecordsRequest;
use Omy\Crius\AppointmentServer\GetCallerRecordsResponse;
use Omy\Crius\AppointmentServer\GetCallersRequest;
use Omy\Crius\AppointmentServer\GetCallersResponse;
use Omy\Crius\AppointmentServer\GetTemplateCalendarRequest;
use Omy\Crius\AppointmentServer\GetTemplateCalendarResponse;
use Omy\Crius\AppointmentServer\GetWechatAppointmentsRequest;
use Omy\Crius\AppointmentServer\GetWechatAppointmentsResponse;
use Omy\Crius\AppointmentServer\JoinSharedAppointmentRequest;
use Omy\Crius\AppointmentServer\PayAppointmentRequest;
use Omy\Crius\AppointmentServer\PBEmpty;
use Omy\Crius\AppointmentServer\RefundAppointmentRequest;
use Omy\Crius\AppointmentServer\Response;
use Omy\Crius\AppointmentServer\SaveAppointmentLockRequest;
use Omy\Crius\AppointmentServer\SaveAppointmentLockResponse;
use Omy\Crius\AppointmentServer\SaveAppointmentTemplateRequest;
use Omy\Crius\AppointmentServer\SaveAppointmentTemplateResponse;
use Omy\Crius\AppointmentServer\SaveAppointmentTempNumberRequest;
use Omy\Crius\AppointmentServer\ShowAppointmentRequest;
use Omy\Crius\AppointmentServer\ShowAppointmentResponse;
use Omy\Crius\AppointmentServer\ShowAppointmentTemplateRequest;
use Omy\Crius\AppointmentServer\ShowAppointmentTemplateResponse;
use Omy\Crius\AppointmentServer\ShowAppointmentThemeCategoryRequest;
use Omy\Crius\AppointmentServer\ShowAppointmentThemeCategoryResponse;
use Omy\Crius\AppointmentServer\ShowAppointmentThemeFeatureRequest;
use Omy\Crius\AppointmentServer\ShowAppointmentThemeFeatureResponse;
use Omy\Crius\AppointmentServer\ShowAppointmentThemeRequest;
use Omy\Crius\AppointmentServer\ShowAppointmentThemeResponse;
use Omy\Crius\AppointmentServer\ShowCallerByPhoneRequest;
use Omy\Crius\AppointmentServer\ShowCallerByPhoneResponse;
use Omy\Crius\AppointmentServer\UpdateAppointmentRequest;
use Omy\Crius\AppointmentServer\UpdateAppointmentResponse;
use Omy\Crius\AppointmentServer\UpdateAppointmentShareRequest;
use Omy\Crius\AppointmentServer\UpdateAppointmentStatusRequest;
use Omy\Crius\AppointmentServer\UpdateAppointmentStatusResponse;
use Omy\Crius\AppointmentServer\UpdateAppointmentTemplateStatusRequest;
use Omy\Crius\AppointmentServer\UpdateAppointmentThemeCategoryRequest;
use Omy\Crius\AppointmentServer\UpdateAppointmentThemeCategoryStatusRequest;
use Omy\Crius\AppointmentServer\UpdateAppointmentThemeFeatureRequest;
use Omy\Crius\AppointmentServer\UpdateAppointmentThemeFeatureStatusRequest;
use Omy\Crius\AppointmentServer\UpdateAppointmentThemeRequest;
use Omy\Crius\AppointmentServer\UpdateAppointmentThemeStatusRequest;
use Omy\Crius\AppointmentServer\UpdateAppointmentTradeInfoRequest;
use Omy\Crius\AppointmentServer\UpdateCallerBlackRequest;
use Omy\Crius\AppointmentServer\UpdateCallerRecordActionRequest;
use Omy\Crius\AppointmentServer\UpdateCallerRequest;
use Omy\Crius\AppointmentServer\UpdateTemplateCalendarRequest;

class AppointmentClient extends GrpcClient
{
    private const BASE_PATH = '/appointment.AppointmentServer/';

    public const CREATE_APPOINTMENT = 'CreateAppointment';

    public const UPDATE_APPOINTMENT = 'UpdateAppointment';

    public const SAVE_APPOINTMENT_LOCK = 'SaveAppointmentLock';

    public const UPDATE_APPOINTMENT_STATUS = 'UpdateAppointmentStatus';

    public const GET_APPOINTMENTS = 'GetAppointments';

    public const GET_APPOINTMENT_REMAINING = 'GetAppointmentRemaining';

    public const GET_APPOINTMENT_RECORD = 'GetAppointmentRecord';

    public const SAVE_APPOINTMENT_TEMP_NUMBER = 'SaveAppointmentTempNumber';

    public const SHOW_APPOINTMENT = 'ShowAppointment';

    public const DELETE_APPOINTMENT_LOCK = 'DeleteAppointmentLock';

    public const UPDATE_APPOINTMENT_TRADE_INFO = 'UpdateAppointmentTradeInfo';

    public const GET_APPOINTMENT_LIMIT = 'GetAppointmentLimit';

    public const GET_APPOINTMENT_REMAINING_FAST = 'GetAppointmentRemainingFast';

    public const GET_CONFIG = 'GetAppointmentConfig';

    public const DELETE_APPOINTMENT = 'DeleteAppointment';

    public const GET_BREACH_APPOINTMENT_LIMIT = 'GetAppointmentBreachLimit';

    public const GET_CALLERS = 'GetCallers';

    public const GET_CALLER_RECORDS = 'GetCallerRecords';

    public const SHOW_CALLER_BY_PHONE = 'ShowCallerByPhone';

    public const CREATE_CALLER_RECORD = 'CreateCallerRecord';

    public const UPDATE_CALLER = 'UpdateCaller';

    public const UPDATE_CALLER_BLACK = 'UpdateCallerBlack';

    public const UPDATE_CALLER_RECORD_ACTION = 'UpdateCallerRecordAction';

    public const GET_APPOINTMENT_TEMPLATES = 'GetAppointmentTemplates';

    public const GET_TEMPLATE_CALENDAR = 'GetTemplateCalendar';

    public const UPDATE_TEMPLATE_CALENDAR = 'UpdateTemplateCalendar';

    public const SAVE_APPOINTMENT_TEMPLATE = 'SaveAppointmentTemplate';


    public const UPDATE_APPOINTMENT_TEMPLATE_STATUS = 'UpdateAppointmentTemplateStatus';

    public const SET_CONFIG = 'UpdateAppointmentConfig';

    public const SHOW_APPOINTMENT_TEMPLATE = 'ShowAppointmentTemplate';

    //主题预约
    public const GET_APPOINTMENT_THEMES_BY_ROOM_TYPE = 'GetAppointmentThemesByRoomType';

    public const SHOW_APPOINTMENT_THEME = 'ShowAppointmentTheme';

    public const GET_APPOINTMENT_THEMES = 'GetAppointmentThemes';

    public const UPDATE_APPOINTMENT_THEME = 'UpdateAppointmentTheme';

    public const UPDATE_APPOINTMENT_THEME_STATUS = 'UpdateAppointmentThemeStatus';

    public const CREATE_APPOINTMENT_THEME = 'CreateAppointmentTheme';

    //主题模板类型
    public const CREATE_APPOINTMENT_THEME_CATEGORY = 'CreateAppointmentThemeCategory';

    public const UPDATE_APPOINTMENT_THEME_CATEGORY = 'UpdateAppointmentThemeCategory';

    public const GET_APPOINTMENT_THEME_CATEGORIES = 'GetAppointmentThemeCategories';

    public const SHOW_APPOINTMENT_THEME_CATEGORY = 'ShowAppointmentThemeCategory';

    public const UPDATE_APPOINTMENT_THEME_CATEGORY_STATUS = 'UpdateAppointmentThemeCategoryStatus';

    public const CREATE_APPOINTMENT_THEME_FEATURE = 'CreateAppointmentThemeFeature';

    public const UPDATE_APPOINTMENT_THEME_FEATURE = 'UpdateAppointmentThemeFeature';

    public const GET_APPOINTMENT_THEME_FEATURES = 'GetAppointmentThemeFeatures';

    public const SHOW_APPOINTMENT_THEME_FEATURE = 'ShowAppointmentThemeFeature';

    public const UPDATE_APPOINTMENT_THEME_FEATURE_STATUS = 'UpdateAppointmentThemeFeatureStatus';

    public const GET_APPOINTMENT_TEMPLATE_CONFIGS = 'GetAppointmentTemplateConfigs';

    //预约分享
    public const JOIN_SHARED_APPOINTMENT = 'JoinSharedAppointment';

    public const UPDATE_APPOINTMENT_SHARE = 'UpdateAppointmentShare';

    public const GET_WECHAT_APPOINTMENTS = 'GetWechatAppointments';

    //对接POS
    public const APPOINTMENT_OPEN_ROOM = 'AppointmentOpenRoom';

    public const APPOINTMENT_CANCEL_OPEN_ROOM = 'AppointmentCancelOpenRoom';

     public const PAY_APPOINTMENT = 'PayAppointment';

     public const CANCEL_APPOINTMENT = 'CancelAppointment';

     public const ARRIVE_APPOINTMENT = 'ArriveAppointment';

     public const REFUND_APPOINTMENT = 'RefundAppointment';


    protected function getBasePath()
    {
        return self::BASE_PATH;
    }

    protected array $actionMap = [
        self::PAY_APPOINTMENT => [
            PayAppointmentRequest::class,
            Response::class
        ],
        self::CANCEL_APPOINTMENT => [
            CancelAppointmentRequest::class,
            CancelAppointmentResponse::class,
        ],
        self::ARRIVE_APPOINTMENT => [
            ArriveAppointmentRequest::class,
            Response::class,
        ],
        self::REFUND_APPOINTMENT => [
            RefundAppointmentRequest::class,
            Response::class,
        ],

        self::APPOINTMENT_OPEN_ROOM => [
            AppointmentOpenRoomRequest::class,
            Response::class
        ],
        self::APPOINTMENT_CANCEL_OPEN_ROOM => [
            AppointmentCancelOpenRoomRequest::class,
            Response::class
        ],
        self::GET_APPOINTMENT_TEMPLATE_CONFIGS => [
            GetAppointmentTemplateConfigsRequest::class,
            GetAppointmentTemplateConfigsResponse::class
        ],
        self::CREATE_APPOINTMENT_THEME_FEATURE => [
            CreateAppointmentThemeFeatureRequest::class,
            Response::class
        ],
        self::UPDATE_APPOINTMENT_THEME_FEATURE => [
            UpdateAppointmentThemeFeatureRequest::class,
            Response::class
        ],
        self::GET_APPOINTMENT_THEME_FEATURES => [
            GetAppointmentThemeFeaturesRequest::class,
            GetAppointmentThemeFeaturesResponse::class
        ],
        self::SHOW_APPOINTMENT_THEME_FEATURE => [
            ShowAppointmentThemeFeatureRequest::class,
            ShowAppointmentThemeFeatureResponse::class
        ],
        self::UPDATE_APPOINTMENT_THEME_FEATURE_STATUS => [
            UpdateAppointmentThemeFeatureStatusRequest::class,
            Response::class
        ],
        self::UPDATE_APPOINTMENT_THEME_CATEGORY_STATUS => [
            UpdateAppointmentThemeCategoryStatusRequest::class,
            Response::class
        ],
        self::CREATE_APPOINTMENT_THEME_CATEGORY => [
            CreateAppointmentThemeCategoryRequest::class,
            Response::class
        ],
        self::UPDATE_APPOINTMENT_THEME_CATEGORY => [
            UpdateAppointmentThemeCategoryRequest::class,
            Response::class
        ],
        self::GET_APPOINTMENT_THEME_CATEGORIES => [
            GetAppointmentThemeCategoriesRequest::class,
            GetAppointmentThemeCategoriesResponse::class
        ],
        self::SHOW_APPOINTMENT_THEME_CATEGORY => [
            ShowAppointmentThemeCategoryRequest::class,
            ShowAppointmentThemeCategoryResponse::class
        ],
        self::SET_CONFIG => [
            AppointmentConfig::class,
            Response::class
        ],
        self::SAVE_APPOINTMENT_TEMP_NUMBER => [
            SaveAppointmentTempNumberRequest::class,
            Response::class
        ],
        self::GET_APPOINTMENT_RECORD => [
            GetAppointmentRecordRequest::class,
            GetAppointmentRecordResponse::class
        ],

        self::UPDATE_APPOINTMENT_TEMPLATE_STATUS => [
            UpdateAppointmentTemplateStatusRequest::class,
            Response::class
        ],
        self::SHOW_APPOINTMENT_TEMPLATE => [
            ShowAppointmentTemplateRequest::class,
            ShowAppointmentTemplateResponse::class
        ],
        self::SAVE_APPOINTMENT_TEMPLATE => [
            SaveAppointmentTemplateRequest::class,
            SaveAppointmentTemplateResponse::class
        ],
        self::UPDATE_TEMPLATE_CALENDAR => [
            UpdateTemplateCalendarRequest::class,
            Response::class
        ],
        self::GET_TEMPLATE_CALENDAR => [
            GetTemplateCalendarRequest::class,
            GetTemplateCalendarResponse::class
        ],
        self::GET_APPOINTMENT_TEMPLATES => [
            GetAppointmentTemplatesRequest::class,
            GetAppointmentTemplatesResponse::class
        ],
        self::UPDATE_CALLER_RECORD_ACTION => [
            UpdateCallerRecordActionRequest::class,
            Response::class
        ],
        self::UPDATE_CALLER_BLACK => [
            UpdateCallerBlackRequest::class,
            Response::class
        ],
        self::UPDATE_CALLER => [
            UpdateCallerRequest::class,
            Response::class
        ],
        self::CREATE_CALLER_RECORD => [
            CreateCallerRecordRequest::class,
            CreateCallerRecordResponse::class
        ],
        self::CREATE_APPOINTMENT => [
            CreateAppointmentRequest::class,
            CreateAppointmentResponse::class
        ],
        self::SHOW_CALLER_BY_PHONE => [
            ShowCallerByPhoneRequest::class,
            ShowCallerByPhoneResponse::class
        ],
        self::GET_CALLER_RECORDS => [
            GetCallerRecordsRequest::class,
            GetCallerRecordsResponse::class
        ],
        self::SHOW_APPOINTMENT => [
            ShowAppointmentRequest::class,
            ShowAppointmentResponse::class
        ],
        self::GET_APPOINTMENTS => [
            GetAppointmentsRequest::class,
            GetAppointmentsResponse::class,
        ],
        self::GET_APPOINTMENT_LIMIT => [
            GetAppointmentLimitRequest::class,
            GetAppointmentLimitResponse::class,
        ],
        self::GET_APPOINTMENT_REMAINING => [
            GetAppointmentRemainingRequest::class,
            GetAppointmentRemainingResponse::class,
        ],
        self::SAVE_APPOINTMENT_LOCK => [
            SaveAppointmentLockRequest::class,
            SaveAppointmentLockResponse::class
        ],
        self::DELETE_APPOINTMENT_LOCK => [
            DeleteAppointmentLockRequest::class,
            Response::class
        ],
        self::GET_APPOINTMENT_REMAINING_FAST => [
            GetAppointmentRemainingFastRequest::class,
            GetAppointmentRemainingFastResponse::class
        ],
        self::UPDATE_APPOINTMENT => [
            UpdateAppointmentRequest::class,
            UpdateAppointmentResponse::class
        ],
        self::GET_CONFIG => [
            PBEmpty::class,
            GetAppointmentConfigResponse::class,
        ],
        self::UPDATE_APPOINTMENT_STATUS => [
            UpdateAppointmentStatusRequest::class,
            UpdateAppointmentStatusResponse::class,
        ],
        self::DELETE_APPOINTMENT => [
            DeleteAppointmentRequest::class,
            Response::class
        ],
        self::GET_BREACH_APPOINTMENT_LIMIT => [
            GetAppointmentBreachLimitRequest::class,
            GetAppointmentBreachLimitResponse::class
        ],
        self::UPDATE_APPOINTMENT_TRADE_INFO => [
            UpdateAppointmentTradeInfoRequest::class,
            Response::class
        ],
        self::GET_CALLERS => [
            GetCallersRequest::class,
            GetCallersResponse::class
        ],
        self::GET_APPOINTMENT_THEMES_BY_ROOM_TYPE => [
            GetAppointmentThemesByRoomTypeRequest::class,
            GetAppointmentThemesByRoomTypeResponse::class
        ],
        self::SHOW_APPOINTMENT_THEME => [
            ShowAppointmentThemeRequest::class,
            ShowAppointmentThemeResponse::class
        ],
        self::CREATE_APPOINTMENT_THEME => [
            CreateAppointmentThemeRequest::class,
            Response::class
        ],
        self::UPDATE_APPOINTMENT_THEME => [
            UpdateAppointmentThemeRequest::class,
            Response::class
        ],
        self::UPDATE_APPOINTMENT_THEME_STATUS => [
            UpdateAppointmentThemeStatusRequest::class,
            Response::class
        ],
        self::GET_APPOINTMENT_THEMES => [
            GetAppointmentThemesRequest::class,
            GetAppointmentThemesResponse::class
        ],
        self::JOIN_SHARED_APPOINTMENT => [
            JoinSharedAppointmentRequest::class,
            Response::class
        ],
        self::UPDATE_APPOINTMENT_SHARE => [
            UpdateAppointmentShareRequest::class,
            Response::class
        ],
        self::GET_WECHAT_APPOINTMENTS => [
            GetWechatAppointmentsRequest::class,
            GetWechatAppointmentsResponse::class
        ]
    ];
}
