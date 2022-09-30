# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/member-account/card.proto](#proto/member-account/card.proto)
    - [ActiveBlankCardRequest](#memberAccount.ActiveBlankCardRequest)
    - [ActiveBlankCardResponse](#memberAccount.ActiveBlankCardResponse)
    - [ActivePrimaryCardRequest](#memberAccount.ActivePrimaryCardRequest)
    - [ActivePrimaryCardResponse](#memberAccount.ActivePrimaryCardResponse)
    - [ActiveSecondaryCardRequest](#memberAccount.ActiveSecondaryCardRequest)
    - [ActiveSecondaryCardResponse](#memberAccount.ActiveSecondaryCardResponse)
    - [BillData](#memberAccount.BillData)
    - [BindCardRequest](#memberAccount.BindCardRequest)
    - [BindCardResponse](#memberAccount.BindCardResponse)
    - [CancelCardRequest](#memberAccount.CancelCardRequest)
    - [Card](#memberAccount.Card)
    - [CardList](#memberAccount.CardList)
    - [GetCardsByAccountRequest](#memberAccount.GetCardsByAccountRequest)
    - [GetCardsByAccountResponse](#memberAccount.GetCardsByAccountResponse)
    - [GetCardsRequest](#memberAccount.GetCardsRequest)
    - [GetCardsResponse](#memberAccount.GetCardsResponse)
    - [MakeCardRequest](#memberAccount.MakeCardRequest)
    - [MakeCardResponse](#memberAccount.MakeCardResponse)
    - [Payments](#memberAccount.Payments)
    - [PromotionCount](#memberAccount.PromotionCount)
    - [RechargeCardRequest](#memberAccount.RechargeCardRequest)
    - [RechargeCardResponse](#memberAccount.RechargeCardResponse)
    - [ReplaceCardRequest](#memberAccount.ReplaceCardRequest)
    - [ShowCardByIDRequest](#memberAccount.ShowCardByIDRequest)
    - [ShowCardByIDResponse](#memberAccount.ShowCardByIDResponse)
    - [ShowCardRequest](#memberAccount.ShowCardRequest)
    - [ShowCardResponse](#memberAccount.ShowCardResponse)
    - [UpdateCardStatusRequest](#memberAccount.UpdateCardStatusRequest)
    - [ValidateCardPasswordRequest](#memberAccount.ValidateCardPasswordRequest)
  
- [proto/member-account/member_account.proto](#proto/member-account/member_account.proto)
    - [Account](#memberAccount.Account)
    - [AccountList](#memberAccount.AccountList)
    - [AccountWithActionReason](#memberAccount.AccountWithActionReason)
    - [AccountWithTag](#memberAccount.AccountWithTag)
    - [AddAccountDeductionRequest](#memberAccount.AddAccountDeductionRequest)
    - [Bill](#memberAccount.Bill)
    - [BillsData](#memberAccount.BillsData)
    - [CardWithAccount](#memberAccount.CardWithAccount)
    - [GetAccountByCardCodeRequest](#memberAccount.GetAccountByCardCodeRequest)
    - [GetAccountByCardCodeResponse](#memberAccount.GetAccountByCardCodeResponse)
    - [GetAccountByMemberIDRequest](#memberAccount.GetAccountByMemberIDRequest)
    - [GetAccountsByCardIDRequest](#memberAccount.GetAccountsByCardIDRequest)
    - [GetAccountsByCardIDResponse](#memberAccount.GetAccountsByCardIDResponse)
    - [GetAccountsRequest](#memberAccount.GetAccountsRequest)
    - [GetAccountsResponse](#memberAccount.GetAccountsResponse)
    - [GetBillsRequest](#memberAccount.GetBillsRequest)
    - [GetBillsResponse](#memberAccount.GetBillsResponse)
    - [GetMemberAccountResponse](#memberAccount.GetMemberAccountResponse)
    - [MemberAccount](#memberAccount.MemberAccount)
    - [NoDataResponse](#memberAccount.NoDataResponse)
    - [Package](#memberAccount.Package)
    - [Product](#memberAccount.Product)
    - [ShowAccountRequest](#memberAccount.ShowAccountRequest)
    - [ShowAccountResponse](#memberAccount.ShowAccountResponse)
    - [ShowBillByCodeRequest](#memberAccount.ShowBillByCodeRequest)
    - [ShowBillByCodeResponse](#memberAccount.ShowBillByCodeResponse)
    - [ShowBillRequest](#memberAccount.ShowBillRequest)
    - [ShowBillResponse](#memberAccount.ShowBillResponse)
    - [Ticket](#memberAccount.Ticket)
    - [TransferCardBalanceRequest](#memberAccount.TransferCardBalanceRequest)
    - [UpdateAccountBalanceRequest](#memberAccount.UpdateAccountBalanceRequest)
    - [UpdateAccountStatusRequest](#memberAccount.UpdateAccountStatusRequest)
  
    - [MemberAccountServer](#memberAccount.MemberAccountServer)
  
- [proto/member-account/report.proto](#proto/member-account/report.proto)
    - [ReportBillDetailRequest](#memberAccount.ReportBillDetailRequest)
    - [ReportBillDetailResponse](#memberAccount.ReportBillDetailResponse)
    - [ReportBillDetailResponse.Data](#memberAccount.ReportBillDetailResponse.Data)
    - [ReportBillDetailResponse.Report](#memberAccount.ReportBillDetailResponse.Report)
    - [ReportBillDetailResponse.Report.PromotionOption](#memberAccount.ReportBillDetailResponse.Report.PromotionOption)
    - [ReportPaymentRequest](#memberAccount.ReportPaymentRequest)
    - [ReportPaymentResponse](#memberAccount.ReportPaymentResponse)
    - [ReportPaymentResponse.Payment](#memberAccount.ReportPaymentResponse.Payment)
    - [ReportPromotionResponse](#memberAccount.ReportPromotionResponse)
    - [ReportPromotionResponse.Data](#memberAccount.ReportPromotionResponse.Data)
    - [ReportPromotionResponse.Report](#memberAccount.ReportPromotionResponse.Report)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto/member-account/card.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/member-account/card.proto



<a name="memberAccount.ActiveBlankCardRequest"></a>

### ActiveBlankCardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| card_id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| payments | [Payments](#memberAccount.Payments) |  |  |
| password | [string](#string) |  |  |
| recommender | [string](#string) |  |  |
| promotions | [PromotionCount](#memberAccount.PromotionCount) | repeated |  |
| recharge_value | [int32](#int32) |  |  |






<a name="memberAccount.ActiveBlankCardResponse"></a>

### ActiveBlankCardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BillData](#memberAccount.BillData) |  |  |






<a name="memberAccount.ActivePrimaryCardRequest"></a>

### ActivePrimaryCardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone | [string](#string) |  |  |
| phone_code | [string](#string) |  |  |
| name | [string](#string) |  |  |
| card_id | [string](#string) |  |  |
| gender | [int32](#int32) |  |  |
| birthday | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| payments | [Payments](#memberAccount.Payments) |  |  |
| password | [string](#string) |  |  |
| recommender | [string](#string) |  |  |
| promotions | [PromotionCount](#memberAccount.PromotionCount) | repeated |  |
| recharge_value | [int32](#int32) |  |  |






<a name="memberAccount.ActivePrimaryCardResponse"></a>

### ActivePrimaryCardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BillData](#memberAccount.BillData) |  |  |






<a name="memberAccount.ActiveSecondaryCardRequest"></a>

### ActiveSecondaryCardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| card_id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| password | [string](#string) |  |  |
| recommender | [string](#string) |  |  |
| recharge_value | [int32](#int32) |  |  |
| primary_id | [string](#string) |  |  |
| primary_password | [string](#string) |  |  |
| primary_verified | [bool](#bool) |  |  |
| primary_account_id | [string](#string) |  |  |






<a name="memberAccount.ActiveSecondaryCardResponse"></a>

### ActiveSecondaryCardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BillData](#memberAccount.BillData) |  |  |






<a name="memberAccount.BillData"></a>

### BillData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bill_id | [string](#string) |  |  |
| bill_code | [string](#string) |  |  |
| account_id | [string](#string) |  |  |
| card_id | [string](#string) |  |  |
| card_code | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| change_value | [int32](#int32) |  |  |
| change_category | [string](#string) |  |  |
| change_type | [string](#string) |  |  |
| base_value | [int32](#int32) |  |  |
| gift_value | [int32](#int32) |  |  |
| products | [string](#string) |  |  |
| packages | [string](#string) |  |  |
| tickets | [string](#string) |  |  |
| payments | [Payments](#memberAccount.Payments) |  |  |
| staff_id | [string](#string) |  |  |
| operator_comment | [string](#string) |  |  |
| merchant_id | [string](#string) |  |  |
| created_at | [int32](#int32) |  |  |
| primary_id | [string](#string) |  |  |
| primary_code | [string](#string) |  |  |
| total_base_value | [int32](#int32) |  |  |
| total_gift_value | [int32](#int32) |  |  |
| total_products | [string](#string) |  |  |
| total_packages | [string](#string) |  |  |
| total_tickets | [string](#string) |  |  |






<a name="memberAccount.BindCardRequest"></a>

### BindCardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| card_id | [string](#string) |  |  |
| account_id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| password | [string](#string) |  |  |
| recommender | [string](#string) |  |  |






<a name="memberAccount.BindCardResponse"></a>

### BindCardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="memberAccount.CancelCardRequest"></a>

### CancelCardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| card_id | [string](#string) |  |  |
| bank_account | [string](#string) |  |  |
| bank_name | [string](#string) |  |  |
| money_receiver | [string](#string) |  |  |
| reason | [string](#string) |  |  |






<a name="memberAccount.Card"></a>

### Card



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| category | [string](#string) |  |  |
| code | [string](#string) |  |  |
| create_branch_id | [string](#string) |  |  |
| create_staff_id | [string](#string) |  |  |
| status | [string](#string) |  |  |
| opened_at | [int32](#int32) |  |  |
| open_operator | [string](#string) |  |  |
| account_id | [string](#string) | repeated |  |
| open_branch_id | [string](#string) |  |  |
| member_id | [string](#string) |  |  |
| sub_category | [string](#string) |  |  |
| primary_id | [string](#string) |  |  |






<a name="memberAccount.CardList"></a>

### CardList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cards | [Card](#memberAccount.Card) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="memberAccount.GetCardsByAccountRequest"></a>

### GetCardsByAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_id | [string](#string) |  |  |






<a name="memberAccount.GetCardsByAccountResponse"></a>

### GetCardsByAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [Card](#memberAccount.Card) | repeated |  |






<a name="memberAccount.GetCardsRequest"></a>

### GetCardsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| category | [string](#string) |  |  |
| status | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="memberAccount.GetCardsResponse"></a>

### GetCardsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [CardList](#memberAccount.CardList) |  |  |






<a name="memberAccount.MakeCardRequest"></a>

### MakeCardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| code | [string](#string) |  |  |






<a name="memberAccount.MakeCardResponse"></a>

### MakeCardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="memberAccount.Payments"></a>

### Payments



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wechat | [int32](#int32) |  |  |
| cash | [int32](#int32) |  |  |
| alipay | [int32](#int32) |  |  |
| card | [int32](#int32) |  |  |






<a name="memberAccount.PromotionCount"></a>

### PromotionCount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| count | [int32](#int32) |  |  |






<a name="memberAccount.RechargeCardRequest"></a>

### RechargeCardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| recharge_value | [int32](#int32) |  |  |
| branch_id | [string](#string) |  |  |
| promotions | [PromotionCount](#memberAccount.PromotionCount) | repeated |  |
| payments | [Payments](#memberAccount.Payments) |  |  |
| recommender | [string](#string) |  |  |
| card_id | [string](#string) |  |  |






<a name="memberAccount.RechargeCardResponse"></a>

### RechargeCardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BillData](#memberAccount.BillData) |  |  |






<a name="memberAccount.ReplaceCardRequest"></a>

### ReplaceCardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| curr_card_id | [string](#string) |  |  |
| new_card_id | [string](#string) |  |  |
| password | [string](#string) |  |  |
| payments | [Payments](#memberAccount.Payments) |  |  |
| open_branch_id | [string](#string) |  |  |






<a name="memberAccount.ShowCardByIDRequest"></a>

### ShowCardByIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="memberAccount.ShowCardByIDResponse"></a>

### ShowCardByIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [Card](#memberAccount.Card) |  |  |






<a name="memberAccount.ShowCardRequest"></a>

### ShowCardRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| card_code | [string](#string) |  |  |






<a name="memberAccount.ShowCardResponse"></a>

### ShowCardResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [Card](#memberAccount.Card) |  |  |






<a name="memberAccount.UpdateCardStatusRequest"></a>

### UpdateCardStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| card_id | [string](#string) |  |  |
| action | [string](#string) |  |  |






<a name="memberAccount.ValidateCardPasswordRequest"></a>

### ValidateCardPasswordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| card_code | [string](#string) |  |  |
| password | [string](#string) |  |  |





 

 

 

 



<a name="proto/member-account/member_account.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/member-account/member_account.proto



<a name="memberAccount.Account"></a>

### Account



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| member_id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| base_value | [int32](#int32) |  |  |
| gift_value | [int32](#int32) |  |  |
| products | [Product](#memberAccount.Product) | repeated |  |
| packages | [Package](#memberAccount.Package) | repeated |  |
| ticket | [Ticket](#memberAccount.Ticket) | repeated |  |
| status | [string](#string) |  |  |
| tag_id | [string](#string) |  |  |
| create_at | [int32](#int32) |  |  |
| category | [string](#string) |  |  |






<a name="memberAccount.AccountList"></a>

### AccountList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accounts | [Account](#memberAccount.Account) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="memberAccount.AccountWithActionReason"></a>

### AccountWithActionReason



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| member_id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| base_value | [int32](#int32) |  |  |
| gift_value | [int32](#int32) |  |  |
| products | [Product](#memberAccount.Product) | repeated |  |
| packages | [Package](#memberAccount.Package) | repeated |  |
| ticket | [Ticket](#memberAccount.Ticket) | repeated |  |
| status | [string](#string) |  |  |
| tag_id | [string](#string) |  |  |
| create_at | [int32](#int32) |  |  |
| action_reason | [string](#string) |  |  |
| category | [string](#string) |  |  |






<a name="memberAccount.AccountWithTag"></a>

### AccountWithTag



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [Account](#memberAccount.Account) |  |  |
| branch_ids | [string](#string) | repeated |  |
| tag_name | [string](#string) |  |  |






<a name="memberAccount.AddAccountDeductionRequest"></a>

### AddAccountDeductionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bill_number | [string](#string) |  |  |
| card_id | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| cost_value | [int32](#int32) |  |  |
| cost_packages | [Package](#memberAccount.Package) | repeated |  |
| cost_tickets | [Ticket](#memberAccount.Ticket) | repeated |  |
| reason | [string](#string) |  |  |
| account_ids | [string](#string) | repeated |  |






<a name="memberAccount.Bill"></a>

### Bill



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bill_code | [string](#string) |  |  |
| create_at | [int32](#int32) |  |  |
| branch | [string](#string) |  |  |
| packages | [Package](#memberAccount.Package) | repeated |  |
| cost_base_value | [int32](#int32) |  |  |
| cost_gift_value | [int32](#int32) |  |  |
| cost_products | [Product](#memberAccount.Product) | repeated |  |
| cost_tickets | [Ticket](#memberAccount.Ticket) | repeated |  |
| curr_base_value | [int32](#int32) |  |  |
| curr_gift_value | [int32](#int32) |  |  |
| curr_products | [Product](#memberAccount.Product) | repeated |  |
| curr_tickets | [Ticket](#memberAccount.Ticket) | repeated |  |
| change_category | [string](#string) |  |  |
| payments | [Payments](#memberAccount.Payments) |  |  |
| staff_id | [string](#string) |  |  |
| id | [string](#string) |  |  |
| base_value_left | [int32](#int32) |  |  |
| gift_value_left | [int32](#int32) |  |  |
| card_id | [string](#string) |  |  |
| card_code | [string](#string) |  |  |






<a name="memberAccount.BillsData"></a>

### BillsData



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bills | [Bill](#memberAccount.Bill) | repeated |  |
| total | [int32](#int32) |  |  |






<a name="memberAccount.CardWithAccount"></a>

### CardWithAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| card | [Card](#memberAccount.Card) |  |  |
| accounts | [Account](#memberAccount.Account) | repeated |  |






<a name="memberAccount.GetAccountByCardCodeRequest"></a>

### GetAccountByCardCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| card_code | [string](#string) |  |  |






<a name="memberAccount.GetAccountByCardCodeResponse"></a>

### GetAccountByCardCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [CardWithAccount](#memberAccount.CardWithAccount) |  |  |






<a name="memberAccount.GetAccountByMemberIDRequest"></a>

### GetAccountByMemberIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| member_id | [string](#string) |  |  |






<a name="memberAccount.GetAccountsByCardIDRequest"></a>

### GetAccountsByCardIDRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="memberAccount.GetAccountsByCardIDResponse"></a>

### GetAccountsByCardIDResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [AccountWithTag](#memberAccount.AccountWithTag) | repeated |  |






<a name="memberAccount.GetAccountsRequest"></a>

### GetAccountsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| status | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="memberAccount.GetAccountsResponse"></a>

### GetAccountsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [AccountList](#memberAccount.AccountList) |  |  |






<a name="memberAccount.GetBillsRequest"></a>

### GetBillsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account_id | [string](#string) |  |  |
| category | [string](#string) |  |  |
| offset | [int32](#int32) |  |  |
| limit | [int32](#int32) |  |  |






<a name="memberAccount.GetBillsResponse"></a>

### GetBillsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [BillsData](#memberAccount.BillsData) |  |  |






<a name="memberAccount.GetMemberAccountResponse"></a>

### GetMemberAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [CardWithAccount](#memberAccount.CardWithAccount) | repeated |  |






<a name="memberAccount.MemberAccount"></a>

### MemberAccount



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cards | [CardWithAccount](#memberAccount.CardWithAccount) | repeated |  |






<a name="memberAccount.NoDataResponse"></a>

### NoDataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |






<a name="memberAccount.Package"></a>

### Package



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| title | [string](#string) |  |  |
| number | [int32](#int32) |  |  |






<a name="memberAccount.Product"></a>

### Product



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| title | [string](#string) |  |  |
| number | [int32](#int32) |  |  |






<a name="memberAccount.ShowAccountRequest"></a>

### ShowAccountRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="memberAccount.ShowAccountResponse"></a>

### ShowAccountResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [AccountWithActionReason](#memberAccount.AccountWithActionReason) |  |  |






<a name="memberAccount.ShowBillByCodeRequest"></a>

### ShowBillByCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bill_code | [string](#string) |  |  |






<a name="memberAccount.ShowBillByCodeResponse"></a>

### ShowBillByCodeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [Bill](#memberAccount.Bill) |  |  |






<a name="memberAccount.ShowBillRequest"></a>

### ShowBillRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bill_id | [string](#string) |  |  |






<a name="memberAccount.ShowBillResponse"></a>

### ShowBillResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [Bill](#memberAccount.Bill) |  |  |






<a name="memberAccount.Ticket"></a>

### Ticket



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| title | [string](#string) |  |  |
| number | [int32](#int32) |  |  |
| type | [string](#string) |  |  |






<a name="memberAccount.TransferCardBalanceRequest"></a>

### TransferCardBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_account_id | [string](#string) |  |  |
| dest_card_id | [string](#string) |  |  |
| amount | [int32](#int32) |  |  |
| operate_branch_id | [string](#string) |  |  |






<a name="memberAccount.UpdateAccountBalanceRequest"></a>

### UpdateAccountBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bill_code | [string](#string) |  |  |
| branch_id | [string](#string) |  |  |
| base_value | [int32](#int32) |  |  |
| gift_value | [int32](#int32) |  |  |
| packages | [Package](#memberAccount.Package) | repeated |  |
| tickets | [Ticket](#memberAccount.Ticket) | repeated |  |
| reason | [string](#string) |  |  |






<a name="memberAccount.UpdateAccountStatusRequest"></a>

### UpdateAccountStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| action | [string](#string) |  |  |
| reason | [string](#string) |  |  |





 

 

 


<a name="memberAccount.MemberAccountServer"></a>

### MemberAccountServer


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| MakeCard | [MakeCardRequest](#memberAccount.MakeCardRequest) | [MakeCardResponse](#memberAccount.MakeCardResponse) |  |
| ActivePrimaryCard | [ActivePrimaryCardRequest](#memberAccount.ActivePrimaryCardRequest) | [ActivePrimaryCardResponse](#memberAccount.ActivePrimaryCardResponse) |  |
| ActiveSecondaryCard | [ActiveSecondaryCardRequest](#memberAccount.ActiveSecondaryCardRequest) | [ActiveSecondaryCardResponse](#memberAccount.ActiveSecondaryCardResponse) |  |
| ActiveBlankCard | [ActiveBlankCardRequest](#memberAccount.ActiveBlankCardRequest) | [ActiveBlankCardResponse](#memberAccount.ActiveBlankCardResponse) |  |
| RechargeCard | [RechargeCardRequest](#memberAccount.RechargeCardRequest) | [RechargeCardResponse](#memberAccount.RechargeCardResponse) |  |
| BindCard | [BindCardRequest](#memberAccount.BindCardRequest) | [BindCardResponse](#memberAccount.BindCardResponse) |  |
| GetCards | [GetCardsRequest](#memberAccount.GetCardsRequest) | [GetCardsResponse](#memberAccount.GetCardsResponse) |  |
| GetCardsByAccount | [GetCardsByAccountRequest](#memberAccount.GetCardsByAccountRequest) | [GetCardsByAccountResponse](#memberAccount.GetCardsByAccountResponse) |  |
| ShowCard | [ShowCardRequest](#memberAccount.ShowCardRequest) | [ShowCardResponse](#memberAccount.ShowCardResponse) |  |
| ShowCardByID | [ShowCardByIDRequest](#memberAccount.ShowCardByIDRequest) | [ShowCardByIDResponse](#memberAccount.ShowCardByIDResponse) |  |
| TransferCardBalance | [TransferCardBalanceRequest](#memberAccount.TransferCardBalanceRequest) | [NoDataResponse](#memberAccount.NoDataResponse) |  |
| UpdateCardStatus | [UpdateCardStatusRequest](#memberAccount.UpdateCardStatusRequest) | [NoDataResponse](#memberAccount.NoDataResponse) |  |
| ReplaceCard | [ReplaceCardRequest](#memberAccount.ReplaceCardRequest) | [NoDataResponse](#memberAccount.NoDataResponse) |  |
| CancelCard | [CancelCardRequest](#memberAccount.CancelCardRequest) | [NoDataResponse](#memberAccount.NoDataResponse) |  |
| ValidateCardPassword | [ValidateCardPasswordRequest](#memberAccount.ValidateCardPasswordRequest) | [NoDataResponse](#memberAccount.NoDataResponse) |  |
| GetAccountsByCardID | [GetAccountsByCardIDRequest](#memberAccount.GetAccountsByCardIDRequest) | [GetAccountsByCardIDResponse](#memberAccount.GetAccountsByCardIDResponse) |  |
| GetAccounts | [GetAccountsRequest](#memberAccount.GetAccountsRequest) | [GetAccountsResponse](#memberAccount.GetAccountsResponse) |  |
| ShowAccount | [ShowAccountRequest](#memberAccount.ShowAccountRequest) | [ShowAccountResponse](#memberAccount.ShowAccountResponse) |  |
| UpdateAccountStatus | [UpdateAccountStatusRequest](#memberAccount.UpdateAccountStatusRequest) | [NoDataResponse](#memberAccount.NoDataResponse) |  |
| UpdateAccountBalance | [UpdateAccountBalanceRequest](#memberAccount.UpdateAccountBalanceRequest) | [NoDataResponse](#memberAccount.NoDataResponse) |  |
| AddAccountDeduction | [AddAccountDeductionRequest](#memberAccount.AddAccountDeductionRequest) | [NoDataResponse](#memberAccount.NoDataResponse) |  |
| GetAccountByMemberID | [GetAccountByMemberIDRequest](#memberAccount.GetAccountByMemberIDRequest) | [GetMemberAccountResponse](#memberAccount.GetMemberAccountResponse) |  |
| GetAccountByCardCode | [GetAccountByCardCodeRequest](#memberAccount.GetAccountByCardCodeRequest) | [GetAccountByCardCodeResponse](#memberAccount.GetAccountByCardCodeResponse) |  |
| GetBills | [GetBillsRequest](#memberAccount.GetBillsRequest) | [GetBillsResponse](#memberAccount.GetBillsResponse) |  |
| ShowBill | [ShowBillRequest](#memberAccount.ShowBillRequest) | [ShowBillResponse](#memberAccount.ShowBillResponse) |  |
| ShowBillByCode | [ShowBillByCodeRequest](#memberAccount.ShowBillByCodeRequest) | [ShowBillByCodeResponse](#memberAccount.ShowBillByCodeResponse) |  |
| ReportPayment | [ReportPaymentRequest](#memberAccount.ReportPaymentRequest) | [ReportPaymentResponse](#memberAccount.ReportPaymentResponse) | ReportPayment 收款汇总 开卡充值收款方式汇总 |
| ReportPromotion | [ReportBillDetailRequest](#memberAccount.ReportBillDetailRequest) | [ReportPromotionResponse](#memberAccount.ReportPromotionResponse) | ReportPromotion 按优惠方案汇总 |
| ReportBillDetail | [ReportBillDetailRequest](#memberAccount.ReportBillDetailRequest) | [ReportBillDetailResponse](#memberAccount.ReportBillDetailResponse) | ReportBillDetail 开卡/充值 详情 |

 



<a name="proto/member-account/report.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/member-account/report.proto



<a name="memberAccount.ReportBillDetailRequest"></a>

### ReportBillDetailRequest
ReportBillDetailRequest 开卡/充值优惠方案汇总，开卡详情，充值详情


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| begin_time | [int64](#int64) |  |  |
| end_time | [int64](#int64) |  |  |
| with_page | [bool](#bool) |  |  |
| report_type | [string](#string) | repeated |  |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |






<a name="memberAccount.ReportBillDetailResponse"></a>

### ReportBillDetailResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ReportBillDetailResponse.Data](#memberAccount.ReportBillDetailResponse.Data) |  |  |






<a name="memberAccount.ReportBillDetailResponse.Data"></a>

### ReportBillDetailResponse.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [ReportBillDetailResponse.Report](#memberAccount.ReportBillDetailResponse.Report) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="memberAccount.ReportBillDetailResponse.Report"></a>

### ReportBillDetailResponse.Report



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| card_id | [string](#string) |  |  |
| card_code | [string](#string) |  |  |
| payments | [Payments](#memberAccount.Payments) |  |  |
| promotion_options | [ReportBillDetailResponse.Report.PromotionOption](#memberAccount.ReportBillDetailResponse.Report.PromotionOption) | repeated |  |
| recharge_value | [int32](#int32) |  |  |






<a name="memberAccount.ReportBillDetailResponse.Report.PromotionOption"></a>

### ReportBillDetailResponse.Report.PromotionOption



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| promotion_option_id | [string](#string) |  |  |
| promotion_option_name | [string](#string) |  |  |
| count | [int32](#int32) |  |  |
| recharge_value | [int32](#int32) |  |  |






<a name="memberAccount.ReportPaymentRequest"></a>

### ReportPaymentRequest
ReportPaymentRequest 收款汇总 开卡充值收款金额按付款方式汇总


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| branch_id | [string](#string) |  |  |
| begin_time | [int64](#int64) |  |  |
| end_time | [int64](#int64) |  |  |






<a name="memberAccount.ReportPaymentResponse"></a>

### ReportPaymentResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ReportPaymentResponse.Payment](#memberAccount.ReportPaymentResponse.Payment) | repeated |  |






<a name="memberAccount.ReportPaymentResponse.Payment"></a>

### ReportPaymentResponse.Payment



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| way | [string](#string) |  |  |
| way_desc | [string](#string) |  |  |
| open_fee | [int32](#int32) |  |  |
| recharge_fee | [int32](#int32) |  |  |






<a name="memberAccount.ReportPromotionResponse"></a>

### ReportPromotionResponse
ReportPromotionResponse 开卡充值汇总 按充值方案汇总


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| error_code | [int32](#int32) |  |  |
| error_message | [string](#string) |  |  |
| data | [ReportPromotionResponse.Data](#memberAccount.ReportPromotionResponse.Data) |  |  |






<a name="memberAccount.ReportPromotionResponse.Data"></a>

### ReportPromotionResponse.Data



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| data | [ReportPromotionResponse.Report](#memberAccount.ReportPromotionResponse.Report) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="memberAccount.ReportPromotionResponse.Report"></a>

### ReportPromotionResponse.Report



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| promotion_option_id | [string](#string) |  |  |
| promotion_option_name | [string](#string) |  |  |
| total | [int32](#int32) |  |  |
| open_total | [int32](#int32) |  |  |
| recharge_total | [int32](#int32) |  |  |
| total_value | [int32](#int32) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

