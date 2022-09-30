<?php

declare(strict_types=1);

namespace Omy\MicroServers\Client;


use Omy\Crius\MemberAccountServer\ActiveBlankCardRequest;
use Omy\Crius\MemberAccountServer\ActiveBlankCardResponse;
use Omy\Crius\MemberAccountServer\ActivePrimaryCardRequest;
use Omy\Crius\MemberAccountServer\ActivePrimaryCardResponse;
use Omy\Crius\MemberAccountServer\ActiveSecondaryCardRequest;
use Omy\Crius\MemberAccountServer\ActiveSecondaryCardResponse;
use Omy\Crius\MemberAccountServer\AddAccountDeductionRequest;
use Omy\Crius\MemberAccountServer\BindCardRequest;
use Omy\Crius\MemberAccountServer\BindCardResponse;
use Omy\Crius\MemberAccountServer\CancelCardRequest;
use Omy\Crius\MemberAccountServer\ConsumeRequest;
use Omy\Crius\MemberAccountServer\ConsumeResponse;
use Omy\Crius\MemberAccountServer\GetAccountByCardCodeRequest;
use Omy\Crius\MemberAccountServer\GetAccountByCardCodeResponse;
use Omy\Crius\MemberAccountServer\GetAccountByMemberIDRequest;
use Omy\Crius\MemberAccountServer\GetAccountsByCardIDRequest;
use Omy\Crius\MemberAccountServer\GetAccountsByCardIDResponse;
use Omy\Crius\MemberAccountServer\GetAccountsRequest;
use Omy\Crius\MemberAccountServer\GetAccountsResponse;
use Omy\Crius\MemberAccountServer\GetBillsRequest;
use Omy\Crius\MemberAccountServer\GetBillsResponse;
use Omy\Crius\MemberAccountServer\GetCardsByAccountRequest;
use Omy\Crius\MemberAccountServer\GetCardsByAccountResponse;
use Omy\Crius\MemberAccountServer\GetCardsRequest;
use Omy\Crius\MemberAccountServer\GetCardsResponse;
use Omy\Crius\MemberAccountServer\GetMemberAccountResponse;
use Omy\Crius\MemberAccountServer\GetMemberAccountsRequest;
use Omy\Crius\MemberAccountServer\GetMemberAccountsResponse;
use Omy\Crius\MemberAccountServer\MakeCardRequest;
use Omy\Crius\MemberAccountServer\MakeCardResponse;
use Omy\Crius\MemberAccountServer\NoDataResponse;
use Omy\Crius\MemberAccountServer\RechargeCardRequest;
use Omy\Crius\MemberAccountServer\RechargeCardResponse;
use Omy\Crius\MemberAccountServer\RefundBillRequest;
use Omy\Crius\MemberAccountServer\ReplaceCardRequest;
use Omy\Crius\MemberAccountServer\ReportBillDetailRequest;
use Omy\Crius\MemberAccountServer\ReportBillDetailResponse;
use Omy\Crius\MemberAccountServer\ReportPaymentRequest;
use Omy\Crius\MemberAccountServer\ReportPaymentResponse;
use Omy\Crius\MemberAccountServer\ReportPromotionResponse;
use Omy\Crius\MemberAccountServer\SearchCardsRequest;
use Omy\Crius\MemberAccountServer\SearchCardsResponse;
use Omy\Crius\MemberAccountServer\ShowAccountByCardIDRequest;
use Omy\Crius\MemberAccountServer\ShowAccountByCardIDResponse;
use Omy\Crius\MemberAccountServer\ShowAccountRequest;
use Omy\Crius\MemberAccountServer\ShowAccountResponse;
use Omy\Crius\MemberAccountServer\ShowBillByCodeRequest;
use Omy\Crius\MemberAccountServer\ShowBillByCodeResponse;
use Omy\Crius\MemberAccountServer\ShowBillRequest;
use Omy\Crius\MemberAccountServer\ShowBillResponse;
use Omy\Crius\MemberAccountServer\ShowCardByIDRequest;
use Omy\Crius\MemberAccountServer\ShowCardByIDResponse;
use Omy\Crius\MemberAccountServer\ShowCardRequest;
use Omy\Crius\MemberAccountServer\ShowCardResponse;
use Omy\Crius\MemberAccountServer\TransferCardBalanceRequest;
use Omy\Crius\MemberAccountServer\UpdateAccountBalanceRequest;
use Omy\Crius\MemberAccountServer\UpdateAccountStatusRequest;
use Omy\Crius\MemberAccountServer\UpdateCardStatusRequest;
use Omy\Crius\MemberAccountServer\ValidateCardPasswordRequest;

class MemberAccountClient extends GrpcClient
{
    private const BASE_PATH = '/memberAccount.MemberAccountServer/';

    // 制卡
    public const MAKE_CARD = 'MakeCard';

    // 开卡
    public const ACTIVE_PRIMARY_CARD = 'ActivePrimaryCard';
    public const ACTIVE_SECONDARY_CARD = 'ActiveSecondaryCard';
    public const ACTIVE_BLANK_CARD = 'ActiveBlankCard';

    // 充值
    public const RECHARGE_CARD = 'RechargeCard';

    // 绑卡
    public const BIND_CARD = 'BindCard';

    // 账户列表
    public const GET_ACCOUNTS = 'GetAccounts';

    // 账户详情
    public const SHOW_ACCOUNT = 'ShowAccount';

    // 根据id获取会员信息
    public const GET_ACCOUNT_BY_MEMBERID = 'GetAccountByMemberID';

    // 根据code获取会员信息
    public const GET_ACCOUNT_BY_CARDCODE = 'GetAccountByCardCode';

    // 根据accountid获取卡列表
    public const GET_CARDS_BY_ACCOUNT = 'GetCardsByAccount';

    //冻结/解冻账户
    public const UPDATE_ACCOUNT_STATUS = 'UpdateAccountStatus';

    //账单流水
    public const GET_BILLS = 'GetBills';

    //账单流水
    public const SHOW_BILL = 'ShowBill';

    //更改账户余额
    public const UPDATE_ACCOUNT_BALANCE = 'UpdateAccountBalance';

    //增加扣款
    public const ADD_ACCOUNT_DEDUCTION = 'AddAccountDeduction';

    //获取卡状态
    public const SHOW_CARD = 'ShowCard';

    //获取带会员id的卡信息
    public const SHOW_CARD_BY_ID = 'ShowCardByID';

    //获取卡列表
    public const GET_CARDS = 'GetCards';

    //卡划账
    public const TRANSFER_CARD_BALANCE = 'TransferCardBalance';

    //卡挂失/找回
    public const UPDATE_CARD_STATUS = 'UpdateCardStatus';

    //补卡
    public const REPLACE_CARD = 'ReplaceCard';

    //注销卡
    public const CANCEL_CARD = 'CancelCard';

    // 验证卡
    public const VALIDATE_CARD = 'ValidateCardPassword';

    // 获取主卡会员信息
    public const SHOW_ACCOUNT_BY_CARD_ID = 'ShowAccountByCardID';

    //ReportPayment 收款汇总
    public const REPORT_PAYMENT = 'ReportPayment';

    //ReportPromotion
    public const REPORT_PROMOTION = 'ReportPromotion';

    //ReportBillDetail
    public const REPORT_BILL_DETAIL = 'ReportBillDetail';

    //通过卡号查流水
    public const SHOW_BILL_BY_CODE = 'ShowBillByCode';

    public const GET_ACCOUNT_BY_CARD_ID = 'GetAccountsByCardID';

    public const GET_MEMBER_ACCOUNTS = 'GetMemberAccounts';

    public const SEARCH_CARDS = 'SearchCards';

    //消费
    public const CONSUME = 'Consume';

    public const REFUND_BILL = 'RefundBill';

    protected array $actionMap = [
        self::REFUND_BILL => [
            RefundBillRequest::class,
            NoDataResponse::class
        ],
        self::CONSUME => [
            ConsumeRequest::class,
            ConsumeResponse::class
        ],
        self::SEARCH_CARDS => [
            SearchCardsRequest::class,
            SearchCardsResponse::class
        ],
        self::GET_ACCOUNT_BY_CARD_ID => [
            GetAccountsByCardIDRequest::class,
            GetAccountsByCardIDResponse::class
        ],
        self::SHOW_BILL_BY_CODE => [
            ShowBillByCodeRequest::class,
            ShowBillByCodeResponse::class
        ],
        self::MAKE_CARD => [
            MakeCardRequest::class,
            MakeCardResponse::class
        ],
        self::ACTIVE_PRIMARY_CARD => [
            ActivePrimaryCardRequest::class,
            ActivePrimaryCardResponse::class
        ],
        self::ACTIVE_SECONDARY_CARD => [
            ActiveSecondaryCardRequest::class,
            ActiveSecondaryCardResponse::class
        ],
        self::ACTIVE_BLANK_CARD => [
            ActiveBlankCardRequest::class,
            ActiveBlankCardResponse::class
        ],
        self::RECHARGE_CARD => [
            RechargeCardRequest::class,
            RechargeCardResponse::class
        ],
        self::BIND_CARD => [
            BindCardRequest::class,
            BindCardResponse::class
        ],
        self::GET_ACCOUNTS => [
            GetAccountsRequest::class,
            GetAccountsResponse::class
        ],
        self::SHOW_ACCOUNT => [
            ShowAccountRequest::class,
            ShowAccountResponse::class
        ],
        self::GET_ACCOUNT_BY_MEMBERID => [
            GetAccountByMemberIDRequest::class,
            GetMemberAccountResponse::class
        ],
        self::GET_ACCOUNT_BY_CARDCODE => [
            GetAccountByCardCodeRequest::class,
            GetAccountByCardCodeResponse::class
        ],
        self::GET_CARDS_BY_ACCOUNT => [
            GetCardsByAccountRequest::class,
            GetCardsByAccountResponse::class
        ],
        self::UPDATE_ACCOUNT_STATUS => [
            UpdateAccountStatusRequest::class,
            NoDataResponse::class
        ],
        self::GET_BILLS => [
            GetBillsRequest::class,
            GetBillsResponse::class
        ],
        self::SHOW_BILL => [
            ShowBillRequest::class,
            ShowBillResponse::class
        ],
        self::UPDATE_ACCOUNT_BALANCE => [
            UpdateAccountBalanceRequest::class,
            NoDataResponse::class
        ],
        self::ADD_ACCOUNT_DEDUCTION => [
            AddAccountDeductionRequest::class,
            NoDataResponse::class
        ],
        self::SHOW_CARD => [
            ShowCardRequest::class,
            ShowCardResponse::class
        ],
        self::SHOW_CARD_BY_ID => [
            ShowCardByIDRequest::class,
            ShowCardByIDResponse::class
        ],
        self::GET_CARDS => [
            GetCardsRequest::class,
            GetCardsResponse::class
        ],
        self::TRANSFER_CARD_BALANCE => [
            TransferCardBalanceRequest::class,
            NoDataResponse::class
        ],
        self::UPDATE_CARD_STATUS => [
            UpdateCardStatusRequest::class,
            NoDataResponse::class
        ],
        self::REPLACE_CARD => [
            ReplaceCardRequest::class,
            NoDataResponse::class
        ],
        self::CANCEL_CARD => [
            CancelCardRequest::class,
            NoDataResponse::class
        ],
        self::VALIDATE_CARD => [
            ValidateCardPasswordRequest::class,
            NoDataResponse::class
        ],
        self::SHOW_ACCOUNT_BY_CARD_ID => [
            ShowAccountByCardIDRequest::class,
            ShowAccountByCardIDResponse::class
        ],
        self::REPORT_PAYMENT => [
            ReportPaymentRequest::class,
            ReportPaymentResponse::class
        ],
        self::REPORT_PROMOTION => [
            ReportBillDetailRequest::class,
            ReportPromotionResponse::class
        ],
        self::REPORT_BILL_DETAIL => [
            ReportBillDetailRequest::class,
            ReportBillDetailResponse::class
        ],
        self::GET_MEMBER_ACCOUNTS => [
            GetMemberAccountsRequest::class,
            GetMemberAccountsResponse::class,
        ]
    ];

    protected function getBasePath()
    {
        return self::BASE_PATH;
    }
}
