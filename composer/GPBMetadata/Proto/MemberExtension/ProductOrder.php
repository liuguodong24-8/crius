<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/member-extension/product_order.proto

namespace GPBMetadata\Proto\MemberExtension;

class ProductOrder
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
?
*proto/member-extension/product_order.protomemberExtension"?
GetProductOrdersRequest
phone (	
status (	

date_begin (
date_end (
offset (
limit (
	member_id (	"z
GetProductOrdersResponse

error_code (
error_message (	3
data (2%.memberExtension.GetProductOrdersData"c
GetProductOrdersData<
orders (2,.memberExtension.ProductOrderWithProductInfo
total ("?
ProductOrderWithProductInfo,
order (2.memberExtension.ProductOrder
product_name (	
product_code (	
point (
price (

line_price (
images (	
videos (	
graphic_detail	 (	
describe
 (	"8
ShowProductOrderRequest

id (	
	member_id (	"?
ShowProductOrderResponse

error_code (
error_message (	:
data (2,.memberExtension.ProductOrderWithProductInfo"?
ProductOrder

id (	

order_code (	
	member_id (	

product_id (	
quantity (
point (
price (
status (	
describe	 (	
express_company
 (	
express_code (	
paid_at (
delivering_at (
trade_id (	
point_bill_id (	

created_at (
receive_name (	
receive_phone (	
receive_phone_code (	
receive_address (	
keep_at (
received_at (
shipping_point (
shipping_price (
express_company_cn (	
extra (	"?
UpdateProductOrderStatusRequest

id (	
status (	
trade_id (	
express_company (	
express_code (	
express_company_cn (	
transaction_id (	"^
CreateProductOrderRequest,
order (2.memberExtension.ProductOrder
ShippingWay (	"t
CreateProductOrderResponse

error_code (
error_message (	+
data (2.memberExtension.ProductOrderB+Z./proto;proto?Omy\\Crius\\ExtensionServerbproto3'
        , true);

        static::$is_initialized = true;
    }
}

